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

var weddingList = []Wedding{
	{
		ID:          1,
		Title:       "Micro-ondas",
		Description: "Micro-ondas Electrolux Prata com Painel Integrado 31l Mi41s 220v",
		Picture:     "https://m.media-amazon.com/images/I/51jWG1vyGyL._AC_SL1500_.jpg",
		Link:        "https://www.amazon.com.br/dp/B076XC1QX7",
		Available:   true,
	},
	{
		ID:          2,
		Title:       "Cafeteira Eletrica",
		Description: "Cafeteira Eletrica Electrolux Programavel Com Timer 30 Xicaras Experience ECM 25-220V",
		Picture:     "https://m.media-amazon.com/images/I/51OyvPACDML._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B09SVPS7S8",
		Available:   true,
	},
	{
		ID:          3,
		Title:       "Torradeira",
		Description: "Torradeira tostador Electrolux 8 niveis de tostagem função descongelar reaquecer cancelar bandeja coletora migalhas botões luminosos ETS25 inox 220v",
		Picture:     "https://m.media-amazon.com/images/I/51pcpIXaj0L._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B09CBT3BY5",
		Available:   true,
	},
	{
		ID:          4,
		Title:       "Sanduicheira",
		Description: "Sanduicheira grill chapa ondulada elétrico Electrolux antiaderente inox ESG20 220v",
		Picture:     "https://m.media-amazon.com/images/I/511Bqbpr12L._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B0CCBY3YMC",
		Available:   true,
	},
	{
		ID:          5,
		Title:       "Xícaras",
		Description: "Conjunto De 6 Xícaras Grandes 220 Ml Com Pires Ryo Pink Sand - Oxford",
		Picture:     "https://m.media-amazon.com/images/I/41MufvzLJ0L._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B07T39LN9D",
		Available:   true,
	},
	{
		ID:          6,
		Title:       "Taças",
		Description: "Jogo de 6 Taças Vinho Tinto Xtra Cristal 460ml A24cm",
		Picture:     "https://m.media-amazon.com/images/I/31EkEkILE5L._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B0B5LMNG3R",
		Available:   true,
	},
	{
		ID:          7,
		Title:       "Airfryer",
		Description: "Fritadeira Elétrica Grand Airfryer 4L Electrolux EAF30 Grafite",
		Picture:     "https://m.media-amazon.com/images/I/51qMgG4EBUL._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B09NRYJ67D",
		Available:   true,
	},
	{
		ID:          8,
		Title:       "Assadeiras",
		Description: "Jogo de assadeiras Marinex Opaline com três peças",
		Picture:     "https://m.media-amazon.com/images/I/51rrb7dMmGL._AC_SL1500_.jpg",
		Link:        "https://www.amazon.com.br/dp/B087699KMS",
		Available:   true,
	},
	{
		ID:          9,
		Title:       "Travessa",
		Description: "Conjunto 3 Travessa Melamina Branca Vasilhas Retangular Pequena Media e Grande de Mesa Posta Jogo de Tigelas Resistentes de Servir Arroz Macarrão e Salada. Ideal para Sobremesa Petiscos Frutas",
		Picture:     "https://m.media-amazon.com/images/I/41BQtWzApnL._AC_SL1000_.jpg",
		Link:        "https://www.amazon.com.br/dp/B0D8NLYL2T",
		Available:   true,
	},
	{
		ID:          10,
		Title:       "Assadeiras",
		Description: "Conjunto de Assadeiras 5 Peças, Assadeira, Assadeira Antiaderente, Forma, Assadeira, Assadeira Antiaderente, Forma, Forma Assadeira, Formas e Assadeiras, Assadeiras, Forma Cupcake",
		Picture:     "https://m.media-amazon.com/images/I/7158Fu0Ey6L._AC_SL1500_.jpg",
		Link:        "https://www.amazon.com.br/dp/B0DQD1GHSP",
		Available:   true,
	},
	{
		ID:          11,
		Title:       "Potes",
		Description: "Kit 10 Potes Herméticos C/Trava de Segurança Vedação de Borracha Porta Mantimentos Empilháveis + Etiquetas Lousa para Identificação",
		Picture:     "https://m.media-amazon.com/images/I/717+18FOhgL._AC_SL1200_.jpg",
		Link:        "https://www.amazon.com.br/dp/B0CZYK2VMG",
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
