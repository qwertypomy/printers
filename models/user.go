package models

import "time"

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `db:"name" bson:"name"`
	Email     string    `db:"email" bson:"email"`
	Password  string    `db:"password" bson:"password"`
	IsAdmin   bool      `db:"is_admin" bson:"isAdmin"`
	CreatedAt time.Time `db:"created_at" bson:"createdAt"`
}
