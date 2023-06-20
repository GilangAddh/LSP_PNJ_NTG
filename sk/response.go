package sk

type SKResponse struct {
	ID                   int    `json:"id"`
	JudulStandar         string `json:"judulStandar"`
	NoStandar            string `json:"noStandar"`
	LegalitasPerundangan string `json:"legalitasPerundangan"`
	Sektor               string `json:"sektor"`
	SubSektor            string `json:"subSektor" `
	Penerbit             string `json:"penerbit"`
	JenisSKID            int    `json:"jenisSKID"`
}
