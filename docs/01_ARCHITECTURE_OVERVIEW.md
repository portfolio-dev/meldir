# 🏛️ Arsitektur & Spesifikasi Sistem meldir.id (PostgreSQL Enterprise System)

> **Catatan Rahasia & Internal**: Dokumen ini merupakan cetak biru arsitektur lengkap platform meldir.id (PT. Melayani Digital Raya) yang mencakup 4 portal domain independen, 5 peranan pengguna (*RBAC*), sistem notifikasi lonceng in-app, broadcast push manual, modul media terlampir, **Universal Multi-Portal PWA Installation**, **7 Modul Bisnis Lanjutan**, serta **Sistem Akuntansi & Perpajakan Digital Perusahaan Terpadu (Double-Entry General Ledger, SAK EMKM COA 5-Digit, Kas Kecil/Petty Cash, PPN 11% Keluaran & Masukan, PPh 21/23/25, dan Estimasi SPT Tahunan Badan 1771)**.

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
│(Landing, SEO │               │  (Corporate  │                        │  (Engineers  │               │   (Clients   │
│ & Leads Form)│               │  Operations) │                        │  Workspace)  │               │   Portal)    │
└───────┬──────┘               └───────┬──────┘                        └───────┬──────┘               └───────┬──────┘
        │ (Install PWA)                │ (Install PWA)                         │ (Install PWA)                │ (Install PWA)
        └──────────────────────────────┴───────────────────┬───────────────────┴──────────────────────────────┘
                                                           │ (HTTPS API Requests /api/v1)
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
                               │  PostgreSQL Database  │       │  Redis Cache & Queue  │
                               └───────────────────────┘       └───────────────────────┘
