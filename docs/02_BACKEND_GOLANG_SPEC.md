# ⚙️ Spesifikasi & Skema Backend Golang (PostgreSQL, Workers & Enterprise Engines)

Dokumen ini merinci arsitektur backend Golang dengan Clean Architecture, database **PostgreSQL 14+**, **Redis Cache & Blacklist**, 17 alur modul operasional, background workers, serta penanganan berkas multi-media (Gambar, Video, PDF).

---

## 1. Endpoints Lengkap per Subdomain & Peranan (*RBAC*)

### A. Autentikasi & Profil Akun (Universal)
* `POST /api/v1/auth/login` -> Autentikasi dengan email & password (Argon2id), menghasilkan JWT Token (24 Jam).
* `POST /api/v1/auth/refresh` -> Refresh token sesi aktif.
* `POST /api/v1/auth/logout` -> Memasukkan token ID ke Redis Blacklist.
* `GET  /api/v1/user/profile` -> Mengambil data profil akun aktif (Nama, Email, WA, Avatar WebP, Role).
* `PUT  /api/v1/user/profile` -> Mengubah data profil (Nama, Email, Nomor WhatsApp).
* `POST /api/v1/user/avatar` -> Mengunggah foto profil (otomatis dikompres ke WebP).
* `PUT  /api/v1/user/password` -> Mengubah kata sandi (Wajib memvalidasi `old_password`).

### B. Inbound Leads CRM Pipeline (`meldir.id` & `office.meldir.id`)
* `POST /api/v1/public/leads` -> Endpoint publik formulir konsultasi & audit gratis (Rate-limited via Redis: maks 5 req/jam/IP).
* `GET  /api/v1/office/leads` -> Mengambil daftar seluruh leads dengan filter status (`new`, `contacted`, `quoted`, `won_contract`, `lost`).
* `PUT  /api/v1/office/leads/:id/status` -> Mengubah status pipeline leads & menambahkan catatan follow-up.
* `POST /api/v1/office/leads/:id/convert` -> Mengonversi lead menjadi akun Klien baru di `portal.meldir.id` dan membuat draf SPK.

### C. Sistem Notifikasi Lonceng In-App (Universal)
* `GET  /api/v1/notifications` -> Mengambil daftar notifikasi pengguna & jumlah unread count.
* `POST /api/v1/notifications/:id/read` -> Menandai satu notifikasi tertentu sebagai telah dibaca.
* `POST /api/v1/notifications/mark-all-read` -> Menandai seluruh notifikasi pengguna sebagai telah dibaca.
* `POST /api/v1/office/notifications/broadcast` -> Mengirim Push Notification massal (Input: `title`, `body`, `target_audience`, `target_url`).

### D. Modul Add-On & Saldo Jam Kerja Metered (`portal.meldir.id` & `office.meldir.id`)
* `GET  /api/v1/addons/packages` -> Mengambil katalog paket Add-On aktif (Hourly, Sprint 10 Jam, Sprint 20 Jam).
* `POST /api/v1/client/addons/order` -> Memesan paket Add-On, otomatis menerbitkan Invoice tagihan.
* `GET  /api/v1/client/addons/balance` -> Mengambil sisa saldo jam kerja aktif (`remaining_hours`) per proyek.
* `GET  /api/v1/office/addons/orders` -> Manajemen seluruh pesanan Add-On klien.

### E. Tracker Kuota Fitur Ringan Bulanan (`portal.meldir.id`)
* `GET  /api/v1/client/quota/minor-feature` -> Mengambil sisa kuota fitur ringan bulan berjalan (e.g. `used: 1, total: 2`).
* Validasi otomatis saat klien membuat tiket dengan opsi `is_minor_feature: true`:
  - Jika kuota bulan ini masih tersedia (`used < 2`), tiket diproses gratis.
  - Jika kuota habis (`used >= 2`), tolak dengan respons HTTP 400 dan tautan untuk memesan paket Add-On.

### F. Engineer Timesheet & Log Waktu Kerja (`jobs.meldir.id` & `portal.meldir.id`)
* `POST /api/v1/engineer/timesheet` -> Engineer menginput log kerja (Tiket ID, Jam kerja, Deskripsi pengerjaan).
  - Jika tiket terhubung dengan Add-On, sistem otomatis memotong `remaining_hours` pada pesanan Add-On klien.
* `GET  /api/v1/client/timesheet` -> Klien melihat riwayat transparansi penggunaan jam kerja per tiket secara real-time.
* `GET  /api/v1/office/timesheets` -> Admin/Direktur memantau rekap jam kerja seluruh engineer per periode.

