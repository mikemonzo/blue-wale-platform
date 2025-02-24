package memory

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
)

type inMemoryUserRepository struct {
	users map[string]*model.User
	mu    sync.RWMutex
}

func NewInMemoryUserRepository() *inMemoryUserRepository {
	return &inMemoryUserRepository{
		users: make(map[string]*model.User),
	}
}

func (r *inMemoryUserRepository) Create(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.Email]; exists {
		return fmt.Errorf("user with email %s already exists", user.Email)
	}

	r.users[user.Email] = user
	log.Printf("Created user in repository: %+v\n", user)
	return nil
}

func (r *inMemoryUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if user, exists := r.users[email]; exists {
		log.Printf("Found user by email: %+v\n", user)
		return user, nil
	}

	log.Printf("User not found by email: %s\n", email)
	return nil, nil
}

func (r *inMemoryUserRepository) Update(ctx context.Context, user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Check if the new email already exists for a different user
	for _, u := range r.users {
		if u.Email == user.Email && u.ID != user.ID {
			log.Printf("Email already exists for a different user: %s\n", user.Email)
			return fmt.Errorf("email already exists")
		}
	}

	// Remove the old email key if the email has changed
	for email, u := range r.users {
		if u.ID == user.ID && email != user.Email {
			delete(r.users, email)
			log.Printf("Removed old email key: %s\n", email)
			break
		}
	}

	// Update user
	r.users[user.Email] = user
	log.Printf("Updated user in repository: %+v\n", user)
	return nil
}
