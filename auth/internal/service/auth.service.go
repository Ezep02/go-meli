package service

import (
	"context"

	"github.com/ezep02/auth/internal/models"
	"github.com/ezep02/auth/internal/repository"
)

type AuthService struct {
	AuthRepository *repository.AuthRepository
}

func NewAuthService(auth_repository *repository.AuthRepository) *AuthService {
	return &AuthService{
		AuthRepository: auth_repository,
	}
}

func (auth_service *AuthService) SignUpService(ctx context.Context, user *models.User) (*models.RegisterUserRes, error) {
	return auth_service.AuthRepository.SignUp(ctx, user)
}

func (auth_service *AuthService) SignInService(ctx context.Context, user *models.LoginUserReq) (*models.LoginUserRes, error) {
	return auth_service.AuthRepository.SignIn(ctx, user)
}
