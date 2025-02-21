package mocks

import (
	"context"
	"sync"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/application/service"
)

// mockMailService implements service.MailService interface for testing
type mockMailService struct {
	mu         sync.RWMutex
	sentEmails map[string]bool
}

// NewMockMailService creates a new instance of mockMailService
func NewMockMailService() service.MailService {
	return &mockMailService{
		sentEmails: make(map[string]bool),
	}
}

// SendWelcomeEmail simulates sending a welcome email and records it
func (s *mockMailService) SendWelcomeEmail(ctx context.Context, email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sentEmails[email] = true
	return nil
}

// VerifyEmailSent checks if an email was sent to the specified address
func (s *mockMailService) VerifyEmailSent(email string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.sentEmails[email]
}
