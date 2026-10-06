-- ==============================================================================
-- DATABASE SCHEMA SPECIFICATION FOR MELDIR.ID (PT. MELAYANI DIGITAL RAYA)
-- Engine: PostgreSQL 14+
-- File: docs/05_POSTGRESQL_SCHEMA.sql
-- Description: Complete enterprise production DDL script covering:
--              1. Users & Profile Management (Argon2id, WebP Avatars)
--              2. Inbound Leads CRM Pipeline
--              3. Contracts, Projects, Media Attachments (Image/Video/PDF)
--              4. Add-On Packages & Metered Hours Balances
--              5. Minor Feature Monthly Quota Tracking (Priority Care 2x/mo)
--              6. Engineer Timesheet & Work Logs
--              7. SLA Tickets & Multi-Media Discussions
--              8. Invoices & Billing Engine
--              9. Corporate Accounting (Chart of Accounts, Double-Entry Journals, General Ledger)
--              10. Corporate Operational Expenses & Petty Cash (Digital Receipts)
--              11. Taxation Engine (PPN 11% Keluaran/Masukan, PPh 21, PPh 23, SPT 1771 PPh Badan)
--              12. Digital BAST (Berita Acara Serah Terima) Engine
--              13. External Server Health Probing (5-Min Cron)
--              14. Monthly Maintenance Executive Reporting
--              15. Credentials Vault (AES-256) & GitHub Collaborators
--              16. Client Offboarding (Retention Flow) & Security Audit Logs
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

-- Business & Operations ENUMs
CREATE TYPE lead_status AS ENUM ('new', 'contacted', 'quoted', 'won_contract', 'lost');
CREATE TYPE addon_package_type AS ENUM ('hourly_rate', 'sprint_10h', 'sprint_20h', 'modular_feature');
CREATE TYPE bast_status AS ENUM ('draft', 'waiting_client_signature', 'signed', 'completed');

-- Corporate Accounting & Taxation Enterprise ENUMs
CREATE TYPE account_category AS ENUM (
    'asset', 
    'liability', 
    'equity', 
    'revenue', 
    'cost_of_sales', 
    'expense'
);
CREATE TYPE normal_balance_type AS ENUM ('debit', 'credit');
CREATE TYPE journal_source_type AS ENUM (
    'invoice_payment', 
    'engineer_payout', 
    'corporate_expense', 
    'tax_settlement', 
    'manual_adjustment', 
    'opening_balance'
);
CREATE TYPE expense_status AS ENUM ('draft', 'pending_approval', 'approved', 'paid', 'rejected');
CREATE TYPE tax_article_type AS ENUM (
    'ppn_keluaran', 
    'ppn_masukan', 
    'pph_21', 
    'pph_23', 
    'pph_4_ayat_2', 
    'pph_25', 
    'pph_badan_1771'
);
CREATE TYPE tax_filing_status AS ENUM ('unfiled', 'drafted', 'paid_ntpn', 'reported_djp');
CREATE TYPE withholding_direction AS ENUM ('dipotong_klien', 'memotong_vendor', 'memotong_engineer');

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
    avatar_url VARCHAR(255) NULL, -- Foto Profil Akun (WebP)
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
    foundation_decree_no VARCHAR(100) NULL, -- Akta/SK Kemenkum untuk yayasan sosial
    pic_name VARCHAR(100) NULL,
    pic_phone VARCHAR(25) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- ==============================================================================
