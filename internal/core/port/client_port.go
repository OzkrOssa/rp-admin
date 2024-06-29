package port

import (
	"context"

	"github.com/OzkrOssa/rp-admin/internal/core/domain"
)

// UserRepository is an interface for interacting with client-related data
type ClientRepository interface {
	// CreateUser inserts a new client into the database
	CreateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	// GetUserByID selects a client by id
	GetClientByID(ctx context.Context, id uint64) (*domain.Client, error)
	// GetUserByDocument selects a document by document
	GetClientByDocument(ctx context.Context, document string) (*domain.Client, error)
	// ListUsers selects a list of clients with pagination
	ListClients(ctx context.Context, skip, limit uint64) ([]domain.Client, error)
	// UpdateUser updates a client
	UpdateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	// DeleteUser deletes a client
	DeleteClient(ctx context.Context, id uint64) error
}

// ClientService is an interface for interacting with client-related business logic
type ClientService interface {
	// Register registers a new client
	Register(ctx context.Context, client *domain.Client) (*domain.Client, error)
	// GetClient returns a client by id
	GetClientID(ctx context.Context, id uint64) (*domain.Client, error)
	// GetClient returns a client by document
	GetClientDocument(ctx context.Context, document string) (*domain.Client, error)
	// ListClients returns a list of clients with pagination
	ListClients(ctx context.Context, skip, limit uint64) ([]domain.Client, error)
	// UpdateClient updates a client
	UpdateClient(ctx context.Context, client *domain.Client) (*domain.Client, error)
	// DeleteClient deletes a client
	DeleteClient(ctx context.Context, id uint64) error
}
