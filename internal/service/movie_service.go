package service

import (
	"context"
	"database/sql"
	q "github.com/core-go/sql"
	. "go-service/internal/filter"
	. "go-service/internal/model"
)

type MovieService interface {
	All(ctx context.Context) ([]Movie, error)
	Load(ctx context.Context, id string) (*User, error)
	Insert(ctx context.Context, movie *Movie) (int64, error)
	Update(ctx context.Context, movie *Movie) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
	Search(ctx context.Context, filter MovieFilter) (*ResultMovie, error)
}

type movieService struct {
	DB         *sql.DB
	BuildParam func(int) string
}

func NewMovieService(db *sql.DB) MovieService {
	buildParam := q.GetBuild(db)
	return &movieService{DB: db, BuildParam: buildParam}
}

func (m movieService) All(ctx context.Context) ([]Movie, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Load(ctx context.Context, id string) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Insert(ctx context.Context, movie *Movie) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Update(ctx context.Context, movie *Movie) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Delete(ctx context.Context, id string) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m movieService) Search(ctx context.Context, filter MovieFilter) (*ResultMovie, error) {
	//TODO implement me
	panic("implement me")
}
