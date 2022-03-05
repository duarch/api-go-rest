package routes

import (
	"log"
	"net/http"

	c "github.com/duarch/go-rest-api/controllers"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", c.Home)
	r.HandleFunc("/api/personalidades", c.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", c.Retornaumapersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
