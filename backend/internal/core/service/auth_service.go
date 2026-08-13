package service

import (
	"context"
	"errors"
	"thekasir/internal/core/domain"
	"thekasir/internal/core/port"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo  port.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo port.UserRepository, secret string) port.AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: secret,
	}
}

func (s *authService) Register(ctx context.Context, email, name, password string) (*domain.User, error) {
	existing, _ := s.userRepo.GetUserByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// For MVP, we can use email as ID or generate a simple ID.
	// In production, use UUID or NanoID. Here we use a simple placeholder.
	id := "usr_" + time.Now().Format("20060102150405")

	user := &domain.User{
		ID:        id,
		Email:     email,
		Name:      name,
		Password:  string(hashed),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, *domain.User, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil || user == nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
