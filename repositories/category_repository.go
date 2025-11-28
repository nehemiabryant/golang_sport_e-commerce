package repositories

import (
	"context"
	"miniproject-nehemia/config"
	"miniproject-nehemia/models"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (cr *CategoryRepository) GetAllCategories(ctx context.Context) ([]models.Category, error) {
	rows, err := config.DB.Query(ctx, `SELECT id, name FROM categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category

	for rows.Next() {
		var c models.Category
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	return categories, nil
}

// get sub-category products by category_id
func (cr *CategoryRepository) GetProductsByCategory(ctx context.Context, categoryID string) ([]models.Product, error) {

	rows, err := config.DB.Query(ctx, `
		SELECT p.id, p.name, p.price, p.stock
		FROM product_categories pc
		JOIN products p ON p.id = pc.product_id
		WHERE pc.category_id = $1
	`, categoryID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (cr *CategoryRepository) AddProductToCategory(ctx context.Context, productID, categoryID string) error {
	_, err := config.DB.Exec(ctx,
		`INSERT INTO product_categories (product_id, category_id) VALUES ($1, $2)`,
		productID, categoryID,
	)
	return err
}
