package lsp

import (
	"encoding/json"
)

type LSPRequest struct {
	Kode               string      `json:"kode" binding:"required"`
	Nama               string      `json:"nama" binding:"required"`
	NamaKetua          string      `json:"namaKetua" binding:"required"`
	NamaDewanPengarah  string      `json:"namaDewanPengarah" binding:"required"`
	NoTelepon          string      `json:"noTelepon" binding:"required"`
	NoWhatsapp         string      `json:"noWhatsapp" binding:"required"`
	Alamat             string      `json:"alamat" binding:"required"`
	Provinsi           string      `json:"provinsi" binding:"required"`
	Kota               string      `json:"kota" binding:"required"`
	Kecamatan          string      `json:"kecamatan" binding:"required"`
	Desa               string      `json:"desa" binding:"required"`
	KodePos            json.Number `json:"kodePos" binding:"required,number"`
	Website            string      `json:"website" binding:"required"`
	NoLisensi          string      `json:"noLisensi" binding:"required"`
	MasaBerlakuLisensi string      `json:"masaBerlakuLisensi" binding:"required"`
	InstitusiInduk     string      `json:"institusiInduk" binding:"required"`
	JenisLSPID         json.Number `json:"JenisLSPID" binding:"required,number"`
}
