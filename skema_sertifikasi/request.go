package skema_sertifikasi

import (
	"encoding/json"
)

type SkemaSertifikasiRequest struct {
	Kode            string      `json:"kode" binding:"required"`
	Judul           string      `json:"judulStandar" binding:"required"`
	JudulInggris    string      `json:"judulInggris" binding:"required"`
	KeteranganBukti string      `json:"keteranganBukti" binding:"required"`
	KedalamanBukti  string      `json:"kedalamanBukti" binding:"required"`
	SKID            json.Number `json:"SKID" binding"required number"`
}
