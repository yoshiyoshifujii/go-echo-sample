package memory

import (
	"context"
	"sync"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/domain"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/repository"
)

type UserRepository struct {
	mu     sync.Mutex
	nextID int64
	users  map[int64]domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		nextID: 1,
		users:  make(map[int64]domain.User),
	}
}

func (r *UserRepository) Save(_ context.Context, user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}

	r.users[user.ID] = user

	return user, nil
}

func (r *UserRepository) FindByID(_ context.Context, id int64) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.users[id]
	if !ok {
		return domain.User{}, repository.ErrUserNotFound
	}

	return user, nil
}
