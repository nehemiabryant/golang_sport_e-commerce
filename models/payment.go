package models

import "time"

type Payment struct {
	ID          int       `db:"id"`
	UserID      int       `db:"user_id"`
	TotalAmount float64   `db:"total_amount"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

type PaymentDetail struct {
	ID        int     `db:"id"`
	PaymentID int     `db:"payment_id"`
	ProductID string  `db:"product_id"`
	Quantity  int     `db:"quantity"`
	Price     float64 `db:"price"`
	Subtotal  float64 `db:"subtotal"`
}

type PaymentDetailFull struct {
	PaymentID   int
	UserID      int
	Status      string
	UserEmail   string
	ProductID   string
	ProductName string
	Quantity    int
	Price       float64
	Subtotal    float64
	TotalAmount float64
	CreatedAt   time.Time
}
