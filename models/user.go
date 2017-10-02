package models

import "time"

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	IsAdmin   bool
	CreatedAt time.Time
}
