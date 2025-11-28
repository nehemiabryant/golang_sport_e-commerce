package models

import "time"

type Cart struct {
	ID        int
	UserID    int
	CreatedAt time.Time
}

type CartItem struct {
	ID        int
	CartID    int
	ProductID string
	Quantity  int
}

type CartItemDetail struct {
	ProductID string  `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
