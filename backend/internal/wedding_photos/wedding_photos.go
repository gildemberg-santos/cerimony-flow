package wedding_photos

type WeddingPhotos struct {
	Pictures []string `json:"pictures"`
}

func NewWeddingPhotos() *WeddingPhotos {
	return &WeddingPhotos{
		Pictures: []string{
			"https://m.media-amazon.com/images/I/61qQCckVsqL._AC_SL1500_.jpg",
		},
	}
}
