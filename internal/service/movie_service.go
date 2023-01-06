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
	Load(ctx context.Context, id string) (*Movie, error)
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

func (m *movieService) All(ctx context.Context) ([]Movie, error) {
	query := "select * from movies"
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []Movie
	for rows.Next() {
		var movie Movie
		err = rows.Scan(&movie.Id, &movie.Name, &movie.Watched)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (m movieService) Load(ctx context.Context, id string) (*Movie, error) {
	query := "select * from movies where id = ? limit 1"
	rows, err := m.DB.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Watched)
		if err != nil {
			return nil, err
		}
		return &movie, nil
	}
	return nil, nil

}

func (m movieService) Insert(ctx context.Context, movie *Movie) (int64, error) {
	query := "insert into movies (id, name, watched) values (?,?.?)"
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err1 := stmt.ExecContext(ctx, movie.Id, movie.Name, movie.Watched)
	if err1 != nil {
		return -1, err1
	}
	return res.RowsAffected()
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
