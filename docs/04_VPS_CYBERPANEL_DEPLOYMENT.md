# 🚀 Panduan Deploy PostgreSQL & Redis Service Config di CyberPanel VPS

Dokumen ini berisi langkah-langkah teknis pengkonfigurasian koneksi PostgreSQL, Redis Token Blacklist, dan Rate Limiter pada VPS AlmaLinux 9.

---

## 1. Konfigurasi Database PostgreSQL & Redis Rate Limiter

Backend Golang terhubung ke instance **PostgreSQL** dan **Redis**:

```ini
# Environment Variables di Systemd Service (/etc/systemd/system/meldir-backend.service)
Environment=DB_DRIVER=postgres
Environment=DB_HOST=127.0.0.1
Environment=DB_PORT=5432
Environment=DB_USER=postgres
Environment=DB_PASSWORD=YOUR_SECURE_POSTGRES_PASSWORD
Environment=DB_NAME=meldir_db
Environment=REDIS_HOST=127.0.0.1
Environment=REDIS_PORT=6379
Environment=JWT_EXPIRATION=24h
Environment=RATE_LIMIT_PUBLIC_SUBMIT=5
```
