package memory

import (
	"context"
	"sync"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repository"
)

type inMemoryUserRepository struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() repository.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.Email] = user
	return nil
}

func (r *inMemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[email]; exists {
		return user, nil
	}

	return nil, nil
}
