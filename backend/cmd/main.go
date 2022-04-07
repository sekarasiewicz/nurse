package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sekarasiewicz/nurse/backend/pkg/handlers"
)

func main() {
	for _, route := range handlers.GetRoutes() {
		http.HandleFunc(route.Path, route.Handler)
	}

	fmt.Println("Listining on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
