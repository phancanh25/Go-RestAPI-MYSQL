package service

import (
	"context"
	. "go-service/internal/model"
)

type MovieService interface {
	All(ctx context.Context) ([]Movie, error)
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, movie *Movie) (int64, error)
	Update(ctx context.Context, movie *Movie) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	Search(ctx context.Context, filter MovieFilter) (*Result, error)
}
