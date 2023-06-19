package account

import "encoding/json"

type AccountRequest struct {
	Nama     string      `json:"nama" binding:"required"`
	NIK      json.Number `json:"nik" binding:"required,number"`
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Role     string      `json:"role" binding:"required"`
}

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
