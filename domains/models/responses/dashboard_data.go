package responses

type GetDashboardData struct {
	HariIni int `json:"hari_ini"` // total members registered today
	Total   int `json:"total"`    // total members overall
}

type GetTimeIntervalData struct {
	Total int    `json:"total"`
	Start string `json:"start"`
	End   string `json:"end"`
}
