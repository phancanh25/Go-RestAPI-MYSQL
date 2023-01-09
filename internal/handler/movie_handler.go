package handler

import (
	"encoding/json"
	sv "github.com/core-go/service"
	"github.com/gorilla/mux"
	. "go-service/internal/filter"
	. "go-service/internal/model"
	. "go-service/internal/service"
	"net/http"
	"reflect"
)

type MovieHandler struct {
	service MovieService
}

func NewMovieHandler(service MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) All(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	res, err := h.service.Load(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.service.Insert(r.Context(), &movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(movie.Id) == 0 {
		movie.Id = id
	} else if id != movie.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	res, er2 := h.service.Update(r.Context(), &movie)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) != 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	var movie Movie
	movieType := reflect.TypeOf(movie)
	_, jsonMap, _ := sv.BuildMapField(movieType)
	body, er1 := sv.BuildMapAndStruct(r, &movie)
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	if len(movie.Id) == 0 {
		movie.Id = id
	} else if id != movie.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}
	json, er2 := sv.BodyToJsonMap(r, movie, body, []string{"id"}, jsonMap)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}

	res, er3 := h.service.Patch(r.Context(), json)
	if er3 != nil {
		http.Error(w, er3.Error(), http.StatusBadRequest)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	res, err := h.service.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Search(w http.ResponseWriter, r *http.Request) {
	var filter MovieFilter
	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := h.service.Search(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)

}