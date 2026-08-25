# 🖥️ Panduan Operasional VPS & CyberPanel Menu (AlmaLinux 9)

> **Catatan Rahasia & Internal**: Dokumen ini merupakan buku panduan operasional yang membedakan secara jelas mana tugas yang dikerjakan melalui **🌐 Menu Web CyberPanel** (klik-visual) dan mana yang dikerjakan melalui **🖥️ Terminal VPS SSH** (command-line).

---

## 📋 Pembagian Lokasi Pengerjaan

| Kategori Tugas | Lokasi Eksekusi | Keterangan |
| :--- | :--- | :--- |
| **Buat Database & User** | 🌐 **CyberPanel GUI** | Menu: `Databases` -> `Create Database` |
| **Buat Child Domain Subdomain** | 🌐 **CyberPanel GUI** | Menu: `Websites` -> `Create Child Domain` |
| **Terbitkan SSL HTTPS** | 🌐 **CyberPanel GUI** | Menu: `SSL` -> `Manage SSL` / `Issue SSL` |
| **Edit Reverse Proxy OpenLiteSpeed** | 🌐 **CyberPanel GUI** | Menu: `Websites` -> `List Websites` -> `vHost Conf` |
| **Upload File Build `/dist`** | 🌐 **CyberPanel GUI** | Menu: `Websites` -> `File Manager` (atau via SCP/SFTP) |
| **Install Golang, Redis, Tools** | 🖥️ **Terminal VPS** | Command `dnf install` via SSH |
| **Impor Skema SQL DDL** | 🖥️ **Terminal VPS** | Command `psql -d meldir_db -f schema.sql` |
| **Setup Systemd Service Golang** | 🖥️ **Terminal VPS** | Command `systemctl enable --now meldir-backend` |
| **Live Log & Monitoring** | 🖥️ **Terminal VPS** | Command `journalctl`, `htop`, `free -m` |

---

## 🌐 BAGIAN I: Pengerjaan Melalui Menu Web CyberPanel

Buka CyberPanel di browser: `https://IP_VPS:8090` (Login dengan user `admin`).

### 1. Membuat Database & User PostgreSQL
1. Masuk ke menu **Databases** -> **Create Database**.
2. Pilih domain utama: `meldir.id`.
3. Database Name: `meldir_db`.
4. Username: `meldir_user`.
5. Password: Masukkan password kuat (misal: `PasswordKuatAnda2026!`).
6. Klik **Create Database**.

### 2. Membuat 3 Child Domain (Subdomain)
1. Masuk ke menu **Websites** -> **Create Child Domain**.
2. Pilih Master Domain: `meldir.id`.
3. Buat 3 subdomain secara bertahap:
   * Subdomain: `office` (Hasil: `office.meldir.id` | Path: `/home/office.meldir.id/public_html`)
   * Subdomain: `jobs` (Hasil: `jobs.meldir.id` | Path: `/home/jobs.meldir.id/public_html`)
   * Subdomain: `portal` (Hasil: `portal.meldir.id` | Path: `/home/portal.meldir.id/public_html`)
4. Centang **Create Mail Domain** jika dibutuhkan, lalu klik **Create Child Domain**.

### 3. Menerbitkan Sertifikat SSL HTTPS
1. Masuk ke menu **SSL** -> **Manage SSL** atau **Issue SSL**.
2. Pilih domain `meldir.id` beserta ketiga subdomainnya (`office`, `jobs`, `portal`).
3. Klik **Issue SSL** (Let's Encrypt akan memasang sertifikat HTTPS otomatis).

### 4. Menambahkan Aturan Reverse Proxy OpenLiteSpeed
1. Masuk ke menu **Websites** -> **List Websites**.
2. Pilih subdomain (misal: `office.meldir.id`) -> Klik tombol **vHost Conf**.
3. Gulir ke baris paling bawah dan tempelkan aturan proxy ke Golang Port `:8080`:

```apache
extprocessor meldir-backend {
  type                    proxy
  address                 127.0.0.1:8080
  maxConns                2000
  initTimeout             60
  retryTimeout            0
}

context /api {
  type                    proxy
  handler                 meldir-backend
  addDefaultCharset       off
}
```
4. Klik **Save**.
5. Ulangi langkah di atas untuk `jobs.meldir.id` dan `portal.meldir.id`.

### 5. Mengunggah File Frontend `/dist` Melalui File Manager
1. Masuk ke menu **Websites** -> **File Manager**.
2. Masuk ke direktori `public_html` pada masing-masing subdomain.
3. Klik **Upload** dan ekstrak file hasil build Vue 3 (`dist/*`).

---

## 🖥️ BAGIAN II: Pengerjaan Melalui Terminal VPS (SSH Root)

Buka terminal SSH (Putty / Terminal Windows / MacOS) dan login sebagai `root`:

### 1. Instalasi Paket & Persiapan Tools
```bash
# Update repo dan install paket
dnf update -y
dnf install -y golang redis epel-release git curl wget nano htop

# Aktifkan dan jalankan Redis Server
systemctl enable --now redis
systemctl status redis
```

### 2. Eksekusi Skema Database SQL DDL
Upload file `05_POSTGRESQL_SCHEMA.sql` ke VPS, lalu jalankan:
```bash
sudo -u postgres psql -d meldir_db -f /root/05_POSTGRESQL_SCHEMA.sql
```

### 3. Pembuatan Folder Backend & Izin Akses
```bash
mkdir -p /home/meldir.id/backend
mkdir -p /home/meldir.id/storage/uploads/ematerai
mkdir -p /home/meldir.id/storage/uploads/proofs
mkdir -p /home/meldir.id/storage/credentials_pdf

chown -R cyberpanel:cyberpanel /home/meldir.id/backend
chown -R cyberpanel:cyberpanel /home/meldir.id/storage
chmod -R 755 /home/meldir.id/storage
```

### 4. Membuat & Menjalankan Systemd Service Backend Golang
Buat file service di `/etc/systemd/system/meldir-backend.service`:
```bash
nano /etc/systemd/system/meldir-backend.service
```

Tempelkan konfigurasi:
```ini
[Unit]
Description=meldir.id Multi-Domain Backend Service (Golang)
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=cyberpanel
WorkingDirectory=/home/meldir.id/backend
ExecStart=/home/meldir.id/backend/meldir-api
Restart=always
RestartSec=3s

# Environment Variables
Environment=PORT=8080
Environment=NODE_ENV=production
Environment=DB_DRIVER=postgres
Environment=DB_HOST=127.0.0.1
Environment=DB_PORT=5432
Environment=DB_USER=meldir_user
Environment=DB_PASSWORD=PasswordKuatAnda2026!
Environment=DB_NAME=meldir_db
Environment=SSL_MODE=disable
Environment=REDIS_HOST=127.0.0.1
Environment=REDIS_PORT=6379
Environment=JWT_EXPIRATION=24h
Environment=RATE_LIMIT_PUBLIC_SUBMIT=5

[Install]
WantedBy=multi-user.target
```
*Simpan (`Ctrl + O`, `Enter`), lalu keluar (`Ctrl + X`).*

Jalankan service:
```bash
systemctl daemon-reload
systemctl enable --now meldir-backend
systemctl status meldir-backend
systemctl restart lsws
```

### 5. Perintah Monitoring Harian (Troubleshooting)
```bash
# Pantau log live backend Golang
journalctl -u meldir-backend -f

# Restart backend setelah compile file baru
systemctl restart meldir-backend

# Cek RAM, Disk & CPU
free -m
df -h
htop
```
