package requests

type GetDashboardData struct {
	Level string `json:"level" validate:"required"`
	For   string `json:"for" validate:"required"`
}

type GetTimeIntervalData struct {
	Level string `json:"level" validate:"required"`
	For   string `json:"for" validate:"required"`
	Start string `json:"start" validate:"required"`
	End   string `json:"end" validate:"required"`
}
