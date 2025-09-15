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
	Title:          fmt.Sprintf("%s & %s", nameBride, nameGroom),
	Description:    "Bem-vindos ao nosso site! Criamos este espaço para compartilhar com vocês todos os detalhes sobre a organização do nosso casamento. Estamos radiantes de felicidade e contamos com a presença de cada um de vocês neste momento tão especial! Pedimos que, por gentileza, confirmem sua presença entrando em contato conosco pelo WhatsApp. A sua confirmação é muito importante para que possamos preparar tudo com muito carinho. Disponibilizamos uma Lista de Casamento. Sintam-se à vontade para escolher qualquer item, seja em lojas físicas ou nos sites indicados, ou até mesmo contribuir por meio de PIX. Estamos ansiosos para celebrar este grande dia ao lado de vocês!",
	WhatsAppGroom:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappGroom, url.QueryEscape(messageGroom)),
	WhatsAppBride:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappBride, url.QueryEscape(messageBride)),
	NameGroom:      nameGroom,
	NameBride:      nameBride,
	CellPhoneGroom: whatsappGroom,
	CellPhoneBride: whatsappBride,
}

func loadWeddingPhotosFromJSON(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var photos []string
	err = json.Unmarshal(data, &photos)
	return photos, err
}

func loadWeddingListFromJSON(path string) ([]Wedding, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var list []Wedding
	err = json.Unmarshal(data, &list)
	return list, err
}

func main() {
	godotenv.Load()

	list, err := loadWeddingListFromJSON("wedding_list.json")
	if err != nil {
		fmt.Println("Erro ao carregar wedding_list.json:", err)
		os.Exit(1)
	}
	weddingList := list

	photos, err := loadWeddingPhotosFromJSON("wedding_photos.json")
	if err != nil {
		fmt.Println("Erro ao carregar wedding_photos.json:", err)
		os.Exit(1)
	}
	weddingPhotos := WeddingPhotos{Pictures: photos}

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
