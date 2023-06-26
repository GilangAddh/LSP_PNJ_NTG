package skema_sertifikasi

type SkemaSertifikasiResponse struct {
	ID              int    `json:"id"`
	Kode            string `json:"kode"`
	Judul           string `json:"judulStandar"`
	JudulInggris    string `json:"judulInggris"`
	KeteranganBukti string `json:"keteranganBukti"`
	KedalamanBukti  string `json:"kedalamanBukti"`
	SKID            int    `json:"SKID" binding"`
}
