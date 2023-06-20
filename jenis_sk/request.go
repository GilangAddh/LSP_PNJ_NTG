package jenis_sk

type JenisSKRequest struct {
	Nama string `json:"nama" binding:"required"`
}
