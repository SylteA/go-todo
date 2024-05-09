package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"todo/api/store"
	"todo/api/utils"
)

func IncludeTodoRouter(prefix string, parent *mux.Router, s store.Store) http.Handler {
	router := parent.PathPrefix(prefix).Subrouter()

	router.Methods("GET").Path("/{id}").Handler(getTodoById(s))
	router.Methods("DELETE").Path("/{id}").Handler(deleteTodoById(s))
	router.Methods("GET").Handler(getTodos(s))
	router.Methods("POST").Handler(createTodo(s))
	return router
}

func getTodos(s store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		todos := s.GetAll()
		err := utils.Encode(w, r, 200, todos)

		if err != nil {
			log.Println(err)
		}
	})
}

func createTodo(s store.Store) http.Handler {
	type request struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := utils.Decode[request](r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		todo, err := s.Create(data.Name, data.Description)

		if err == nil {
			_ = utils.Encode(w, r, 201, todo)
			return
		} else {
			_ = utils.Encode(w, r, 409, err.Error())
		}
	})
}

func deleteTodoById(s store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id := params["id"]
		integer, err := strconv.Atoi(id)

		var todo *store.Todo
		var identifier = "name"

		if err == nil {
			todo = s.GetById(integer)
			identifier = "id"
		} else {
			todo = s.GetByName(id)
		}

		if todo == nil {
			err = utils.Encode(
				w, r, 404,
				map[string]string{"error": fmt.Sprintf("todo with %s %s not found", identifier, id)},
			)
			return
		} else {
			err = s.Delete(todo.ID)
		}

		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(202)
	})
}

func getTodoById(s store.Store) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id := params["id"]
		integer, err := strconv.Atoi(id)

		var todo *store.Todo
		var identifier = "name"

		if err == nil {
			todo = s.GetById(integer)
			identifier = "id"
		} else {
			todo = s.GetByName(id)
		}

		if todo == nil {
			err = utils.Encode(
				w, r, 404,
				map[string]string{"error": fmt.Sprintf("todo with %s %s not found", identifier, id)},
			)
		} else {
			err = utils.Encode(w, r, 200, todo)
		}

		if err != nil {
			log.Println(err)
		}
	})
}
