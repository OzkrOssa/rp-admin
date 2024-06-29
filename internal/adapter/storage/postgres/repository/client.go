package repository

import (
	"context"

	"github.com/OzkrOssa/rp-admin/internal/adapter/storage/postgres"
	"github.com/OzkrOssa/rp-admin/internal/core/domain"
)

type ClientRepository struct {
	db *postgres.DB
}

func NewClientRepository(db *postgres.DB) *ClientRepository {
	return &ClientRepository{db}
}

func (cr *ClientRepository) CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {

	return nil, nil
}

func (cr *ClientRepository) GetClientByID(ctx context.Context, id uint64) (*domain.Client, error) {
	return nil, nil
}

func (cr *ClientRepository) GetClientByDocument(ctx context.Context, id uint64) (*domain.Client, error) {
	return nil, nil
}

func (cr *ClientRepository) ListClients(ctx context.Context, id uint64) ([]domain.Client, error) {
	return nil, nil
}

func (cr *ClientRepository) UpdateClient(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	return nil, nil
}

func (cr *ClientRepository) DeleteClient(ctx context.Context, id uint64) error {
	return nil
}
