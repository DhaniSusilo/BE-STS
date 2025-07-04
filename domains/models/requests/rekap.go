package requests

type RekapitulasiRequest struct {
	Level       string `json:"level"`
	Wilayah     string `json:"wilayah"`
	Page        int    `json:"page"`
	RowsPerPage int    `json:"rowsPerPage"`
}
