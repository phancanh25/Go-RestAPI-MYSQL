package handler

import (
	"github.com/gorilla/mux"
	. "go-service/internal/service"
	"net/http"
)

type MovieHandler struct {
	servie MovieService
}

func NewMovieHandler(service MovieService) *MovieHandler {
	return &MovieHandler{servie: service}
}

func (h *MovieHandler) All(w http.ResponseWriter, r *http.Request) {
	res, err := h.servie.All(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	JSON(w, http.StatusOK, res)
}

func (h *MovieHandler) Load(w http.ResponseWriter, r http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
}
