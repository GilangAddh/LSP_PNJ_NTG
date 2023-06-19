package account

type AccountResponse struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	NIK      int    `json:"nik"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
