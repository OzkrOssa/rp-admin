package service

import (
	"context"

	"github.com/OzkrOssa/rp-admin/internal/core/domain"
	"github.com/OzkrOssa/rp-admin/internal/core/port"
	"github.com/OzkrOssa/rp-admin/internal/core/util"
)

type ClientService struct {
	repo  port.ClientRepository
	cache port.CacheRepository
}

func NewClientService(repo port.ClientRepository, cache port.CacheRepository) *ClientService {
	return &ClientService{
		repo,
		cache,
	}
}

func (cs *ClientService) Register(ctx context.Context, c *domain.Client) (*domain.Client, error) {

	client, err := cs.repo.CreateClient(ctx, c)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	cacheKey := util.GenerateCacheKey("client", client.ID)
	clientSerialized, err := util.Serialize(client)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.Set(ctx, cacheKey, clientSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = cs.cache.DeleteByPrefix(ctx, "clients:*")
	if err != nil {
		return nil, domain.ErrInternal
	}

	return client, nil
}
