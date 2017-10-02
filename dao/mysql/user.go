package mysql

import (
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

type UserDaoImpl struct {
}

func (UserDaoImpl) Create(user *models.User) (err error) {
	_, err = Db.Exec("insert into user (uuid, name, email, is_admin, password) values (?, ?, ?, ?, ?)",
		utils.CreateUUID(), user.Name, user.Email, user.IsAdmin, utils.Encrypt(user.Password))
	if err != nil {
		return
	}
	err = Db.QueryRow("SELECT id, uuid, created_at FROM user where name = ?", user.Name).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}

func (UserDaoImpl) Delete(user *models.User) (err error) {
	_, err = Db.Exec("delete from user where id = ?", user.Id)
	return
}

func (UserDaoImpl) Update(user *models.User) (err error) {
	_, err = Db.Exec("update users set email = ? where id = ?", user.Email)
	return
}

func (UserDaoImpl) UserDeleteAll() (err error) {
	_, err = Db.Exec("delete from user")
	return
}

func (UserDaoImpl) UserList() (users []models.User, err error) {
	rows, err := Db.Query("SELECT id, uuid, name, email, is_admin, password, created_at FROM user")
	if err != nil {
		return
	}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.IsAdmin, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

func (UserDaoImpl) UserByName(name string) (user models.User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, is_admin, password, created_at FROM user WHERE name = ?", name).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.IsAdmin, &user.Password, &user.CreatedAt)
	return
}

func (UserDaoImpl) UserByUUID(uuid string) (user models.User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, is_admin, password, created_at FROM user WHERE uuid = ?", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.IsAdmin, &user.Password, &user.CreatedAt)
	return
}

func (UserDaoImpl) UserById(id string) (user models.User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, is_admin, password, created_at FROM user WHERE id = ?", id).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.IsAdmin, &user.Password, &user.CreatedAt)
	return
}
