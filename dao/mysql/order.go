package mysql

import (
	"github.com/qwertypomy/printers/models"
)

type OrderDaoImpl struct {
}

// Order functions
func (OrderDaoImpl) CreateOrder(order *models.Order) (err error) {
	res, err := Db.NamedExec("INSERT INTO `order` (user_id, status) VALUES (:user_id, :status)", &order)
	if err != nil {
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			order.ID = string(id)
		}
	}
	return
}

func (OrderDaoImpl) UpdateOrder(order *models.Order) (err error) {
	_, err = Db.NamedExec("update `order` set status=:status, where id =:id", &order)
	return
}

func (OrderDaoImpl) DeleteOrderByID(id string) (err error) {
	_, err = Db.Exec("DELETE FROM `order` WHERE id=?", id)
	return
}

func (OrderDaoImpl) GetOrderByID(id string) (order *models.Order, err error) {
	err = Db.QueryRowx("SELECT * FROM `order` WHERE id=?", id).StructScan(&order)
	return
}

func (OrderDaoImpl) OrderList() (orderList []models.Order, err error) {
	rows, err := Db.Queryx("SELECT * FROM `order`")
	if err != nil {
		return
	}
	for rows.Next() {
		var order models.Order
		err = rows.StructScan(&order)
		if err != nil {
			return
		}
		orderList = append(orderList, order)
	}
	rows.Close()
	return
}

func (OrderDaoImpl) DeleteAllOrders() (err error) {
	_, err = Db.Exec("DELETE FROM `order`")
	return
}

// OrderHasPrinter functions
func (OrderDaoImpl) CreateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	_, err = Db.NamedExec("INSERT INTO order_has_printer (printer_id, order_id, amount) VALUES (:printer_id, :order_id, :amount)", &orderHasPrinter)
	return
}

func (OrderDaoImpl) UpdateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	_, err = Db.NamedExec("update order_has_printer set amount=:amount,  where id =:id", &orderHasPrinter)
	return
}

func (OrderDaoImpl) DeleteOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	_, err = Db.Exec("DELETE FROM order_has_printer WHERE order_id=? and printer_id=?", orderHasPrinter.OrderID, orderHasPrinter.PrinterID)
	return
}

func (OrderDaoImpl) GetOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error) {
	err = Db.QueryRowx("SELECT * FROM order_has_printer WHERE order_id=? and printer_id=?", orderHasPrinter.OrderID, orderHasPrinter.PrinterID).StructScan(&orderHasPrinter)
	return
}

func (OrderDaoImpl) OrderHasPrinterList() (orderHasPrinterList []models.OrderHasPrinter, err error) {
	rows, err := Db.Queryx("SELECT * FROM order_has_printer")
	if err != nil {
		return
	}
	for rows.Next() {
		var orderHasPrinter models.OrderHasPrinter
		err = rows.StructScan(&orderHasPrinter)
		if err != nil {
			return
		}
		orderHasPrinterList = append(orderHasPrinterList, orderHasPrinter)
	}
	rows.Close()
	return
}

func (OrderDaoImpl) OrderHasPrinterListByOrderId(orderID string) (orderHasPrinterList []models.OrderHasPrinter, err error) {
	rows, err := Db.Queryx("SELECT * FROM order_has_printer WHERE order_id=?", orderID)
	if err != nil {
		return
	}
	for rows.Next() {
		var orderHasPrinter models.OrderHasPrinter
		err = rows.StructScan(&orderHasPrinter)
		if err != nil {
			return
		}
		orderHasPrinterList = append(orderHasPrinterList, orderHasPrinter)
	}
	rows.Close()
	return
}

func (OrderDaoImpl) DeleteAllOrderHasPrinters() (err error) {
	_, err = Db.Exec("DELETE FROM order_has_printer")
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
