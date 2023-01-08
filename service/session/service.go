package session

import (
	"contactsList/repository/cache"
)

type Service struct {
	cacheRepo cache.Repository
}

func NewService(
	cacheRepo cache.Repository,
) *Service {
	return &Service{
		cacheRepo: cacheRepo,
	}
}
