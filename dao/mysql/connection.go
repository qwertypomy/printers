package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/models"
	"log"
)

var Db *gorm.DB

func init() {
	Db = getDB()
}

func getDB() *gorm.DB {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Server, cfg.Port, cfg.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	db.AutoMigrate(
		&models.User{},
		&models.Printer{},
		&models.PrintSize{},
		&models.Brand{},
		&models.ConnectivityType{},
		&models.FunctionType{},
		&models.PrintingTechnology{},
		&models.PrintResolution{},
	)

	return db
}
