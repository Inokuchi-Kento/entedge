package repository

import (
	"context"
	"entedge/domain"
	"entedge/ent"
)

type PetRepository interface {
	CreatePet(ctx context.Context, p *domain.Pet, owner *ent.User) (domain.Pet, error)
	GetPet(ctx context.Context, id int) (*ent.Pet, error)
}

type petRepository struct {
	DB *ent.Client
}

func NewPetRepository(db *ent.Client) PetRepository {
	return &petRepository{
		DB: db,
	}
}

func (pr *petRepository) CreatePet(ctx context.Context, p *domain.Pet, owner *ent.User) (pet domain.Pet, err error) {
	createPetRow, err := pr.DB.Pet.Create().
		SetName(p.Name).
		SetOwner(owner).
		Save(ctx)
	if err != nil {
		return pet, err
	}
	pet = domain.Pet{
		ID:   createPetRow.ID,
		Name: createPetRow.Name,
	}
	return pet, nil
}

func (pr *petRepository) GetPet(ctx context.Context, id int) (*ent.Pet, error) {
	p, err := pr.DB.Pet.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
