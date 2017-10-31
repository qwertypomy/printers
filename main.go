package main

import (
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/dao/factory"
	//"github.com/qwertypomy/printers/db"
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/utils"
)

var Config models.Config

func main() {
	Config, err := config.GetConfig()
	utils.FatalError(err)

	//db.Populate(Config)

	factoryDao := factory.FactoryDao{Engine: Config.Engine}
	userDao := factoryDao.GetUserDaoInterface()
	//printerDao := factoryDao.GetPrinterDaoInterface()

	err = userDao.DeleteAllUsers()
	utils.FatalError(err)
	//err = printerDao.DeleteAllData()
	//utils.FatalError(err)
}
