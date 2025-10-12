package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gildemberg-santos/cerimony-flow/internal/model"
)

func WeddingListHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Println("Requisição para /wedding-list")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(model.NewWeddingList())
	_, err := w.Write(resp)
	if err != nil {
		return
	}
}
