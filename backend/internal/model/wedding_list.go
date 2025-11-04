package model

import (
	"sort"
	"strings"
)

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

func NewWeddingList() *WeddingList {
	list := []Wedding{
		{
			Title:       "# Pix ou Dinheiro",
			Description: "Contribuição em dinheiro para ajudar na realização do nosso sonho! Qualquer valor é bem-vindo e será muito apreciado. Agradecemos de coração pela sua generosidade e por fazer parte deste momento tão especial em nossas vidas.",
			Picture:     "https://simsoft.com.br/storage/app/uploads/public/421/_pi/x-i/421_pix-interno.jpg",
			Link:        "",
			Available:   true,
		},
		{
			Title:       "Micro-ondas",
			Description: "Micro-ondas Electrolux Prata com Painel Integrado 31l Mi41s 220v",
			Picture:     "https://m.media-amazon.com/images/I/51jWG1vyGyL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B076XC1QX7",
			Available:   true,
		},
		{
			Title:       "Cafeteira Eletrica",
			Description: "Cafeteira Eletrica Electrolux Programavel Com Timer 30 Xicaras Experience ECM 25-220V",
			Picture:     "https://m.media-amazon.com/images/I/51OyvPACDML._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09SVPS7S8",
			Available:   false,
		},
		{
			Title:       "Torradeira",
			Description: "Torradeira tostador Electrolux 8 niveis de tostagem função descongelar reaquecer cancelar bandeja coletora migalhas botões luminosos ETS25 inox 220v",
			Picture:     "https://m.media-amazon.com/images/I/51pcpIXaj0L._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09CBT3BY5",
			Available:   false,
		},
		{
			Title:       "Sanduicheira",
			Description: "Sanduicheira grill chapa ondulada elétrico Electrolux antiaderente inox ESG20 220v",
			Picture:     "https://m.media-amazon.com/images/I/511Bqbpr12L._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0CCBY3YMC",
			Available:   false,
		},
		{
			Title:       "Xícaras",
			Description: "Conjunto De 6 Xícaras Grandes 220 Ml Com Pires Ryo Pink Sand - Oxford",
			Picture:     "https://m.media-amazon.com/images/I/41MufvzLJ0L._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B07T39LN9D",
			Available:   false,
		},
		{
			Title:       "Taças",
			Description: "Jogo de 6 Taças Vinho Tinto Xtra Cristal 460ml A24cm",
			Picture:     "https://m.media-amazon.com/images/I/31EkEkILE5L._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0B5LMNG3R",
			Available:   true,
		},
		{
			Title:       "Airfryer",
			Description: "Fritadeira Elétrica Grand Airfryer 4L Electrolux EAF30 Grafite",
			Picture:     "https://m.media-amazon.com/images/I/51qMgG4EBUL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09NRYJ67D",
			Available:   false,
		},
		{
			Title:       "Assadeiras",
			Description: "Jogo de assadeiras Marinex Opaline com três peças",
			Picture:     "https://m.media-amazon.com/images/I/51rrb7dMmGL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B087699KMS",
			Available:   true,
		},
		{
			Title:       "Travessa",
			Description: "Conjunto 3 Travessa Melamina Branca Vasilhas Retangular Pequena Media e Grande de Mesa Posta Jogo de Tigelas Resistentes de Servir Arroz Macarrão e Salada. ideal para Sobremesa Petiscos Frutas",
			Picture:     "https://m.media-amazon.com/images/I/41BQtWzApnL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0D8NLYL2T",
			Available:   true,
		},
		{
			Title:       "Assadeiras",
			Description: "Conjunto de Assadeiras 5 Peças, Assadeira, Assadeira Antiaderente, Forma, Assadeira, Assadeira Antiaderente, Forma, Forma Assadeira, Formas e Assadeiras, Assadeiras, Forma Cupcake",
			Picture:     "https://m.media-amazon.com/images/I/7158Fu0Ey6L._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DQD1GHSP",
			Available:   true,
		},
		{
			Title:       "Potes",
			Description: "Kit 10 Potes Herméticos C/Trava de Segurança Vedação de Borracha Porta Mantimentos Empilháveis + Etiquetas Lousa para Identificação",
			Picture:     "https://m.media-amazon.com/images/I/717+18FOhgL._AC_SL1200_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0CZYK2VMG",
			Available:   true,
		},
		{
			Title:       "Sapateira",
			Description: "Sapateira Articulada Safira 2 Portas Basculante Mell (Branco)",
			Picture:     "https://m.media-amazon.com/images/I/31picHZt-QL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0CC673K8D",
			Available:   true,
		},
		{
			Title:       "Travesseiro",
			Description: "Kit Travesseiro 04 Peças Fiber Team Antialérgico com Fibra Siliconizada - Ortobom",
			Picture:     "https://m.media-amazon.com/images/I/71fxsaGvfYL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B08TMDFMN1",
			Available:   true,
		},
		{
			Title:       "Colcha",
			Description: "Colcha de Cama Casal 3 peças Luxo Algodão Cinza Menina Dupla Face Masculino Feminino (Bege Rosada, Casal)",
			Picture:     "https://m.media-amazon.com/images/I/71zTINHgKqL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0FF9CCB3X",
			Available:   false,
		},
		{
			Title:       "Colcha",
			Description: "Kit Colcha Balle Para Cama King 400 Fios Dupla Face 3 Peças Cores Lisas Tecido Piquet Toque Mácio Acompanha Cobre Leito e 2 Porta Travesseiros (Tiffany, Casal King (260x280cm))",
			Picture:     "https://m.media-amazon.com/images/I/613JI8rSTML._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DPGPGY5B",
			Available:   false,
		},
		{
			Title:       "Colchão",
			Description: "Colchão Casal de espuma D28 Emma Basics 17cm – com tecnologia alemã, 100 noites de teste e 5 anos de garantia",
			Picture:     "https://m.media-amazon.com/images/I/7102jT-oGNL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0B8ZF9DR3",
			Available:   true,
		},
		{
			Title:       "Kit Banheiro",
			Description: "Kit Banheiro Lavabo Conjunto Acessórios lixeira escova 6 Peças Bambu Luxo Preto/Branco",
			Picture:     "https://down-br.img.susercontent.com/file/br-11134207-7r98o-mcjiyijarzua36.webp",
			Link:        "https://shopee.com.br/product/397571861/23393873729",
			Available:   true,
		},
		{
			Title:       "Aspirador",
			Description: "WAP Aspirador de Pó Robô ROBOT W90 3 em 1, Automático, 250ml, Sistema Antiqueda e Rodas Emborrachadas, 30W 3,6VDC Bivolt",
			Picture:     "https://m.media-amazon.com/images/I/716GA5MQQkL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0B9PSBNYL",
			Available:   false,
		},
		{
			Title:       "Batedeira",
			Description: "Batedeira Planetária, Mondial, Preto/Inox, 700W, 110V - BP-01P-B",
			Picture:     "https://m.media-amazon.com/images/I/811erab5-PL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B07WDCZWFW",
			Available:   false,
		},
		{
			Title:       "Panela pressão elétrica",
			Description: "Panela pressão elétrica Electrolux digital capacidade 6L silenciosa segura 10 travas segurança 15 receitas pré-programadas 3 níveis pressão PCC20 inox preto 127v por Rita Lobo",
			Picture:     "https://m.media-amazon.com/images/I/51zXHW5nVjL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B076HYKFL7",
			Available:   true,
		},
		{
			Title:       "Panificadora",
			Description: "Panificadora Automática Master Bread, Mondial, Preto, 700W, 220V - NPF-53",
			Picture:     "https://m.media-amazon.com/images/I/81G6rXDd9iL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B07L8X97SR",
			Available:   true,
		},
		{
			Title:       "Toalhas de Banho",
			Description: "Buddemeyer Jogo de Toalhas Amanhecer Banho Branco 5 peças",
			Picture:     "https://m.media-amazon.com/images/I/511VoVajAVL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DWY1CHR4",
			Available:   true,
		},
		{
			Title:       "Faqueiro",
			Description: "Faqueiro Tramontina Laguna em Aço Inox com Facas para Churrasco Acabamento em Brilho 24 Peças",
			Picture:     "https://m.media-amazon.com/images/I/614p7dWdJ0L._AC_SL1200_.jpg",
			Link:        "https://www.amazon.com.br/dp/B07KRNHSXV",
			Available:   false,
		},
		{
			Title:       "Kit Churrasco",
			Description: "Kit Churrasco com Tabua, Kit Churrasco, Kit Churrasco Inox, Kit Churrasco Tradicional Inox 6 Peças, Kit Tábua para Churrasco, Kit de Churrasco 6 Peças com Tábua de Bambu",
			Picture:     "https://m.media-amazon.com/images/I/61y3C7GbLlL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DQLBWH8R",
			Available:   false,
		},
		{
			Title:       "Tapete",
			Description: "Tapete Felpudo 1,40 X 2,00 Peludo 40mm de Pelagem - Cinza Mesclado",
			Picture:     "https://m.media-amazon.com/images/I/71mYs37EHML._AC_SL1280_.jpg",
			Link:        "https://www.amazon.com.br/dp/B08BYZRKTP",
			Available:   true,
		},
		{
			Title:       "Samsung Lava e Seca",
			Description: "Samsung Lava e Seca WD11M com Digital Inverter WD11M4473PX Inox Look 11/7kg - 220V",
			Picture:     "https://m.media-amazon.com/images/I/51dBND271dL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0B61WJRY5",
			Available:   true,
		},
		{
			Title:       "Escrivaninha",
			Description: "Escrivaninha 5 Gavetas Berlim Mavaular Off White",
			Picture:     "https://m.media-amazon.com/images/I/51yzZ0Ttk6L._AC_.jpg",
			Link:        "https://www.amazon.com.br/dp/B08P63687L",
			Available:   true,
		},
		{
			Title:       "Cadeira",
			Description: "Cadeira, Pétala Decorativa, Jantar, Penteadeira, Recepção e Manicure. (Rose)",
			Picture:     "https://m.media-amazon.com/images/I/51PGiMfpWhL._AC_SL1093_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0FPMN7MR7",
			Available:   true,
		},
		{
			Title:       "Kit Utensílios de Cozinha",
			Description: "Kit Utensílios de Cozinha Branco 19 Peças Colher de Silicone e Madeira Conjunto Jogo | Organizar e Completar a Cozinha",
			Picture:     "https://m.media-amazon.com/images/I/51sn-39DnFL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DL7DH9L5",
			Available:   true,
		},
		{
			Title:       "Guarda Roupa",
			Description: "Guarda-roupa Casal 8 Portas Batentes Branco/branco/rustic Netuno Madesa",
			Picture:     "https://m.media-amazon.com/images/I/51SN0TEEhnL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0FGDJX4M7",
			Available:   true,
		},
		{
			Title:       "Cesto de Roupas",
			Description: "Cesto de Roupas grid com tampa 49 litros - branco",
			Picture:     "https://m.media-amazon.com/images/I/61kfENWSHaL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0BCKWVHD9",
			Available:   true,
		},
		{
			Title:       "Jogo com 6 Tigela",
			Description: "Kit com 6 Bowl Tigela Pote Cumbucas Redonda de Melamina Para Sobremesas Petiscos Molhos Brancas 250ml - Resistentes e Elegantes - Jogo de Sobremesa Melanina Bellagio Gourmet",
			Picture:     "https://m.media-amazon.com/images/I/51c9o8nsflL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0D1RHZL3V",
			Available:   true,
		},
		{
			Title:       "Jogo com 5 Potes de Vidro",
			Description: "Oikos - Kit 5 Potes de Vidro com Tampa de Bambu Herméticos Redondos",
			Picture:     "https://m.media-amazon.com/images/I/5139KTvprnL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09KNY2MVC",
			Available:   true,
		},
		{
			Title:       "Jogo de Jantar",
			Description: "1 Ap. De Jantar/cha 20pcs - M. Order - Jardim Secreto - Aw20-5953",
			Picture:     "https://m.media-amazon.com/images/I/81wLBCuPjLL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0CYJ3FQ9B",
			Available:   false,
		},
		{
			Title:       "Ferro de passar",
			Description: "Electrolux, Ferro de passar seco e vapor, base antiaderente vapor vertical extra 1200W cabo 1,35m azul ESI10 220v",
			Picture:     "https://m.media-amazon.com/images/I/515DMOLkl4L._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09LJYZWX7",
			Available:   true,
		},
		{
			Title:       "Ventilador",
			Description: "Ventilador de Mesa Mondial 220V, 40cm, 6 pás, Super Power - VSP-40-B",
			Picture:     "https://m.media-amazon.com/images/I/81r2JL+Ru5L._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B09B15RWPS",
			Available:   false,
		},
		{
			Title:       "Jogo com 6 Copos",
			Description: "Jogo com 6 Copos Conjunto de Peças de Vidro Transparente, Linha 423ml para Suco e Água",
			Picture:     "https://m.media-amazon.com/images/I/71lFvdnngbL._AC_SL1500_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0DJ6DRH9J",
			Available:   true,
		},
		{
			Title:       "Kit com 6 Tigela",
			Description: "Kit com 6 Bowl Tigela Pote Cumbucas Redonda de Melamina Para Sobremesas Petiscos Molhos Brancas 250ml - Resistentes e Elegantes - Jogo de Sobremesa Melanina Bellagio Gourmet",
			Picture:     "https://m.media-amazon.com/images/I/51c9o8nsflL._AC_SL1000_.jpg",
			Link:        "https://www.amazon.com.br/dp/B0D1RHZL3V",
			Available:   true,
		},
	}

	sort.SliceStable(list, func(i, j int) bool {
		return strings.ToLower(list[i].Title) < strings.ToLower(list[j].Title)
	})

	for i := range list {
		list[i].ID = i + 1
	}

	return &WeddingList{
		WeddingList: list,
	}
}
