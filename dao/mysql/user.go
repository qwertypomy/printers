package mysql

import (
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

type UserDaoImpl struct {
}

func (UserDaoImpl) CreateUser(user *models.User) (err error) {
	rowPassword := user.Password
	user.Password = utils.Encrypt(user.Password)
	res, err := Db.NamedExec("INSERT INTO user (name, email, password, is_admin) VALUES (:name, :email, :password, :is_admin)", &user)
	user.Password = rowPassword
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			user.ID = string(id)
		}
	}
	return
}

func (UserDaoImpl) DeleteUserByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM user WHERE id=?", id)
	return
}

func (UserDaoImpl) DeleteAllUsers() (err error) {
	_, err = Db.Exec("DELETE FROM user")
	return
}

func (UserDaoImpl) UserList() (userList []models.User, err error) {
	rows, err := Db.Queryx("SELECT * FROM user")
	if err != nil {
		return
	}
	for rows.Next() {
		var user models.User
		err = rows.StructScan(&user)
		if err != nil {
			return
		}
		userList = append(userList, user)
	}
	rows.Close()
	return
}

func (UserDaoImpl) UserListByName(name string) (userList []models.User, err error) {
	rows, err := Db.NamedQuery("SELECT * FROM  user WHERE name=?", name)
	if err != nil {
		return
	}
	for rows.Next() {
		var user models.User
		err = rows.StructScan(&user)
		if err != nil {
			return
		}
		userList = append(userList, user)
	}
	rows.Close()
	return
}

func (UserDaoImpl) UserByID(id string) (user *models.User, err error) {
	err = Db.QueryRowx("SELECT * FROM  user WHERE id=?", id).StructScan(&user)
	return
}
