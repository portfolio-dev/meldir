# 🎨 Panduan & Cetak Biru Desain UI/UX Multi-Portal meldir.id

> **Catatan Rahasia & Internal**: Dokumen ini merinci cetak biru desain antarmuka (*User Interface*) dan pengalaman pengguna (*User Experience*) untuk 4 Portal Domain (`meldir.id`, `office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`), Base Layout 80% Compact Density, PWA Universal, dan 7 instrumen bisnis baru.

---

## 1. Tata Letak Desktop Header, Lonceng Notifikasi & Tombol Install PWA

Header dan Sidebar pada ketiga dashboard dilengkapi dengan **Tombol Unduh / Install PWA**:

```text
┌────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│ [Logo] │ [< Toggle] │ [Search Bar / Subdomain Badge]  [📱 Unduh PWA]  [🔔 (3)]  [📢]  [🌗 Theme] [👤] │ (Header Bar)
├────────┬───────────────────────────────────────────────────────────────────────────────────────────────┤
│ 📌 (h) │                                                                                               │
│ 📜     │                                                                                               │
│ 🎫     │                                  COMPACT MAIN CONTENT AREA                                    │
│ 🖥️     │                               (Base Scale 80% High Density UI)                                │
│ 🧾     │                                                                                               │
│ 📱 PWA │ ➔ (Tombol Unduh Aplikasi di Bagian Bawah Sidebar)                                             │
└────────┴───────────────────────────────────────────────────────────────────────────────────────────────┘
```

### Karakteristik Tombol Unduh PWA:
* **Sidebar Footer Placement**: Tombol "Unduh Aplikasi PWA" terpasang rapi di bagian bawah sidebar desktop.
* **Mobile Action Sheet Placement**: Menu unduh aplikasi juga tersedia di dalam Bottom Sheet "Lainnya" saat dibuka di HP.

---

## 2. Navigasi Mobile PWA (5-Item Bottom Bar & Bottom Sheet "Lainnya")

Tampilan pada perangkat seluler (HP) menggunakan **Native 5-Item Bottom Navigation Bar**:

```text
┌─────────────────────────────────────────┐
│                                         │
│          MOBILE CONTENT VIEW            │
│                                         │
├─────────────────────────────────────────┤
│  [🏠]   [🎫]   [🔔]   [🧾]   [☰]  │ (Bottom Nav Bar)
│ Home   Tiket  Notif  Invoice Lainnya   │
└─────────────────────────────────────────┘
```

---

## 3. Komponen Lampiran Media (Preview Gambar, Video, PDF)

* **Gambar**: Thumbnail preview interaktif dengan opsi perbesar (*lightbox modal*).
* **Video**: Pemutar video bawaan (*HTML5 Video Player*) untuk memutar rekaman bug atau video demo tanpa unduh file.
* **PDF**: Penampil dokumen PDF terintegrasi (*PDF Viewer*) untuk melihat kontrak, BAST, invoice, atau laporan bulanan langsung.

---

## 4. Visualisasi Desain 7 Modul Bisnis Baru

### 🎯 A. Papan Kanban Leads CRM (`office.meldir.id`)
* **5 Kolom Pipeline**: `Baru` (Slate) | `Dihubungi` (Blue) | `Kirim RAB` (Indigo) | `Menang/Kontrak` (Emerald) | `Batal` (Rose).
* **Kartu Prospek**: Nama PIC, Nama Perusahaan, Estimasi Anggaran, dan tombol aksi "Detail" & "Buat Kontrak SPK".

### 🧩 B. Saldo Jam Kerja & Katalog Add-On (`portal.meldir.id`)
* **Widget Speedometer / Progress Bar Saldo Jam**:
  - Menampilkan: `Sisa 7.5 Jam dari 10 Jam (Masa Aktif s/d 15 Des 2026)`.
  - Indikator Warna: Hijau (> 5 Jam), Kuning (2–5 Jam), Merah (< 2 Jam).
* **Kartu Pilihan Add-On**: 3 kartu rapi dengan tombol pemesanan instan yang otomatis menerbitkan tagihan invoice.

