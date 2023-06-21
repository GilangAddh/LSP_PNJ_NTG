package entity

import "gorm.io/gorm"

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

type JenisSK struct {
	ID   int
	Nama string
}

type SK struct {
	gorm.Model
	ID                   int
	JudulStandar         string
	NoStandar            string
	LegalitasPerundangan string
	Sektor               string
	SubSektor            string
	Penerbit             string
	JenisSKID            int
	JenisSK              JenisSK
}

type LSP_SK struct {
	gorm.Model
	ID    int
	LSPID int
	SKID  int
	LSP   LSP
	SK    SK
}

type Accounts struct {
	gorm.Model
	ID       int
	Nama     string
	NIK      int
	Email    string
	Password string
	Role     string
}

type Skema_Sertifikasi struct {
	gorm.Model
	ID              int
	Kode            string
	Judul           string
	JudulInggris    string
	KeteranganBukti string
	KedalamanBukti  string
	SKID            int
	SK              SK
}

type Persyaratan_Spesifikasi struct {
	gorm.Model
	ID                  int
	Jenis               string
	Nama                string
	Skema_SertifikasiID int
	Skema_Sertifikasi   Skema_Sertifikasi
}
