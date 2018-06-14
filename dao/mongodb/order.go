package mongodb

import (
	"github.com/globalsign/mgo/bson"
	"github.com/qwertypomy/printers/models"
)

type OrderDaoImpl struct {
}

// Order functions
func (OrderDaoImpl) CreateOrder(order *models.Order) (err error) {
	order.ID = bson.NewObjectId().Hex()
	err = Db.C("order").Insert(&order)
	return
}

func (OrderDaoImpl) UpdateOrder(order *models.Order) (err error) {
	err = Db.C("order").UpdateId(order.ID, order)
	return
}

func (OrderDaoImpl) DeleteOrderByID(id string) (err error) {
	err = Db.C("order").RemoveId(id)
	return
}

func (OrderDaoImpl) GetOrderByID(id string) (order *models.Order, err error) {
	err = Db.C("order").FindId(bson.ObjectIdHex(id)).One(&order)
	return
}

func (OrderDaoImpl) OrderList() (orderList []models.Order, err error) {
	err = Db.C("order").Find(bson.M{}).All(&orderList)
	return
}

func (OrderDaoImpl) DeleteAllOrders() (err error) {
	_, err = Db.C("order").RemoveAll(bson.M{})
	return
}

// OrderHasPrinter functions
func (OrderDaoImpl) CreateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	err = Db.C("orderHasPrinter").Insert(&orderHasPrinter)
	return
}

func (OrderDaoImpl) UpdateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	err = Db.C("orderHasPrinter").UpdateId(bson.M{"PrinterID": orderHasPrinter.PrinterID, "OrderID": orderHasPrinter.OrderID}, orderHasPrinter)
	return
}

func (OrderDaoImpl) DeleteOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	err = Db.C("orderHasPrinter").Remove(bson.M{"PrinterID": orderHasPrinter.PrinterID, "OrderID": orderHasPrinter.OrderID})
	return
}

func (OrderDaoImpl) GetOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	err = Db.C("orderHasPrinter").Find(bson.M{"PrinterID": orderHasPrinter.PrinterID, "OrderID": orderHasPrinter.OrderID}).One(&orderHasPrinter)
	return
}

func (OrderDaoImpl) OrderHasPrinterList() (orderHasPrinterList []models.OrderHasPrinter, err error) {
	err = Db.C("orderHasPrinter").Find(bson.M{}).All(&orderHasPrinterList)
	return
}

func (OrderDaoImpl) OrderHasPrinterListByOrderId(orderID string) (orderHasPrinterList []models.OrderHasPrinter, err error) {
	err = Db.C("orderHasPrinter").Find(bson.M{"OrderID": orderID}).All(&orderHasPrinterList)
	return
}

func (OrderDaoImpl) DeleteAllOrderHasPrinters() (err error) {
	_, err = Db.C("orderHasPrinter").RemoveAll(bson.M{})
	return
}

// All Tables functions
func (odi OrderDaoImpl) DeleteAllData() (err error) {
	err = odi.DeleteAllOrderHasPrinters()
	if err != nil {
		return
	}
	err = odi.DeleteAllOrders()
	return
}
