package routes

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todo/api/utils"
)

func IncludeGreetingHandler(prefix string, parent *mux.Router) {
	parent.Methods("POST").Path(prefix).Handler(handleGreeting())
}

func handleGreeting() http.Handler {
	type request struct {
		Name string
	}
	type response struct {
		Greeting string `json:"greeting"`
	}

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			data, err := utils.Decode[request](r)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = utils.Encode(w, r, 200, response{Greeting: fmt.Sprintf("Hello, %s!", data.Name)})

			if err != nil {
				log.Println(err)
			}
		})
}
