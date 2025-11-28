package services

import (
	"context"
	"errors"
	"miniproject-nehemia/models"
	"miniproject-nehemia/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(r *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: r}
}

func (s *ProductService) GetByID(ctx context.Context, id string) (*models.Product, error) {
	return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) DecreaseStock(ctx context.Context, id string, qty int) error {
	return s.repo.DecreaseStock(ctx, id, qty)
}

func (s *ProductService) IncreaseStock(ctx context.Context, id string, qty int) error {
	return s.repo.IncreaseStock(ctx, id, qty)
}

// Add new product - Admin function
func (s *ProductService) AddProduct(ctx context.Context, id string, name string, price float64, stock int) error {

	// cek jika ID sudah ada
	existing, _ := s.repo.GetProductByID(ctx, id)
	if existing != nil {
		return errors.New("produk dengan ID tersebut sudah ada")
	}

	product := models.Product{
		ID:    id,
		Name:  name,
		Price: price,
		Stock: stock,
	}

	return s.repo.Create(ctx, product)
}

// Restock product - Admin function
func (s *ProductService) RestockProduct(ctx context.Context, id string, qty int) error {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	product.Stock += qty
	return s.repo.UpdateStock(ctx, id, product.Stock)
}

// Decrease stock - Admin function
func (s *ProductService) AdminDecreaseStock(ctx context.Context, id string, qty int) error {
	product, err := s.repo.GetProductByID(ctx, id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	if product.Stock < qty {
		return errors.New("stok tidak mencukupi")
	}

	product.Stock -= qty
	return s.repo.UpdateStock(ctx, id, product.Stock)
}

// All products - Admin function
func (s *ProductService) GetAllProducts(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) AssignProductToCategory(ctx context.Context, productID, categoryID string) error {
	return s.repo.AssignProductToCategory(ctx, productID, categoryID)
}

// Create product and assign to category - Admin function
func (s *ProductService) CreateProduct(ctx context.Context, p models.Product, categoryID string) error {
	// 1. Create product
	err := s.repo.Create(ctx, p)
	if err != nil {
		return err
	}

	// 2. Assign to category
	return s.repo.AssignProductToCategory(ctx, p.ID, categoryID)
}

func (s *ProductService) GetProductsGroupedByCategory(ctx context.Context) (map[string][]models.Product, error) {
	return s.repo.GetProductsGroupedByCategory(ctx)
}

func (s *PaymentService) GetCartTotal(ctx context.Context, userID int) (float64, error) {
	items, err := s.cartRepo.GetCartItemsDetail(ctx, userID)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, it := range items {
		total += it.Subtotal
	}

	return total, nil
}

func (s *PaymentService) GetAllPayments(ctx context.Context) ([]models.PaymentDetailFull, error) {
	return s.paymentRepo.GetAllPayments(ctx)
}
