package repositories

import (
	"context"
	"miniproject-nehemia/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Insert payment header
func (r *PaymentRepository) CreatePayment(ctx context.Context, userID int, total float64) (int, error) {
	query := `
		INSERT INTO payment (user_id, total_amount, status)
		VALUES ($1, $2, 'PAID')
		RETURNING id
	`
	var id int
	err := r.db.QueryRow(ctx, query, userID, total).Scan(&id)
	return id, err
}

// Insert payment detail
func (r *PaymentRepository) InsertPaymentDetail(ctx context.Context, detail models.PaymentDetail) error {
	query := `
		INSERT INTO payment_detail (payment_id, product_id, quantity, price, subtotal)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(ctx,
		query,
		detail.PaymentID,
		detail.ProductID,
		detail.Quantity,
		detail.Price,
		detail.Subtotal,
	)

	return err
}

// Get all payment with details (for admin)
func (r *PaymentRepository) GetAllPayments(ctx context.Context) ([]models.PaymentDetailFull, error) {
	query := `
        SELECT 
            p.id AS payment_id,
            p.user_id,
            p.total_amount,
            p.status,
            p.created_at,
            pd.product_id,
            pr.name AS product_name,
            pd.quantity,
            pd.price,
            pd.subtotal
        FROM payment p
        JOIN payment_detail pd ON p.id = pd.payment_id
        JOIN products pr ON pr.id = pd.product_id
        ORDER BY p.created_at DESC;
    `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.PaymentDetailFull

	for rows.Next() {
		var d models.PaymentDetailFull

		err := rows.Scan(
			&d.PaymentID,
			&d.UserID,
			&d.TotalAmount,
			&d.Status,
			&d.CreatedAt,
			&d.ProductID,
			&d.ProductName,
			&d.Quantity,
			&d.Price,
			&d.Subtotal,
		)
		if err != nil {
			return nil, err
		}

		results = append(results, d)
	}

	return results, nil
}
