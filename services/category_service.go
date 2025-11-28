package services

import (
	"context"
	"miniproject-nehemia/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(r *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: r}
}

func (cs *CategoryService) GetCategories(ctx context.Context) ([]string, []string, error) {
	cats, err := cs.repo.GetAllCategories(ctx)
	if err != nil {
		return nil, nil, err
	}

	var ids []string
	var names []string

	for _, c := range cats {
		ids = append(ids, c.ID)
		names = append(names, c.Name)
	}

	return ids, names, nil
}

func (cs *CategoryService) GetProductsByCategory(ctx context.Context, id string) (
	[]string, []string, []float64, []int, error) {

	products, err := cs.repo.GetProductsByCategory(ctx, id)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	var ids []string
	var names []string
	var prices []float64
	var stocks []int

	for _, p := range products {
		ids = append(ids, p.ID)
		names = append(names, p.Name)
		prices = append(prices, p.Price)
		stocks = append(stocks, p.Stock)
	}

	return ids, names, prices, stocks, nil
}

func (cs *CategoryService) AssignProductToCategory(ctx context.Context, productID, categoryID string) error {
	return cs.repo.AddProductToCategory(ctx, productID, categoryID)
}
