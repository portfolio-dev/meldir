-- ==============================================================================
-- DATABASE SCHEMA SPECIFICATION FOR MELDIR.ID (PT. MELAYANI DIGITAL RAYA)
-- Engine: PostgreSQL 14+
-- File: docs/05_POSTGRESQL_SCHEMA.sql
-- Description: Complete production DDL script covering 12 operational modules,
--              Profile Management, Media Attachments (Image/Video/PDF),
--              In-App Notification Bell, and Manual Broadcast Push Engine.
-- ==============================================================================

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ==============================================================================
-- 1. ENUM TYPES CREATION
-- ==============================================================================
CREATE TYPE user_role AS ENUM ('direktur', 'admin', 'engineer', 'klien', 'audit');
CREATE TYPE engineer_type AS ENUM ('internal', 'external', 'none');
CREATE TYPE client_type AS ENUM ('individual', 'company', 'foundation', 'none');
CREATE TYPE contract_status AS ENUM (
    'draft', 
    'waiting_signature', 
    'waiting_ematerai', 
    'active', 
    'retention_offering', 
    'terminated', 
    'expired'
);
CREATE TYPE service_category AS ENUM (
    'custom_dev', 
    'managed_care', 
    'scaling', 
    'mobile_pwa', 
    'free_audit_social'
);
CREATE TYPE ticket_priority AS ENUM ('p1_critical', 'p2_medium', 'p3_low');
CREATE TYPE ticket_status AS ENUM ('open', 'in_progress', 'resolved', 'closed');
CREATE TYPE invoice_status AS ENUM ('unpaid', 'verifying', 'paid', 'overdue', 'cancelled');
CREATE TYPE offboard_status AS ENUM ('submitted', 'counter_offered', 'approved', 'rejected');
CREATE TYPE server_health_status AS ENUM ('online', 'degraded', 'critical_down');
CREATE TYPE media_file_type AS ENUM ('image', 'video', 'pdf', 'archive', 'other');
CREATE TYPE broadcast_target AS ENUM ('all_users', 'clients_only', 'engineers_only', 'admins_only');

-- ==============================================================================
-- 2. USERS & PROFILE MANAGEMENT TABLES
-- ==============================================================================
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    email VARCHAR(150) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    role user_role NOT NULL DEFAULT 'klien',
    engineer_type engineer_type DEFAULT 'none',
    client_type client_type DEFAULT 'none',
    phone_wa VARCHAR(25) NOT NULL,
    avatar_url VARCHAR(255) NULL, -- Foto Profil Akun
    github_username VARCHAR(100) NULL,
    status VARCHAR(30) NOT NULL DEFAULT 'pengajuan', -- 'pengajuan', 'aktif', 'suspended'
    last_login_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_status ON users(status);

