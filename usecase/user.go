package usecase

import (
	"context"
	"entedge/domain"
	"entedge/repository"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

type UserUseCase interface {
	CreateUser(ctx context.Context, user *domain.User) error
	// GetUser(id int) error
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (uc *userUseCase) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := uc.userRepo.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
