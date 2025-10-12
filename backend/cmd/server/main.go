package main

import (
	"log"
	"net/http"

	"github.com/gildemberg-santos/cerimony-flow/internal/router"
)

func main() {
	mux := http.NewServeMux()
	router.SetupRoutes(mux)

	log.Default().Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Default().Println("Erro ao carregar servidor.", err)
	}
}