### ⏱️ C. Indikator Kuota Fitur Ringan Bulanan (`portal.meldir.id`)
* Terpasang di atas form pengajuan tiket SLA:
  - Kotak badge: `⚡ Kuota Fitur Ringan Bulan Ini: [ 1 / 2 Terpakai ] (Reset 1 Nov)`.
  - Jika kuota habis (2/2): Badge berubah menjadi amber dengan tombol switch "Gunakan Saldo Jam Add-On".

### 📝 D. Log Jam Kerja Engineer & Transparansi Klien (`jobs` & `portal`)
* **Portal Jobs**: Modal ringkas input tiket, durasi (slider 0.5 jam s/d 8 jam), dan rincian pekerjaan.
* **Portal Klien**: Tabel kronologis riwayat jam kerja yang transparan sehingga klien melihat jelas pemakaian saldo jamnya.

### 📊 E. Kartu Laporan Kinerja Bulanan Eksekutif (`portal.meldir.id`)
* **Kartu Ringkasan Periode**:
  - Menampilkan metrik utama: `Uptime 99.98%` | `4x Backup Sukses` | `3 Tiket Selesai` | `0 Insiden Keamanan`.
  - Tombol aksi utama: `📥 Unduh Laporan Eksekutif (PDF Resmi)`.

### 🤝 F. Canvas Tanda Tangan Digital & E-BAST Hub (`portal.meldir.id`)
* **Canvas iPad Style**: Pilihan Warna Tinta: **Hitam** (`#000000`) dan **Biru Tua** (`#1e3a8a`).
* Tombol: **"Bersihkan Canvas"** (*Clear*) dan **"Simpan & Lanjutkan E-Materai"** (*Save & Proceed*).
* Berlaku identik untuk pengesahan Kontrak SPK maupun Berita Acara Serah Terima (BAST).

### 🧾 G. Badge Unduh e-Faktur Pajak PPN 11% (`portal.meldir.id`)
* Pada tabel invoice yang berstatus `PAID`, jika admin telah mengunggah e-Faktur:
  - Tampil badge hijau: `📄 e-Faktur Tersedia (No: 010.002-26.XXXXXXXX)`.
  - Tombol unduh file PDF e-Faktur resmi untuk arsip pajak klien.

---

## 5. Visualisasi Desain Modul Akuntansi & Perpajakan Digital (`office.meldir.id`)

### 📊 H. Buku Besar & Jurnal Umum Berpasangan (General Ledger & Double-Entry Journal)
* **Tabel Jurnal Ringkas & Presisi (Compact 80% Density)**:
  - Kolom: `No. Jurnal` (Monospace), `Tanggal`, `Memo / Keterangan Transaksi`, `Akun Debit` & `Akun Kredit` (Kode 5-Digit & Nama Akun), `Nilai (Rp)`, `Status Posting`.
  - **Live Balance Validator Header**:
    - Jika Seimbang: Badge hijau emerald solid `✓ Jurnal Seimbang (Total: Rp 84.500.000)`.
    - Jika Tidak Seimbang: Badge merah berkedip `⚠️ PERINGATAN: Selisih Rp 2.500.000 (Debit != Credit)`.
* **Pencarian & Filter Multi-Kriteria**: Filter rentang tanggal, filter akun spesifik (misal: hanya Kas Kecil `11010` atau Bank Mandiri `11020`), dan filter sumber (`Invoice`, `Honor Engineer`, `Beban Kantor`, `Pajak`).

### 📑 I. Tampilan Laporan Keuangan Digital (Financial Statements Hub)
* **Navigasi 3 Tab Dinamis**:
  1. **Laba Rugi (Profit & Loss / P&L)**:
     - Struktur rapi: Pendapatan Usaha (41xxx) -> HPP Jasa & Cloud (51xxx) = Laba Kotor.
     - Beban Operasional OPEX (61xxx) = Laba Bersih Komersial.
     - Bagian Bawah: Panel Rekonsiliasi Fiskal (Koreksi Fiskal Positif & Negatif) = Penghasilan Neto Fiskal.
  2. **Neraca Keuangan (Balance Sheet)**:
     - Dua kolom berdampingan: **Aset** (Lancar & Tetap) vs **Kewajiban + Ekuitas Modal**.
     - Indikator otomatis di bawah: `Total Aset == Total Kewajiban & Ekuitas` (Status Validasi Hijau).
  3. **Neraca Saldo (Trial Balance)**:
     - Tabel komprehensif 5 kolom: `Kode Akun`, `Nama Akun`, `Saldo Awal`, `Debit`, `Kredit`, `Saldo Akhir`.
