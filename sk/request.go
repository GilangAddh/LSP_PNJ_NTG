package sk

import (
	"encoding/json"
)

type SKRequest struct {
	JudulStandar         string      `json:"judulStandar" binding:"required"`
	NoStandar            string      `json:"noStandar" binding:"required"`
	LegalitasPerundangan string      `json:"legalitasPerundangan" binding:"required"`
	Sektor               string      `json:"sektor" binding:"required"`
	SubSektor            string      `json:"subSektor" binding:"required"`
	Penerbit             string      `json:"penerbit" binding:"required"`
	JenisSKID            json.Number `json:"jenisSKID" binding"required number"`
}
