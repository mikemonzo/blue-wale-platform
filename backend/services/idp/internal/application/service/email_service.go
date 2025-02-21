package service

import "context"

type MailService interface {
	SendWelcomeEmail(ctx context.Context, email string) error
}
