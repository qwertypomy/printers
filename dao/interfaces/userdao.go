package interfaces

import "github.com/qwertypomy/printers/models"

type UserDao interface {
	CreateUser(user *models.User)
	DeleteUser(user *models.User)
	UpdateUser(user *models.User)
	DeleteAllUsers()
	UserList() []models.User
	UserByName(name string) models.User
	UserByUUID(uuid string) models.User
	UserById(id uint) models.User
}
