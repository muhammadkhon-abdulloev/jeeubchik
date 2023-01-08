package session

import "context"

func (s *Service) DeleteSessionByKey(ctx context.Context, key string) error {
	return s.cacheRepo.DeleteByKey(ctx, key)
}
