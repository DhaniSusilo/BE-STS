package requests

type GetDashboardData struct {
	Level string `json:"level" validate:"required"`
	For   string `json:"for" validate:"required"` // E.g. "Jawa Barat" for provinsi, "Bandung" for kabupaten, etc.
}
