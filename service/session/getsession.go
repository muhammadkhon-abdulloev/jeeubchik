package session

import (
	"context"
	"fmt"
	"github.com/google/uuid"
)

func (s *Service) GetSession(ctx context.Context, sessionKey string) (*uuid.UUID, error) {
	id, err := s.cacheRepo.GetStringByKey(ctx, sessionKey)
	if err != nil {
		return nil, fmt.Errorf("s.cacheRepo.GetStringByKey: %w", err)
	}
	accountID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("uuid.Parse: %w", err)
	}
	return &accountID, nil
}
