package repositories

import (
	"context"
	"errors"
	"miniproject-nehemia/config"
	"miniproject-nehemia/models"
)

type CartRepository struct{}

func NewCartRepository() *CartRepository {
	return &CartRepository{}
}

// GetOrCreateCart returns cart id for a user (creates cart row if not exists)
func (r *CartRepository) GetOrCreateCart(ctx context.Context, userID int) (int, error) {
	var cartID int

	// Try get
	row := config.DB.QueryRow(ctx, `SELECT id FROM cart WHERE user_id = $1`, userID)
	err := row.Scan(&cartID)
	if err == nil {
		return cartID, nil
	}

	// If not exists, create and return id
	err = config.DB.QueryRow(ctx, `
		INSERT INTO cart (user_id) VALUES ($1) RETURNING id
	`, userID).Scan(&cartID)
	if err != nil {
		return 0, err
	}

	return cartID, nil
}

// AddOrIncrementItem: add new item or increment existing quantity
func (r *CartRepository) AddOrIncrementItem(ctx context.Context, cartID int, productID string, qty int) error {
	// Insert or update quantity (ON CONFLICT requires a unique constraint on cart_id+product_id)
	_, err := config.DB.Exec(ctx, `
		INSERT INTO cart_item (cart_id, product_id, quantity)
		VALUES ($1, $2, $3)
		ON CONFLICT (cart_id, product_id)
		DO UPDATE SET quantity = cart_item.quantity + EXCLUDED.quantity
	`, cartID, productID, qty)
	return err
}

func (r *CartRepository) GetCartItemsDetail(ctx context.Context, userID int) ([]models.CartItemDetail, error) {
	rows, err := config.DB.Query(ctx, `
		SELECT ci.product_id, p.name, p.price, ci.quantity
		FROM cart_item ci
		JOIN cart c ON c.id = ci.cart_id
		JOIN products p ON p.id = ci.product_id
		WHERE c.user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.CartItemDetail
	for rows.Next() {
		var it models.CartItemDetail
		err := rows.Scan(&it.ProductID, &it.Name, &it.Price, &it.Quantity)
		if err != nil {
			return nil, err
		}
		it.Subtotal = it.Price * float64(it.Quantity)
		items = append(items, it)
	}
	return items, nil
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, cartID int, productID string, qty int) error {
	if qty <= 0 {
		return errors.New("quantity must be > 0")
	}
	_, err := config.DB.Exec(ctx, `
		UPDATE cart_item SET quantity = $1
		WHERE cart_id = $2 AND product_id = $3
	`, qty, cartID, productID)
	return err
}

func (r *CartRepository) RemoveItem(ctx context.Context, cartID int, productID string) error {
	_, err := config.DB.Exec(ctx, `
		DELETE FROM cart_item
		WHERE cart_id = $1 AND product_id = $2
	`, cartID, productID)
	return err
}

func (r *CartRepository) ClearCart(ctx context.Context, cartID int) error {
	_, err := config.DB.Exec(ctx, `
		DELETE FROM cart_item
		WHERE cart_id = $1`, cartID)
	return err
}

// GetItemQuantity returns current quantity of a product in a cart
func (r *CartRepository) GetItemQuantity(ctx context.Context, cartID int, productID string) (int, error) {
	var qty int
	row := config.DB.QueryRow(ctx, `
		SELECT quantity FROM cart_item WHERE cart_id = $1 AND product_id = $2
	`, cartID, productID)
	if err := row.Scan(&qty); err != nil {
		// if no row found, return 0 and the error
		return 0, err
	}
	return qty, nil
}

func (r *CartRepository) AddQty(ctx context.Context, cartID int, productID string) error {
	_, err := config.DB.Exec(ctx, `
		UPDATE cart_item
		SET quantity = quantity + 1
		WHERE cart_id = $1 AND product_id = $2
	`, cartID, productID)
	return err
}

func (r *CartRepository) ReduceQty(ctx context.Context, cartID int, productID string) error {
	tx, err := config.DB.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			tx.Commit(ctx)
		}
	}()

	// 1. Kurangi quantity dalam cart
	_, err = tx.Exec(ctx, `
        UPDATE cart_item
        SET quantity = quantity - 1
        WHERE cart_id = $1 AND product_id = $2 AND quantity > 1
    `, cartID, productID)
	if err != nil {
		return err
	}

	// 2. Tambah stok produk kembali
	_, err = tx.Exec(ctx, `
        UPDATE product
        SET stock = stock + 1
        WHERE id = $1
    `, productID)
	if err != nil {
		return err
	}

	return nil
}
