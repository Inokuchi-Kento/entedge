package usecase

import (
	"context"
	"entedge/domain"
	"entedge/repository"
)

type PetUseCase interface {
	CreatePet(ctx context.Context, pet *domain.Pet, id int) error
}

type petUseCase struct {
	petRepo  repository.PetRepository
	userRepo repository.UserRepository
}

func NewPetUseCase(
	petRepo repository.PetRepository,
	userRepo repository.UserRepository,
) PetUseCase {
	return &petUseCase{
		petRepo:  petRepo,
		userRepo: userRepo,
	}
}

func (uc *petUseCase) CreatePet(ctx context.Context, pet *domain.Pet, id int) error {
	owner, err := uc.userRepo.GetUser(ctx, id)
	if err != nil {
		return err
	}
	_, err = uc.petRepo.CreatePet(ctx, pet, owner)
	if err != nil {
		return err
	}
	return nil
}