```

---

## 2. 18 Alur Kerja Spesifik & Protokol Bisnis Sistem

### 📱 A. Dukungan PWA Universal (Installable di 4 Domain)
* Setiap subdomain memiliki konfigurasi Web App Manifest & Service Worker tersendiri:
  - **`meldir.id`**: PWA Public Portal (Landing Page, Portfolio & Form Audit).
  - **`office.meldir.id`**: PWA Office Admin App (Icon Khusus Office Admin untuk akses cepat Direktur/Admin di HP/Desktop).
  - **`jobs.meldir.id`**: PWA Engineer Work App (Icon Khusus Engineer Workspace untuk tracking tiket & task di HP).
  - **`portal.meldir.id`**: PWA Client Portal App (Icon Khusus Client Portal untuk monitoring proyek & tiket di HP).
* Tombol **"Install / Unduh Aplikasi PWA"** tersedia di sidebar dan header masing-masing dashboard.

### 👤 B. Manajemen Profil & Akun Dasar (Profile & Security Settings)
* Tersedia di ketiga portal (`office`, `jobs`, `portal`) untuk mengubah:
  - Nama Lengkap, Alamat Email Resmi, dan Nomor WhatsApp.
  - Foto Profil (Avatar Upload dengan kompresi otomatis WebP).
  - Ganti Password Aman (Wajib memasukkan kata sandi lama untuk verifikasi keamanan Argon2id).

### 🔔 C. Sistem Notifikasi Lonceng In-App (Interactive Bell & Deep-Linking)
* Header bar ketiga dashboard dilengkapi **Icon Lonceng (Notification Bell)**:
  - **Unread Badge Counter**: Menampilkan jumlah notifikasi belum dibaca secara real-time.
  - **Dropdown Pop-Up**: Menampilkan riwayat notifikasi (update tiket SLA, invoice baru, status server, kontrak, laporan bulanan).
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

### 🎯 K. Inbound Leads CRM Pipeline (`meldir.id` -> `office.meldir.id`)
* Formulir konsultasi dan audit gratis dari landing page otomatis masuk ke tabel `inbound_leads`.
* Admin Office mengelola pipeline visual (Papan Kanban):
  - `Baru` -> `Dihubungi` -> `Kirim RAB` -> `Menang (Kontrak Dibuat)` -> `Gagal/Batal`.
* Mencegah prospek potensial bernilai puluhan juta tercecer akibat chat WA tertimbun.

### 🧩 L. Modul Add-On & Saldo Jam Kerja Metered (`portal.meldir.id`)
* Klien dapat memesan paket pengembangan on-demand langsung dari portal:
  - *On-Demand Extension* (Rp 175.000/jam)
  - *Sprint Block 10 Jam* (Rp 1.500.000 / 60 hari aktif)
  - *Sprint Block 20 Jam* (Rp 2.800.000 / 90 hari aktif)
* Setelah invoice terbayar, saldo jam kerja (`remaining_hours`) aktif di akun klien dan berkurang otomatis setiap kali log kerja engineer di-submit.

### ⏱️ M. Tracker Kuota Fitur Ringan Bulanan (Priority Care 2x/Bulan)
* Pada form pembuatan tiket di `portal.meldir.id`, terdapat widget status kuota:  
  **"Kuota Fitur Ringan Bulan Ini: [ 1 / 2 Terpakai ]"**.
* Kuota otomatis me-reset ke 2/2 setiap tanggal 1 awal bulan.
* Jika kuota telah habis, sistem otomatis menyarankan pemakaian saldo jam kerja Add-On.

### 📝 N. Engineer Timesheet & Log Waktu Kerja Terukur (`jobs.meldir.id`)
* Setiap menyelesaikan tugas atau tiket, Engineer mengisi entri log jam kerja:
  - Input: Tiket ID, Jam yang Digunakan (e.g. 2.5 jam), Deskripsi Pekerjaan.
* Log ini memotong saldo Add-On klien secara transparan dan tampil di riwayat pemakaian jam kerja klien.

### 🧾 O. Lampiran e-Faktur Pajak PPN (11%)
* Pada modul invoice di `office.meldir.id`, Admin dapat mengunggah file **PDF e-Faktur resmi DJP** dan menginput Nomor Seri Faktur Pajak.
* Klien korporat dapat mengunduh invoice dan e-Faktur secara mandiri dari `portal.meldir.id` untuk pembukuan akuntansi legal.

### 🤝 P. Digital BAST (Berita Acara Serah Terima) Sign-off
* Saat proyek custom development mencapai milestone akhir, Admin menerbitkan dokumen BAST.
* Klien menandatangani BAST secara digital via kanvas web dan mengunggah versi e-Materai.
* Penandatanganan BAST menjadi prasyarat sah pelunasan termin akhir (100%) dan dimulainya masa garansi/pemeliharaan.

### 📊 Q. Laporan Kinerja Bulanan Eksekutif Otomatis (Monthly Executive Report)
* Worker background berjalan otomatis setiap tanggal 1 (00:00 WIB) untuk merangkum kinerja sistem per proyek selama 1 bulan penuh:
  - Persentase Uptime Server (e.g. 99.98%).
  - Total durasi downtime dan rekaman insiden.
  - Jumlah pencadangan (backup) database yang berhasil.
  - Jumlah tiket yang diselesaikan & pembaruan keamanan.
* Hasil laporan otomatis dikonversi ke PDF eksekutif yang dapat diunduh klien di `portal.meldir.id`, mengeliminasi alasan klien berhenti berlangganan (*anti-churn mechanism*).

### 🏛️ R. Sistem Akuntansi & Perpajakan Digital Terpadu Perusahaan (`office.meldir.id`)
Modul akuntansi operasional dan kepatuhan perpajakan (*Corporate Accounting & Tax Compliance*) yang dirancang khusus untuk operasional digital 100% PT. Melayani Digital Raya:
1. **Bagan Akun Standar (Standard Chart of Accounts / COA 5-Digit SAK EMKM)**:
   - Terstruktur penuh: `10000 Aset`, `20000 Kewajiban`, `30000 Ekuitas`, `40000 Pendapatan`, `50000 HPP / Beban Pokok`, `60000 Beban Operasional (OPEX)`.
2. **Mesin Penjurnalan Otomatis (Auto-Journaling Engine)**:
   - **Invoice Klien Terbayar**: Otomatis mendebit Kas/Bank (`11020`) & Kredit Pajak PPh 23 (`11050`), serta mengkredit Pendapatan Jasa (`41010/41020`) & Utang PPN Keluaran (`21030`).
   - **Pencairan Honor Engineer**: Otomatis mendebit Biaya Honor Proyek (`51010`), mengkredit Utang PPh 21 (`21040`), dan mengkredit Kas/Bank (`11020`).
   - **Pengeluaran Operasional**: Otomatis mendebit Akun Beban Terkait (`61010 - 61060`) & PPN Masukan (`11060`), serta mengkredit Kas Kecil/Bank.
3. **Beban Operasional & Kas Kecil Digital (Petty Cash & Digital Receipts)**:
   - Pencatatan seluruh pengeluaran kantor harian dengan bukti nota/kuitansi digital (WebP/PDF).
   - Klasifikasi biaya: *Fiscal Deductible* (diakui pajak) vs *Non-Deductible* (wajib koreksi positif).
4. **Hub Kepatuhan Pajak Pertambahan Nilai (PPN 11%)**:
   - **PPN Keluaran**: Integrasi invoice klien dengan Nomor Seri Faktur Pajak resmi (NSFP) DJP dan unggah PDF e-Faktur.
   - **PPN Masukan**: Pencatatan faktur pajak dari vendor hosting/cloud/tools untuk dikreditkan.
   - **Kalkulator SPT Masa PPN 1111**: Otomatis menghitung status PPN Kurang/Lebih Bayar per masa pajak dan pencatatan kode NTPN setoran negara.
5. **Hub Pajak Penghasilan (PPh 21, PPh 23, PPh 25, PPh Badan 1771)**:
   - **PPh 21 Engineer/Tenaga Ahli**: Perhitungan otomatis tarif TER / Pasal 17, penerbitan data e-Bupot 21, dan arsip bukti potong.
   - **PPh 23 Jasa TI**: Pelacakan kredit pajak atas pemotongan 2% oleh klien institusi (diakui sebagai aset uang muka pajak pengurang PPh Badan).
   - **Simulasi & Rekonsiliasi SPT Tahunan PPh Badan 1771**: Kalkulasi Laba Komersial -> Koreksi Fiskal Positif/Negatif -> Penghasilan Kena Pajak (PKP) -> Penerapan Fasilitas Pasal 31E UU HPP (diskon tarif 50% menjadi 11%) -> Pengurangan Kredit PPh 23 & PPh 25 -> PPh Kurang Bayar Pasal 29 akhir tahun.
6. **Laporan Finansial Real-Time Interaktif**:
   - Buku Besar (General Ledger), Neraca Saldo (Trial Balance), Laporan Laba Rugi Komersial & Fiskal (P&L), serta Neraca Keuangan (Balance Sheet) siap cetak dan ekspor.

