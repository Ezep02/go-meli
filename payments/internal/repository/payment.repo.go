package repository

import (
	"context"

	"github.com/ezep02/payments/internal/models"
)

type PaymentRepository struct {
	items []models.Item
}

func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		items: []models.Item{},
	}
}

// TODO: crear estructuras para crear las ordenes y para solicitarlas para devolverlas al front

func (pay_repo *PaymentRepository) CreateOrder(ctx context.Context, item models.Item) (*models.Item, error) {
	println("Hello, you add a new order")
	return &item, nil
}
