# 📋 Matriks QA, Akses Pengguna & Interaksi Sistem meldir.id

> **Catatan Rahasia & Internal**: Dokumen ini merupakan laporan spesifikasi fungsional dan pengujian *Quality Assurance* (QA Matrix) yang merinci secara presisi apa saja yang dapat **Dilihat (View)**, **Diakses (Access)**, **Dipilih/Dioperasikan (Interact/Operate)**, serta batasan wewenang setiap jenis pengguna pada 4 domain platform `meldir.id` termasuk fitur instalasi PWA di seluruh portal.

---

## 1. 🌐 Public Experience: `meldir.id` (Pengunjung Umum & Calon Klien)

### A. Apa Saja yang Dapat Dilihat (View)
* **Hero Section**: Headline representatif Managed IT & Development, subheadline keunggulan, badge status PT legal.
* **4 Layanan Komersial Utama**: Custom Web/App Dev, Managed IT Care, Scaling & Upgrade, Mobile/PWA Management.
* **Tabel Komparasi Solusi**: Komparasi 5 poin antara *Tanpa Tim Ahli / Mandiri (AI)* vs *Dalam Pengelolaan meldir.id* (termasuk efisiensi biaya SDM hingga 70%).
* **Banner Free Audit Keamanan & Performa**: Call-to-action gratis analisa sistem eksisting.
* **Pilar Keamanan & Modal Detail**: 6 pilar keamanan data lengkap dengan modal penjelasannya.
* **Fitur & Teknologi Modern (2026)**: Grid interaktif fitur (AI Vision OCR, AI Code Audit, Multi-Tenant Cloud, Wasm Native Core, 24/7 Auto-Scaling, Immutable Log).
* **Mockup Device Interaktif**: Simulasi antarmuka Desktop SaaS, Tablet Adaptive, Mobile PWA, dan Smartwatch OS.
* **Alur Kerja Proyek & Program Yayasan**: 5 tahapan kerja & program pro-bono sosial.
* **Footer & Halaman Legal**: Identitas resmi PT, SK Kemenkum, dan halaman legal (`/terms`, `/privacy`, `/refund`, `/support`).

### B. Apa Saja yang Dapat Dipilih & Diinteraksikan (Interact / Select)
* **Bilingual Switcher**: Beralih bahasa instan antara Bahasa Indonesia (`ID`) dan Bahasa Inggris (`EN`).
* **Filter Kategori Fitur**: Memilih filter tab (*Semua, Mobile & PWA, Otomatisasi & Chat, Transaksi & Lokasi, Manajemen Data*).
* **Tombol Konsultasi WhatsApp & Program Sosial**: Menghubungkan langsung ke WhatsApp resmi `+62 821-3173-357`.
* **PWA Install Button**: Menginstal aplikasi PWA `meldir.id` langsung ke homescreen HP/Desktop.

---

## 2. 🏢 Office Experience: `office.meldir.id` (Direktur & Admin Office)

### A. Apa Saja yang Dapat Dilihat (View)
* **Executive Summary Dashboard**: Total Revenue, status invoice, klien aktif, & tiket SLA.
* **Icon Lonceng Notifikasi Header**: Badge counter merah live & dropdown pop-up list notifikasi.
* **Kartu Status WhatsApp API Engine**: Status indikator *Connected* atau *Disconnected* secara live.
* **Grid Monitoring Uptime Server External Klien**: Indikator warna live (Hijau: Uptime 99.9%, Merah: Server Down).
* **Papan Pipeline Klien (Kanban)** & Manajemen Kontrak E-Materai.
* **Manajemen Keuangan PT** & Form Manual Payout Engineer (termasuk Rp 0).
* **Vault Kredensial Server**, Tabel Sesi Aktif, & Security Audit Logs.
* **Lampiran Media Proyek/Tiket**: Preview gambar (lightbox), video player rekaman bug, dan penampil dokumen PDF.

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Unduh / Install PWA Office**: Mengunduh aplikasi PWA `office.meldir.id` ke desktop / smartphone admin.
* **Manajemen Akun & Profil Pribadi**: Ubah Nama, Email, WhatsApp, Foto Profil (Avatar), dan Ganti Password.
* **Kirim Push Notifikasi Manual (Broadcast)**: Membuka modal broadcast untuk kirim push notifikasi ke Semua User / Klien / Engineer.
* **Manajemen Lonceng Notifikasi**: Klik tombol **"Tandai Sudah Dibaca Semua"** atau klik item notifikasi untuk langsung menuju halaman terkait.
* **Manajemen Akun Klien & Contract Builder**: Buat akun baru, terbitkan kontrak TTD canvas & verifikasi E-Materai.
* **Kontrol WhatsApp Engine**: Scan QR Code live, Putuskan Sesi, Bersihkan Cache WA.
* **Verifikasi Pembayaran & Offboarding**: Verifikasi bukti transfer bank & kirim penawaran banding (*counter-offer*) retensi.
* **Force Logout Sesi**: Menekan tombol Force Logout untuk mencabut sesi token aktif pengguna dari jarak jauh.

