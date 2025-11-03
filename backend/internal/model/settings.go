package model

import (
	"fmt"
	"net/url"
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
	KeyPixGroom    string `json:"key_pix_groom"`
	KeyPixBride    string `json:"key_pix_bride"`
}

func NewSettings() *Settings {
	var nameGroom = "Gildemberg"
	var nameBride = "Bruna"
	var whatsappGroom = "5585991365507"
	var whatsappBride = "5585994510355"
	var messageGroom = fmt.Sprintf("Oi %s, tudo bem?", nameGroom)
	var messageBride = fmt.Sprintf("Oi %s, tudo bem?", nameBride)
	var keyPixGroom = "5585991365507"
	var keyPixBride = "09861125345"

	return &Settings{
		Title:          fmt.Sprintf("%s & %s", nameBride, nameGroom),
		Description:    "Bem-vindos ao nosso site! Criamos este espaço para compartilhar com vocês todos os detalhes sobre a organização do nosso casamento. Estamos radiantes de felicidade e contamos com a presença de cada um de vocês neste momento tão especial! Pedimos que, por gentileza, confirmem sua presença entrando em contato conosco pelo WhatsApp. A sua confirmação é muito importante para que possamos preparar tudo com muito carinho. Disponibilizamos uma Lista de Casamento. Sintam-se à vontade para escolher qualquer item, seja em lojas físicas ou nos sites indicados, ou até mesmo contribuir por meio de PIX. Estamos ansiosos para celebrar este grande dia ao lado de vocês!",
		WhatsAppGroom:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappGroom, url.QueryEscape(messageGroom)),
		WhatsAppBride:  fmt.Sprintf("https://api.whatsapp.com/send?phone=%s&text=%s", whatsappBride, url.QueryEscape(messageBride)),
		NameGroom:      nameGroom,
		NameBride:      nameBride,
		CellPhoneGroom: whatsappGroom,
		CellPhoneBride: whatsappBride,
		KeyPixGroom:    keyPixGroom,
		KeyPixBride:    keyPixBride,
	}
}
