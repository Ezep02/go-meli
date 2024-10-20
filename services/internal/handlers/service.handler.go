package handlers

import (
	"context"

	"github.com/ezep02/services/internal/service"
)

type ServiceHandler struct {
	Svc_ser *service.NewServicesService
	Ctx     context.Context
}

func NewServiceHandler(svc_ser *service.NewServicesService) *ServiceHandler {
	return &ServiceHandler{
		Svc_ser: svc_ser,
		Ctx:     context.Background(),
	}
}
