package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Settings struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	WhatsAppGroom  string `json:"whatsapp_groom"`
	WhatsAppBride  string `json:"whatsapp_bride"`
	NameGroom      string `json:"name_groom"`
	NameBride      string `json:"name_bride"`
	CellPhoneGroom string `json:"cell_phone_groom"`
	CellPhoneBride string `json:"cell_phone_bride"`
}

type Wedding struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Link        string `json:"link"`
	Available   bool   `json:"available"`
}

type WeddingList struct {
	WeddingList []Wedding `json:"wedding_list"`
}

type WeddingPhotos struct {
	Pictures []string `json:"pictures"`
}

var nameGroom = "Gildemberg"
var nameBride = "Bruna"
var whatsappGroom = "5585991365507"
var whatsappBride = "5585994510355"
var messageGroom = fmt.Sprintf("Oi %s, tudo bem?", nameGroom)
var messageBride = fmt.Sprintf("Oi %s, tudo bem?", nameBride)

var settings = Settings{
	Title:          fmt.Sprintf("%s & %s", nameGroom, nameBride),
	Description:    "Bem-vindo ao nosso site de casamento. Estamos muito animados para celebrar nosso dia especial com você. Por favor, confira nossa lista de casamento e confirme sua presença. ",
	WhatsAppGroom:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappGroom, url.QueryEscape(messageGroom)),
	WhatsAppBride:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappBride, url.QueryEscape(messageBride)),
	NameGroom:      nameGroom,
	NameBride:      nameBride,
	CellPhoneGroom: whatsappGroom,
	CellPhoneBride: whatsappBride,
}

var weddingList = []Wedding{
	{
		ID:          1,
		Title:       "Fire TV Stick HD",
		Description: "Fire TV Stick HD | Com controle remoto por voz com Alexa (inclui comandos de TV), controles de casa inteligente e streaming em HD",
		Picture:     "https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		Link:        "https://www.amazon.com.br/fire-tv-stick-hd/dp/B0CQMT33WX/",
		Available:   true,
	},
	{
		ID:          2,
		Title:       "Gift 2",
		Description: "Gift 2 description",
		Picture:     "https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		Link:        "https://www.google.com",
		Available:   true,
	},
	{
		ID:          3,
		Title:       "Gift 3",
		Description: "Gift 3 description",
		Picture:     "https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		Link:        "https://www.google.com",
		Available:   true,
	},
	{
		ID:          4,
		Title:       "Gift 4",
		Description: "Gift 4 description",
		Picture:     "https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		Link:        "https://www.google.com",
		Available:   true,
	},
	{
		ID:          5,
		Title:       "Gift 5",
		Description: "Gift 5 description",
		Picture:     "https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		Link:        "https://www.google.com",
		Available:   true,
	},
}

var weddingPhotos = WeddingPhotos{
	Pictures: []string{
		"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
	},
}

func main() {
	godotenv.Load()

	fs := http.FileServer(http.Dir(os.Getenv("PATH_REACT")))
	http.Handle("/", fs)

	http.HandleFunc("/settings", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(settings)
		w.Write(resp)
	})

	http.HandleFunc("/wedding-list", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		resp, _ := json.Marshal(WeddingList{WeddingList: weddingList})
		w.Write(resp)
	})

	http.HandleFunc("/wedding-list/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		idStr := r.URL.Path[len("/wedding-list/"):]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		for _, wedding := range weddingList {
			if wedding.ID == id {
				w.WriteHeader(http.StatusOK)
				resp, _ := json.Marshal(wedding)
				w.Write(resp)
				return
			}
		}

		http.Error(w, "Wedding not found", http.StatusNotFound)
	})

	http.HandleFunc("/wedding-photos", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		resp, _ := json.Marshal(weddingPhotos)
		w.Write(resp)
	})

	fmt.Println("Server is running on port 8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
