package routes

import (
	"log"
	"net/http"

	c "github.com/duarch/go-rest-api/controllers"
)

func HandleRequests() {
	http.HandleFunc("/", c.Home)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
