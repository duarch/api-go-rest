package main

import (
	m "github.com/duarch/go-rest-api/models"
	r "github.com/duarch/go-rest-api/routes"
)

func main() {
	m.Personalidades = []m.Personalidade{
		{Nome: "Darth Vader", Historia: "O senhor dos Sith"},
		{Nome: "Obi-Wan Kenobi", Historia: "O mais poderoso Jedi"},
	}

	println("Iniciando servidor rest com Go...")
	r.HandleRequests()
}
