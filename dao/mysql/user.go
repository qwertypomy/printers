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
			user.ID = uint(id)
		}
	}
	return
}

func (UserDaoImpl) DeleteUserByID(id uint) (err error) {
	_, err = Db.Exec("DELETE FROM user WHERE id=?", id)
	return
}

func (UserDaoImpl) UpdateUserEmail(user *models.User) (err error) {
	_, err = Db.Exec("UPDATE users SET email = ? WHERE id =?", user.Email, user.ID)
	return
}

func (UserDaoImpl) UpdateUserPassword(user *models.User) (err error) {
	_, err = Db.Exec("UPDATE users SET password = ? WHERE id =?", utils.Encrypt(user.Password), user.ID)
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

func (UserDaoImpl) UserByName(name string) (user *models.User, err error) {
	err = Db.QueryRowx("SELECT * FROM  user WHERE name=?", name).StructScan(&user)
	return
}

func (UserDaoImpl) UserByID(id uint) (user *models.User, err error) {
	err = Db.QueryRowx("SELECT * FROM  user WHERE id=?", id).StructScan(&user)
	return
}
