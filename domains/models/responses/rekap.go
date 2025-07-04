package responses

import "time"

type RekapitulasiResult struct {
	No           int
	Wilayah      string
	TotalAnggota int
}

type MemberDetail struct {
	Nama          string    `json:"nama"`
	Nik           string    `json:"nik"`
	NoHp          string    `json:"noHp"`
	Provinsi      string    `json:"provinsi"`
	Kabupaten     string    `json:"kabupaten"`
	Kecamatan     string    `json:"kecamatan"`
	Kelurahan     string    `json:"kelurahan"`
	TanggalDaftar time.Time `json:"tanggalDaftar"`
}


// output can be either aggregated or detailed data based on level
type RekapitulasiResponse struct {
	Aggregated []RekapitulasiResult
	Members    []MemberDetail
	TotalCount int
}

