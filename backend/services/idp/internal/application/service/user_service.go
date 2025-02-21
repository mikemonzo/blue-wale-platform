package service

import (
	"context"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
)

type UserService interface {
	Create(ctx context.Context, email, username, firstName, lastName, password string) (*model.User, error)
}
