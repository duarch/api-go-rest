package main

import (
	"github.com/duarch/go-rest-api/database"
	m "github.com/duarch/go-rest-api/models"
	r "github.com/duarch/go-rest-api/routes"
)

func main() {
	m.Personalidades = []m.Personalidade{
		{Id: 1, Nome: "Darth Vader", Historia: "O senhor dos Sith"},
		{Id: 2, Nome: "Obi-Wan Kenobi", Historia: "O mais poderoso Jedi"},
	}
	database.ConectaComBancoDeDados()
	println("Iniciando servidor rest com Go...")
	r.HandleRequests()
}