### G. Invoicing & Lampiran e-Faktur Pajak PPN 11% (`office` & `portal`)
* `GET  /api/v1/invoices` -> Daftar invoice per klien atau seluruh invoice (Office).
* `POST /api/v1/client/invoices/:id/proof` -> Klien mengunggah bukti transfer manual (JPG/PDF).
* `POST /api/v1/office/invoices/:id/verify` -> Admin memverifikasi pembayaran -> Status menjadi `paid`.
* `POST /api/v1/office/invoices/:id/tax-invoice` -> Admin mengunggah PDF e-Faktur resmi DJP dan nomor seri faktur.
* `GET  /api/v1/client/invoices/:id/tax-invoice` -> Klien mengunduh PDF e-Faktur resmi.

### H. Digital BAST (Berita Acara Serah Terima) Engine (`office` & `portal`)
* `POST /api/v1/office/bast` -> Admin menerbitkan draf dokumen BAST per milestone proyek.
* `POST /api/v1/client/bast/:id/sign` -> Klien membubuhkan tanda tangan digital canvas -> Generate PDF BAST awal.
* `POST /api/v1/client/bast/:id/ematerai` -> Klien mengunggah ulang PDF BAST yang telah dibubuhi E-Materai resmi.
* `GET  /api/v1/bast/:id/download` -> Mengunduh salinan berkas BAST sah.

### I. Laporan Kinerja Bulanan Eksekutif (Monthly Executive Reports)
* `GET  /api/v1/client/reports/monthly` -> Klien melihat daftar arsip laporan bulanan sistem per periode.
* `GET  /api/v1/client/reports/monthly/:id/download` -> Mengunduh berkas PDF laporan resmi.
* `POST /api/v1/office/reports/monthly/generate` -> Pemicu manual generasi laporan bulanan (jika diperlukan sebelum tanggal 1).

### J. Multi-Media Universal Uploader
* `POST /api/v1/media/upload` -> Endpoint upload berkas universal:
  - Validasi MIME-type: Gambar (`jpg`, `png`, `webp` maks 5MB), Video (`mp4`, `webm` maks 50MB), Dokumen (`pdf` maks 20MB).
  - Menyimpan file ke `/home/meldir.id/storage/uploads/` dengan penamaan UUID unik.

### K. Akuntansi Perusahaan (Chart of Accounts & Double-Entry General Ledger) (`office.meldir.id`)
* `GET  /api/v1/office/accounting/coa` -> Mengambil daftar seluruh akun COA 5-digit beserta saldo berjalan terkini.
* `POST /api/v1/office/accounting/coa` -> Menambahkan sub-akun baru (Kode Akun, Nama, Kategori, Saldo Normal).
* `PUT  /api/v1/office/accounting/coa/:id` -> Memperbarui nama akun, deskripsi, atau status aktif/nonaktif.
* `GET  /api/v1/office/accounting/journals` -> Mengambil daftar jurnal umum dengan filter tanggal, sumber transaksi, dan pencarian memo.
* `POST /api/v1/office/accounting/journals` -> Entri jurnal umum manual (Validasi mutlak: `Sum(Debit) == Sum(Credit)`).
* `GET  /api/v1/office/accounting/ledger/:account_id` -> Buku Besar (General Ledger mutasi kartu akun per periode).
* `GET  /api/v1/office/accounting/reports/trial-balance` -> Neraca Saldo (Trial Balance) per periode tanggal.
* `GET  /api/v1/office/accounting/reports/profit-loss` -> Laporan Laba Rugi Komersial & Rekonsiliasi Fiskal (P&L).
* `GET  /api/v1/office/accounting/reports/balance-sheet` -> Neraca Keuangan (Aset = Kewajiban + Ekuitas).

### L. Beban Operasional & Kas Kecil (Corporate Expenses & Petty Cash) (`office.meldir.id`)
* `GET  /api/v1/office/expenses` -> Mengambil seluruh daftar beban operasional kantor dengan filter tanggal & kategori.
* `POST /api/v1/office/expenses` -> Mencatat beban baru (Vendor, Nomor Nota, Akun Beban, Akun Kas/Bank, PPN Masukan, Potongan Pajak, Upload Bukti Nota WebP/PDF, Flag Fiscal Deductible). Otomatis memicu entri jurnal berpasangan.
* `PUT  /api/v1/office/expenses/:id/approve` -> Persetujuan pengeluaran oleh Direktur Utama.
* `DELETE /api/v1/office/expenses/:id` -> Pembatalan pengeluaran beban & pembalikan otomatis entri jurnal (*reversing entry*).

