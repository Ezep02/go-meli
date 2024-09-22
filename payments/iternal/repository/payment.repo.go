package repository

import (
	"context"

	"github.com/ezep/go-meli/payments/iternal/models"
)

type PaymentRepository struct {
	items []models.Item
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		items: []models.Item{},
	}
}

func (pay_repo *PaymentRepository) CreateOrder(ctx context.Context, item models.Item) (*models.Item, error) {
	println("Hello, you add a new order")
	return &item, nil
}
