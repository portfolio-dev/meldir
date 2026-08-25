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
- [ ] 🖥️ 👤 **[1.4]** Eksekusi file DDL [05_POSTGRESQL_SCHEMA.sql](file:///c:/Users/Acer/Documents/Pekerjaan/Triple_C_Online/meldir/meldir/docs/05_POSTGRESQL_SCHEMA.sql) ke PostgreSQL via terminal (`sudo -u postgres psql -d meldir_db -f /path/05_POSTGRESQL_SCHEMA.sql`).
- [ ] 🤖 **[1.5]** Verifikasi integritas relasi tabel dan tipe data ENUM database.
- [ ] 🌐 👤 **[1.6]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **File Manager** -> Buat folder `/home/meldir.id/backend` dan direktori penyimpanan `/home/meldir.id/storage/uploads/` (atau jalankan `mkdir -p` & `chown` di terminal).

---

## 📌 Fase 2: Fondasi Backend Golang (Core API Engine)

- [ ] 🤖 **[2.1]** Inisialisasi modul Golang (`go mod init meldir-backend`) dan setup Clean Architecture (`/cmd`, `/internal`, `/pkg`).
- [ ] 🤖 **[2.2]** Buat modul koneksi database PostgreSQL menggunakan `pgx / GORM` dan koneksi `go-redis`.
- [ ] 🤖 **[2.3]** Buat middleware Host-Based Subdomain Routing (`office`, `jobs`, `portal`, `public`).
- [ ] 🤖 **[2.4]** Buat middleware Authentication JWT (Masa berlaku 24 Jam) & Redis Token Blacklist untuk fitur **Force Logout**.
- [ ] 🤖 **[2.5]** Buat middleware Redis Rate-Limiter untuk proteksi form publik & anti-brute force.
- [ ] 🤖 **[2.6]** Implementasi modul Auth (Login, Refresh Token, Logout, Profile) untuk 5 peranan (*RBAC*).
- [ ] 🤖 **[2.7]** Implementasi modul Manajemen Profil Akun (Ubah Nama, Email, WA, Upload Avatar WebP, & Ganti Password dengan verifikasi password lama).
- [ ] 🤖 **[2.8]** Implementasi Universal Multi-Media Upload Handler (Validasi & simpan Gambar, Video, PDF).

---

## 📌 Fase 3: Logika Bisnis 12 Modul Operasional & Background Workers

- [ ] 🤖 **[3.1]** Implementasi modul Manajemen Kontrak (Generasi PDF awal dari TTD canvas & verifikasi upload E-Materai).
- [ ] 🤖 **[3.2]** Implementasi modul In-App Notification Bell System (Unread count, mark as read, & deep-link routing).
- [ ] 🤖 **[3.3]** Implementasi modul Manual Push Broadcast Engine (Kirim notifikasi massal per target audiens).
- [ ] 🤖 **[3.4]** Implementasi modul WhatsApp API Engine Unofficial (Status handler, QR Code stream, disconnect, & clear cache).
- [ ] 🤖 **[3.5]** Implementasi Worker External Server Health Probe (Cron ping/HTTP status check setiap 5 menit).
- [ ] 🤖 **[3.6]** Implementasi Worker SLA Auto-Escalation (Deteksi tiket kritis < 30 menit & kirim alert WA darurat ke Direktur).
- [ ] 🤖 **[3.7]** Implementasi modul Tagihan & Invoice (Upload bukti transfer manual & verifikasi admin).
- [ ] 🤖 **[3.8]** Implementasi modul Keuangan PT & Form Input Manual Honorarium Engineer (termasuk Rp 0 pro-bono).
- [ ] 🤖 **[3.9]** Implementasi modul Vault Kredensial Server (Enkripsi AES-256) & Otomatisasi Hak Push GitHub Engineer.
- [ ] 🤖 **[3.10]** Implementasi modul Offboarding Klien (Pengajuan, Penawaran Banding Office, Transfer GitHub, & PDF Ringkasan Kredensial).
- [ ] 🤖 **[3.11]** Implementasi Security Audit Logger (Pencatatan jejak digital *tamper-proof* ke database).

---

## 📌 Fase 4: Fondasi Frontend Vue 3 + Vite Multi-Portal

- [ ] 🤖 **[4.1]** Inisialisasi project Vue 3 + Vite dengan Composition API dan Tailwind / Custom Design System.
- [ ] 🤖 **[4.2]** Setup Pinia State Management (`authStore`, `notificationStore`, `projectStore`, `ticketStore`, `financeStore`, `langStore`).
- [ ] 🤖 **[4.3]** Setup Vue Router dengan Subdomain Detection Guard & Role Enforcement (Kunci form untuk role `audit`).
- [ ] 🤖 **[4.4]** Setup Konfigurasi Multi-Manifest PWA (`office-manifest.json`, `jobs-manifest.json`, `portal-manifest.json`) dan Service Worker Caching.
- [ ] 🤖 **[4.5]** Implementasi Base Layout Shell dengan Skala Tampilan **80% Compact Density** dan **Universal Light/Dark Toggle**.
- [ ] 🤖 **[4.6]** Implementasi Header Bar Universal lengkap dengan **Icon Lonceng Notifikasi Dropdown**, **Tombol Unduh PWA**, dan **Avatar Profile Menu**.
- [ ] 🤖 **[4.7]** Implementasi Sidebar Desktop (Collapsible Icon-Only Mode) dan Mobile PWA (5-Item Bottom Bar + Bottom Sheet "Lainnya").

---

## 📌 Fase 5: Antarmuka UI/UX 3 Portal Dashboard

### 🏢 A. Portal Office (`office.meldir.id`)
- [ ] 🤖 **[5.1]** Bangun Executive Summary Cards & Grafik Keuangan PT (Smooth Area Wave Chart Solid Color).
- [ ] 🤖 **[5.2]** Bangun Tombol Unduh PWA Office & Modal Broadcast Push Notifikasi Manual.
- [ ] 🤖 **[5.3]** Bangun Komponen Kartu WhatsApp API Engine (Live QR Scanner & Action Buttons).
- [ ] 🤖 **[5.4]** Bangun Widget Monitoring Uptime Server External Klien (Grid Hijau/Kuning/Merah).
- [ ] 🤖 **[5.5]** Bangun Papan Kanban Pipeline Klien & Contract Builder.
- [ ] 🤖 **[5.6]** Bangun Form Manual Komisi Engineer & Modal Penawaran Banding Offboarding.
- [ ] 🤖 **[5.7]** Bangun Tabel Sesi Aktif dengan Tombol **"Force Logout All"**.

### ⚙️ B. Portal Jobs (`jobs.meldir.id`)
- [ ] 🤖 **[5.8]** Bangun Developer Taskboard dengan Penghitung Waktu Mundur SLA Real-Time & Tombol Unduh PWA Jobs.
- [ ] 🤖 **[5.9]** Bangun Panel GitHub Repository Integration & History Payout Log.
- [ ] 🤖 **[5.10]** Bangun Halaman Diskusi Tiket Bug SLA dengan Multi-Media Uploader & Viewer (Gambar/Video/PDF).

### 💼 C. Portal Klien (`portal.meldir.id`)
- [ ] 🤖 **[5.11]** Bangun Client Control Center, Grafik Live Server Uptime & Tombol Unduh PWA Portal Klien.
- [ ] 🤖 **[5.12]** Bangun Form Tiket SLA Baru dengan Multi-Media Attachment (Screenshot/Video Bug).
- [ ] 🤖 **[5.13]** Bangun Modal Tanda Tangan Canvas Digital ala iPad Drawing Pad (Pilihan Tinta Hitam/Biru).
- [ ] 🤖 **[5.14]** Bangun Hub E-Materai (Download PDF awal & Upload Ulang PDF E-Materai).
- [ ] 🤖 **[5.15]** Bangun Billing Center (Unduh Invoice PDF & Upload Bukti Transfer Manual).
- [ ] 🤖 **[5.16]** Bangun Wizard Offboarding Klien & Download PDF Kredensial Server.

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
- [ ] 🤖 & 👤 **[7.3]** Uji coba alur pendaftaran Klien (Pengajuan -> Kontrak TTD Canvas -> E-Materai -> Klien Aktif).
- [ ] 🤖 & 👤 **[7.4]** Uji coba Icon Lonceng Notifikasi (Unread count, mark all read, deep link navigation).
- [ ] 🤖 & 👤 **[7.5]** Uji coba kirim Push Notifikasi Broadcast Manual dari Office ke Klien/Engineer.
- [ ] 🤖 & 👤 **[7.6]** Uji coba upload & preview berkas media (Gambar, Video, PDF) di tiket SLA.
- [ ] 🤖 & 👤 **[7.7]** Uji coba scan QR Code WhatsApp API dan pengiriman pesan notifikasi otomatis.
- [ ] 🤖 & 👤 **[7.8]** Uji coba simulasi Server External Down (Pastikan alert WA & Push PWA terkirim dalam < 5 menit).
- [ ] 🤖 & 👤 **[7.9]** Uji coba Force Logout Sesi dari dashboard Office.
- [ ] 🤖 & 👤 **[7.10]** Uji coba akun role `audit` untuk memastikan seluruh tombol mutasi ter-disable.

---

## 📌 Fase 8: Peluncuran Resmi (Go-Live)

- [ ] 🤖 **[8.1]** Build production bundle Vue 3 Frontend (`npm run build`).
- [ ] 🌐 👤 **[8.2]** Buka **CyberPanel Menu**: Masuk ke **Websites** -> **File Manager** -> Unggah hasil build `/dist` ke folder `public_html` masing-masing subdomain.
- [ ] 🖥️ 👤 **[8.3]** Buka **Terminal VPS**: Compile biner akhir Golang (`go build -o meldir-api`) di `/home/meldir.id/backend` dan restart service (`systemctl restart meldir-backend`).
- [ ] 🤖 & 👤 **[8.4]** Verifikasi live traffic di seluruh domain (`meldir.id`, `office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`).
- [ ] 🚀 **[8.5]** Sistem Resmi Beroperasi Penuh!