---

## 3. ⚙️ Jobs Experience: `jobs.meldir.id` (Engineers Internal & Eksternal)

### A. Apa Saja yang Dapat Dilihat (View)
* **Papan Tugas Developer & SLA Timer**: Daftar tugas & countdown waktu SLA (Merah Berkedip jika kritis).
* **Icon Lonceng Notifikasi Header**: Notifikasi tugas baru, komentar tiket, dan status pencairan honorarium.
* **Panel Repositori GitHub**: Tautan repositori GitHub proyek terkait beserta status izin *collaborator push access*.
* **Lampiran Media Bug**: Gambar screenshot error, video rekaman bug, atau file log PDF dari klien.

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Unduh / Install PWA Jobs**: Mengunduh aplikasi PWA `jobs.meldir.id` ke HP engineer untuk notifikasi dan update tugas bergerak.
* **Manajemen Akun & Profil Pribadi**: Ubah Nama, Email, WhatsApp, Foto Profil, dan Ganti Password.
* **Lonceng Notifikasi**: Klik notifikasi untuk langsung membuka tiket yang di-assign.
* **Update Status & Diskusi Tiket**: Ubah status pekerjaan dan kirim pesan balasan dengan lampiran file media.
* **Push ke GitHub**: Melakukan `git push` ke repositori proyek selama tugas berstatus `in_progress`.

---

## 4. 💼 Client Experience: `portal.meldir.id` (Klien Perorangan, PT, & Yayasan)

### A. Apa Saja yang Dapat Dilihat (View)
* **Client Control Center & Live Server Uptime**: Status proyek & grafik monitoring server external.
* **Icon Lonceng Notifikasi Header**: Notifikasi progres tiket SLA, invoice terbit, atau pengingat E-Materai.
* **Hub Kontrak & Dokumen**: Kontrak resmi yang telah dibubuhi E-Materai.
* **Pusat Tagihan & Invoice**: Riwayat invoice & rekening tujuan PT.

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Unduh / Install PWA Client Portal**: Mengunduh aplikasi PWA `portal.meldir.id` ke smartphone klien untuk monitoring proyek secara instan.
* **Manajemen Akun & Profil Pribadi**: Ubah Nama, Email, WhatsApp, Foto Profil, dan Ganti Password.
* **Buat Tiket SLA dengan Lampiran Media**: Isi form kendala & unggah screenshot gambar/video rekaman kendala teknis.
* **Tanda Tangan Kontrak Online & E-Materai**: TTD canvas iPad style & upload ulang PDF E-Materai.
* **Upload Bukti Transfer Bank**: Unggah foto/file bukti transfer untuk diverifikasi admin.
* **Wizard Offboarding**: Pengajuan pemutusan layanan & unduh PDF Ringkasan Kredensial.

---

## 5. 👁️ Audit Role Experience (Third-Party Watcher / Reviewer)
* **Cakupan Akses**: Login ke seluruh dashboard (`office`, `jobs`, `portal`) dengan status wewenang **READ-ONLY**.
* **Batasan**: Seluruh tombol mutasi data (Simpan, Edit, Hapus, Kirim, TTD, Force Logout, Upload, Broadcast Push) dinonaktifkan otomatis.
