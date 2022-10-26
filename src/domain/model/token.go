package model

type Token struct {
	Id           uint   `json:"token_id"`
	UserId       uint   `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Exp          int64  `json:"exp"`
}
