package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router
	shoppingItems []Item
}

func NewServer() *Server {
	var s = &Server{
		Router:        mux.NewRouter(),
		shoppingItems: []Item{},
	}
	return s
}

func (s *Server) createShoppingItem() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil
	}
}
