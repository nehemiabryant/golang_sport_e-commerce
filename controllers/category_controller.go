package controllers

import (
	"context"
	"miniproject-nehemia/models"
	"miniproject-nehemia/services"
)

type CategoryController struct {
	service *services.CategoryService
}

func NewCategoryController(service *services.CategoryService) *CategoryController {
	return &CategoryController{
		service: service,
	}
}

// Wrapper: Ambil semua kategori
func (cc *CategoryController) GetAllCategories() ([]models.Category, error) {
	ctx := context.Background()

	ids, names, err := cc.service.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	var result []models.Category
	for i := range ids {
		result = append(result, models.Category{
			ID:   ids[i],
			Name: names[i],
		})
	}

	return result, nil
}

// Wrapper: Ambil produk berdasarkan categoryID
func (cc *CategoryController) GetProductsByCategory(categoryID string) ([]models.Product, error) {
	ctx := context.Background()

	ids, names, prices, stocks, err := cc.service.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		return nil, err
	}

	var result []models.Product
	for i := range ids {
		result = append(result, models.Product{
			ID:    ids[i],
			Name:  names[i],
			Price: prices[i],
			Stock: stocks[i],
		})
	}

	return result, nil
}
