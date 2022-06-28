package repository

import (
	_ "context"
	//"kevinPicon/go/rest-ws/models"
)

type UserRepository interface {
	//InsertUser(ctx context.Context, user *models.User) error
	Close() error
}

var impl UserRepository

func SetRepository(repo UserRepository) {
	impl = repo
}

//func ListPost(ctx context.Context, page uint64) ([]*models.Post, error) {
//	return impl.ListPost(ctx, page)
//}

func Close() error {
	return impl.Close()
}
