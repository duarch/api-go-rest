package main

import r "github.com/duarch/go-rest-api/routes"

func main() {
	println("Iniciando servidor rest com Go...")
	r.HandleRequests()
}
