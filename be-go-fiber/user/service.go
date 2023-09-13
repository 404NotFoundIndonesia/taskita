package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type service struct{ repository Repository }

func NewService(repository Repository) service {
	return service{repository: repository}
}

func (s service) Register(ctx context.Context, request RegisterRequest) (RegisterResponse, error) {
	// Hashing password
	hashed, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		return RegisterResponse{}, ErrRegisterFailedToHashPassword
	}

	// Check if user with registered email already exists
	_, err = s.repository.Login(ctx, request.Email)
	if !errors.Is(err, sql.ErrNoRows) {
		return RegisterResponse{}, ErrRegisterAccountAlreadyExists
	}

	// Register new account
	user, err := s.repository.Register(ctx, User{
		Name:      request.Name,
		Username:  request.Email,
		Email:     request.Email,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return RegisterResponse{}, ErrRegisterFailedToRegisterAccount
	}

	// Set token expired time
	expiredAt := time.Now().Add(time.Hour * 24)

	// Generate json web token
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["username"] = user.Username
	claims["expired"] = expiredAt.Unix()
	token, _ := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return RegisterResponse{
		ID:             user.ID,
		Name:           user.Name,
		Username:       user.Username,
		Email:          user.Email,
		Token:          token,
		TokenExpiredAt: expiredAt,
	}, nil
}
