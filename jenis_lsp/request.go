package jenis_lsp

type JenisLSPRequest struct {
	Nama string `json:"nama" binding:"required"`
}
