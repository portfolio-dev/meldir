# 🎨 Spesifikasi & Frontend Components Vue 3 + Vite

Dokumen ini merinci komponen antarmuka Vue 3 khusus untuk Sistem Lonceng Notifikasi, Manajemen Profil Akun, Broadcast Push Notifikasi, Multi-Media Uploader, dan **PWA Multi-Portal Installer**.

---

## 1. Komponen Antarmuka Kunci Baru (Universal & Portal-Specific)

### A. Universal Components (Digunakan di `office`, `jobs`, `portal`, `public`)
* **`PWAInstallPrompt.vue`**:
  - Tombol **"Install / Unduh Aplikasi PWA"** yang terpasang di sidebar dan modal pop-up interaktif saat dibuka di browser HP/Desktop.
  - Mengelola event `beforeinstallprompt` untuk mengunduh PWA sesuai domain yang sedang dibuka (`office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`).
* **`NotificationBellDropdown.vue`**:
  - Icon lonceng pada header dengan badge angka merah (*unread count*).
  - Pop-up panel daftar notifikasi dengan tombol **"Tandai Sudah Dibaca Semua"** dan *deep-link redirection*.
* **`ProfileSettingsModal.vue`**:
  - Form ubah Nama Lengkap, Email, Nomor WhatsApp, Upload Foto Profil (Avatar preview), dan Ganti Password Aman.
* **`MultiMediaUploader.vue`**:
  - Komponen upload drag-and-drop yang mendukung berkas Gambar, Video, dan PDF.

### B. Konfigurasi Multi-Domain Web App Manifests
* `public/manifests/office-manifest.json` -> Nama: "Meldir Office", Theme: Navy (`#0f172a`), Icon: Badge Admin.
* `public/manifests/jobs-manifest.json` -> Nama: "Meldir Jobs", Theme: Slate (`#090d16`), Icon: Badge Dev.
* `public/manifests/portal-manifest.json` -> Nama: "Meldir Client Portal", Theme: Emerald (`#0f172a`), Icon: Badge Client.
