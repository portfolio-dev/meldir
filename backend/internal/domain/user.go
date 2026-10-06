package domain

import "time"

type UserRole string

const (
	RoleDirektur UserRole = "direktur"
	RoleAdmin    UserRole = "admin"
	RoleEngineer UserRole = "engineer"
	RoleKlien    UserRole = "klien"
	RoleAudit    UserRole = "audit"
)

type User struct {
	ID             int64      `json:"id"`
	Name           string     `json:"name"`
	Email          string     `json:"email"`
	PasswordHash   string     `json:"-"`
	Role           UserRole   `json:"role"`
	EngineerType   string     `json:"engineer_type"`
	ClientType     string     `json:"client_type"`
	PhoneWA        string     `json:"phone_wa"`
	AvatarURL      *string    `json:"avatar_url"`
	GitHubUsername *string    `json:"github_username"`
	Status         string     `json:"status"`
	LastLoginAt    *time.Time `json:"last_login_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type CreateUserRequest struct {
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password"`
	Role           UserRole `json:"role"`
	EngineerType   string   `json:"engineer_type"`
	ClientType     string   `json:"client_type"`
	PhoneWA        string   `json:"phone_wa"`
	Status         string   `json:"status"`
	GitHubUsername *string  `json:"github_username,omitempty"`
}

type UpdateUserRequest struct {
	ID             int64    `json:"id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Password       string   `json:"password,omitempty"`
	Role           UserRole `json:"role"`
	EngineerType   string   `json:"engineer_type"`
	ClientType     string   `json:"client_type"`
	PhoneWA        string   `json:"phone_wa"`
	Status         string   `json:"status"`
	GitHubUsername *string  `json:"github_username,omitempty"`
}
