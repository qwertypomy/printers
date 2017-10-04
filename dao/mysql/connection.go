package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qwertypomy/printers/config"
	"log"
)

var Db *sql.DB

func init() {
	Db = getDB()
}

func getDB() *sql.DB {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&parseTime=true", cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return db
}
