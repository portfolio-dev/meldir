package repository

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"meldir-backend/internal/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id int64) (*domain.User, error)
	UpdateLastLogin(ctx context.Context, id int64) error
}

type userRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) UserRepository {
	return &userRepo{pool: pool}
}

func (r *userRepo) FindByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
		SELECT id, name, email, password_hash, role, engineer_type, client_type,
		       phone_wa, avatar_url, github_username, status, last_login_at, created_at, updated_at
		FROM users
		WHERE LOWER(email) = LOWER($1)
		LIMIT 1
	`
	var user domain.User
	var roleStr string
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &roleStr,
		&user.EngineerType, &user.ClientType, &user.PhoneWA, &user.AvatarURL,
		&user.GitHubUsername, &user.Status, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("pengguna dengan email tersebut tidak ditemukan")
		}
		return nil, err
	}
	user.Role = domain.UserRole(roleStr)
	return &user, nil
}

func (r *userRepo) FindByID(ctx context.Context, id int64) (*domain.User, error) {
	query := `
		SELECT id, name, email, password_hash, role, engineer_type, client_type,
		       phone_wa, avatar_url, github_username, status, last_login_at, created_at, updated_at
		FROM users
		WHERE id = $1
		LIMIT 1
	`
	var user domain.User
	var roleStr string
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.PasswordHash, &roleStr,
		&user.EngineerType, &user.ClientType, &user.PhoneWA, &user.AvatarURL,
		&user.GitHubUsername, &user.Status, &user.LastLoginAt, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		return nil, err
	}
	user.Role = domain.UserRole(roleStr)
	return &user, nil
}

func (r *userRepo) UpdateLastLogin(ctx context.Context, id int64) error {
	now := time.Now().UTC()
	_, err := r.pool.Exec(ctx, "UPDATE users SET last_login_at = $1 WHERE id = $2", now, id)
	return err
}
