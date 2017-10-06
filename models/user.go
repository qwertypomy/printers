package models

import "time"

type User struct {
	ID        uint
	Name      string `gorm:"size:16;unique;not null"`
	Uuid      string `gorm:"size:36;not null"`
	Email     string
	Password  string `gorm:"size:40;not null"`
	IsAdmin   bool   `gorm:"not null;default:0"`
	CreatedAt time.Time
}
