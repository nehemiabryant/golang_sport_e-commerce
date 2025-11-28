package repositories

import (
	"context"
	"miniproject-nehemia/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(ctx, `
		SELECT id, name, price, stock
		FROM products
		WHERE id = $1
	`, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) DecreaseStock(ctx context.Context, productID string, qty int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE products
		SET stock = stock - $1
		WHERE id = $2 AND stock >= $1
	`, qty, productID)
	return err
}

func (r *ProductRepository) IncreaseStock(ctx context.Context, productID string, qty int) error {
	_, err := r.db.Exec(ctx, `
		UPDATE products
		SET stock = stock + $1
		WHERE id = $2
	`, qty, productID)
	return err
}

func (r *ProductRepository) Create(ctx context.Context, p models.Product) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO products (id, name, price, stock) VALUES ($1, $2, $3, $4)`,
		p.ID, p.Name, p.Price, p.Stock,
	)
	return err
}

func (r *ProductRepository) UpdateStock(ctx context.Context, id string, stock int) error {
	_, err := r.db.Exec(ctx,
		`UPDATE products SET stock = $1 WHERE id = $2`,
		stock, id,
	)
	return err
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]models.Product, error) {
	rows, err := r.db.Query(ctx, `SELECT id, name, price, stock FROM products ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

// Assign product ke category
func (r *ProductRepository) AssignProductToCategory(ctx context.Context, productID, categoryID string) error {
	_, err := r.db.Exec(ctx,
		`INSERT INTO product_categories (product_id, category_id) VALUES ($1, $2)`,
		productID, categoryID,
	)
	return err
}

func (r *ProductRepository) GetProductsGroupedByCategory(ctx context.Context) (map[string][]models.Product, error) {
	rows, err := r.db.Query(ctx, `
		SELECT c.name, p.id, p.name, p.price, p.stock
		FROM product_categories pc
		JOIN products p ON p.id = pc.product_id
		JOIN categories c ON c.id = pc.category_id
		ORDER BY c.name, p.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grouped := make(map[string][]models.Product)

	for rows.Next() {
		var categoryName string
		var p models.Product

		err := rows.Scan(&categoryName, &p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}

		grouped[categoryName] = append(grouped[categoryName], p)
	}

	return grouped, nil
}
