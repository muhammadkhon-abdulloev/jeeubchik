package main

import (
	"contactsList/config"
	accountHTTP "contactsList/http/account"
	contactHTTP "contactsList/http/contact"
	mdw "contactsList/http/middleware"
	"contactsList/pkg/logger"
	"contactsList/pkg/storage"
	"contactsList/repository/account"
	"contactsList/repository/cache"
	"contactsList/repository/contact"
	"contactsList/repository/item"
	account2 "contactsList/service/account"
	"contactsList/service/session"
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	echoApp := echo.New()
	echoApp.Use(middleware.Recover())
	psql, err := storage.InitPsqlDB(config.GetConfig().Postgres)
	if err != nil {
		panic(fmt.Errorf("storage.InitPsqlDB: %w", err))
	}
	redisClient, err := storage.InitRedisClient(config.GetConfig().Redis)
	if err != nil {
		panic(fmt.Errorf("storage.InitRedisClient: %w", err))

	}

	accountRepo := account.NewPGRepository(psql)
	contactRepo := contact.NewPGRepository(psql)
	cacheRepo := cache.NewRedisRepository(redisClient)
	item.NewPGRepository(psql)

	accountService := account2.NewService(accountRepo)
	sessionService := session.NewService(cacheRepo)

	accountHandler := accountHTTP.NewHandler(accountService, sessionService)
	contactHandler := contactHTTP.NewHandler(contactRepo)
	mw := mdw.NewMiddleware(accountService, sessionService)

	apiGroup := echoApp.Group("/api")
	accountGroup := apiGroup.Group("/account")
	contactsGroup := apiGroup.Group("/contacts")
	accountHTTP.MapAccountRoutes(accountGroup, accountHandler, mw)
	contactHTTP.MapContactRoutes(contactsGroup, contactHandler, mw)
	go func() {
		if err = echoApp.Start(config.GetConfig().Server.Port); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(fmt.Errorf("srv.Run(): %w", err))
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err = echoApp.Shutdown(context.Background()); err != nil {
		panic(fmt.Errorf("srv.stop: %w", err))
	}
}

func init() {
	err := config.Init()
	if err != nil {
		panic(fmt.Errorf("config.Init: %w", err))
	}

	logger.InitApiLogger(config.GetConfig().Logger)
	logger.GetLogger().InitLogger()
}
