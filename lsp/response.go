package lsp

type LSPResponse struct {
	ID                 int    `json:"id"`
	Kode               string `json:"kode"`
	Nama               string `json:"nama"`
	NamaKetua          string `json:"namaKetua"`
	NamaDewanPengarah  string `json:"namaDewanPengarah"`
	NoTelepon          string `json:"noTelepon"`
	NoWhatsapp         string `json:"noWhatsapp"`
	Alamat             string `json:"alamat"`
	Provinsi           string `json:"provinsi"`
	Kota               string `json:"kota"`
	Kecamatan          string `json:"kecamatan"`
	Desa               string `json:"desa"`
	KodePos            int    `json:"kodePos"`
	Website            string `json:"website"`
	NoLisensi          string `json:"noLisensi"`
	MasaBerlakuLisensi string `json:"masaBerlakuLisensi"`
	InstitusiInduk     string `json:"institusiInduk"`
	JenisLSPID         int    `json:"jenisLSPID"`
}
