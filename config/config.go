package config

import (
	"github.com/spf13/viper"
	"time"
)

var cfg = &Config{}

func Init() error {
	v := viper.New()

	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	return v.Unmarshal(&cfg)
}

func GetConfig() *Config {
	return cfg
}

type Config struct {
	Cookie   *Cookie   `json:"cookie"`
	Logger   *Logger   `json:"logger"`
	Postgres *Postgres `json:"postgres"`
	Redis    *Redis    `json:"redis"`
	Server   *Server   `json:"server"`
}

type Cookie struct {
	Name     string        `json:"name"`
	MaxAge   int           `json:"maxAge"`
	Secure   bool          `json:"secure"`
	HTTPOnly bool          `json:"httpOnly"`
	Expire   time.Duration `json:"expire"`
}

type Logger struct {
	Development       bool   `json:"development" validate:"required"`
	DisableCaller     bool   `json:"DisableCaller" validate:"required"`
	DisableStacktrace bool   `json:"DisableStacktrace" validate:"required"`
	Encoding          string `json:"Encoding" validate:"required"`
	Level             string `json:"Level" validate:"required"`
}

type Postgres struct {
	Host     string           `json:"host" validate:"required"`
	Port     string           `json:"port" validate:"required"`
	User     string           `json:"user" validate:"required"`
	Password string           `json:"-" validate:"required"`
	DBName   string           `json:"DBName" validate:"required"`
	SSLMode  string           `json:"sslMode" validate:"required"`
	PgDriver string           `json:"pgDriver" validate:"required"`
	Settings PostgresSettings `json:"settings"`
}

type PostgresSettings struct {
	MaxOpenConns    int           `json:"maxOpenConns" validate:"required,min=1"`
	ConnMaxLifetime time.Duration `json:"connMaxLifetime" validate:"required,min=1"`
	MaxIdleConns    int           `json:"maxIdleConns" validate:"required,min=1"`
	ConnMaxIdleTime time.Duration `json:"connMaxIdleTime" validate:"required,min=1"`
}

type Redis struct {
	Host         string `validate:"required"`
	Port         string `validate:"required"`
	MinIdleConns int    `validate:"required"`
	PoolSize     int    `validate:"required"`
	PoolTimeout  int    `validate:"required"`
	Password     string `validate:"required"`
	DB           int
}

type Server struct {
	Port string `json:"port" validate:"required"`
}
