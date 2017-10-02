package interfaces

import "github.com/qwertypomy/printers/models"

type UserDao interface {
	Create(user *models.User) error
	Delete(user *models.User) error
	Update(user *models.User) error
	UserDeleteAll() error
	UserList() ([]models.User, error)
	UserByName(name string) (models.User, error)
	UserByUUID(uuid string) (models.User, error)
	UserById(id string) (models.User, error)
}
