# 🏛️ Arsitektur & Spesifikasi Sistem meldir.id (PostgreSQL Enterprise System)

> **Catatan Rahasia & Internal**: Dokumen ini merupakan cetak biru arsitektur lengkap platform meldir.id (PT. Melayani Digital Raya) yang mencakup 3 portal domain independen, 5 peranan pengguna (*RBAC*), sistem notifikasi lonceng in-app, broadcast push manual, modul media terlampir, dan dukungan **Universal Multi-Portal PWA Installation**.

---

## 1. Topologi Multi-Domain & Database PostgreSQL Engine

```text
                               ┌────────────────────────────────────────────────────────┐
                               │             Public Visitor / Client Access             │
                               └───────────────────────────┬────────────────────────────┘
                                                           │
        ┌──────────────────────────────┬───────────────────┴───────────────────┬──────────────────────────────┐
        │                              │                                       │                              │
        ▼                              ▼                                       ▼                              ▼
┌──────────────┐               ┌──────────────┐                        ┌──────────────┐               ┌──────────────┐
│  meldir.id   │               │office.meldir │                        │ jobs.meldir  │               │portal.meldir │
│ (Landing/SEO)│               │  (Office)    │                        │  (Engineers) │               │  (Clients)   │
└───────┬──────┘               └───────┬──────┘                        └───────┬──────┘               └───────┬──────┘
        │ (Install PWA)                │ (Install PWA)                         │ (Install PWA)                │ (Install PWA)
        └──────────────────────────────┴───────────────────┬───────────────────┴──────────────────────────────┘
                                                           │ (HTTPS API Requests)
                                                           ▼
                                       ┌───────────────────────────────────────┐
                                       │     OpenLiteSpeed (CyberPanel VPS)    │
                                       └───────────────────┬───────────────────┘
                                                           │ (Proxy Pass /api)
                                                           ▼
                                       ┌───────────────────────────────────────┐
                                       │     Golang RESTful API Service        │
                                       │       (Listening at port :8080)       │
                                       └───────────────────┬───────────────────┘
                                                           │
                                           ┌───────────────┴───────────────┐
                                           ▼                               ▼
                               ┌───────────────────────┐       ┌───────────────────────┐
                               │ PostgreSQL Database   │       │ Redis Cache & Queue   │
                               └───────────────────────┘       └───────────────────────┘
```

---

## 2. 10 Alur Kerja Spesifik & Protokol Keamanan Sistem

### 📱 A. Dukungan PWA Universal (Installable di 4 Domain)
* Setiap subdomain memiliki konfigurasi Web App Manifest & Service Worker tersendiri:
  - **`meldir.id`**: PWA Public Portal (Landing Page & Portfolio).
  - **`office.meldir.id`**: PWA Office Admin App (Icon Khusus Office Admin untuk akses cepat Direktur/Admin di HP/Desktop).
  - **`jobs.meldir.id`**: PWA Engineer Work App (Icon Khusus Engineer Workspace untuk tracking tiket & task di HP).
  - **`portal.meldir.id`**: PWA Client Portal App (Icon Khusus Client Portal untuk monitoring proyek & tiket di HP).
* Tombol **"Install / Unduh Aplikasi PWA"** tersedia di sidebar dan header masing-masing dashboard.

### 👤 B. Manajemen Profil & Akun Dasar (Profile & Security Settings)
* Tersedia di ketiga portal (`office`, `jobs`, `portal`) untuk mengubah:
  - Nama Lengkap, Alamat Email Resmi, dan Nomor WhatsApp.
  - Foto Profil (Avatar Upload dengan kompresi otomatis WebP).
  - Ganti Password Aman (Wajib memasukkan kata sandi lama untuk verifikasi keamanan).

### 🔔 C. Sistem Notifikasi Lonceng In-App (Interactive Bell & Deep-Linking)
* Header bar ketiga dashboard dilengkapi **Icon Lonceng (Notification Bell)**:
  - **Unread Badge Counter**: Menampilkan jumlah notifikasi belum dibaca secara real-time.
  - **Dropdown Pop-Up**: Menampilkan riwayat notifikasi (update tiket SLA, invoice baru, status server, kontrak).
  - **Tombol "Tandai Sudah Dibaca Semua"** (*Mark All as Read*).
  - **Deep-Linking Redirection**: Mengklik notifikasi langsung mengarahkan pengguna ke halaman spesifik terkait.

