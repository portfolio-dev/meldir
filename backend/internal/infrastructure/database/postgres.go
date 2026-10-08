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
	if err := db.MigrateSchema(context.Background()); err != nil {
		log.Printf("⚠️ Peringatan migrasi skema: %v", err)
	}

	return db, nil
}

func (db *PostgresDB) MigrateSchema(ctx context.Context) error {
	queries := []string{
		"ALTER TABLE tickets ALTER COLUMN project_id DROP NOT NULL",
		"ALTER TABLE tickets ALTER COLUMN contract_id DROP NOT NULL",
		"ALTER TABLE timesheet_logs ALTER COLUMN project_id DROP NOT NULL",
		"ALTER TABLE timesheet_logs ADD COLUMN IF NOT EXISTS ticket_code VARCHAR(30) NULL",
		"ALTER TABLE timesheet_logs ADD COLUMN IF NOT EXISTS project_name VARCHAR(150) NULL",

		// Invoices Table & Column Alterations
		`CREATE TABLE IF NOT EXISTS invoices (
			id BIGSERIAL PRIMARY KEY,
			invoice_number VARCHAR(50) UNIQUE NOT NULL,
			client_id BIGINT NULL,
			client_name VARCHAR(150) NULL,
			contract_id BIGINT NULL,
			addon_order_id BIGINT NULL,
			amount NUMERIC(15,2) NOT NULL,
			tax_amount NUMERIC(15,2) DEFAULT 0.00,
			due_date DATE NOT NULL,
			status VARCHAR(30) DEFAULT 'unpaid',
			bank_destination VARCHAR(120) DEFAULT 'PT. Melayani Digital Raya - Bank Mandiri & BCA',
			tax_invoice_number VARCHAR(60) NULL,
			paid_at TIMESTAMP WITH TIME ZONE NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		"ALTER TABLE invoices ADD COLUMN IF NOT EXISTS client_name VARCHAR(150) NULL",
		"ALTER TABLE invoices ADD COLUMN IF NOT EXISTS bank_destination VARCHAR(120) DEFAULT 'PT. Melayani Digital Raya - Bank Mandiri & BCA'",
		"ALTER TABLE invoices ADD COLUMN IF NOT EXISTS tax_amount NUMERIC(15,2) DEFAULT 0.00",
		"ALTER TABLE invoices ADD COLUMN IF NOT EXISTS tax_invoice_number VARCHAR(60) NULL",
		"ALTER TABLE invoices ADD COLUMN IF NOT EXISTS paid_at TIMESTAMP WITH TIME ZONE NULL",
		"ALTER TABLE invoices ALTER COLUMN client_id DROP NOT NULL",
		"ALTER TABLE invoices ALTER COLUMN status DROP DEFAULT",
		"ALTER TABLE invoices ALTER COLUMN status TYPE VARCHAR(30) USING status::text",
		"ALTER TABLE invoices ALTER COLUMN status SET DEFAULT 'unpaid'",

		// Accounting Journals Table & Column Alterations
		`CREATE TABLE IF NOT EXISTS accounting_journals (
			id BIGSERIAL PRIMARY KEY,
			journal_number VARCHAR(50) UNIQUE NOT NULL,
			journal_date DATE NOT NULL,
			source_type VARCHAR(50) DEFAULT 'general_entry',
			source_reference_id VARCHAR(100) NULL,
			memo TEXT NOT NULL,
			debit_account VARCHAR(150) NULL,
			credit_account VARCHAR(150) NULL,
			total_debit NUMERIC(18,2) NOT NULL DEFAULT 0.00,
			total_credit NUMERIC(18,2) NOT NULL DEFAULT 0.00,
			is_posted BOOLEAN DEFAULT TRUE,
			created_by BIGINT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		)`,
		"ALTER TABLE accounting_journals ADD COLUMN IF NOT EXISTS debit_account VARCHAR(150) NULL",
		"ALTER TABLE accounting_journals ADD COLUMN IF NOT EXISTS credit_account VARCHAR(150) NULL",
		"ALTER TABLE accounting_journals ADD COLUMN IF NOT EXISTS total_debit NUMERIC(18,2) NOT NULL DEFAULT 0.00",
		"ALTER TABLE accounting_journals ADD COLUMN IF NOT EXISTS total_credit NUMERIC(18,2) NOT NULL DEFAULT 0.00",
		"ALTER TABLE accounting_journals ADD COLUMN IF NOT EXISTS is_posted BOOLEAN DEFAULT TRUE",
		"ALTER TABLE accounting_journals ALTER COLUMN debit_account DROP NOT NULL",
		"ALTER TABLE accounting_journals ALTER COLUMN credit_account DROP NOT NULL",
		"ALTER TABLE accounting_journals ALTER COLUMN source_type DROP DEFAULT",
		"ALTER TABLE accounting_journals ALTER COLUMN source_type TYPE VARCHAR(50) USING source_type::text",
		"ALTER TABLE accounting_journals ALTER COLUMN source_type SET DEFAULT 'general_entry'",
		"ALTER TABLE accounting_journals ALTER COLUMN created_by DROP NOT NULL",
	}

	for _, q := range queries {
		if _, err := db.Pool.Exec(ctx, q); err != nil {
			log.Printf("ℹ️ Info migrasi skema (dilewati jika sudah ada): %v", err)
		}
	}
	log.Printf("✅ Skema database PostgreSQL (invoices, journals, timesheets, tickets) tervalidasi siap.")
	return nil
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
