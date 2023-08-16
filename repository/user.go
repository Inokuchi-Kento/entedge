package repository

import (
	"context"
	"entedge/domain"
	"entedge/ent"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (domain.User, error)
	GetUser(ctx context.Context, id int) (*ent.User, error)
}

type userRepository struct {
	DB *ent.Client
}

func NewUserRepository(db *ent.Client) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, u *domain.User) (user domain.User, err error) {
	createUserRow, err := ur.DB.User.Create().
		SetAge(u.Age).
		SetName(u.Name).
		Save(ctx)
	if err != nil {
		return user, err
	}
	user = domain.User{
		ID:   createUserRow.ID,
		Name: createUserRow.Name,
		Age:  createUserRow.Age,
	}

	return user, nil
}

func (ur *userRepository) GetUser(ctx context.Context, id int) (*ent.User, error) {
	u, err := ur.DB.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
