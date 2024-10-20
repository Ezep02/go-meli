package repository

import (
	"context"
	"time"

	"github.com/ezep02/services/internal/models"
	"gorm.io/gorm"
)

type ServiceRepository struct {
	Connection *gorm.DB
}

func NewServicesRepository(DATABASE *gorm.DB) *ServiceRepository {
	return &ServiceRepository{
		Connection: DATABASE,
	}
}

func (Svs_repo *ServiceRepository) CreateService(ctx context.Context, service *models.ServiceModel) (*models.ServiceModel, error) {

	return &models.ServiceModel{
		Model:            service.Model,
		Title:            service.Title,
		User_id:          service.User_id,
		Description:      service.Description,
		Price:            service.Price,
		Service_Duration: service.Service_Duration,
	}, nil
}

func (Svs_repo *ServiceRepository) GetAllServices(ctx context.Context) (*[]models.ServiceModel, error) {

	return &[]models.ServiceModel{
		{
			Model: gorm.Model{
				ID:        2,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Title:            "Corte de pelo",
			User_id:          21,
			Description:      "Corte con shaver",
			Price:            6700,
			Service_Duration: 50,
		},
	}, nil
}
