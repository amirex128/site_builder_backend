package user_dto

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshDto struct {
	RefreshToken string `json:"refresh_token"`
}
