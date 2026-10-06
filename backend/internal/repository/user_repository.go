package repository

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"meldir-backend/internal/domain"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	FindByID(ctx context.Context, id int64) (*domain.User, error)
	UpdateLastLogin(ctx context.Context, id int64) error
	ListUsers(ctx context.Context, role, search string) ([]domain.User, error)
	CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error)
	UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx context.Context, id int64) error
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

func (r *userRepo) ListUsers(ctx context.Context, role, search string) ([]domain.User, error) {
	query := `
		SELECT id, name, email, role, engineer_type, client_type,
		       phone_wa, avatar_url, github_username, status, last_login_at, created_at, updated_at
		FROM users
		WHERE ($1 = '' OR role::text = $1)
		  AND ($2 = '' OR LOWER(name) LIKE '%' || LOWER($2) || '%' OR LOWER(email) LIKE '%' || LOWER($2) || '%')
		ORDER BY id ASC
	`
	rows, err := r.pool.Query(ctx, query, role, strings.TrimSpace(search))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		var u domain.User
		var roleStr string
		err := rows.Scan(
			&u.ID, &u.Name, &u.Email, &roleStr,
			&u.EngineerType, &u.ClientType, &u.PhoneWA, &u.AvatarURL,
			&u.GitHubUsername, &u.Status, &u.LastLoginAt, &u.CreatedAt, &u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		u.Role = domain.UserRole(roleStr)
		users = append(users, u)
	}

	return users, rows.Err()
}

func (r *userRepo) CreateUser(ctx context.Context, req domain.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi kata sandi")
	}

	if req.EngineerType == "" {
		req.EngineerType = "none"
	}
	if req.ClientType == "" {
		req.ClientType = "none"
	}
	if req.Status == "" {
		req.Status = "aktif"
	}

	query := `
		INSERT INTO users (
			name, email, password_hash, role, engineer_type, client_type, phone_wa, status, github_username
		) VALUES (
			$1, $2, $3, $4::user_role, $5::engineer_type, $6::client_type, $7, $8, $9
		) RETURNING id, created_at, updated_at
	`

	var user domain.User
	user.Name = req.Name
	user.Email = req.Email
	user.Role = req.Role
	user.EngineerType = req.EngineerType
	user.ClientType = req.ClientType
	user.PhoneWA = req.PhoneWA
	user.Status = req.Status
	user.GitHubUsername = req.GitHubUsername

	err = r.pool.QueryRow(ctx, query,
		req.Name, req.Email, string(hashedPassword),
		string(req.Role), req.EngineerType, req.ClientType,
		req.PhoneWA, req.Status, req.GitHubUsername,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New("email sudah terdaftar dalam sistem")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, req domain.UpdateUserRequest) (*domain.User, error) {
	if req.EngineerType == "" {
		req.EngineerType = "none"
	}
	if req.ClientType == "" {
		req.ClientType = "none"
	}
	if req.Status == "" {
		req.Status = "aktif"
	}

	var query string
	var err error

	if req.Password != "" {
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if hashErr != nil {
			return nil, errors.New("gagal mengenkripsi kata sandi baru")
		}

		query = `
			UPDATE users SET
				name = $1,
				email = $2,
				password_hash = $3,
				role = $4::user_role,
				engineer_type = $5::engineer_type,
				client_type = $6::client_type,
				phone_wa = $7,
				status = $8,
				github_username = $9,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $10
			RETURNING updated_at
		`
		var updatedAt time.Time
		err = r.pool.QueryRow(ctx, query,
			req.Name, req.Email, string(hashedPassword),
			string(req.Role), req.EngineerType, req.ClientType,
			req.PhoneWA, req.Status, req.GitHubUsername, req.ID,
		).Scan(&updatedAt)
	} else {
		query = `
			UPDATE users SET
				name = $1,
				email = $2,
				role = $3::user_role,
				engineer_type = $4::engineer_type,
				client_type = $5::client_type,
				phone_wa = $6,
				status = $7,
				github_username = $8,
				updated_at = CURRENT_TIMESTAMP
			WHERE id = $9
			RETURNING updated_at
		`
		var updatedAt time.Time
		err = r.pool.QueryRow(ctx, query,
			req.Name, req.Email,
			string(req.Role), req.EngineerType, req.ClientType,
			req.PhoneWA, req.Status, req.GitHubUsername, req.ID,
		).Scan(&updatedAt)
	}

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("pengguna tidak ditemukan")
		}
		if strings.Contains(err.Error(), "unique") || strings.Contains(err.Error(), "duplicate") {
			return nil, errors.New("email sudah digunakan oleh akun lain")
		}
		return nil, err
	}

	return r.FindByID(ctx, req.ID)
}

func (r *userRepo) DeleteUser(ctx context.Context, id int64) error {
	cmd, err := r.pool.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("pengguna tidak ditemukan")
	}
	return nil
}
