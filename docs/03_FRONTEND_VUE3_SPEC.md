# 🎨 Spesifikasi & Frontend Components Vue 3 + Vite Multi-Portal

Dokumen ini merinci komponen antarmuka Vue 3 (Composition API + TypeScript/JavaScript + Tailwind CSS) untuk 4 portal domain (`meldir.id`, `office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`), mencakup Base Shell, Pinia Stores, dan 17 modul fungsional.

---

## 1. Komponen Antarmuka Kunci (Universal & Portal-Specific)

### A. Universal Components (Digunakan di Seluruh Portal)
* **`PWAInstallPrompt.vue`**:
  - Tombol **"Install / Unduh Aplikasi PWA"** di sidebar desktop dan action sheet mobile.
  - Menangani event `beforeinstallprompt` untuk mengunduh PWA sesuai domain aktif.
* **`NotificationBellDropdown.vue`**:
  - Ikon lonceng header dengan unread badge count real-time.
  - Dropdown daftar notifikasi, tombol "Tandai Dibaca Semua", dan deep-link router navigation.
* **`ProfileSettingsModal.vue`**:
  - Form Nama, Email, WhatsApp, Avatar WebP Uploader, dan Ganti Password Argon2id.
* **`MultiMediaUploader.vue`**:
  - Komponen drag-and-drop pendukung Gambar (5MB), Video (50MB), dan Dokumen PDF (20MB) lengkap dengan baris progress upload.
* **`DocumentSignatureCanvas.vue`**:
  - Kanvas web penandatanganan digital (*iPad drawing pad style*) dengan pilihan tinta hitam (`#000000`) dan biru tua (`#1e3a8a`), tombol clear, dan export PNG transparan.

### B. Komponen Khusus `office.meldir.id` (Corporate Operations)
* **`LeadsPipelineKanban.vue`**:
  - Papan Kanban drag-and-drop untuk prospek baru (`Baru`, `Dihubungi`, `Kirim RAB`, `Menang`, `Gagal`).
  - Modal detail prospek dengan tombol cepat "Konversi Menjadi Klien & Buat SPK".
* **`AddonPackageManager.vue`**:
  - Pengelolaan katalog paket Add-On (ubah harga, kuota jam, masa aktif hari, deskripsi).
* **`WhatsAppStatusCard.vue`**:
  - Kartu pemantau koneksi WhatsApp API, live QR Code Stream, tombol Disconnect, dan Clear Session.
* **`TaxInvoiceUploadModal.vue`**:
  - Modal input Nomor Seri e-Faktur Pajak dan upload file PDF faktur resmi DJP untuk dilampirkan ke invoice klien.
* **`BASTBuilderModal.vue`**:
  - Form penerbitan draf Berita Acara Serah Terima (BAST) per milestone proyek.
* **`ExecutiveFinanceChart.vue`**:
  - Grafik *Smooth Area Wave Chart* untuk tren pemasukan invoice vs pengeluaran honorarium engineer.
* **`GeneralLedgerTable.vue`**:
  - Tabel interaktif Buku Besar & Jurnal Umum Berpasangan (Double-Entry).
  - Filter rentang tanggal, pencarian memo, pemilahan akun COA, dan validasi seimbang (Badge Hijau: `Balanced` vs Badge Merah: `Unbalanced`).
* **`FinancialReportStatements.vue`**:
  - Tampilan tab dinamis Laporan Finansial Resmi:
    1. **Laba Rugi (Profit & Loss)**: Pendapatan Jasa, HPP Jasa, Laba Kotor, Beban Operasional OPEX, Laba Bersih Komersial, dan Rekonsiliasi Fiskal.
    2. **Neraca Keuangan (Balance Sheet)**: Total Aset, Total Kewajiban, Ekuitas Modal & Saldo Laba Ditahan.
    3. **Neraca Saldo (Trial Balance)**: Ringkasan Saldo Awal, Mutasi Debit, Mutasi Kredit, dan Saldo Akhir akun 5-digit.
  - Tombol cetak PDF resmi kop surat PT dan ekspor spreadsheet.
* **`CorporateExpenseModal.vue`**:
  - Modal entri beban kantor harian & petty cash digital.
  - Input Vendor, NPWP, No. Nota, Akun Beban (6xxxx), Akun Pembayar (Kas/Mandiri/BCA), PPN Masukan (11%), Potongan Pajak, Upload Berkas Bukti (WebP/PDF), dan toggle Biaya Fiskal (*Deductible*).
* **`VATManagementHub.vue`**:
  - Hub Kepatuhan PPN 11%: Rekap Faktur Pajak Keluaran (dari invoice klien) vs Faktur Pajak Masukan (dari vendor server/cloud/tools).
  - Kalkulator SPT Masa PPN 1111: Otomatis menghitung status PPN Kurang/Lebih Bayar per bulan dan form input kode NTPN setoran negara.
* **`WithholdingTaxBupotHub.vue`**:
  - Hub Bukti Potong Pajak:
    1. **PPh 21 Tenaga Ahli**: Perhitungan tarif TER / Pasal 17 atas honorarium engineer, cetak slip bukti potong, dan input NTPN.
    2. **PPh 23 Jasa TI**: Arsip kredit pajak pemotongan 2% oleh klien (pengurang PPh Badan) & pemotongan vendor pihak ketiga.
