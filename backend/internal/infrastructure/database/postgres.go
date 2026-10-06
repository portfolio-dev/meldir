package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"meldir-backend/internal/config"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("gagal mem-parsing DSN PostgreSQL: %w", err)
	}

	poolConfig.ConnConfig.User = cfg.DBUser
	poolConfig.ConnConfig.Password = cfg.DBPassword

	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat connection pool PostgreSQL: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("gagal melakukan ping PostgreSQL: %w", err)
	}

	log.Printf("🐘 Terhubung sukses ke basis data PostgreSQL (%s:%s/%s)", cfg.DBHost, cfg.DBPort, cfg.DBName)

	db := &PostgresDB{Pool: pool}
	if err := db.SeedDefaultSuperadmin(context.Background()); err != nil {
		log.Printf("⚠️ Peringatan seed admin: %v", err)
	}

	return db, nil
}

func (db *PostgresDB) SeedDefaultSuperadmin(ctx context.Context) error {
	var count int
	err := db.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return fmt.Errorf("gagal memeriksa tabel users: %w", err)
	}

	if count > 0 {
		return nil // Sudah ada akun terdaftar
	}

	defaultPassword := "MeldirAdmin2026!"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("gagal melakukan hash kata sandi: %w", err)
	}

	query := `
		INSERT INTO users (
			name, email, password_hash, role, engineer_type, client_type, phone_wa, status
		) VALUES (
			$1, $2, $3, 'direktur'::user_role, 'none'::engineer_type, 'none'::client_type, $4, 'aktif'
		)
	`
	_, err = db.Pool.Exec(ctx, query,
		"Direktur Utama PT. Melayani Digital Raya",
		"direktur@meldir.id",
		string(hashedPassword),
		"+628213173357",
	)
	if err != nil {
		return fmt.Errorf("gagal membuat akun direktur awal: %w", err)
	}

	log.Printf("👑 Berhasil auto-seed akun Direktur Utama awal:")
	log.Printf("   Email   : direktur@meldir.id")
	log.Printf("   Password: %s", defaultPassword)
	log.Printf("   Status  : Aktif (Superadmin)")
	return nil
}

func (db *PostgresDB) Close() {
	if db.Pool != nil {
		db.Pool.Close()
	}
}
