package app

import (
	"context"
	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func Route(r *mux.Router, ctx context.Context, config Config) error {
	app, err := NewApp(ctx, config)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	userPath := "/users"
	r.HandleFunc(userPath, app.UserHandler.All).Methods(GET)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Load).Methods(GET)
	r.HandleFunc(userPath, app.UserHandler.Insert).Methods(POST)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Update).Methods(PUT)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Patch).Methods(PATCH)
	r.HandleFunc(userPath+"/{id}", app.UserHandler.Delete).Methods(DELETE)
	r.HandleFunc(userPath+"/search", app.UserHandler.Search).Methods(POST)
	moviePath := "/movies"
	r.HandleFunc(moviePath, app.MovieHandler.All).Methods(GET)
	r.HandleFunc(moviePath+"/{id}", app.MovieHandler.Load).Methods(GET)
	r.HandleFunc(moviePath+"/{id}", app.MovieHandler.Insert).Methods(PUT)
	r.HandleFunc(moviePath+"/{id}", app.MovieHandler.Patch).Methods(PATCH)
	r.HandleFunc(moviePath+"/{id}", app.MovieHandler.Delete).Methods(DELETE)
	r.HandleFunc(moviePath+"/search", app.MovieHandler.Search).Methods(POST)
	return nil
}
