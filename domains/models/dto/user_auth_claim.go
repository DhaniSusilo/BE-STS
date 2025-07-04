package dto

type AuthUserDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type AuthAccessToken struct {
	Id string `json:"token_id"`
}
