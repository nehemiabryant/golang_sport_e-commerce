package services

import (
	"context"
	"errors"
	"miniproject-nehemia/models"
	"miniproject-nehemia/repositories"
)

type CartService struct {
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewCartService(c *repositories.CartRepository, p *repositories.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    c,
		productRepo: p,
	}
}

func (s *CartService) AddToCart(ctx context.Context, userID int, productID string, qty int) error {
	product, err := s.productRepo.GetProductByID(ctx, productID)
	if err != nil {
		return err
	}
	if product.Stock < qty {
		return errors.New("stok tidak cukup")
	}

	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}

	if err := s.cartRepo.AddOrIncrementItem(ctx, cartID, productID, qty); err != nil {
		return err
	}

	return s.productRepo.DecreaseStock(ctx, productID, qty)
}

func (s *CartService) GetCartItems(ctx context.Context, userID int) ([]models.CartItemDetail, error) {
	return s.cartRepo.GetCartItemsDetail(ctx, userID)
}

func (s *CartService) UpdateQuantity(ctx context.Context, userID int, productID string, qty int) error {
	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}
	return s.cartRepo.UpdateItemQuantity(ctx, cartID, productID, qty)
}

func (s *CartService) RemoveItem(ctx context.Context, userID int, productID string) error {
	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}
	return s.cartRepo.RemoveItem(ctx, cartID, productID)
}

func (s *CartService) ClearCart(ctx context.Context, userID int) error {
	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}
	return s.cartRepo.ClearCart(ctx, cartID)
}

func (s *CartService) AddQty(ctx context.Context, userID int, productID string) error {
	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}

	return s.cartRepo.AddOrIncrementItem(ctx, cartID, productID, 1)
}

func (s *CartService) ReduceQty(ctx context.Context, userID int, productID string) error {
	cartID, err := s.cartRepo.GetOrCreateCart(ctx, userID)
	if err != nil {
		return err
	}

	// cek jumlah saat ini
	qty, err := s.cartRepo.GetItemQuantity(ctx, cartID, productID)
	if err != nil {
		return err
	}

	if qty <= 1 {
		return errors.New("tidak bisa mengurangi, qty minimal 1")
	}

	return s.cartRepo.UpdateItemQuantity(ctx, cartID, productID, qty-1)
}
