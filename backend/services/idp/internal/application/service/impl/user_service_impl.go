package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repository"
)

type userService struct {
	userRepo    repository.UserRepository
	mailService service.MailService
}

func NewUserService(userRepo repository.UserRepository, mailService service.MailService) service.UserService {
	return &userService{
		userRepo:    userRepo,
		mailService: mailService,
	}
}

func (s *userService) Create(ctx context.Context, email, username, firstName, lastName, password string) (*model.User, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("error checking existing user: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}

	// Create new user
	user := &model.User{
		ID:        uuid.New().String(),
		Email:     email,
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
		Status:    model.UserStatusInactive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save user to repository
	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}

	// Send welcome email
	if err := s.mailService.SendWelcomeEmail(ctx, user.Email); err != nil {
		// Log error but do not fail the creation of the user
		fmt.Printf("Failed to send welcome email to user %s: %v\n", user.Email, err)
	}

	return user, nil
}

func (s *userService) Update(ctx context.Context, user *model.User) error {
	// First get the existing user by ID to check old email
	oldUser, err := s.userRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return fmt.Errorf("error retrieving user: %w", err)
	}

	// Check if new email already exists for a different user
	if oldUser != nil && oldUser.ID != user.ID {
		return fmt.Errorf("email already exists")
	}

	// Update user
	user.UpdatedAt = time.Now()
	if err := s.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}

	return nil
}
