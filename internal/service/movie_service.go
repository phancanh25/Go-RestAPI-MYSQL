package service

import (
	"context"
	"database/sql"
	"fmt"
	q "github.com/core-go/sql"
	. "go-service/internal/filter"
	. "go-service/internal/model"
	"reflect"
	"strings"
)

type MovieService interface {
	All(ctx context.Context) ([]Movie, error)
	Load(ctx context.Context, id string) (*Movie, error)
	Insert(ctx context.Context, movie *Movie) (int64, error)
	Update(ctx context.Context, movie *Movie) (int64, error)
	Patch(ctx context.Context, movie map[string]interface{}) (int64, error)
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

func (m *movieService) Load(ctx context.Context, id string) (*Movie, error) {
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

func (m *movieService) Insert(ctx context.Context, movie *Movie) (int64, error) {
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

func (m *movieService) Update(ctx context.Context, movie *Movie) (int64, error) {
	query := "update movies set name = ?, watched = ?"
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err1 := stmt.ExecContext(ctx, movie.Name, movie.Watched)
	if err1 != nil {
		return -1, err1
	}
	return res.RowsAffected()
}

func (m *movieService) Patch(ctx context.Context, movie map[string]interface{}) (int64, error) {
	movieType := reflect.TypeOf(Movie{})
	jsonColumnMap := q.MakeJsonColumnMap(movieType)
	colMap := q.JSONToColumns(movie, jsonColumnMap)
	keys, _ := q.FindPrimaryKeys(movieType)
	query, agrs := q.BuildToPatch("movies", colMap, keys, q.BuildParam)
	res, err := m.DB.ExecContext(ctx, query, agrs...)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (m movieService) Delete(ctx context.Context, id string) (int64, error) {
	query := "delete from movie where id = ?"
	stmt, err := m.DB.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return -1, err
	}
	return res.RowsAffected()
}

func (m movieService) Search(ctx context.Context, filter MovieFilter) (*ResultMovie, error) {
	query, param := BuildMovieQuery(filter, m.BuildParam)
	rows, err := m.DB.QueryContext(ctx, query, param...)
	if err != nil {
		return nil, err
	}
	var movies []Movie
	for rows.Next() {
		movie := Movie{}
		err := rows.Scan(&movie.Id, &movie.Name, &movie.Watched)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	query, params := BuildMovieCount(filter, m.BuildParam)
	rows, err = m.DB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, err
	}
	var total int64
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			return nil, err
		}
	}
	return &ResultMovie{List: movies, Total: total}, nil
}

func BuildMovieCount(filter MovieFilter, buildParam func(int) string) (string, []interface{}) {
	query := "select count* from movie"
	where, params := BuildMovieFilter(filter, buildParam)
	if len(where) > 0 {
		query = query + " where " + where
	}
	return query, params
}

func BuildMovieQuery(filter MovieFilter, buildParam func(int) string) (string, []interface{}) {
	query := "select * from movies"
	where, params := BuildMovieCount(filter, buildParam)
	if len(where) > 0 {
		query = query + " where " + where
	}
	return query, params
}

func BuildMovieFilter(filter MovieFilter, buildParam func(int) string) (string, []interface{}) {
	var condition []string
	var params []interface{}
	i := 1
	if len(filter.Id) > 0 {
		params = append(params, filter.Id)
		condition = append(condition, fmt.Sprint(`id = %s`, buildParam(i)))
		i++
	}
	if len(filter.Name) > 0 {
		q := "%" + filter.Name + "%"
		params = append(params, q)
		condition = append(condition, fmt.Sprintf(`name like %s`, buildParam(i)))
		i++
	}
	if len(condition) > 0 {
		return strings.Join(condition, " and "), params
	} else {
		return "", params
	}
}
