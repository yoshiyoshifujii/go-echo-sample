package repository

import (
	"context"
	"errors"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/domain"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, id int64) (domain.User, error)
}
