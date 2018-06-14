package models

import "time"

type Order struct {
	ID        string    `bson:"_id"`
	UserID    string    `db:"user_id"`
	Status    uint      `db:"status"`
	Cost      uint      `db:"cost"`
	CreatedAt time.Time `db:"created_at"`
}

type OrderHasPrinter struct {
	PrinterID string `db:"printer_id"`
	OrderID   string `db:"order_id"`
	Amount    uint   `db:"amount"`
}
