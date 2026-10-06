package usecase

import (
	"context"
	"errors"
	"strings"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/repository"
)

type UserUsecase interface {
	ListUsers(ctx context.Context, role, search string) ([]domain.User, error)
	CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error)
	UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64, requesterID int64) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
}

func (u *userUsecase) ListUsers(ctx context.Context, role, search string) ([]domain.User, error) {
	if u.userRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}
	return u.userRepo.ListUsers(ctx, role, search)
}

func (u *userUsecase) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	if u.userRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.PhoneWA = strings.TrimSpace(req.PhoneWA)

	if req.Name == "" {
		return nil, errors.New("nama lengkap wajib diisi")
	}
	if req.Email == "" {
		return nil, errors.New("alamat email wajib diisi")
	}
	if req.Password == "" || len(req.Password) < 6 {
		return nil, errors.New("kata sandi wajib minimal 6 karakter")
	}
	if req.PhoneWA == "" {
		return nil, errors.New("nomor WhatsApp resmi wajib diisi")
	}
	if req.Role == "" {
		req.Role = domain.RoleKlien
	}

	return u.userRepo.CreateUser(ctx, req)
}

func (u *userUsecase) UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (*domain.User, error) {
	if u.userRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if req.ID <= 0 {
		return nil, errors.New("ID pengguna tidak valid")
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.PhoneWA = strings.TrimSpace(req.PhoneWA)

	if req.Name == "" {
		return nil, errors.New("nama lengkap tidak boleh kosong")
	}
	if req.Email == "" {
		return nil, errors.New("alamat email tidak boleh kosong")
	}
	if req.PhoneWA == "" {
		return nil, errors.New("nomor WhatsApp resmi tidak boleh kosong")
	}

	if req.Password != "" && len(req.Password) < 6 {
		return nil, errors.New("kata sandi baru minimal 6 karakter")
	}

	return u.userRepo.UpdateUser(ctx, req)
}

func (u *userUsecase) DeleteUser(ctx context.Context, id int64, requesterID int64) error {
	if u.userRepo == nil {
		return nil, errors.New("koneksi basis data belum siap")
	}

	if id <= 0 {
		return nil, errors.New("ID pengguna tidak valid")
	}

	if id == requesterID {
		return errors.New("Anda tidak dapat menghapus akun Anda sendiri yang sedang aktif")
	}

	return u.userRepo.DeleteUser(ctx, id)
}
