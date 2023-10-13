package repositories

import (
	"context"

	"github.com/dungnh3/mfv-codingchallenge/internal/models"
)

type Repository interface {
	Ping() error
	Transaction(txFunc func(Repository) error) error
	UserRepository
}

type UserRepository interface {
	UpsertUser(ctx context.Context, user *models.User) error
	UpsertUserAccount(ctx context.Context, userAccount *models.UserAccount) error
	GetUser(ctx context.Context, userId int64) (*models.User, error)
	GetAccount(ctx context.Context, accountId int64) (*models.UserAccount, error)
	ListAccounts(ctx context.Context, userId int64) ([]*models.UserAccount, error)
}
