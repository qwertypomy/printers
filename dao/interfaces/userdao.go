package interfaces

import "github.com/qwertypomy/printers/models"

type UserDao interface {
	CreateUser(user *models.User) (err error)
	DeleteUserByID(id string) (err error)
	DeleteAllUsers() (err error)
	UserList() (userList []models.User, err error)
	UserListByName(name string) (userList []models.User, err error)
	UserByID(id string) (user *models.User, err error)
}
