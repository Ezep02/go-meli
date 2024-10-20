package repository

import (
	"context"
	"log"

	"github.com/ezep02/auth/internal/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	Connection *gorm.DB
}

func NewAuthRepository(DATABASE *gorm.DB) *AuthRepository {
	return &AuthRepository{
		Connection: DATABASE,
	}
}

func (auth_repo *AuthRepository) SignUp(ctx context.Context, user *models.User) (*models.RegisterUserRes, error) {

	log.Printf("Sign up")

	result := auth_repo.Connection.WithContext(ctx).Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.RegisterUserRes{
		Model:        user.Model,
		Name:         user.Name,
		Email:        user.Email,
		Is_admin:     user.IsAdmin,
		Surname:      user.Surname,
		Phone_number: user.Phone_number,
	}, nil
}

func (auth_repo *AuthRepository) SignIn(ctx context.Context, user *models.LoginUserReq) (*models.LoginUserRes, error) {

	log.Println("Log in user: ", user)

	var u models.User

	result := auth_repo.Connection.WithContext(ctx).Where("email = ?", user.Email).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return &models.LoginUserRes{
		User: models.User{
			Model:    u.Model,
			Name:     u.Name,
			Email:    u.Email,
			Password: u.Password,
			IsAdmin:  u.IsAdmin,
			Surname:  u.Surname,
		},
	}, nil

}
