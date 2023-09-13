package user

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository interface {
	Register(ctx context.Context, user User) (User, error)
	Login(ctx context.Context, username string) (User, error)
	User(ctx context.Context, id int) (User, error)
}

type userRepository struct {
	database *sql.DB
}

func (r *userRepository) Register(ctx context.Context, user User) (User, error) {
	query := "INSERT INTO users (name, username, email, password) VALUES (?, ?, ?, ?)"
	result, err := r.database.ExecContext(ctx, query, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return user, fmt.Errorf("repository.register: %v", err)
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return user, fmt.Errorf("repository.register: %v", err)
	}

	user.ID = int(lastInsertedID)

	return user, nil
}

func (r *userRepository) Login(ctx context.Context, username string) (User, error) {
	var user User

	query := "SELECT id, name, username, email, password, created_at, updated_at FROM users WHERE username = ? OR email = ?"
	if err := r.database.QueryRowContext(ctx, query, username, username).Scan(
		&user.ID, &user.Name, &user.Username, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		return User{}, fmt.Errorf("repository.login: %v", err)
	}

	return user, nil
}

func (r *userRepository) User(ctx context.Context, id int) (User, error) {
	var user User

	query := "SELECT id, name, username, email, password, created_at, updated_at FROM users WHERE id = ?"
	if err := r.database.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Username, &user.Email,
		&user.Password, &user.CreatedAt, &user.UpdatedAt,
	); err != nil {
		return User{}, fmt.Errorf("repository.user: %v", err)
	}

	return user, nil
}