CREATE TABLE clients_metadata (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    company_name VARCHAR(150) NULL,
    tax_id_npwp VARCHAR(50) NULL,
    address TEXT NULL,
    foundation_decree_no VARCHAR(100) NULL, -- For non-profit foundations
    pic_name VARCHAR(100) NULL,
    pic_phone VARCHAR(25) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 3. IN-APP NOTIFICATIONS & MANUAL BROADCAST ENGINE
-- ==============================================================================
CREATE TABLE user_notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(150) NOT NULL,
    message TEXT NOT NULL,
    target_url VARCHAR(255) NOT NULL, -- Deep link (e.g. '/tickets/TCK-102')
    icon_type VARCHAR(50) DEFAULT 'bell', -- 'ticket', 'invoice', 'server', 'contract'
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_user_notifications_user_read ON user_notifications(user_id, is_read);

CREATE TABLE manual_broadcast_notifications (
    id BIGSERIAL PRIMARY KEY,
    sender_id BIGINT NOT NULL REFERENCES users(id),
    title VARCHAR(150) NOT NULL,
    body TEXT NOT NULL,
    target_audience broadcast_target NOT NULL DEFAULT 'all_users',
    target_url VARCHAR(255) DEFAULT '/',
    sent_count INT DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 4. CONTRACTS & PROJECTS TABLES
-- ==============================================================================
CREATE TABLE contracts (
    id BIGSERIAL PRIMARY KEY,
    contract_number VARCHAR(50) UNIQUE NOT NULL,
    party_a_id BIGINT NOT NULL REFERENCES users(id), -- PT. Melayani Digital Raya (Direktur)
    party_b_id BIGINT NOT NULL REFERENCES users(id), -- Klien atau Engineer
    service_category service_category NOT NULL,
    contract_terms TEXT NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NULL,
    monthly_fee NUMERIC(15,2) DEFAULT 0.00,
    project_fee NUMERIC(15,2) DEFAULT 0.00,
    status contract_status DEFAULT 'draft',
    canvas_signature_path VARCHAR(255) NULL,
    initial_pdf_path VARCHAR(255) NULL,
    ematerai_pdf_path VARCHAR(255) NULL,
    signed_at TIMESTAMP WITH TIME ZONE NULL,
    ematerai_uploaded_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_contracts_status ON contracts(status);
CREATE INDEX idx_contracts_party_b ON contracts(party_b_id);

CREATE TABLE projects (
    id BIGSERIAL PRIMARY KEY,
    contract_id BIGINT NOT NULL REFERENCES contracts(id) ON DELETE CASCADE,
    client_id BIGINT NOT NULL REFERENCES users(id),
    project_name VARCHAR(150) NOT NULL,
    domain_url VARCHAR(200) NULL,
    repository_url VARCHAR(255) NULL,
    status VARCHAR(30) DEFAULT 'in_progress', -- 'in_progress', 'maintenance', 'completed'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Lampiran Media Proyek (Image, Video, PDF, Docs)
CREATE TABLE project_attachments (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    uploader_id BIGINT NOT NULL REFERENCES users(id),
    file_name VARCHAR(200) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    file_type media_file_type NOT NULL, -- 'image', 'video', 'pdf', 'archive'
    file_size_kb INT NOT NULL,
    notes TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 5. EXTERNAL SERVER HEALTH MONITORING TABLE
-- ==============================================================================
CREATE TABLE external_servers (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    client_id BIGINT NOT NULL REFERENCES users(id),
    server_name VARCHAR(100) NOT NULL,
    ip_or_domain VARCHAR(200) NOT NULL,
    health_endpoint VARCHAR(200) DEFAULT '/health',
    check_interval_minutes INT DEFAULT 5,
    current_status server_health_status DEFAULT 'online',
    last_ping_at TIMESTAMP WITH TIME ZONE NULL,
    last_response_time_ms INT NULL,
    last_error_message TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_external_servers_status ON external_servers(current_status);

-- ==============================================================================
-- 6. SLA TICKETS & MULTI-MEDIA REPLIES TABLES
-- ==============================================================================
CREATE TABLE tickets (
    id BIGSERIAL PRIMARY KEY,
    ticket_code VARCHAR(30) UNIQUE NOT NULL,
    project_id BIGINT NOT NULL REFERENCES projects(id),
    client_id BIGINT NOT NULL REFERENCES users(id),
    contract_id BIGINT NOT NULL REFERENCES contracts(id),
    title VARCHAR(200) NOT NULL,
    description TEXT NOT NULL,
    priority ticket_priority NOT NULL DEFAULT 'p3_low',
    status ticket_status DEFAULT 'open',
    assigned_engineer_id BIGINT NULL REFERENCES users(id),
    sla_deadline TIMESTAMP WITH TIME ZONE NOT NULL,
    is_sla_breached BOOLEAN DEFAULT FALSE,
    resolved_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_tickets_status ON tickets(status);
CREATE INDEX idx_tickets_priority ON tickets(priority);

CREATE TABLE ticket_replies (
    id BIGSERIAL PRIMARY KEY,
    ticket_id BIGINT NOT NULL REFERENCES tickets(id) ON DELETE CASCADE,
    sender_id BIGINT NOT NULL REFERENCES users(id),
    message TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabel Lampiran Berkas Tiket (Mendukung Multi-Upload Gambar, Video, PDF)
CREATE TABLE ticket_reply_attachments (
    id BIGSERIAL PRIMARY KEY,
    reply_id BIGINT NOT NULL REFERENCES ticket_replies(id) ON DELETE CASCADE,
    file_name VARCHAR(200) NOT NULL,
    file_path VARCHAR(255) NOT NULL,
    file_type media_file_type NOT NULL, -- 'image', 'video', 'pdf'
    file_size_kb INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 7. INVOICES & CORPORATE FINANCE TABLES
-- ==============================================================================
CREATE TABLE invoices (
    id BIGSERIAL PRIMARY KEY,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    client_id BIGINT NOT NULL REFERENCES users(id),
    contract_id BIGINT NULL REFERENCES contracts(id),
    amount NUMERIC(15,2) NOT NULL,
    due_date DATE NOT NULL,
    status invoice_status DEFAULT 'unpaid',
    payment_proof_path VARCHAR(255) NULL, -- Bukti transfer gambar/PDF
    bank_destination VARCHAR(120) DEFAULT 'PT. Melayani Digital Raya - Bank Mandiri',
    paid_at TIMESTAMP WITH TIME ZONE NULL,
    verified_by BIGINT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_invoices_status ON invoices(status);

CREATE TABLE pt_financial_ledger (
    id BIGSERIAL PRIMARY KEY,
    transaction_type VARCHAR(20) NOT NULL, -- 'income', 'expense'
    category VARCHAR(80) NOT NULL, -- 'client_invoice', 'engineer_payout', 'hosting_cost', 'office_overhead'
    invoice_id BIGINT NULL REFERENCES invoices(id),
    amount NUMERIC(15,2) NOT NULL,
    description TEXT NOT NULL,
    transaction_date DATE NOT NULL,
    created_by BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE engineer_payouts (
    id BIGSERIAL PRIMARY KEY,
    engineer_id BIGINT NOT NULL REFERENCES users(id),
    contract_id BIGINT NULL REFERENCES contracts(id),
    amount NUMERIC(15,2) DEFAULT 0.00, -- Supports Rp 0 for pro-bono foundation tasks
    description TEXT NOT NULL,
    payout_date DATE NOT NULL,
    created_by BIGINT NOT NULL REFERENCES users(id), -- Admin who input payout
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 8. VAULT CREDENTIALS & GITHUB ASSIGNMENTS
-- ==============================================================================
CREATE TABLE credentials_vault (
    id BIGSERIAL PRIMARY KEY,
    client_id BIGINT NOT NULL REFERENCES users(id),
    service_name VARCHAR(120) NOT NULL, -- e.g. "Production Server Hostinger"
    host_address VARCHAR(150) NOT NULL,
    username VARCHAR(100) NOT NULL,
    encrypted_password TEXT NOT NULL, -- AES-256 Encrypted
    access_notes TEXT NULL,
    created_by BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE github_assignments (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    engineer_id BIGINT NOT NULL REFERENCES users(id),
    repository_name VARCHAR(150) NOT NULL,
    permission_level VARCHAR(20) DEFAULT 'push', -- 'pull', 'push', 'admin'
    granted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    revoked_at TIMESTAMP WITH TIME ZONE NULL
);

-- ==============================================================================
-- 9. OFFBOARDING REQUESTS & SECURITY AUDIT LOGS
-- ==============================================================================
CREATE TABLE offboarding_requests (
    id BIGSERIAL PRIMARY KEY,
    client_id BIGINT NOT NULL REFERENCES users(id),
    contract_id BIGINT NOT NULL REFERENCES contracts(id),
    reason TEXT NOT NULL,
    counter_offer_text TEXT NULL,
    counter_offer_discount NUMERIC(15,2) DEFAULT 0.00,
    status offboard_status DEFAULT 'submitted',
    verified_email VARCHAR(150) NULL,
    verified_phone VARCHAR(25) NULL,
    generated_credential_pdf VARCHAR(255) NULL,
    github_transferred_at TIMESTAMP WITH TIME ZONE NULL,
    completed_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE security_audit_logs (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NULL REFERENCES users(id),
    portal_origin VARCHAR(30) NOT NULL, -- 'office', 'jobs', 'portal', 'public'
    action_event VARCHAR(150) NOT NULL, -- e.g. "VIEW_CREDENTIAL", "FORCE_LOGOUT", "SEND_BROADCAST_PUSH"
    ip_address VARCHAR(45) NOT NULL,
    user_agent TEXT NOT NULL,
    payload_hash VARCHAR(64) NOT NULL, -- SHA-256 Hash
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_security_audit_logs_event ON security_audit_logs(action_event);
CREATE INDEX idx_security_audit_logs_ip ON security_audit_logs(ip_address);

-- ==============================================================================
-- END OF SCHEMA DDL SCRIPT
-- ==============================================================================