### M. Hub Kepatuhan Perpajakan Terpadu (Tax Compliance Engine) (`office.meldir.id`)
* `GET  /api/v1/office/tax/vat` -> Rekapitulasi Pajak Pertambahan Nilai (PPN Keluaran & PPN Masukan) per masa pajak.
* `POST /api/v1/office/tax/vat/masukan` -> Mencatat Faktur Pajak Masukan vendor (NSFP, DPP, PPN 11%, upload PDF/QR).
* `GET  /api/v1/office/tax/vat/spt-masa` -> Kalkulasi formulir SPT Masa PPN 1111 (Keluaran - Masukan = Kurang/Lebih Bayar).
* `POST /api/v1/office/tax/vat/settle-ntpn` -> Menginput kode NTPN penyetoran PPN Kurang Bayar ke kas negara.
* `GET  /api/v1/office/tax/withholding` -> Mengambil daftar bukti potong PPh (PPh 21 Engineer & PPh 23 Klien/Vendor).
* `POST /api/v1/office/tax/withholding/bupot-21` -> Menerbitkan Bukti Potong PPh 21 honorarium engineer (Tarif TER / Pasal 17 & slip e-Bupot).
* `POST /api/v1/office/tax/withholding/kredit-23` -> Mencatat Bukti Potong PPh 23 dari klien (pengurang PPh Badan / Prepaid Tax).
* `GET  /api/v1/office/tax/corporate-estimate/:year` -> Mengambil simulasi SPT Tahunan PPh Badan 1771 (Omset, HPP, OPEX, Koreksi Fiskal Positif/Negatif, PKP, Fasilitas Pasal 31E UU HPP diskon 50%, Kredit Pajak PPh 23 & 25, PPh Kurang Bayar Pasal 29).
* `POST /api/v1/office/tax/corporate-estimate/:year/finalize` -> Mengesahkan dan mengunci laporan SPT Tahunan Badan 1771.

---

## 2. Background Workers (Golang Goroutines & Cron Scheduler)

### A. External Server Health Probe Worker (`internal/worker/health_probe.go`)
* **Interval**: Berjalan otomatis setiap 5 menit.
* **Mekanisme**: Melakukan HTTP GET/PING ke server external klien. Jika respons code `>= 500` atau `timeout > 10s`, ubah status menjadi `CRITICAL_DOWN`, rekam log incident, serta picu notifikasi lonceng in-app, WhatsApp darurat, dan PWA Push.

### B. SLA Breach Auto-Escalation Worker (`internal/worker/sla_escalator.go`)
* **Interval**: Berjalan setiap 10 menit.
* **Mekanisme**: Memeriksa tiket `open` / `in_progress` yang sisa waktu SLA-nya `< 30 menit` tanpa update dari engineer.
* **Tindakan**: Kirim notifikasi WA darurat ke Direktur dan Admin.

### C. Monthly Executive Report Generator Worker (`internal/worker/monthly_report.go`)
* **Jadwal**: Berjalan otomatis pada tanggal 1 setiap bulan (00:00 WIB).
* **Mekanisme**: Mengumpulkan metrik 30 hari ke belakang (Uptime %, durasi downtime, backup sukses, tiket terselesaikan), mengkalkulasi skor kesehatan sistem, menyimpan record ke `monthly_maintenance_reports`, dan menggenerasi dokumen PDF eksekutif.

### D. Monthly Quota Reset Worker (`internal/worker/quota_reset.go`)
* **Jadwal**: Berjalan otomatis pada tanggal 1 setiap bulan (00:01 WIB).
* **Mekanisme**: Menginisialisasi baris baru di `contract_monthly_quotas` dengan `minor_feature_quota_used: 0` dan `minor_feature_quota_total: 2` untuk seluruh kontrak aktif `Priority Care`.

### E. Redis Token Blacklist Middleware (`internal/delivery/http/middleware/auth.go`)
* Memeriksa token JWT 24-Jam. Jika token ID ada dalam Redis Blacklist (di-force logout oleh Admin), tolak request dengan HTTP 401 Unauthorized.

### F. Tax Deadline & Settlement Alert Worker (`internal/worker/tax_alert.go`)
* **Jadwal**: Berjalan setiap tanggal 8, 13, dan 25 setiap bulan (08:00 WIB).
* **Mekanisme**: Memeriksa tenggat jatuh tempo pembayaran PPh/PPN (Tanggal 10 dan 15) serta batas akhir pelaporan SPT Masa (Tanggal 20 dan akhir bulan). Mengirim notifikasi lonceng in-app dan pesan WhatsApp pengingat kepatuhan pajak ke Direktur dan Tim Finance.

### G. Daily General Ledger Balance Validator Worker (`internal/worker/ledger_validator.go`)
* **Jadwal**: Berjalan setiap hari pada pukul 00:05 WIB.
* **Mekanisme**: Memvalidasi seluruh jurnal aktif pada `accounting_journals`. Memastikan `Sum(total_debit) == Sum(total_credit)`. Jika terdeteksi ketidakseimbangan (*unbalanced journal entry*), otomatis membuat notifikasi audit kritis ke Direktur.

