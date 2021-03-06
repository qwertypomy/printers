package factory

import (
	"log"

	"github.com/qwertypomy/printers/dao/interfaces"
	"github.com/qwertypomy/printers/dao/mongodb"
	"github.com/qwertypomy/printers/dao/mysql"
)

type FactoryDao struct {
	Engine string
}

func (factory *FactoryDao) GetUserDaoInterface() interfaces.UserDao {
	switch factory.Engine {
	case "mysql":
		return mysql.UserDaoImpl{}
	case "mongodb":
		return mongodb.UserDaoImpl{}
	default:
		log.Fatalf("No UserDao implementation for %s", factory.Engine)
		return nil
	}
}

func (factory *FactoryDao) GetPrinterDaoInterface() interfaces.PrinterDao {
	switch factory.Engine {
	case "mysql":
		return mysql.PrinterDaoImpl{}
	case "mongodb":
		return mongodb.PrinterDaoImpl{}
	default:
		log.Fatalf("No PrinterDao implementation for %s", factory.Engine)
		return nil
	}
}

func (factory *FactoryDao) GetOrderDaoInterface() interfaces.OrderDao {
	switch factory.Engine {
	case "mysql":
		return mysql.OrderDaoImpl{}
	case "mongodb":
		return mongodb.OrderDaoImpl{}
	default:
		log.Fatalf("No OrderDao implementation for %s", factory.Engine)
		return nil
	}
}
