# 🎨 Panduan & Cetak Biru Desain UI/UX Multi-Portal meldir.id

> **Catatan Rahasia & Internal**: Dokumen ini merinci cetak biru desain antarmuka (*User Interface*) dan pengalaman pengguna (*User Experience*) untuk 3 Portal Dashboard (`office.meldir.id`, `jobs.meldir.id`, `portal.meldir.id`) serta **Tombol Unduh PWA Per Portal**.

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
* **PDF**: Penampil dokumen PDF terintegrasi (*PDF Viewer*) untuk melihat kontrak atau invoice langsung.

---

## 4. Visualisasi Grafik & Canvas TTD Digital

### 📊 A. Smooth Area Wave Chart (Solid Clean Palette)
* **Gaya Grafik**: Wave Chart kurva halus (*Smooth Spline Area*).
* **Warna Palette**: Warna solid/doff profesional (*Matte Palette*: Navy `#1e293b`, Cyan `#0284c7`, Emerald `#059669`) **tanpa efek gradien transparan**.

### ✒️ B. Canvas Tanda Tangan Digital (iPad Drawing Pad Style)
* Pilihan Warna Tinta: **Hitam** (`#000000`) dan **Biru Tua** (`#1e3a8a`).
* Tombol **"Bersihkan Canvas"** (*Clear*) dan **"Simpan & Lanjutkan E-Materai"** (*Save & Proceed*).
