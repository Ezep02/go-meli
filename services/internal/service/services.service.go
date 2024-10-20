package service

import (
	"context"

	"github.com/ezep02/services/internal/models"
	"github.com/ezep02/services/internal/repository"
)

type NewServicesService struct {
	Svs_repo *repository.ServiceRepository
}

func ServiceService(Svs_repo *repository.ServiceRepository) *NewServicesService {
	return &NewServicesService{
		Svs_repo: Svs_repo,
	}
}

func (Svs_ser *NewServicesService) CreateService(ctx context.Context, service *models.ServiceModel) (*models.ServiceModel, error) {
	return Svs_ser.Svs_repo.CreateService(ctx, service)
}

func (Svs_ser *NewServicesService) GetAllServices(ctx context.Context) (*[]models.ServiceModel, error) {
	return Svs_ser.Svs_repo.GetAllServices(ctx)
}
