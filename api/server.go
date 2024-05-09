package api

import (
	"github.com/gorilla/mux"
	"todo/api/routes"
	"todo/api/store"
)

type Server struct {
	Router *mux.Router
	Store  store.Store
}

func NewServer(storageType store.StorageType) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Store:  nil,
	}

	switch storageType {
	default:
		newStore, _ := store.NewInMemoryStore()
		s.Store = newStore
	}

	return s
}

func (server *Server) AddRoutes() {
	routes.IncludeGreetingHandler("/greet", server.Router)
	routes.IncludeTodoRouter("/api/todo", server.Router, server.Store)
	routes.IncludeIndexRouter("/", server.Router)
}
