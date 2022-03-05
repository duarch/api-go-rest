package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	d "github.com/duarch/go-rest-api/database"
	m "github.com/duarch/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []m.Personalidade
	d.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
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
