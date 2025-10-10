package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gildemberg-santos/cerimony-flow/internal/settings"
	"github.com/gildemberg-santos/cerimony-flow/internal/wedding_list"
	"github.com/gildemberg-santos/cerimony-flow/internal/wedding_photos"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Default().Println("Erro ao carregar .env file", err)
	}

	fs := http.FileServer(http.Dir(os.Getenv("PATH_REACT")))
	http.Handle("/", fs)

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Println("Requisição para /settings")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		resp, _ := json.Marshal(settings.NewSettings())
		_, err := w.Write(resp)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/wedding-list", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Println("Requisição para /wedding-list")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		resp, _ := json.Marshal(wedding_list.NewWeddingList())
		_, err := w.Write(resp)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/wedding-photos", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Println("Requisição para /wedding-photos")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		resp, _ := json.Marshal(wedding_photos.NewWeddingPhotos())
		_, err = w.Write(resp)
		if err != nil {
			return
		}
	})

	log.Default().Println("Server is running on port 8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Default().Println("Erro ao carregar servidor.", err)
	}
}
