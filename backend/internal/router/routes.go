package router

import (
	"net/http"
	"os"

	"github.com/gildemberg-santos/cerimony-flow/internal/handler"
	"github.com/joho/godotenv"
)

func SetupRoutes(r *http.ServeMux) {
	_ = godotenv.Load()
	fs := http.FileServer(http.Dir(os.Getenv("PATH_REACT")))
	r.Handle("/", fs)
	r.HandleFunc("/settings", handler.SettingsHandler)
	r.HandleFunc("/wedding-list", handler.WeddingListHandler)
	r.HandleFunc("/wedding-photos", handler.WeddingPhotosHandler)
}
