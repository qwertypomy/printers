package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/utils"
)

var Db *sqlx.DB

func init() {
	Db = getDB()
}

func getDB() *sqlx.DB {
	cfg, err := config.GetConfig()
	utils.FatalError(err)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
	db, err := sqlx.Open("mysql", dsn)
	utils.FatalError(err)
	return db
}
