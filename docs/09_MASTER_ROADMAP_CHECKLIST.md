# 🗺️ Master Roadmap & Checklist Pengerjaan Webapp meldir.id

> **Catatan Rahasia & Internal**: Dokumen ini merupakan daftar periksa (*master checklist*) rencana eksekusi pembuatan webapp dan multi-dashboard meldir.id.
> 
> **Keterangan Penanggung Jawab & Lokasi Eksekusi**:
> * 🤖 **AI Assistant (Antigravity)**: Penulisan kode sumber (Golang, Vue 3, SQL DDL, Unit Testing, File Konfigurasi).
> * 🖥️ 👤 **User (Terminal VPS)**: Eksekusi perintah command-line via SSH (dnf, systemctl, redis, systemd).
> * 🌐 👤 **User (CyberPanel GUI)**: Eksekusi konfigurasi visual via Dashboard Web CyberPanel (Websites, Database, SSL, vHost, File Manager).

---

## 📌 Fase 1: Persiapan Server, Database & Direktori

- [ ] 🖥️ 👤 **[1.1]** Akses SSH ke VPS AlmaLinux 9 Hostinger sebagai `root` dan jalankan instalasi paket (`dnf install -y golang redis epel-release git htop`).
- [ ] 🖥️ 👤 **[1.2]** Aktifkan dan jalankan Redis Server di terminal (`systemctl enable --now redis`).
- [ ] 🌐 👤 **[1.3]** Buka **CyberPanel Menu**: Masuk ke **Databases** -> **Create Database** -> Buat database `meldir_db` dan user `meldir_user`.
- [ ] 🖥️ 👤 **[1.4]** Eksekusi file DDL [05_POSTGRESQL_SCHEMA.sql](file:///c:/Users/YSPJ/Documents/Works/meldir/docs/05_POSTGRESQL_SCHEMA.sql) ke PostgreSQL via terminal (`sudo -u postgres psql -d meldir_db -f /root/05_POSTGRESQL_SCHEMA.sql`).
- [ ] 🤖 **[1.5]** Verifikasi integritas relasi tabel, foreign keys, dan tipe data ENUM database.
- [ ] 🌐 👤 **[1.6]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **File Manager** -> Buat folder `/home/meldir.id/backend` dan direktori penyimpanan:
  - `/home/meldir.id/storage/uploads/ematerai`
  - `/home/meldir.id/storage/uploads/proofs`
  - `/home/meldir.id/storage/uploads/tax_invoices`
  - `/home/meldir.id/storage/uploads/expenses`
  - `/home/meldir.id/storage/uploads/tax_bupot`
  - `/home/meldir.id/storage/uploads/tax_vat_in`
  - `/home/meldir.id/storage/uploads/spt_reports`
  - `/home/meldir.id/storage/uploads/bast`
  - `/home/meldir.id/storage/uploads/reports`
  - `/home/meldir.id/storage/uploads/media`
  - `/home/meldir.id/storage/credentials_pdf`

---

## 📌 Fase 2: Fondasi Backend Golang (Core API Engine)

- [ ] 🤖 **[2.1]** Inisialisasi modul Golang (`go mod init meldir-backend`) dan setup Clean Architecture (`/cmd`, `/internal`, `/pkg`).
- [ ] 🤖 **[2.2]** Buat modul koneksi database PostgreSQL menggunakan `pgx / GORM` dan koneksi `go-redis`.
- [ ] 🤖 **[2.3]** Buat middleware Host-Based Subdomain Routing (`office`, `jobs`, `portal`, `public`).
- [ ] 🤖 **[2.4]** Buat middleware Authentication JWT (Masa berlaku 24 Jam) & Redis Token Blacklist untuk fitur **Force Logout**.
- [ ] 🤖 **[2.5]** Buat middleware Redis Rate-Limiter untuk proteksi form publik & anti-brute force.
- [ ] 🤖 **[2.6]** Implementasi modul Auth (Login, Refresh Token, Logout, Profile) untuk 5 peranan (*RBAC*).
- [ ] 🤖 **[2.7]** Implementasi modul Manajemen Profil Akun (Ubah Nama, Email, WA, Upload Avatar WebP, & Ganti Password Argon2id).
- [ ] 🤖 **[2.8]** Implementasi Universal Multi-Media Upload Handler (Validasi & simpan Gambar, Video, PDF).

---

## 📌 Fase 3: Logika Bisnis 18 Modul Operasional & Background Workers

- [ ] 🤖 **[3.1]** Implementasi modul Inbound Leads CRM (Submit form publik & manajemen pipeline status di Office).
- [ ] 🤖 **[3.2]** Implementasi modul Manajemen Kontrak SPK (Generasi PDF awal dari TTD canvas & verifikasi upload E-Materai).
- [ ] 🤖 **[3.3]** Implementasi modul In-App Notification Bell System (Unread count, mark as read, & deep-link routing).
- [ ] 🤖 **[3.4]** Implementasi modul Manual Push Broadcast Engine (Kirim notifikasi massal per target audiens).
- [ ] 🤖 **[3.5]** Implementasi modul WhatsApp API Engine Unofficial (Status handler, QR Code stream, disconnect, & clear cache).
- [ ] 🤖 **[3.6]** Implementasi modul Add-On & Saldo Jam Kerja Metered (Katalog paket, order add-on, auto-invoice, & saldo hours).
- [ ] 🤖 **[3.7]** Implementasi modul Tracker Kuota Fitur Ringan (Validasi tiket minor feature limit 2x/bulan & auto-suggest Add-On).
- [ ] 🤖 **[3.8]** Implementasi modul Engineer Timesheet & Work Logs (Input jam kerja per tiket & auto-deduct saldo Add-On klien).
- [ ] 🤖 **[3.9]** Implementasi modul Invoicing & Pelampiran e-Faktur Pajak PPN (Upload PDF e-Faktur DJP & download mandiri klien).
- [ ] 🤖 **[3.10]** Implementasi modul Digital BAST Engine (Draf BAST, TTD canvas digital, & upload BAST e-Materai).
- [ ] 🤖 **[3.11]** Implementasi modul Keuangan PT & Form Input Manual Honorarium Engineer (termasuk kalkulasi PPh 21 & opsi Rp 0 pro-bono).
- [ ] 🤖 **[3.12]** Implementasi modul Vault Kredensial Server (Enkripsi AES-256) & Otomatisasi Hak Push GitHub Engineer.
- [ ] 🤖 **[3.13]** Implementasi modul Offboarding Klien (Pengajuan, Penawaran Banding Office, Transfer GitHub, & PDF Kredensial).
- [ ] 🤖 **[3.14]** Implementasi Security Audit Logger (Pencatatan jejak digital *tamper-proof* ke database).
- [ ] 🤖 **[3.15]** Implementasi Worker External Server Health Probe (Cron ping/HTTP status check setiap 5 menit).
- [ ] 🤖 **[3.16]** Implementasi Worker SLA Auto-Escalation (Deteksi tiket kritis < 30 menit & kirim alert WA darurat ke Direktur).
- [ ] 🤖 **[3.17]** Implementasi Worker Monthly Executive Report Generator & Quota Reset (Cron tanggal 1 setiap bulan 00:00 WIB).
- [ ] 🤖 **[3.18]** Implementasi Modul Akuntansi Double-Entry (Bagan Akun COA 5-Digit, Jurnal Umum, Validasi Otomatis Keseimbangan Debit/Kredit, & Mutasi Buku Besar General Ledger).
- [ ] 🤖 **[3.19]** Implementasi Mesin Penjurnalan Otomatis (Auto-Journaling Engine saat Invoice lunas, Honorarium dicairkan, & Beban operasional dicatat).
- [ ] 🤖 **[3.20]** Implementasi Modul Beban Operasional & Kas Kecil Digital (Petty Cash, upload kuitansi WebP/PDF, & klasifikasi biaya fiskal *Deductible*).
- [ ] 🤖 **[3.21]** Implementasi Modul Kepatuhan PPN 11% (Rekap Faktur Keluaran & Masukan, Formulir SPT Masa PPN 1111 Kurang/Lebih Bayar, & Bukti Setor NTPN).
- [ ] 🤖 **[3.22]** Implementasi Modul Bukti Potong Pajak (PPh 21 Engineer e-Bupot & Kredit Pajak PPh 23 Klien 2%).
- [ ] 🤖 **[3.23]** Implementasi Mesin Simulasi & Estimasi SPT Tahunan PPh Badan 1771 (Koreksi Fiskal Positif/Negatif, Fasilitas Diskon 50% Tarif Pasal 31E UU HPP, Kredit PPh 23 & 25, serta PPh Kurang Bayar Pasal 29).
- [ ] 🤖 **[3.24]** Implementasi Worker Pengingat Batas Waktu Pajak (Tax Deadline Alert) & Validator Keseimbangan Jurnal Harian.

---

## 📌 Fase 4: Fondasi Frontend Vue 3 + Vite Multi-Portal

- [ ] 🤖 **[4.1]** Inisialisasi project Vue 3 + Vite dengan Composition API dan Tailwind / Custom Design System.
- [ ] 🤖 **[4.2]** Setup Pinia State Management (`authStore`, `leadStore`, `notificationStore`, `addonStore`, `quotaStore`, `ticketStore`, `timesheetStore`, `financeStore`, `accountingStore`, `expenseStore`, `taxStore`, `reportStore`, `langStore`).
- [ ] 🤖 **[4.3]** Setup Vue Router dengan Subdomain Detection Guard & Role Enforcement (Kunci form untuk role `audit`).
- [ ] 🤖 **[4.4]** Setup Konfigurasi Multi-Manifest PWA (`office-manifest.json`, `jobs-manifest.json`, `portal-manifest.json`) dan Service Worker Caching.
- [ ] 🤖 **[4.5]** Implementasi Base Layout Shell dengan Skala Tampilan **80% Compact Density** dan **Universal Light/Dark Toggle**.
- [ ] 🤖 **[4.6]** Implementasi Header Bar Universal lengkap dengan **Icon Lonceng Notifikasi Dropdown**, **Tombol Unduh PWA**, dan **Avatar Profile Menu**.
- [ ] 🤖 **[4.7]** Implementasi Sidebar Desktop (Collapsible Icon-Only Mode) dan Mobile PWA (5-Item Bottom Bar + Bottom Sheet "Lainnya").

---

## 📌 Fase 5: Antarmuka UI/UX 3 Portal Dashboard

### 🏢 A. Portal Office (`office.meldir.id`)
- [ ] 🤖 **[5.1]** Bangun Executive Summary Cards & Grafik Keuangan PT (Smooth Area Wave Chart Solid Color).
- [ ] 🤖 **[5.2]** Bangun Papan Pipeline Prospek (Leads Kanban) dengan modal detail & tombol konversi ke Klien SPK.
- [ ] 🤖 **[5.3]** Bangun Manajemen Katalog Add-On & Rekapitulasi Timesheet Engineer.
- [ ] 🤖 **[5.4]** Bangun Modal Unggah e-Faktur Pajak PPN (11%) pada modul Invoice.
- [ ] 🤖 **[5.5]** Bangun BAST Document Builder & Verifikasi BAST e-Materai.
- [ ] 🤖 **[5.6]** Bangun Komponen Kartu WhatsApp API Engine (Live QR Scanner & Action Buttons).
- [ ] 🤖 **[5.7]** Bangun Widget Monitoring Uptime Server External Klien (Grid Hijau/Kuning/Merah).
- [ ] 🤖 **[5.8]** Bangun Form Manual Komisi Engineer & Modal Penawaran Banding Offboarding.
- [ ] 🤖 **[5.9]** Bangun Tabel Sesi Aktif dengan Tombol **"Force Logout All"**.
- [ ] 🤖 **[5.10]** Bangun Menu Navigasi & Tabel Buku Besar Berpasangan (General Ledger & Double-Entry Journals) dengan Live Balance Validator.
- [ ] 🤖 **[5.11]** Bangun Laporan Finansial Komprehensif (Laba Rugi Komersial & Fiskal, Neraca Keuangan, Neraca Saldo, & Tombol Cetak PDF Kop PT).
- [ ] 🤖 **[5.12]** Bangun Modal Entri Beban Operasional & Kas Kecil Digital (Petty Cash Receipts & Fiscal Deductible Switch).
- [ ] 🤖 **[5.13]** Bangun Hub Perpajakan Terpadu (Papan PPN 11%, Hub e-Bupot PPh 21/23, & Simulator Interaktif SPT Tahunan PPh Badan 1771).

### ⚙️ B. Portal Jobs (`jobs.meldir.id`)
- [ ] 🤖 **[5.14]** Bangun Developer Taskboard dengan Penghitung Waktu Mundur SLA Real-Time & Tombol Unduh PWA Jobs.
- [ ] 🤖 **[5.15]** Bangun Modal Input Timesheet Log Jam Kerja Teknis per tiket/task.
- [ ] 🤖 **[5.16]** Bangun Panel GitHub Repository Integration & History Payout Log.
- [ ] 🤖 **[5.17]** Bangun Halaman Diskusi Tiket Bug SLA dengan Multi-Media Uploader & Viewer (Gambar/Video/PDF).
- [ ] 🤖 **[5.18]** Bangun Panel Bukti Potong Pajak PPh 21 Engineer (Unduh Slip e-Bupot 21 Resmi).

### 💼 C. Portal Klien (`portal.meldir.id`)
- [ ] 🤖 **[5.19]** Bangun Client Control Center, Grafik Live Server Uptime & Tombol Unduh PWA Portal Klien.
- [ ] 🤖 **[5.20]** Bangun Widget Saldo Jam Kerja Add-On & Katalog Pemesanan Add-On Instan.
- [ ] 🤖 **[5.21]** Bangun Badge Indikator Kuota Fitur Ringan Bulanan (2/2) pada form tiket.
- [ ] 🤖 **[5.22]** Bangun Tabel Transparansi Timesheet Log Pekerjaan Engineer.
- [ ] 🤖 **[5.23]** Bangun Hub Arsip Laporan Kinerja Bulanan Eksekutif (Download PDF Uptime & Security).
- [ ] 🤖 **[5.24]** Bangun Billing Center dengan Tombol Unduh e-Faktur Pajak PPN resmi.
- [ ] 🤖 **[5.25]** Bangun Modal Tanda Tangan Canvas Digital untuk SPK & BAST (iPad Drawing Pad Style).
- [ ] 🤖 **[5.26]** Bangun Hub E-Materai (Download PDF awal & Upload Ulang PDF E-Materai).
- [ ] 🤖 **[5.27]** Bangun Wizard Offboarding Klien & Download PDF Kredensial Server.

---

## 📌 Fase 6: Konfigurasi CyberPanel Menu & Systemd Service di VPS

- [ ] 🌐 👤 **[6.1]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **Create Child Domain** -> Buat 3 subdomain (`office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`).
- [ ] 🌐 👤 **[6.2]** Buka **CyberPanel Menu**: Masuk ke **SSL** -> **Manage SSL** / **Issue SSL** -> Terbitkan SSL untuk masing-masing subdomain / Wildcard SSL `*.meldir.id`.
- [ ] 🌐 👤 **[6.3]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **List Websites** -> Pilih subdomain -> Klik **vHost Conf** -> Tempelkan aturan Reverse Proxy OpenLiteSpeed (`extprocessor` & `context /api` -> `127.0.0.1:8080`).
- [ ] 🖥️ 👤 **[6.4]** Buka **Terminal VPS**: Buat file `/etc/systemd/system/meldir-backend.service`, lalu jalankan `systemctl daemon-reload && systemctl enable --now meldir-backend`.
- [ ] 🖥️ 👤 **[6.5]** Buka **Terminal VPS**: Restart OpenLiteSpeed web server (`systemctl restart lsws`).

---

## 📌 Fase 7: Pengujian, QA Checklist & UAT (User Acceptance Testing)

- [ ] 🤖 **[7.1]** Jalankan Automated Unit Tests & API Integration Tests untuk endpoint Golang.
- [ ] 🤖 & 👤 **[7.2]** Uji coba instalasi PWA di ketiga domain (`office`, `jobs`, `portal`).
- [ ] 🤖 & 👤 **[7.3]** Uji coba formulir Inbound Leads dari landing page -> pastikan kartu otomatis muncul di Kanban Office.
- [ ] 🤖 & 👤 **[7.4]** Uji coba alur pendaftaran Klien (Pengajuan -> Kontrak TTD Canvas -> E-Materai -> Klien Aktif).
- [ ] 🤖 & 👤 **[7.5]** Uji coba pemesanan Add-On Sprint Jam -> terbit invoice -> saldo hours bertambah setelah bayar.
- [ ] 🤖 & 👤 **[7.6]** Uji coba kuota fitur ringan (Tiket 1 & 2 gratis -> Tiket 3 diblokir dan diminta potong saldo Add-On).
- [ ] 🤖 & 👤 **[7.7]** Uji coba input Timesheet oleh engineer -> saldo hours klien berkurang otomatis dan riwayat tampil di portal.
- [ ] 🤖 & 👤 **[7.8]** Uji coba penerbitan & TTD digital BAST di akhir proyek.
- [ ] 🤖 & 👤 **[7.9]** Uji coba pelampiran e-Faktur Pajak oleh Admin & unduh file oleh Klien.
- [ ] 🤖 & 👤 **[7.10]** Uji coba pemicu Worker Laporan Bulanan -> verifikasi file PDF ter-generate dan tersimpan di database.
- [ ] 🤖 & 👤 **[7.11]** Uji coba simulasi Server External Down (< 5 menit kirim alert WA & Push PWA).
- [ ] 🤖 & 👤 **[7.12]** Uji coba Force Logout Sesi dari dashboard Office.
- [ ] 🤖 & 👤 **[7.13]** Uji coba Auto-Journaling Akuntansi (Invoice lunas -> terbit jurnal debit kas, kredit piutang/pendapatan & PPN keluaran).
- [ ] 🤖 & 👤 **[7.14]** Uji coba pencatatan beban operasional kantor (Petty Cash) & verifikasi keseimbangan Buku Besar.
- [ ] 🤖 & 👤 **[7.15]** Uji coba rekonsiliasi SPT Masa PPN 1111 (Keluaran - Masukan) & input NTPN.
- [ ] 🤖 & 👤 **[7.16]** Uji coba penerbitan Bukti Potong PPh 21 Engineer & unduh slip oleh engineer di Portal Jobs.
- [ ] 🤖 & 👤 **[7.17]** Uji coba Simulator SPT Tahunan PPh Badan 1771 & validasi penerapan diskon 50% Fasilitas Pasal 31E UU HPP.
- [ ] 🤖 & 👤 **[7.18]** Uji coba akun role `audit` untuk memastikan seluruh modul akuntansi & operasional berstatus READ-ONLY tanpa tombol manipulasi data.

---

## 📌 Fase 8: Peluncuran Resmi (Go-Live)

- [ ] 🤖 **[8.1]** Build production bundle Vue 3 Frontend (`npm run build`).
- [ ] 🌐 👤 **[8.2]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **File Manager** -> Unggah hasil build `/dist` ke folder `public_html` masing-masing subdomain.
- [ ] 🖥️ 👤 **[8.3]** Buka **Terminal VPS**: Compile biner akhir Golang (`go build -o meldir-api`) di `/home/meldir.id/backend` dan restart service (`systemctl restart meldir-backend`).
- [ ] 🤖 & 👤 **[8.4]** Verifikasi live traffic di seluruh domain (`meldir.id`, `office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`).
- [ ] 🚀 **[8.5]** Sistem Enterprise meldir.id Resmi Beroperasi Penuh!

