package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gildemberg-santos/cerimony-flow/internal/model"
)

func SettingsHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Requisição para /settings")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(model.NewSettings())
	_, err := w.Write(resp)
	if err != nil {
		return
	}
}
