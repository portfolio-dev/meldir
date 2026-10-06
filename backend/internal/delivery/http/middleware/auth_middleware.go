package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"meldir-backend/internal/domain"
	"meldir-backend/internal/infrastructure/cache"
	"meldir-backend/internal/pkg/token"
)

type contextKey string

const (
	ContextKeyUserID contextKey = "user_id"
	ContextKeyEmail  contextKey = "email"
	ContextKeyRole   contextKey = "role"
	ContextKeyToken  contextKey = "raw_token"
)

type AuthMiddleware struct {
	jwtManager *token.JWTManager
	redis      *cache.RedisClient
}

func NewAuthMiddleware(jwtManager *token.JWTManager, redis *cache.RedisClient) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
		redis:      redis,
	}
}

func (m *AuthMiddleware) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			writeJSONError(w, http.StatusUnauthorized, "Sesi login tidak ditemukan (Bearer token missing)")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Periksa apakah token ada di blacklist Redis (hasil force logout)
		if m.redis != nil {
			isBlacklisted, err := m.redis.IsTokenBlacklisted(r.Context(), tokenString)
			if err == nil && isBlacklisted {
				writeJSONError(w, http.StatusUnauthorized, "Sesi telah dicabut atau kedaluwarsa. Silakan login kembali.")
				return
			}
		}

		// Validasi token JWT
		claims, err := m.jwtManager.ValidateToken(tokenString)
		if err != nil {
			writeJSONError(w, http.StatusUnauthorized, "Token otentikasi tidak valid atau telah kedaluwarsa")
			return
		}

		// Suntikkan data user ke request context
		ctx := context.WithValue(r.Context(), ContextKeyUserID, claims.UserID)
		ctx = context.WithValue(ctx, ContextKeyEmail, claims.Email)
		ctx = context.WithValue(ctx, ContextKeyRole, claims.Role)
		ctx = context.WithValue(ctx, ContextKeyToken, tokenString)

		next(w, r.WithContext(ctx))
	}
}

func (m *AuthMiddleware) RequireRoles(allowedRoles ...domain.UserRole) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return m.Authenticate(func(w http.ResponseWriter, r *http.Request) {
			roleVal := r.Context().Value(ContextKeyRole)
			if roleVal == nil {
				writeJSONError(w, http.StatusForbidden, "Hak akses peranan tidak terdefinisi")
				return
			}

			userRole := roleVal.(domain.UserRole)
			hasAccess := false
			for _, allowed := range allowedRoles {
				if userRole == allowed {
					hasAccess = true
					break
				}
			}

			if !hasAccess {
				writeJSONError(w, http.StatusForbidden, "Anda tidak memiliki wewenang untuk mengakses sumber daya ini")
				return
			}

			next(w, r)
		})
	}
}

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": false,
		"error":   message,
	})
}
