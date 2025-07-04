package requests

type MemberRegistration struct {
	NIK       string `json:"nik" validate:"required"`       // No. KTP
	Nama      string `json:"nama" validate:"required"`      // Nama
	NoHp      string `json:"no_hp" validate:"required"`     // No. Handphone
	Provinsi  string `json:"provinsi" validate:"required"`  // Provinsi
	Kabupaten string `json:"kabupaten" validate:"required"` // Kabupaten
	Kecamatan string `json:"kecamatan" validate:"required"` // Kecamatan
	Kelurahan string `json:"kelurahan" validate:"required"` // Kelurahan
}