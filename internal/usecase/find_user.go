package usecase

import (
	"context"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/domain"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/repository"
)

type FindUser struct {
	userRepo repository.UserRepository
}

func NewFindUser(userRepo repository.UserRepository) *FindUser {
	return &FindUser{
		userRepo: userRepo,
	}
}

func (uc *FindUser) Execute(ctx context.Context, id int64) (domain.User, error) {
	return uc.userRepo.FindByID(ctx, id)
}
