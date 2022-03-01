package routes

import (
	"log"
	"net/http"

	c "github.com/duarch/go-rest-api/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", c.Home)
	http.HandleFunc("/api/personalidades", c.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
