package factory

import (
	"github.com/qwertypomy/printers/dao/interfaces"
	"github.com/qwertypomy/printers/dao/mysql"
	"log"
)

type FactoryDao struct {
	Engine string
}

func (factory *FactoryDao) GetUserDaoInterface() interfaces.UserDao {
	switch factory.Engine {
	case "mysql":
		return mysql.UserDaoImpl{}
	default:
		log.Fatalf("No UserDao implementation for %s", factory.Engine)
		return nil
	}
}

func (factory *FactoryDao) GetPrinterDaoInterface() interfaces.PrinterDao {
	switch factory.Engine {
	case "mysql":
		return mysql.PrinterDaoImpl{}
	default:
		log.Fatalf("No PrinterDao implementation for %s", factory.Engine)
		return nil
	}
}
