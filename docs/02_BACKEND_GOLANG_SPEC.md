# ⚙️ Spesifikasi & Skema Backend Golang (PostgreSQL, Workers & Media Engine)

Dokumen ini merinci arsitektur backend Golang dengan database **PostgreSQL**, modul manajemen akun/profil, mesin broadcast notifikasi, serta penanganan upload berkas multi-media (Gambar, Video, PDF).

---

## 1. Endpoints Profil Pengguna, Notifikasi & Media Upload

### A. Manajemen Profil Akun (Universal Seluruh Subdomain)
* `GET  /api/v1/user/profile` -> Mengambil data profil akun aktif (Nama, Email, WA, Avatar).
* `PUT  /api/v1/user/profile` -> Mengubah data nama, email, dan nomor WhatsApp.
* `POST /api/v1/user/avatar` -> Mengunggah foto profil baru (otomatis dikompres ke format WebP).
* `PUT  /api/v1/user/password` -> Mengubah kata sandi (Wajib memvalidasi `old_password` dengan Argon2id).

### B. Sistem Notifikasi Lonceng In-App (Universal Seluruh Subdomain)
* `GET  /api/v1/notifications` -> Mengambil daftar notifikasi pengguna & jumlah unread count.
* `POST /api/v1/notifications/:id/read` -> Menandai satu notifikasi tertentu sebagai telah dibaca.
* `POST /api/v1/notifications/mark-all-read` -> Menandai seluruh notifikasi pengguna sebagai telah dibaca.

### C. Mesin Broadcast Push Notifikasi Manual (Khusus `office.meldir.id`)
* `POST /api/v1/office/notifications/broadcast` -> Mengirim Push Notification massal (Input: `title`, `body`, `target_audience`, `target_url`).

### D. Multi-Media Upload Handler (Gambar, Video, PDF)
* `POST /api/v1/media/upload` -> Endpoint upload berkas universal:
  - Validasi MIME-type: Images (`jpg`, `png`, `webp`), Videos (`mp4`, `webm`), Documents (`pdf`).
  - Maksimal ukuran berkas: Gambar (5MB), Video (50MB), PDF (20MB).
  - Menyimpan file ke `/home/meldir.id/storage/uploads/` dengan nama unik UUID.

---

## 2. Background Workers (Golang Goroutines & Cron)

### A. External Server Health Probe Worker (`internal/worker/health_probe.go`)
* **Interval**: Berjalan otomatis setiap 5 menit.
* **Mekanisme**: Melakukan HTTP GET/PING ke daftar domain/IP server external klien yang terdaftar di database.
* **Tindakan Error**: Jika respons `code >= 500` atau `timeout > 10s`, ubah status server menjadi `CRITICAL_DOWN`, catat timestamp incident, serta picu notifikasi lonceng in-app, WhatsApp darurat, dan Push Notification.

### B. SLA Breach Auto-Escalation Worker (`internal/worker/sla_escalator.go`)
* **Interval**: Berjalan setiap 10 menit.
* **Mekanisme**: Memeriksa tiket `open` / `in_progress` yang sisa SLA timernya `< 30 menit` tanpa ada update dari engineer.
* **Tindakan**: Kirim notifikasi WA darurat, notifikasi lonceng, & PWA Push langsung ke WA Direktur & Admin.

### C. Redis Token Blacklist Middleware (`internal/delivery/http/middleware/auth.go`)
* Memeriksa token JWT 24-Jam. Jika token ID ada dalam Redis Blacklist (di-force logout oleh Admin), batalkan request dengan HTTP 401 Unauthorized.
