package usecase

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"meldir-backend/internal/domain"
	"meldir-backend/internal/infrastructure/cache"
	"meldir-backend/internal/pkg/token"
	"meldir-backend/internal/repository"
)

type AuthUsecase interface {
	Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error)
	GetProfile(ctx context.Context, userID int64) (*domain.User, error)
	Logout(ctx context.Context, tokenString string) error
}

type authUsecase struct {
	userRepo   repository.UserRepository
	jwtManager *token.JWTManager
	redis      *cache.RedisClient
}

func NewAuthUsecase(
	userRepo repository.UserRepository,
	jwtManager *token.JWTManager,
	redis *cache.RedisClient,
) AuthUsecase {
	return &authUsecase{
		userRepo:   userRepo,
		jwtManager: jwtManager,
		redis:      redis,
	}
}

func (u *authUsecase) Login(ctx context.Context, req domain.LoginRequest) (*domain.LoginResponse, error) {
	if u.userRepo == nil {
		return nil, errors.New("koneksi basis data PostgreSQL belum terhubung. Periksa kredensial database di server")
	}

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email dan password wajib diisi")
	}

	user, err := u.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("kombinasi email atau kata sandi tidak valid")
	}

	if user.Status != "aktif" {
		return nil, errors.New("status akun Anda belum aktif atau ditangguhkan")
	}

	// Verifikasi kata sandi dengan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("kombinasi email atau kata sandi tidak valid")
	}

	// Update timestamp login terakhir
	_ = u.userRepo.UpdateLastLogin(ctx, user.ID)

	// Terbitkan JWT Token 24 Jam
	jwtToken, err := u.jwtManager.GenerateToken(user)
	if err != nil {
		return nil, errors.New("gagal menerbitkan token otentikasi")
	}

	return &domain.LoginResponse{
		Token: jwtToken,
		User:  *user,
	}, nil
}

func (u *authUsecase) GetProfile(ctx context.Context, userID int64) (*domain.User, error) {
	return u.userRepo.FindByID(ctx, userID)
}

func (u *authUsecase) Logout(ctx context.Context, tokenString string) error {
	if u.redis != nil && tokenString != "" {
		// Masukkan token ke Redis blacklist selama 24 jam
		return u.redis.BlacklistToken(ctx, tokenString, 24*time.Hour)
	}
	return nil
}
