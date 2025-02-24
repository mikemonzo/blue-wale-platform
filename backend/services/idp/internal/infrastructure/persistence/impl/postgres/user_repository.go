package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/model"
	"github.com/mikemonzo/blue-whale-platform/backend/services/idp/internal/domain/repository"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) repository.UserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *model.User) error {
	query := `
	INSERT INTO users (id, email, username, first_name, last_name, password, status, created_at, updated_at) 
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password, status, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRowContext(ctx, query, email)

	var user model.User
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error getting user by email: %w", err)
	}
	return &user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET email = $1, username = $2, first_name = $3, last_name = $4, password = $5, status = $6, updated_at = $7 WHERE id = $8`
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.FirstName, user.LastName, user.Password, user.Status, user.UpdatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

func (r *PostgresUserRepository) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	query := `SELECT id, email, username, first_name, last_name, password, status, created_at, updated_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting all users: %w", err)
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.FirstName, &user.LastName, &user.Password, &user.Status, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, fmt.Errorf("error scanning user: %w", err)
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users: %w", err)
	}

	return users, nil
}
