package usecase

import (
	"context"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/domain"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/repository"
)

type CreateUser struct {
	userRepo repository.UserRepository
}

func NewCreateUser(userRepo repository.UserRepository) *CreateUser {
	return &CreateUser{
		userRepo: userRepo,
	}
}

func (uc *CreateUser) Execute(ctx context.Context, name, email string) (domain.User, error) {
	user := domain.User{
		Name:  name,
		Email: email,
	}

	return uc.userRepo.Save(ctx, user)
}
