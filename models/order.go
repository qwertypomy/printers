package models

import "time"

type Order struct {
	ID        uint
	UserID    uint      `db:"user_id"`
	Status    uint      `db:"status"`
	Cost      uint      `db:"cost"`
	CreatedAt time.Time `db:"created_at"`
}

type OrderHasPrinter struct {
	PrinterID uint `db:"printer_id"`
	OrderID   uint `db:"order_id"`
	Amount    uint `db:"amount"`
}
