package requests

type UpdateUser struct {
	UserName  string `json:"username" validate:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Level     string `json:"level" validate:"required"`
	For       string `json:"for"`
}
