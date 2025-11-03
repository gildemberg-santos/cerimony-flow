package model

type WeddingPhotos struct {
	Pictures []string `json:"pictures"`
}

func NewWeddingPhotos() *WeddingPhotos {
	return &WeddingPhotos{
		Pictures: []string{
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7245.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7351.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7234.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7227.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7236A.jpg",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7749.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7800.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7953.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7179.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7181.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7210.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7291.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7254.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7299.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7447.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7596.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7705.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7715.JPG",
			"https://pub-d4dd48e211d74a8eb80ddc1aa60888ca.r2.dev/IMG_7802.JPG",
		},
	}
}
