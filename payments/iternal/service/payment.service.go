package service

import (
	"context"

	"github.com/ezep/go-meli/payments/iternal/models"
	"github.com/ezep/go-meli/payments/iternal/repository"
)

// Definición del servicio PaymentService
type PaymentService struct {
	PaymentRepository *repository.PaymentRepository
}

// Constructor para inicializar el PaymentService
func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{
		PaymentRepository: repo,
	}
}

// Método del servicio para crear una orden
func (payService *PaymentService) CreateOrderService(ctx context.Context, item models.Item) (*models.Item, error) {

	return payService.PaymentRepository.CreateOrder(ctx, item)
}
