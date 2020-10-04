package main

import (
	"api-rest-go/routes"
	"fmt"
	"log"
	"net/http"
)

// função principal para executar a api
func main() {
	fmt.Println("Listening on :8000")
	log.Fatal(http.ListenAndServe(":8000", routes.LoadRoat()))
}