### 📢 D. Mesin Broadcast Push Notifikasi Manual (Khusus `office.meldir.id`)
* Fitur di dashboard Direktur/Admin untuk mengirimkan Push Notification manual ke perangkat pengguna:
  - Input Judul (*Title*), Pesan (*Body*), dan Link Tujuan (*Target URL*).
  - Pilihan Target Audiens: **Semua Pengguna**, **Hanya Klien**, **Hanya Engineer**, atau **Hanya Tim Admin**.

### 📎 E. Manajemen Lampiran Media Proyek & Tiket (Gambar, Video, PDF)
* Sistem mendukung pengunggahan berkas multi-format pada modul:
  - **Tiket SLA & Balasan Chat**: Unggah screenshot gambar (PNG/JPG/WEBP), video rekaman bug (MP4/WebM), atau log PDF.
  - **Milestone Proyek**: Unggah dokumentasi PDF rilis fitur, gambar demo, dan video panduan operasional.
  - **Kontrak & Tagihan**: Unggah berkas E-Materai PDF dan foto bukti transfer bank.

### 📜 F. Alur Kontrak Online & E-Materai
1. **Drafting & TTD Canvas**: Admin `office.meldir.id` menerbitkan kontrak. Klien melanggam TTD digital melalui *canvas web* di `portal.meldir.id`.
2. **Generasi PDF Awal**: Sistem menggabungkan isi kontrak & TTD canvas menjadi dokumen PDF awal.
3. **E-Materai Stamping**: Klien mengunduh PDF, membubuhkan **E-Materai resmi**, dan mengunggah (*re-upload*) kembali PDF final yang telah dibubuhi E-Materai ke portal.

### 💬 G. Manajemen WhatsApp API Engine (Unofficial Card)
* Pada dashboard `office.meldir.id`, terdapat **Kartu Status WA API**:
  - **Status Indicator**: Menampilkan status *Connected* / *Disconnected*.
  - **QR Code Scanner**: Jika *Disconnected*, menampilkan barcode live untuk mentautkan nomor WhatsApp pengirim PT.
  - **Tindakan Pemulihan**: Tombol **"Putuskan Sesi"** dan **"Bersihkan Cache WA API"** untuk mengatasi kendala koneksi tanpa perlu me-restart server.

### 🔐 H. Isolasi Keamanan Akses Engineer (GitHub Collaborative Only)
* **Zero Direct Server Access**: Engineer (internal & eksternal) **TIDAK PERNAH** diberikan akses kredensial server/hosting/database secara langsung. Kredensial server tersimpan aman di vault `office.meldir.id`.
* **GitHub Integration**: Engineer hanya diberikan akses kolaborator pada repositori GitHub proyek yang ditugaskan selama status proyek `in_progress`.

### 💰 I. Skema Bagi Hasil Engineer Fleksibel (Manual Payout Admin)
* Form pembayaran komisi/honorarium engineer diisi manual oleh Admin Office.
* Mendukung alokasi nilai dari **Rp 0** (khusus proyek sosial/yayasan pro-bono jika engineer bersedia) hingga nilai kustom per proyek/milestone.
* Sistem menghitung dan mengakumulasi total pengeluaran honorarium per engineer secara otomatis di dashboard keuangan PT.

### 🚫 J. Alur Offboarding Klien Ringan & Penawaran Banding (Retention Flow)
1. **Inisiasi Offboarding**: Klien mengajukan pemutusan layanan di `portal.meldir.id`.
2. **Penawaran Banding (*Counter-Offer*)**: Pihak Office dapat memberikan penawaran banding (diskon/penyesuaian scope) untuk mempertahankan klien. Jika tidak disepakati atau Office tidak sanggup, permohonan disetujui.
3. **Verifikasi Ulang**: Klien memverifikasi email & nomor telepon aktif.
4. **Transfer GitHub & Generasi PDF Kredensial**:
   - Sistem mentransfer kepemilikan repositori GitHub ke akun GitHub milik klien.
   - Sistem menggenerasi **Dokumen PDF Kredensial Server Lengkap** terenkripsi untuk diunduh klien.
   - *Tidak ada pengiriman file berat (.zip/dump)* agar proses berjalan sangat efisien dan ringan.
