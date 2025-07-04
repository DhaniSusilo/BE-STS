package requests

type SignUpRequest struct {
	UserName        string `json:"username" validate:"required"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Level           string `json:"level" validate:"required"`
	For 			string `json:"for"`
}
