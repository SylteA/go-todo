package api

import (
	"flag"
	"github.com/gorilla/mux"
	"todo/api/routes"
	"todo/api/store"
)

type Server struct {
	Router *mux.Router
	Store  store.Store
}

var jsonFilePath = flag.String(
	"jsonFilePath", "data.json",
	"Used to override the file todos are stored in.",
)

func NewServer(storageType store.StorageType) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Store:  nil,
	}

	switch storageType {
	case store.JSON:
		newStore, _ := store.NewJsonStore(*jsonFilePath)
		s.Store = newStore
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