-- 3. INBOUND LEADS & CRM PIPELINE TABLE
-- ==============================================================================
CREATE TABLE inbound_leads (
    id BIGSERIAL PRIMARY KEY,
    lead_code VARCHAR(30) UNIQUE NOT NULL, -- e.g. "LEAD-2026-001"
    name VARCHAR(120) NOT NULL,
    email VARCHAR(150) NULL,
    phone_wa VARCHAR(25) NOT NULL,
    company_name VARCHAR(150) NULL,
    service_interest service_category NOT NULL DEFAULT 'managed_care',
    budget_range VARCHAR(80) NULL,
    message TEXT NULL,
    status lead_status DEFAULT 'new',
    source VARCHAR(50) DEFAULT 'website_landing', -- 'website_landing', 'audit_form', 'whatsapp_direct'
    assigned_admin_id BIGINT NULL REFERENCES users(id),
    notes TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_inbound_leads_status ON inbound_leads(status);
CREATE INDEX idx_inbound_leads_created ON inbound_leads(created_at);

-- ==============================================================================
-- 4. IN-APP NOTIFICATIONS & MANUAL BROADCAST ENGINE
-- ==============================================================================
CREATE TABLE user_notifications (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title VARCHAR(150) NOT NULL,
    message TEXT NOT NULL,
    target_url VARCHAR(255) NOT NULL, -- Deep link (e.g. '/tickets/TCK-102')
    icon_type VARCHAR(50) DEFAULT 'bell', -- 'ticket', 'invoice', 'server', 'contract', 'addon', 'report'
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
-- 5. CONTRACTS & PROJECTS TABLES
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
-- 6. MINOR FEATURE MONTHLY QUOTA TRACKER (Priority Care 2x/Bulan)
-- ==============================================================================
CREATE TABLE contract_monthly_quotas (
    id BIGSERIAL PRIMARY KEY,
    contract_id BIGINT NOT NULL REFERENCES contracts(id) ON DELETE CASCADE,
    period_month INT NOT NULL, -- 1 - 12
    period_year INT NOT NULL,  -- e.g. 2026
    minor_feature_quota_total INT DEFAULT 2, -- Default 2x per bulan untuk Priority Care
    minor_feature_quota_used INT DEFAULT 0,
    notes TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(contract_id, period_month, period_year)
);

CREATE INDEX idx_monthly_quotas_period ON contract_monthly_quotas(contract_id, period_month, period_year);

-- ==============================================================================
-- 7. ADD-ON PACKAGES & CLIENT METERED HOURS ORDERS
-- ==============================================================================
CREATE TABLE addon_packages (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(120) NOT NULL, -- e.g. "Sprint Block 10 Jam", "On-Demand Extension"
    package_type addon_package_type NOT NULL,
    allocated_hours NUMERIC(6,2) DEFAULT 0.00,
    price NUMERIC(15,2) NOT NULL,
    active_days INT DEFAULT 60, -- Masa aktif saldo jam (e.g. 60 hari kalender)
    description TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE client_addon_orders (
    id BIGSERIAL PRIMARY KEY,
    order_code VARCHAR(40) UNIQUE NOT NULL, -- e.g. "ADDON-2026-001"
    client_id BIGINT NOT NULL REFERENCES users(id),
    project_id BIGINT NOT NULL REFERENCES projects(id),
    package_id BIGINT NOT NULL REFERENCES addon_packages(id),
    invoice_id BIGINT NULL, -- Forward-referenced to invoices table
    total_hours NUMERIC(6,2) NOT NULL,
    remaining_hours NUMERIC(6,2) NOT NULL,
    price NUMERIC(15,2) NOT NULL,
    expires_at DATE NULL,
    status VARCHAR(30) DEFAULT 'unpaid', -- 'unpaid', 'active', 'exhausted', 'expired'
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_addon_orders_client ON client_addon_orders(client_id, status);

-- ==============================================================================
-- 8. EXTERNAL SERVER HEALTH MONITORING TABLE (5-Min Probe)
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
-- 9. SLA TICKETS & MULTI-MEDIA REPLIES TABLES
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
    is_minor_feature BOOLEAN DEFAULT FALSE, -- Flag jika permintaan adalah kuota fitur ringan
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
-- 10. ENGINEER TIMESHEET & WORK LOGS TABLE
-- ==============================================================================
CREATE TABLE timesheet_logs (
    id BIGSERIAL PRIMARY KEY,
    engineer_id BIGINT NOT NULL REFERENCES users(id),
    project_id BIGINT NOT NULL REFERENCES projects(id),
    ticket_id BIGINT NULL REFERENCES tickets(id),
    addon_order_id BIGINT NULL REFERENCES client_addon_orders(id),
    hours_spent NUMERIC(5,2) NOT NULL,
    work_description TEXT NOT NULL,
    log_date DATE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_timesheet_engineer ON timesheet_logs(engineer_id, log_date);
CREATE INDEX idx_timesheet_project ON timesheet_logs(project_id);

-- ==============================================================================
-- 11. INVOICES & CORPORATE FINANCE TABLES (e-Faktur PPN 11%)
-- ==============================================================================
CREATE TABLE invoices (
    id BIGSERIAL PRIMARY KEY,
    invoice_number VARCHAR(50) UNIQUE NOT NULL,
    client_id BIGINT NOT NULL REFERENCES users(id),
    contract_id BIGINT NULL REFERENCES contracts(id),
    addon_order_id BIGINT NULL REFERENCES client_addon_orders(id),
    amount NUMERIC(15,2) NOT NULL,
    tax_amount NUMERIC(15,2) DEFAULT 0.00, -- Nilai PPN (11%)
    due_date DATE NOT NULL,
    status invoice_status DEFAULT 'unpaid',
    payment_proof_path VARCHAR(255) NULL, -- Bukti transfer gambar/PDF
    bank_destination VARCHAR(120) DEFAULT 'PT. Melayani Digital Raya - Bank Mandiri',
    tax_invoice_number VARCHAR(60) NULL, -- Nomor Seri e-Faktur Pajak resmi DJP
    tax_invoice_pdf_path VARCHAR(255) NULL, -- Berkas PDF e-Faktur resmi untuk diunduh klien
    is_tax_invoice_issued BOOLEAN DEFAULT FALSE,
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
    gross_amount NUMERIC(15,2) DEFAULT 0.00, -- Nilai bruto honorarium (bisa Rp 0 untuk tugas sosial/pro-bono)
    pph21_rate NUMERIC(5,2) DEFAULT 0.00,    -- Tarif PPh 21 (TER / Tarif Pasal 17)
    pph21_amount NUMERIC(15,2) DEFAULT 0.00, -- Nilai potongan PPh 21 yang disetor ke kas negara
    net_payout NUMERIC(15,2) DEFAULT 0.00,   -- Nilai bersih yang ditransfer ke rekening engineer
    description TEXT NOT NULL,
    payout_date DATE NOT NULL,
    payment_proof_path VARCHAR(255) NULL,   -- Bukti transfer honorarium (WebP/PDF)
    created_by BIGINT NOT NULL REFERENCES users(id), -- Admin pencatat payout
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_engineer_payouts_engineer ON engineer_payouts(engineer_id, payout_date);

-- ==============================================================================
-- 12. SISTEM AKUNTANSI DIGITAL & PERPAJAKAN OPERASIONAL PT (ENTERPRISE COA & TAX)
-- ==============================================================================

-- A. Bagan Akun Standar (Chart of Accounts / COA 5-Digit Standar Indonesia SAK EMKM)
CREATE TABLE accounting_chart_of_accounts (
    id BIGSERIAL PRIMARY KEY,
    account_code VARCHAR(20) UNIQUE NOT NULL, -- e.g. '11010', '11020', '21030', '41010'
    account_name VARCHAR(150) NOT NULL,       -- e.g. 'Kas Operasional / Petty Cash'
    category account_category NOT NULL,       -- 'asset', 'liability', 'equity', 'revenue', 'cost_of_sales', 'expense'
    normal_balance normal_balance_type NOT NULL, -- 'debit', 'credit'
    parent_code VARCHAR(20) NULL,             -- e.g. '11000' untuk akun induk
    is_active BOOLEAN DEFAULT TRUE,
    description TEXT NULL,
    current_balance NUMERIC(18,2) DEFAULT 0.00,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_coa_category ON accounting_chart_of_accounts(category);
CREATE INDEX idx_coa_code ON accounting_chart_of_accounts(account_code);

-- B. Jurnal Umum Berpasangan (Double-Entry General Journals)
CREATE TABLE accounting_journals (
    id BIGSERIAL PRIMARY KEY,
    journal_number VARCHAR(50) UNIQUE NOT NULL, -- e.g. 'JRN/2026/10/0001'
    journal_date DATE NOT NULL,
    source_type journal_source_type NOT NULL,  -- 'invoice_payment', 'engineer_payout', 'corporate_expense', etc.
    source_reference_id VARCHAR(100) NULL,      -- ID atau Nomor Invoice / Payout / Expense
    memo TEXT NOT NULL,
    total_debit NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    total_credit NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    is_posted BOOLEAN DEFAULT TRUE,
    posted_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    created_by BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_journals_date ON accounting_journals(journal_date);
CREATE INDEX idx_journals_source ON accounting_journals(source_type, source_reference_id);

-- C. Rincian Item Jurnal (Journal Items / Debit-Credit Rows)
CREATE TABLE accounting_journal_items (
    id BIGSERIAL PRIMARY KEY,
    journal_id BIGINT NOT NULL REFERENCES accounting_journals(id) ON DELETE CASCADE,
    account_id BIGINT NOT NULL REFERENCES accounting_chart_of_accounts(id),
    debit_amount NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    credit_amount NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    description TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_journal_items_journal ON accounting_journal_items(journal_id);
CREATE INDEX idx_journal_items_account ON accounting_journal_items(account_id);

-- D. Beban Operasional & Kas Kecil Digital (Corporate Operational Expenses & Petty Cash)
CREATE TABLE corporate_expenses (
    id BIGSERIAL PRIMARY KEY,
    expense_number VARCHAR(50) UNIQUE NOT NULL, -- e.g. 'EXP/2026/10/0001'
    expense_date DATE NOT NULL,
    account_id BIGINT NOT NULL REFERENCES accounting_chart_of_accounts(id), -- Akun Beban (6xxxx atau 5xxxx)
    paid_from_account_id BIGINT NOT NULL REFERENCES accounting_chart_of_accounts(id), -- Akun Kas/Bank (11010 / 11020)
    vendor_name VARCHAR(150) NOT NULL,
    vendor_npwp VARCHAR(50) NULL,
    vendor_invoice_number VARCHAR(100) NULL,
    gross_amount NUMERIC(15,2) NOT NULL,
    vat_amount NUMERIC(15,2) DEFAULT 0.00, -- PPN Masukan jika vendor menerbitkan Faktur Pajak
    tax_withholding_amount NUMERIC(15,2) DEFAULT 0.00, -- Nilai PPh 21/23 yang dipotong dari vendor
    net_amount_paid NUMERIC(15,2) NOT NULL,
    is_fiscal_deductible BOOLEAN DEFAULT TRUE, -- Biaya dapat dikurangkan secara fiskal (Deductible Expense)
    receipt_file_path VARCHAR(255) NULL, -- Berkas kuitansi/struk/nota digital (WebP/PDF)
    status expense_status DEFAULT 'approved',
    approved_by BIGINT NULL REFERENCES users(id),
    created_by BIGINT NOT NULL REFERENCES users(id),
    notes TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_expenses_date ON corporate_expenses(expense_date);
CREATE INDEX idx_expenses_status ON corporate_expenses(status);

-- E. Buku Catatan Pajak PPN 11% (Keluaran & Masukan - SPT Masa PPN 1111)
CREATE TABLE tax_records_vat (
    id BIGSERIAL PRIMARY KEY,
    tax_period_month INT NOT NULL, -- 1 - 12
    tax_period_year INT NOT NULL,  -- e.g. 2026
    vat_type VARCHAR(20) NOT NULL, -- 'keluaran', 'masukan'
    invoice_id BIGINT NULL REFERENCES invoices(id), -- Terhubung jika PPN Keluaran
    expense_id BIGINT NULL REFERENCES corporate_expenses(id), -- Terhubung jika PPN Masukan
    counterpart_name VARCHAR(150) NOT NULL, -- Nama Klien atau Vendor
    counterpart_npwp VARCHAR(50) NOT NULL,
    tax_invoice_nsfp VARCHAR(60) NOT NULL,  -- Nomor Seri Faktur Pajak resmi DJP (e.g. 010.002-26.XXXXXXXX)
    tax_invoice_date DATE NOT NULL,
    dpp_amount NUMERIC(15,2) NOT NULL,      -- Dasar Pengenaan Pajak
    vat_rate NUMERIC(4,2) DEFAULT 11.00,    -- 11.00%
    vat_amount NUMERIC(15,2) NOT NULL,      -- Nilai PPN (11%)
    is_creditable BOOLEAN DEFAULT TRUE,     -- Masukan dapat dikreditkan
    filing_status tax_filing_status DEFAULT 'unfiled',
    ntpn_settlement_code VARCHAR(50) NULL,  -- Kode NTPN jika PPN Kurang Bayar disetor ke kas negara
    ebupot_qr_url VARCHAR(255) NULL,
    tax_pdf_path VARCHAR(255) NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_vat_period ON tax_records_vat(tax_period_year, tax_period_month, vat_type);

-- F. Bukti Pemotongan Pajak PPh (PPh 21 Engineer, PPh 23 Klien, PPh 23 Vendor)
CREATE TABLE tax_withholding_slips (
    id BIGSERIAL PRIMARY KEY,
    slip_number VARCHAR(60) UNIQUE NOT NULL, -- e.g. 'BUPOT-21/2026/10/001' atau 'BUPOT-23/2026/10/001'
    tax_article tax_article_type NOT NULL,   -- 'pph_21', 'pph_23', 'pph_4_ayat_2'
    direction withholding_direction NOT NULL, -- 'memotong_engineer', 'dipotong_klien', 'memotong_vendor'
    counterpart_name VARCHAR(150) NOT NULL,
    counterpart_npwp VARCHAR(50) NULL,
    counterpart_nik VARCHAR(50) NULL,
    tax_period_month INT NOT NULL,
    tax_period_year INT NOT NULL,
    gross_amount NUMERIC(15,2) NOT NULL,
    tax_base_dpp NUMERIC(15,2) NOT NULL,
    effective_tax_rate NUMERIC(5,2) NOT NULL, -- e.g. 5.00%, 2.00%
    tax_withheld_amount NUMERIC(15,2) NOT NULL,
    payout_id BIGINT NULL REFERENCES engineer_payouts(id),
    invoice_id BIGINT NULL REFERENCES invoices(id),
    expense_id BIGINT NULL REFERENCES corporate_expenses(id),
    ntpn_payment_code VARCHAR(50) NULL,      -- Bukti setor NTPN ke kas negara
    ebupot_pdf_path VARCHAR(255) NULL,        -- Salinan berkas bukti potong e-Bupot unifikasi DJP
    filing_status tax_filing_status DEFAULT 'unfiled',
    created_by BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_withholding_period ON tax_withholding_slips(tax_period_year, tax_period_month, tax_article);
CREATE INDEX idx_withholding_direction ON tax_withholding_slips(direction);

-- G. Estimasi & Rekonsiliasi SPT Tahunan PPh Badan 1771 & Angsuran PPh 25
CREATE TABLE corporate_annual_tax_estimates (
    id BIGSERIAL PRIMARY KEY,
    tax_year INT UNIQUE NOT NULL,
    gross_revenue NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    cost_of_services_hpp NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    gross_profit NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    operating_expenses_opex NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    commercial_net_income NUMERIC(18,2) NOT NULL DEFAULT 0.00,
    positive_fiscal_adjustments NUMERIC(18,2) DEFAULT 0.00, -- Beban non-deductible (sumbangan, representasi tanpa nominatif)
    negative_fiscal_adjustments NUMERIC(18,2) DEFAULT 0.00,
    fiscal_net_income_pkp NUMERIC(18,2) NOT NULL DEFAULT 0.00, -- Penghasilan Kena Pajak (PKP)
    tax_facility_applied VARCHAR(80) DEFAULT 'pasal_31e_uu_hpp_50_pct', -- Fasilitas Pasal 31E UU HPP diskon 50% tarif (11%)
    corporate_tax_payable NUMERIC(18,2) NOT NULL DEFAULT 0.00, -- Total PPh Terutang
    tax_credits_pph23 NUMERIC(18,2) DEFAULT 0.00,             -- Kredit Pajak PPh 23 bukti potong dari klien
    tax_credits_pph25 NUMERIC(18,2) DEFAULT 0.00,             -- Setoran angsuran bulanan PPh 25
    pph_article_29_underpaid NUMERIC(18,2) NOT NULL DEFAULT 0.00, -- PPh Kurang Bayar (Pasal 29)
    next_year_monthly_pph25 NUMERIC(18,2) DEFAULT 0.00,      -- Kewajiban angsuran PPh 25 bulanan tahun berikutnya
    status tax_filing_status DEFAULT 'unfiled',
    spt_1771_pdf_path VARCHAR(255) NULL,
    ntpn_settlement_code VARCHAR(50) NULL,
    calculated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- H. Data Awal Standar Bagan Akun (Standard Chart of Accounts Seeding SAK EMKM)
INSERT INTO accounting_chart_of_accounts (account_code, account_name, category, normal_balance, parent_code, description) VALUES
-- 10000 ASET (ASSETS)
('11010', 'Kas Operasional / Petty Cash', 'asset', 'debit', '11000', 'Kas kecil operasional kantor harian'),
('11020', 'Bank Mandiri Operasional PT (Rekening Utama)', 'asset', 'debit', '11000', 'Rekening penerimaan invoice & pengeluaran utama'),
('11030', 'Bank BCA Operasional PT (Rekening Cadangan)', 'asset', 'debit', '11000', 'Rekening operasional sekunder'),
('11040', 'Piutang Usaha Klien (Accounts Receivable)', 'asset', 'debit', '11000', 'Tagihan invoice yang belum dilunasi klien'),
('11050', 'Uang Muka Pajak PPh 23 (Kredit Pajak Potongan Klien)', 'asset', 'debit', '11000', 'Bukti potong PPh 23 dari klien (pengurang PPh Badan)'),
('11060', 'Pajak Pertambahan Nilai (PPN) Masukan', 'asset', 'debit', '11000', 'PPN 11% yang dibayar ke vendor (dapat dikreditkan)'),
('12010', 'Aset Tetap - Perangkat Keras Komputer & Server', 'asset', 'debit', '12000', 'Peralatan workstation kantor dan server lokal'),
('12020', 'Akumulasi Penyusutan Aset Tetap', 'asset', 'credit', '12000', 'Akumulasi depresiasi aset tetap'),
-- 20000 KEWAJIBAN (LIABILITIES)
('21010', 'Utang Usaha Vendor (Accounts Payable)', 'liability', 'credit', '21000', 'Kewajiban tagihan vendor pihak ketiga'),
('21020', 'Utang Honorarium & Gaji Engineer', 'liability', 'credit', '21000', 'Honorarium engineer terakru yang belum ditransfer'),
('21030', 'Utang Pajak PPN Keluaran', 'liability', 'credit', '21000', 'PPN 11% dipungut dari klien atas invoice yang terbit'),
('21040', 'Utang Pajak PPh Pasal 21', 'liability', 'credit', '21000', 'Pajak pemotongan honorarium engineer wajib setor kas negara'),
('21050', 'Utang Pajak PPh Pasal 23', 'liability', 'credit', '21000', 'Pajak pemotongan jasa vendor luar'),
('21060', 'Utang Pajak PPh Badan (Pasal 29)', 'liability', 'credit', '21000', 'PPh Badan terutang akhir tahun'),
('21070', 'Pendapatan Diterima di Muka (Unearned Retainer)', 'liability', 'credit', '21000', 'Pembayaran termin/deposit sebelum pengerjaan'),
-- 30000 EKUITAS (EQUITY)
('31010', 'Modal Disetor Pendiri', 'equity', 'credit', '30000', 'Modal saham disetor resmi sesuai Akta Pendirian'),
('32010', 'Saldo Laba Ditahan (Retained Earnings)', 'equity', 'credit', '30000', 'Akumulasi laba bersih tahun-tahun sebelumnya'),
('33010', 'Laba / (Rugi) Tahun Berjalan', 'equity', 'credit', '30000', 'Laba bersih operasional periode berjalan'),
-- 40000 PENDAPATAN OPERASIONAL (REVENUE)
('41010', 'Pendapatan Jasa Managed Maintenance Care', 'revenue', 'credit', '40000', 'Pendapatan paket pemeliharaan bulanan'),
('41020', 'Pendapatan Jasa Custom Software Development', 'revenue', 'credit', '40000', 'Pendapatan pembuatan aplikasi baru & sistem kustom'),
('41030', 'Pendapatan Jasa Add-On On-Demand Sprint Hours', 'revenue', 'credit', '40000', 'Pendapatan pembelian saldo jam kerja tambahan klien'),
('41040', 'Pendapatan Jasa Infrastruktur & Server Scaling', 'revenue', 'credit', '40000', 'Pendapatan migrasi, optimasi, dan scaling arsitektur server'),
-- 50000 BEBAN POKOK PENDAPATAN / HPP (COST OF SERVICES)
('51010', 'Biaya Honorarium Tenaga Ahli / Engineer Proyek', 'cost_of_sales', 'debit', '50000', 'Pembayaran honor langsung engineer per proyek/tiket'),
('51020', 'Biaya Cloud Hosting, VPS, CDN & Third-Party API', 'cost_of_sales', 'debit', '50000', 'Langganan server cloud staging, VPS client, API gateway'),
-- 60000 BEBAN OPERASIONAL & UMUM (OPEX)
('61010', 'Beban Operasional Kantor & Petty Cash Harian', 'expense', 'debit', '60000', 'Biaya utilitas, listrik, air, internet, ATK kantor'),
('61020', 'Beban Perizinan, Notaris & Legalitas Perusahaan', 'expense', 'debit', '60000', 'Biaya pembaruan akta, perizinan OSS, konsultan legal'),
('61030', 'Beban Pembelian E-Materai & Sertifikasi Digital', 'expense', 'debit', '60000', 'Biaya kuota e-Materai Peruri untuk SPK & BAST'),
('61040', 'Beban Pemasaran, Promosi & Biaya Domain', 'expense', 'debit', '60000', 'Biaya iklan digital, pembaruan domain resmi meldir.id'),
('61050', 'Beban Administrasi Bank & Payment Gateway Fee', 'expense', 'debit', '60000', 'Biaya transfer antarbank, biaya transaksi Midtrans/Xendit'),
('61060', 'Beban Penyusutan Peralatan & Inventaris Komputer', 'expense', 'debit', '60000', 'Beban depresiasi inventaris kantor harian');

-- ==============================================================================
-- 13. DIGITAL BAST (BERITA ACARA SERAH TERIMA) ENGINE
-- ==============================================================================
CREATE TABLE bast_documents (
    id BIGSERIAL PRIMARY KEY,
    bast_number VARCHAR(50) UNIQUE NOT NULL, -- e.g. "BAST/MELDIR/2026/001"
    project_id BIGINT NOT NULL REFERENCES projects(id),
    contract_id BIGINT NOT NULL REFERENCES contracts(id),
    client_id BIGINT NOT NULL REFERENCES users(id),
    title VARCHAR(200) NOT NULL,
    deliverables_summary TEXT NOT NULL,
    sign_date DATE NULL,
    canvas_signature_path VARCHAR(255) NULL,
    initial_pdf_path VARCHAR(255) NULL,
    ematerai_pdf_path VARCHAR(255) NULL,
    status bast_status DEFAULT 'draft',
    signed_at TIMESTAMP WITH TIME ZONE NULL,
    ematerai_uploaded_at TIMESTAMP WITH TIME ZONE NULL,
    notes TEXT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_bast_project ON bast_documents(project_id);
CREATE INDEX idx_bast_status ON bast_documents(status);

-- ==============================================================================
-- 14. MONTHLY MAINTENANCE EXECUTIVE REPORTING
-- ==============================================================================
CREATE TABLE monthly_maintenance_reports (
    id BIGSERIAL PRIMARY KEY,
    project_id BIGINT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    client_id BIGINT NOT NULL REFERENCES users(id),
    period_month INT NOT NULL, -- 1 - 12
    period_year INT NOT NULL,  -- e.g. 2026
    uptime_percentage NUMERIC(5,2) DEFAULT 100.00,
    total_downtime_minutes INT DEFAULT 0,
    successful_backups_count INT DEFAULT 0,
    security_incidents_count INT DEFAULT 0,
    tickets_resolved_count INT DEFAULT 0,
    minor_features_delivered INT DEFAULT 0,
    summary_notes TEXT NULL,
    generated_pdf_path VARCHAR(255) NULL,
    published_at TIMESTAMP WITH TIME ZONE NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(project_id, period_month, period_year)
);

CREATE INDEX idx_monthly_reports_client ON monthly_maintenance_reports(client_id, period_year, period_month);

-- ==============================================================================
-- 15. VAULT CREDENTIALS & GITHUB ASSIGNMENTS
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
-- 16. OFFBOARDING REQUESTS & SECURITY AUDIT LOGS
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
-- END OF COMPLETE ENTERPRISE SCHEMA DDL SCRIPT
-- ==============================================================================
