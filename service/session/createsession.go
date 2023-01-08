package session

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

func (s *Service) CreateSession(accountID uuid.UUID, exp time.Duration) (string, error) {
	sessionKey := uuid.NewString()
	err := s.cacheRepo.StoreToCache(context.Background(), sessionKey, accountID.String(), exp)
	if err != nil {
		return "", fmt.Errorf("s.cacheRepo.StoreToCache: %w", err)
	}

	return sessionKey, nil
}
