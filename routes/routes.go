package routes

import (
	"log"
	"net/http"

	c "github.com/duarch/go-rest-api/controllers"
	"github.com/duarch/go-rest-api/middleware"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", c.Home)
	r.HandleFunc("/api/personalidades", c.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", c.RetornaUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades", c.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", c.DeletaUmaPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", c.EditaPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
