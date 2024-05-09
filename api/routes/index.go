package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo/api/utils"
)

func IncludeIndexRouter(prefix string, parent *mux.Router) {
	parent.Methods("GET").Path(prefix).Handler(index(parent))
}

func index(root *mux.Router) http.Handler {
	type Endpoint struct {
		Methods []string `json:"methods"`
		Path    string   `json:"path"`
		Name    string   `json:"name"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var endpoints []*Endpoint

		err := root.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
			path, err := route.GetPathTemplate()
			if err != nil {
				return err
			}

			methods, _ := route.GetMethods()

			endpoints = append(endpoints, &Endpoint{
				Path:    path,
				Methods: methods,
			})

			return nil
		})

		if err != nil {
			log.Println(err)
		}
		err = utils.Encode(w, r, 200, endpoints)
		if err != nil {
			log.Println(err)
		}
	})
}
