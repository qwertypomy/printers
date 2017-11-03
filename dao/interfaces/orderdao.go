package interfaces

import (
	"github.com/qwertypomy/printers/models"
)

type OrderDao interface {
	// Order functions
	CreateOrder(order *models.Order) (err error)
	UpdateOrder(order *models.Order) (err error)
	DeleteOrderByID(id uint) (err error)
	GetOrderByID(id uint) (order *models.Order, err error)
	OrderList() (orderList []models.Order, err error)
	DeleteAllOrders() (err error)
	// OrderHasPrinter functions
	CreateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error)
	UpdateOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error)
	DeleteOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error)
	GetOrderHasPrinter(orderHasPrinter *models.OrderHasPrinter) (err error)
	OrderHasPrinterList() (orderHasPrinterList []models.OrderHasPrinter, err error)
	OrderHasPrinterListByOrderId(orderID uint) (orderHasPrinterList []models.OrderHasPrinter, err error)
	DeleteAllOrderHasPrinters() (err error)

	DeleteAllData() (err error)
}
