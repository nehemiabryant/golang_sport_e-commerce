package services

import (
	"context"
	"miniproject-nehemia/models"
	"miniproject-nehemia/repositories"
)

type PaymentService struct {
	paymentRepo *repositories.PaymentRepository
	cartRepo    *repositories.CartRepository
	productRepo *repositories.ProductRepository
}

func NewPaymentService(
	p *repositories.PaymentRepository,
	c *repositories.CartRepository,
	pr *repositories.ProductRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepo: p,
		cartRepo:    c,
		productRepo: pr,
	}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, userID int) (int, error) {
	// Ambil isi keranjang
	items, err := s.cartRepo.GetCartItemsDetail(ctx, userID)
	if err != nil {
		return 0, err
	}

	var total float64
	for _, it := range items {
		total += it.Subtotal
	}

	if total == 0 {
		return 0, nil
	}

	// Buat payment header
	paymentID, err := s.paymentRepo.CreatePayment(ctx, userID, total)
	if err != nil {
		return 0, err
	}

	// Insert detail + Kurangi stok
	for _, item := range items {
		detail := models.PaymentDetail{
			PaymentID: paymentID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
			Subtotal:  item.Subtotal,
		}

		// simpan detail
		err = s.paymentRepo.InsertPaymentDetail(ctx, detail)
		if err != nil {
			return 0, err
		}

		// kurangi stok
		err = s.productRepo.DecreaseStock(ctx, item.ProductID, item.Quantity)
		if err != nil {
			return 0, err
		}
	}

	// Kosongkan cart
	cartID, _ := s.cartRepo.GetOrCreateCart(ctx, userID)
	s.cartRepo.ClearCart(ctx, cartID)

	return paymentID, nil
}
