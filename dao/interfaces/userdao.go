package interfaces

import "github.com/qwertypomy/printers/models"

type UserDao interface {
	CreateUser(user *models.User) (err error)
	DeleteUserByID(id uint) (err error)
	UpdateUserEmail(user *models.User) (err error)
	UpdateUserPassword(user *models.User) (err error)
	DeleteAllUsers() (err error)
	UserList() (userList []models.User, err error)
	UserByName(name string) (user *models.User, err error)
	UserByID(id uint) (user *models.User, err error)
}
