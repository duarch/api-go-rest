package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	m "github.com/duarch/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(m.Personalidades)
}

func Retornaumapersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for _, item := range m.Personalidades {
		if strconv.Itoa(item.Id) == id {
			json.NewEncoder(w).Encode(item)
		}
	}

}
