package mysql

import (
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

type UserDaoImpl struct {
}

func (UserDaoImpl) CreateUser(user *models.User) {
	user.Uuid = utils.CreateUUID()
	user.Password = utils.Encrypt(user.Password)
	Db.Create(&user)
}

func (UserDaoImpl) DeleteUser(user *models.User) {
	Db.Delete(&user)
}

func (UserDaoImpl) UpdateUser(user *models.User) {
	Db.Model(&user).Updates(map[string]interface{}{"name": user.Name, "email": user.Email, "password": utils.Encrypt(user.Password)})
}

func (UserDaoImpl) DeleteAllUsers() {
	Db.Exec("truncate table users")
}

func (UserDaoImpl) UserList() (users []models.User) {
	Db.Find(&users)
	return
}

func (UserDaoImpl) UserByName(name string) (user models.User) {
	Db.Where("name = ?", name).First(&user)
	return
}

func (UserDaoImpl) UserByUUID(uuid string) (user models.User) {
	Db.Where("uuid = ?", uuid).First(&user)
	return
}

func (UserDaoImpl) UserById(id uint) (user models.User) {
	Db.Where("id = ?", id).First(&user)
	return
}
