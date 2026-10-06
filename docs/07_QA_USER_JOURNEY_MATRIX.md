# 📋 Matriks QA, Akses Pengguna & Interaksi Sistem meldir.id

> **Catatan Rahasia & Internal**: Dokumen ini merupakan laporan spesifikasi fungsional dan pengujian *Quality Assurance* (QA Matrix) yang merinci secara presisi apa saja yang dapat **Dilihat (View)**, **Diakses (Access)**, **Dipilih/Dioperasikan (Interact/Operate)**, serta batasan wewenang setiap jenis pengguna pada 4 domain platform `meldir.id` termasuk 17 alur kerja bisnis baru.

---

## 1. 🌐 Public Experience: `meldir.id` (Pengunjung Umum & Calon Klien)

### A. Apa Saja yang Dapat Dilihat (View)
* **Hero Section**: Headline representatif Managed IT & Development, subheadline keunggulan, badge status PT legal.
* **4 Layanan Komersial Utama**: Custom Web/App Dev, Managed IT Care, Scaling & Upgrade, Mobile/PWA Management.
* **Tabel Komparasi Solusi**: Komparasi 5 poin antara *Tanpa Tim Ahli / Mandiri (AI)* vs *Dalam Pengelolaan meldir.id*.
* **Seksi Skema Add-On & Matriks Fitur**: Rate card Hourly Rp 175rb, Sprint 10 Jam Rp 1.5jt, Sprint 20 Jam Rp 2.8jt, dan tabel batasan ruang lingkup kerja.
* **Seksi Transparansi Biaya Cloud & Kepemilikan Aset 100%**: Jaminan anti-vendor lock-in.
* **Seksi FAQ Layanan**: 5 pertanyaan terpopuler seputar pengambilalihan sistem, add-on, kontrak, dan faktur pajak.
* **Footer & Halaman Legal**: Identitas resmi PT, SK Kemenkum, dan halaman legal (`/terms`, `/privacy`, `/refund`, `/support`).

### B. Apa Saja yang Dapat Dipilih & Diinteraksikan (Interact / Select)
* **Bilingual Switcher**: Beralih bahasa instan antara Bahasa Indonesia (`ID`) dan Bahasa Inggris (`EN`).
* **Formulir Konsultasi & Audit Sistem Gratis**: Mengisi nama, email, nomor WA, dan kebutuhan sistem -> data otomatis masuk ke tabel `inbound_leads` di `office.meldir.id`.
* **Tombol Konsultasi WhatsApp Cepat**: Menghubungkan langsung ke WhatsApp resmi `+62 821-3173-357`.
* **PWA Install Button**: Menginstal aplikasi PWA `meldir.id` langsung ke homescreen HP/Desktop.

---

## 2. 🏢 Office Experience: `office.meldir.id` (Direktur & Admin Office)

### A. Apa Saja yang Dapat Dilihat (View)
* **Executive Summary Dashboard**: Total Revenue, status invoice, klien aktif, & tiket SLA.
* **Papan Pipeline Prospek (Leads Kanban)**: 5 kolom status (`Baru`, `Dihubungi`, `Kirim RAB`, `Menang/Kontrak`, `Batal`).
* **Katalog & Pesanan Add-On**: Rekap paket add-on aktif dan sisa saldo jam kerja per klien.
* **Rekapitulasi Timesheet Engineer**: Rincian jam kerja terpakai per engineer dan per proyek.
* **Icon Lonceng Notifikasi Header**: Badge counter merah live & dropdown pop-up list notifikasi.
* **Kartu Status WhatsApp API Engine**: Status indikator *Connected* atau *Disconnected* secara live.
* **Grid Monitoring Uptime Server External Klien**: Indikator warna live (Hijau: Uptime 99.9%, Merah: Server Down).
* **Bagan Akun Standar (COA 5-Digit)** & Buku Besar (General Ledger): Rincian debit/kredit per akun SAK EMKM.
* **Laporan Finansial Resmi**: Laba Rugi Komersial & Fiskal (P&L), Neraca Keuangan (Balance Sheet), dan Neraca Saldo (Trial Balance).
* **Hub Pajak Terpadu (Tax Hub)**:
  - Monitoring PPN 11% (Faktur Pajak Keluaran & Masukan, SPT Masa PPN 1111 Kurang/Lebih Bayar).
  - Hub Bukti Potong PPh 21 Engineer & Rekap Kredit Pajak PPh 23 Klien.
  - Simulator & Rekonsiliasi SPT Tahunan PPh Badan 1771 (Fasilitas Pasal 31E UU HPP, Kredit PPh 23/25, PPh Kurang Bayar Pasal 29).
* **Vault Kredensial Server**, Tabel Sesi Aktif, & Security Audit Logs.

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Manajemen Pipeline Leads**: Memindahkan kartu prospek, menambah catatan follow-up, dan klik "Konversi Menjadi Klien & Buat SPK".
* **Penerbitan Kontrak SPK & Dokumen BAST**: Membuat draf SPK, menerbitkan draf BAST di akhir proyek, dan verifikasi E-Materai.
* **Pelampiran e-Faktur Pajak PPN**: Menginput nomor seri e-Faktur dan mengunggah berkas PDF faktur resmi DJP ke invoice klien.
* **Operasional Akuntansi & Jurnal Umum**:
  - Input entri jurnal umum manual (validasi ketat `Debit == Credit`).
  - Pencatatan Beban Operasional & Kas Kecil (Petty Cash) dengan bukti struk WebP/PDF dan klasifikasi fiskal (*Deductible* vs *Non-Deductible*).
  - Ekspor dokumen laporan keuangan resmi ber-kop surat PT (PDF) dan format spreadsheet (CSV/XLSX).
