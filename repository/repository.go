package repository

import (
	"context"
	"kevinPicon/go/src/CvPro/models"
)

type UserRepository interface {
	GetUserByName(ctx context.Context, userName string) (bool, error)
	InsertUser(ctx context.Context, user *models.User) error
	Close() error
}

var impl UserRepository

func SetRepository(repo UserRepository) {
	impl = repo
}
func GetUserByName(ctx context.Context, userName string) (bool, error) {
	return impl.GetUserByName(ctx, userName)
}
func InsertUser(ctx context.Context, user *models.User) error {
	return impl.InsertUser(ctx, user)
}

func Close() error {
	return impl.Close()
}
