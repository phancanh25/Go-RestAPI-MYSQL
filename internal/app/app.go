package app

import (
	"context"
	"github.com/core-go/health"
	s "github.com/core-go/health/sql"
	"github.com/core-go/sql"
	_ "github.com/go-sql-driver/mysql"

	"go-service/internal/handler"
	"go-service/internal/service"
)

const (
	CreateTableUser = `
	create table if not exists users (
	  id varchar(40) not null,
	  username varchar(120),
	  email varchar(120),
	  phone varchar(45),
	  date_of_birth date,
	  primary key (id)
	)`

	CreateTableMovie = `
	create table if not exists movies (
	  id varchar(40) not null,
	  name varchar(120),
	  watched tinyint,
	  primary key (id)
	)`
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	UserHandler   *handler.UserHandler
	MovieHandler  *handler.MovieHandler
}

func NewApp(ctx context.Context, config Config) (*ApplicationContext, error) {
	db, err := sql.OpenByConfig(config.Sql)
	if err != nil {
		return nil, err
	}

	stmtCreate := "create database if not exists masterdata"
	_, err = db.ExecContext(ctx, stmtCreate)
	if err != nil {
		return nil, err
	}

	stmtUseDB := "use masterdata"
	_, err = db.ExecContext(ctx, stmtUseDB)
	if err != nil {
		return nil, err
	}

	_, err = db.ExecContext(ctx, CreateTableUser)
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, CreateTableMovie)
	if err != nil {
		return nil, err
	}

	userService := service.NewUserService(db)
	userHandler := handler.NewUserHandler(userService)

	movieService := service.NewMovieService(db)
	movieHandler := handler.NewMovieHandler(movieService)

	sqlChecker := s.NewHealthChecker(db)
	healthHandler := health.NewHandler(sqlChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		UserHandler:   userHandler,
		MovieHandler:  movieHandler,
	}, nil
}
