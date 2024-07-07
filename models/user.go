package models

import "gorm.io/gorm"

// Struktur User yang menggunakan GORM
type User struct {
	gorm.Model
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	EmailTemp string `json:"email_temp"`
}
