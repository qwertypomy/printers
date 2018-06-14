package mongodb

import (
	"github.com/globalsign/mgo/bson"

	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

type UserDaoImpl struct {
}

func (UserDaoImpl) CreateUser(user *models.User) (err error) {
	rowPassword := user.Password
	user.Password = utils.Encrypt(user.Password)
	user.ID = bson.NewObjectId().Hex()
	err = Db.C("user").Insert(&user)
	user.Password = rowPassword
	return
}

func (UserDaoImpl) DeleteUserByID(id string) (err error) {
	err = Db.C("user").RemoveId(id)
	return
}

func (UserDaoImpl) DeleteAllUsers() (err error) {
	Db.C("user").RemoveAll(bson.M{})
	return
}

func (UserDaoImpl) UserList() (userList []models.User, err error) {
	err = Db.C("user").Find(bson.M{}).All(&userList)
	return
}

func (UserDaoImpl) UserListByName(name string) (userList []models.User, err error) {
	err = Db.C("user").Find(bson.M{name: name}).All(&userList)
	return
}

func (UserDaoImpl) UserByID(id string) (user *models.User, err error) {
	err = Db.C("user").FindId(bson.ObjectIdHex(id)).One(&user)
	return
}