* **`AnnualTaxSimulator1771.vue`**:
  - Simulator interaktif SPT Tahunan PPh Badan 1771:
    - Input koreksi fiskal positif/negatif.
    - Kalkulasi fasilitas diskon tarif 50% Pasal 31E UU HPP (tarif efektif 11% untuk omset s/d Rp 4.8 Milyar).
    - Pemotongan kredit pajak PPh 23 & setoran angsuran PPh 25.
    - Proyeksi PPh Kurang Bayar Pasal 29 dan angsuran bulanan PPh 25 tahun berikutnya.

### C. Komponen Khusus `portal.meldir.id` (Client Control Center)
* **`AddonCatalogCard.vue` & `HoursBalanceWidget.vue`**:
  - Kartu pilihan paket Add-On (On-Demand, Sprint 10 Jam, Sprint 20 Jam) dengan tombol order instan.
  - Widget speedometer/bar ringkasan saldo jam kerja aktif (`sisa 7.5 dari 10 jam`).
* **`MinorFeatureQuotaBadge.vue`**:
  - Indikator visual di form pembuatan tiket:  
    `[ Kuota Fitur Ringan Bulan Ini: 1/2 Terpakai ]` (reset otomatis tiap awal bulan).
* **`TimesheetHistoryTable.vue`**:
  - Tabel transparansi jam kerja teknis engineer yang memotong saldo Add-On klien (Tanggal, Engineer, Durasi Jam, Rincian Pekerjaan).
* **`MonthlyReportViewer.vue`**:
  - Hub arsip Laporan Kinerja Bulanan Eksekutif (Uptime %, durasi downtime, backup log, tiket selesai) dengan tombol unduh PDF resmi.
* **`TaxInvoiceDownloadBtn.vue`**:
  - Tombol unduh berkas e-Faktur Pajak (PPN 11%) resmi langsung dari baris tabel invoice.
* **`BASTSigningModal.vue`**:
  - Modal review draf BAST, penandatanganan canvas digital, dan re-upload PDF ber-e-Materai.

### D. Komponen Khusus `jobs.meldir.id` (Engineer Workspace)
* **`TimesheetEntryModal.vue`**:
  - Form pencatatan waktu pengerjaan tugas/tiket (Input Tiket ID, Durasi Jam, Catatan pekerjaan).
* **`SLACountdownTimer.vue`**:
  - Widget hitung mundur sisa waktu penanganan tiket SLA (Warna dinamis: Hijau > 6 Jam, Kuning 2–6 Jam, Merah berkedip < 2 Jam).
* **`GitHubRepoCard.vue`**:
  - Kartu info repositori GitHub tempat engineer diundang sebagai kolaborator aktif.
* **`EngineerTaxWithholdingSlipModal.vue`**:
  - Modal bagi engineer untuk melihat rincian pemotongan PPh 21 dan mengunduh berkas e-Bupot resmi bukti setor pajak mereka.

---

## 2. Struktur Pinia State Management Stores

* **`authStore.ts`**: Menyimpan token JWT, data user profil aktif, role, dan izin RBAC.
* **`leadStore.ts`**: Menyimpan data pipeline prospek masuk, filter status, dan fungsi update status.
* **`notificationStore.ts`**: Menyimpan unread count lonceng, list notifikasi in-app, dan fungsi mark as read.
* **`addonStore.ts`**: Menyimpan katalog add-on, riwayat pesanan add-on, dan saldo jam kerja aktif.
* **`quotaStore.ts`**: Menyimpan status kuota fitur ringan bulanan (`used` vs `total`).
* **`ticketStore.ts`**: Menyimpan tiket aktif, status SLA deadline, dan balasan chat tiket.
* **`timesheetStore.ts`**: Menyimpan entri log jam kerja engineer dan riwayat penggunaan jam klien.
* **`financeStore.ts`**: Menyimpan data invoice, status pembayaran, e-Faktur, dan ledger keuangan ringkas.
* **`accountingStore.ts`**: Menyimpan bagan akun COA, daftar jurnal umum, mutasi buku besar, neraca saldo, P&L, dan neraca keuangan.
* **`expenseStore.ts`**: Menyimpan daftar pengeluaran operasional perusahaan, nota digital, kas kecil, dan verifikasi persetujuan.
* **`taxStore.ts`**: Menyimpan rekapitulasi PPN Keluaran & Masukan, SPT Masa PPN 1111, bukti potong PPh 21/23, status NTPN, dan kalkulasi SPT Tahunan 1771.
* **`reportStore.ts`**: Menyimpan arsip dokumen laporan bulanan eksekutif per proyek.

---

## 3. Konfigurasi Multi-Domain Web App Manifests

* `public/manifests/office-manifest.json` -> Nama: "Meldir Office", Theme: Navy (`#0f172a`), Icon: Admin Badge.
* `public/manifests/jobs-manifest.json` -> Nama: "Meldir Jobs", Theme: Slate (`#090d16`), Icon: Dev Terminal.
* `public/manifests/portal-manifest.json` -> Nama: "Meldir Client Portal", Theme: Emerald (`#0f172a`), Icon: Client Shield.