* **Manajemen Kepatuhan Perpajakan**:
  - Catat Faktur Pajak Masukan vendor cloud/server untuk mengkreditkan PPN.
  - Input kode NTPN penyetoran kas negara untuk PPN dan PPh.
  - Terbitkan Bukti Potong e-Bupot PPh 21 atas pembayaran honor engineer.
  - Unggah dan catat Bukti Potong PPh 23 dari klien korporat sebagai aset kredit pajak (Prepaid Tax).
  - Simulasikan dan kunci perhitungan SPT Tahunan PPh Badan 1771.
* **Manajemen Katalog Add-On**: Tambah/ubah paket jam kerja (tarif, durasi jam, masa aktif hari).
* **Kirim Push Notifikasi Manual (Broadcast)**: Membuka modal broadcast untuk kirim push notifikasi ke Semua User / Klien / Engineer.
* **Verifikasi Pembayaran Invoice & Kontrol WhatsApp**: Verifikasi transfer bank, scan live QR WhatsApp, disconnect session.
* **Force Logout Sesi**: Menekan tombol Force Logout untuk mencabut sesi token aktif pengguna dari jarak jauh.

---

## 3. ⚙️ Jobs Experience: `jobs.meldir.id` (Engineers Internal & Eksternal)

### A. Apa Saja yang Dapat Dilihat (View)
* **Papan Tugas Developer & SLA Timer**: Daftar tugas & countdown waktu SLA (Merah Berkedip jika kritis < 2 Jam).
* **Icon Lonceng Notifikasi Header**: Notifikasi tugas baru, komentar tiket, dan status pencairan honorarium.
* **Panel Repositori GitHub**: Tautan repositori GitHub proyek terkait beserta status izin *collaborator push access*.
* **Lampiran Media Bug**: Gambar screenshot error, video rekaman bug, atau file log PDF dari klien.
* **Panel Bukti Potong Pajak PPh 21**: Rincian pemotongan pajak honorarium dan tombol unduh berkas resmi bukti potong e-Bupot 21 untuk SPT Tahunan Pribadi engineer.

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Unduh / Install PWA Jobs**: Mengunduh aplikasi PWA `jobs.meldir.id` ke HP engineer.
* **Input Timesheet Log Waktu Kerja**: Mengisi modal log jam kerja per tiket (Durasi jam + deskripsi pekerjaan terukur).
* **Update Status & Diskusi Tiket SLA**: Ubah status pekerjaan dan kirim pesan balasan dengan lampiran file media.
* **Unduh Bukti Potong PPh 21**: Mengunduh berkas slip bukti potong pajak penghasilan yang telah diterbitkan perusahaan.
* **Push ke GitHub**: Melakukan `git push` ke repositori proyek selama tugas berstatus `in_progress`.

---

## 4. 💼 Client Experience: `portal.meldir.id` (Klien Perorangan, PT, & Yayasan)

### A. Apa Saja yang Dapat Dilihat (View)
* **Client Control Center & Live Server Uptime**: Status proyek & grafik monitoring server external.
* **Widget Saldo Jam Kerja Add-On**: Sisa jam aktif terukur (e.g. `Sisa 7.5 dari 10 Jam`).
* **Widget Kuota Fitur Ringan Bulanan**: Indikator kuota gratis `[ 1 / 2 Terpakai ]` untuk paket Priority Care.
* **Tabel Riwayat Transparansi Jam Kerja**: Log pekerjaan riil engineer yang memotong saldo Add-On.
* **Hub Laporan Kinerja Bulanan Eksekutif**: Arsip PDF laporan bulanan (Uptime %, backup sukses, tiket selesai).
* **Pusat Tagihan & e-Faktur**: Invoice resmi, status bayar, rekening PT, dan tombol unduh e-Faktur PPN (PDF).

### B. Apa Saja yang Dapat Dipilih & Dioperasikan (Interact / Operate)
* **Unduh / Install PWA Client Portal**: Mengunduh aplikasi PWA `portal.meldir.id` ke smartphone klien.
* **Pesan Paket Add-On Jam Kerja**: Memilih paket jam kerja langsung dari portal untuk fitur non-ringan.
* **Buat Tiket SLA dengan Opsi Fitur Ringan**: Membuat tiket perbaikan/fitur dengan attachment media.
* **Tanda Tangan Kontrak SPK & Dokumen BAST**: Menandatangani draf SPK & BAST secara digital pada kanvas iPad style & upload PDF E-Materai.
* **Upload Bukti Transfer Bank**: Unggah foto/file bukti transfer untuk diverifikasi admin.
* **Unduh Laporan Bulanan & e-Faktur**: Unduh dokumen PDF resmi kapan saja secara mandiri.
* **Wizard Offboarding**: Pengajuan pemutusan layanan & unduh PDF Ringkasan Kredensial.

---

## 5. 👁️ Audit Role Experience (Third-Party Watcher / Reviewer)
* **Cakupan Akses**: Login ke seluruh dashboard (`office`, `jobs`, `portal`) dengan status wewenang **READ-ONLY**.
* **Batasan**: Seluruh tombol mutasi data (Simpan, Edit, Hapus, Kirim, TTD, Force Logout, Upload, Broadcast Push, Order Add-On, Input Timesheet) dinonaktifkan otomatis.
