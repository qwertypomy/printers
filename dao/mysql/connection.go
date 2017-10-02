package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qwertypomy/printers/utils"
	"log"
)

func getDB() *sql.DB {
	config, err := utils.GetConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&parseTime=true", config.User, config.Password, config.Server, config.Port, config.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}

var Db *sql.DB

func init() {
	Db = getDB()
}
