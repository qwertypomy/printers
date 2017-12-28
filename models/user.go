package models

import "time"

type User struct {
	ID        uint
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	IsAdmin   bool      `db:"is_admin"`
	CreatedAt time.Time `db:"created_at"`
}
