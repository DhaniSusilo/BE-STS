package responses



type LoginResponse struct {
	AccessToken string `json:"accessToken"`
	Name     	string `json:"name"`
	Level     	string `json:"level"`
	For       	string `json:"for"`
}

type CurrentUserResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
