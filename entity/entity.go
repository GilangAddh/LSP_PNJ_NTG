package entity

import "gorm.io/gorm"

type AdminAccount struct {
	ID       int
	Nama     string
	Username string
	Password string
}

type LSP struct {
	gorm.Model
	ID                 int
	Kode               string
	Nama               string
	NamaKetua          string
	NamaDewanPengarah  string
	NoTelepon          string
	NoWhatsapp         string
	Alamat             string
	Provinsi           string
	Kota               string
	Kecamatan          string
	Desa               string
	KodePos            int
	Website            string
	NoLisensi          string
	MasaBerlakuLisensi string
	InstitusiInduk     string
	JenisLSPID         int
	JenisLSP           JenisLSP
}

type JenisLSP struct {
	ID   int
	Nama string
}
