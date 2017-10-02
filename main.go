package main

import (
	"fmt"
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	config, err := utils.GetConfig()
	HandleErr(err)

	// factoryDao usage example

	factoryDao := factory.FactoryDao{config.Engine}
	userDao := factoryDao.GetUserDaoInterface()
	printerDao := factoryDao.GetPrinterDaoInterface()

	user := models.User{
		Name:     "admin",
		Email:    "admin@admin.admin",
		Password: "12345",
		IsAdmin:  true,
	}

	err = userDao.Create(&user)
	HandleErr(err)

	fmt.Println(user)

	err = userDao.Delete(&user)
	HandleErr(err)

	printerResolution, err := printerDao.CreatePrinterPrintResolution(800, 600)
	HandleErr(err)

	fmt.Println(printerResolution)
}