* **Toolbar Ekspor**: Tombol `📄 Cetak PDF Resmi Kop PT` dan `📊 Ekspor XLSX / CSV`.

### 🧾 J. Modal Pencatatan Beban Kantor & Kas Kecil (Petty Cash Entry)
* **Input Cepat & Komprehensif**:
  - Pilihan Akun Beban (Dropdown pintar dengan autocomplete nama akun, e.g. `61010 Operasional Kantor`, `61030 E-Materai`).
  - Akun Sumber Dana: Radio pill `Kas Kecil (11010)` | `Bank Mandiri PT (11020)` | `Bank BCA (11030)`.
  - Vendor & NPWP: Input nama vendor rekanan dan NPWP/NIK jika ada.
  - Komponen Pajak: Checkbox `Memiliki Faktur PPN 11%` (otomatis memisahkan DPP & PPN Masukan) dan Checkbox `Potong PPh 21/23` (menghitung utang pajak terpotong).
  - Klasifikasi Pajak: Switch `Biaya Diakui Fiskal (Deductible)` vs `Non-Deductible` (untuk bahan rekonsiliasi SPT 1771).
  - Dropzone Berkas Bukti: Upload struk/kuitansi/nota dengan auto-preview thumbnail WebP/PDF.

### 🏛️ K. Hub Kepatuhan Perpajakan Lengkap (Tax Compliance Hub)
* **Tab 1: Manajemen PPN 11% (SPT Masa PPN 1111)**:
  - Kartu Ringkasan Bulanan: `PPN Keluaran (Rp A)` - `PPN Masukan (Rp B)` = `Status: Kurang Bayar (Rp C)`.
  - Tabel Faktur Keluaran (sinkron langsung dari Invoice Klien lunas + Nomor Seri Faktur Pajak resmi DJP).
  - Tabel Faktur Masukan (input manual belanja server AWS, Google, Hostinger, perlengkapan IT).
  - Tombol aksi: `Setor & Input Kode NTPN` -> Status berubah menjadi `Lunas / Dilaporkan DJP`.
* **Tab 2: Hub Bukti Pemotongan Pajak (e-Bupot Unifikasi)**:
  - **Sub-Tab PPh 21 Tenaga Ahli**: Rekapitulasi honorarium engineer, tarif TER / Pasal 17, nilai pajak dipotong, tombol cetak slip e-Bupot resmi untuk engineer.
  - **Sub-Tab PPh 23 Klien**: Arsip berkas bukti potong 2% yang diterima dari klien korporat (otomatis terdata sebagai aset kredit pajak pengurang PPh Badan di akun `11050`).
* **Tab 3: Simulator & Proyeksi SPT Tahunan PPh Badan 1771**:
  - Widget Proaktif Proyeksi Pajak Tahunan:
    - Peredaran Bruto (Omset): Ditampilkan real-time dari seluruh invoice paid.
    - Laba Komersial & Koreksi Fiskal Positif/Negatif.
    - Penghasilan Kena Pajak (PKP).
    - **Fasilitas Pasal 31E UU HPP**: Diskon 50% tarif PPh Badan (efektif 11% jika omset di bawah Rp 4.8 Milyar).
    - Kredit Pajak: Dikurangi akumulasi bukti potong PPh 23 dari klien dan angsuran bulanan PPh 25.
    - Estimasi Akhir: PPh Kurang Bayar Pasal 29 yang harus disetor saat pelaporan SPT 1771 di bulan April, serta perkiraan besaran angsuran bulanan PPh 25 tahun berikutnya.

