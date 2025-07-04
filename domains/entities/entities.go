package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	Username  string
	FirstName string
	LastName  string
	Password  string
	Level     string
	For       string
	Enabled   int `gorm:"default:1"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Member struct {
	NIK        string `gorm:"type:varchar(16);uniqueIndex;not null" json:"nik" validate:"required,len=16,numeric"`
	Nama       string `gorm:"type:varchar(100);not null" json:"nama" validate:"required"`                         
	NoHp       string `gorm:"type:varchar(20);not null" json:"no_hp" validate:"required"`     
	Provinsi   string `gorm:"type:varchar(100);not null" json:"provinsi" validate:"required"`                     
	Kabupaten  string `gorm:"type:varchar(100);not null" json:"kabupaten" validate:"required"`                    
	Kecamatan  string `gorm:"type:varchar(100);not null" json:"kecamatan" validate:"required"`                    
	Kelurahan  string `gorm:"type:varchar(100);not null" json:"kelurahan" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

