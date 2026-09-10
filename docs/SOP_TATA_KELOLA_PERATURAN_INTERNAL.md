PERATURAN INTERNAL & STANDAR OPERASIONAL PROSEDUR (SOP)
TATA KELOLA PERUSAHAAN & PLATFORM DIGITAL
PT. MELAYANI DIGITAL RAYA (MELDIR)
SK Kemenkumham: AHU-A104016.AH.01.30.Tahun 2026

================================================================================

PANDUAN CETAK & LEGALITAS:
Dokumen ini merupakan instrumen peraturan internal resmi (Internal Corporate Bylaws & SOP) yang mengikat seluruh jajaran pimpinan, karyawan staf, tenaga ahli (engineer), serta menjadi rujukan operasional dalam melayani klien PT. Melayani Digital Raya. Dokumen ini dirancang untuk dicetak secara fisik, dibubuhi tanda tangan pengesahan Direksi, dan disimpan sebagai arsip tata kelola perusahaan (corporate governance handbook).

================================================================================

DAFTAR ISI

1. BAB I: KETENTUAN UMUM & RUANG LINGKUP
2. BAB II: STRUKTUR PERANAN & HAK AKSES PENGGUNA (ROLE-BASED ACCESS CONTROL)
3. BAB III: STANDAR OPERASIONAL PENGEMBANGAN SISTEM (CUSTOM DEVELOPMENT)
4. BAB IV: STANDAR OPERASIONAL PEMELIHARAAN & DUKUNGAN SISTEM (MAINTENANCE & SLA)
5. BAB V: TATA KELOLA TRANSAKSI KEUANGAN, INVOICING, & PAYMENT GATEWAY
6. BAB VI: KEAMANAN INFORMASI, PRIVASI DATA, & KERAHASIAAN (NDA)
7. BAB VII: KODE ETIK KERJA, INTEGRITAS, & LARANGAN BENTURAN KEPENTINGAN
8. BAB VIII: SANKSI PELANGGARAN & MEKANISME RESOLUSI
9. BAB IX: KETENTUAN PERALIHAN & PENGESAHAN

================================================================================

BAB I: KETENTUAN UMUM & RUANG LINGKUP

Pasal 1 — Definisi Istilah
Dalam Peraturan Internal dan SOP ini, yang dimaksud dengan:
1. Perusahaan adalah PT. Melayani Digital Raya (MELDIR), badan hukum resmi berbentuk Perseroan Terbatas yang bergerak di bidang jasa konsultasi, pengembangan, pemeliharaan, dan tata kelola teknologi informasi.
2. Platform Webapp MELDIR adalah ekosistem aplikasi digital terpadu milik Perusahaan yang beroperasi di domain utama meldir.id beserta seluruh portal penunjangnya (office.meldir.id, jobs.meldir.id, portal.meldir.id).
3. Superadmin adalah Direktur Utama, Direksi, atau Founder yang memegang hak kontrol mutlak (root access) atas seluruh data, kebijakan finansial, dan infrastruktur Perusahaan.
4. Admin / Staf Pegawai adalah karyawan operasional Perusahaan yang bertugas mengelola administrasi, verifikasi transaksi keuangan, layanan pelanggan (customer support), dan operasional harian di portal office.meldir.id.
5. Software Engineer adalah tenaga ahli pengembang perangkat lunak, baik staf internal maupun mitra profesional terikat kontrak, yang bertugas merancang, menguji, dan memelihara kode program di portal jobs.meldir.id.
6. Customer / Klien adalah entitas bisnis, instansi pemerintah, yayasan sosial, atau individu terdaftar di portal portal.meldir.id yang memanfaatkan jasa digital Perusahaan.
7. SPK / PKS adalah Surat Perintah Kerja atau Perjanjian Kerja Sama tertulis yang mengikat hak dan kewajiban antara Perusahaan dan Klien.
8. SLA (Service Level Agreement) adalah standar komitmen kecepatan penanganan kendala teknis maksimal 1x24 jam kerja.
9. RAB (Rencana Anggaran Biaya) adalah rincian estimasi biaya pengembangan perangkat lunak yang disetujui bersama sebelum pengerjaan dimulai.

Pasal 2 — Tujuan & Landasan Pemberlakuan
1. Menjamin transparansi, akuntabilitas, dan keamanan operasional ekosistem webapp Perusahaan.
2. Mencegah kebocoran data rahasia (trade secret), aset digital, dan penyalahgunaan wewenang.
3. Memberikan kepastian standar mutu layanan (Anti-Ghosting, SLA 24 Jam, Legalitas PT) kepada seluruh Klien.
4. Mengatur hak, kewajiban, dan pembagian tugas antar organ operasional Perusahaan.

================================================================================

BAB II: STRUKTUR PERANAN & HAK AKSES PENGGUNA (RBAC)

Pasal 3 — Peranan Superadmin (Direksi / Founder)
1. Wewenang Mutlak:
   - Memiliki akses penuh terhadap seluruh basis data, konfigurasi server, log audit, dan laporan keuangan komprehensif.
   - Mengesahkan RAB bernilai di atas ambang batas wewenang Admin (di atas Rp 25.000.000).
   - Memegang kendali tunggal atas Server Credential Vault (kunci SSH, kata sandi root database, token API perbankan/payment gateway).
   - Menetapkan pembagian honorarium/komisi bagi Software Engineer.
   - Mengambil keputusan final dalam perkara banding (counter-offer) pemutusan layanan Klien (offboarding).
2. Kewajiban Keamanan: Wajib mengaktifkan autentikasi dua faktor (2FA) dan dilarang membagikan kredensial Superadmin kepada pihak mana pun tanpa persetujuan tertulis resmi.

Pasal 4 — Peranan Admin / Staf (Pegawai Operasional)
1. Tugas & Wewenang:
   - Memproses penerimaan prospek (leads), verifikasi data legalitas Klien, dan penyusunan draf SPK/PKS.
   - Melakukan validasi pembayaran transfer manual rekening PT dan memantau status notifikasi webhook Payment Gateway.
   - Menerbitkan Invoice resmi dan mengoordinasikan penerbitan Faktur Pajak bersama bagian akuntansi/pajak.
   - Mendistribusikan tiket kendala (tickets) dari Klien kepada Software Engineer yang relevan.
   - Mengoperasikan mesin siaran push notifikasi (broadcast push) dan memantau status WhatsApp API Perusahaan.
2. Batasan Ketat: Admin dilarang keras mengubah kode sumber (source code) aplikasi atau mengakses basis data teknis secara langsung tanpa mandat teknis tertulis.

Pasal 5 — Peranan Software Engineer (Teknis & DevOps)
1. Tugas & Wewenang:
   - Menulis kode program berkualitas tinggi, bersih (clean code), dan terdokumentasi dengan baik pada repositori GitHub resmi Perusahaan.
   - Menyelesaikan tiket bug dan permintaan pemeliharaan sesuai batas waktu SLA yang ditentukan.
   - Membangun antarmuka dan sistem backend sesuai spesifikasi rancangan UI/UX dan arsitektur database yang telah disetujui.
2. Protokol Zero Direct Server Access:
   - Engineer TIDAK DIBERIKAN akses langsung ke kata sandi root server VPS, database produksi, atau hosting Klien.
   - Akses pengerjaan hanya diberikan melalui kolaborasi repositori GitHub selama status proyek berjalan (in_progress).
   - Seluruh penerapan (deployment) ke server produksi dilakukan melalui pipa otomatis (CI/CD) atau dijalankan langsung oleh Superadmin/DevOps Lead resmi.

Pasal 6 — Peranan Customer / Klien
1. Hak Klien:
   - Mengakses portal.meldir.id untuk memantau kemajuan proyek, riwayat pembayaran, dokumen kontrak, dan status kesehatan server.
   - Menandatangani SPK/PKS secara digital melalui kanvas web dan membubuhkan E-Materai resmi.
   - Menerima salinan aset kode (source code) dan hak milik hukum setelah proyek selesai dan terlunasi 100%.
   - Mengajukan tiket kendala teknis 24 jam dengan jaminan respons maksimal 1x24 jam hari kerja.
2. Kewajiban Klien:
   - Melakukan pembayaran tepat waktu sesuai tagihan invoice / termin yang disepakati.
   - Menyediakan materi konten, data bisnis, dan akses penunjang yang diperlukan dalam proses pengerjaan.
   - Menjaga kerahasiaan kredensial akun portal klien milik perusahaannya.

================================================================================

BAB III: STANDAR OPERASIONAL PENGEMBANGAN SISTEM (CUSTOM DEVELOPMENT)

Pasal 7 — Alur Siklus Pengerjaan Proyek (Development Lifecycle)
Setiap proyek pengembangan sistem baru wajib melalui 5 tahapan berurutan:
1. Tahap Konsultasi & Analisis Kebutuhan (Discovery Phase):
   - Admin/Superadmin mendengarkan SOP unik bisnis Klien dan memetakan alur kerja.
   - Perusahaan menerbitkan dokumen Rencana Anggaran Biaya (RAB) dan estimasi jadwal kerja (timeline).
2. Tahap Pengikatan Kontrak (Legal Binding):
   - Penandatanganan SPK/PKS resmi berbadan hukum PT. Melayani Digital Raya.
   - Pembayaran uang muka (DP) minimal 40% (atau termin pertama sesuai kesepakatan) masuk ke rekening resmi PT.
3. Tahap Desain & Pemrograman (Development Phase):
   - Pembuatan purwarupa antarmuka (UI/UX prototype) dan skema database PostgreSQL/MySQL.
   - Penulisan kode program di repositori privat GitHub Perusahaan oleh tim Engineer.
4. Tahap Pengujian & Uji Coba Klien (User Acceptance Testing - UAT):
   - Demonstrasi fungsionalitas sistem di lingkungan staging.
   - Klien menguji alur kerja dan mencatat perbaikan minor yang masih berada di dalam ruang lingkup RAB.
5. Tahap Serah Terima & Pelunasan (Handover & BAST):
   - Penandatanganan Berita Acara Serah Terima (BAST).
   - Pelunasan sisa tagihan 100%.
   - Penyerahan kredensial administratif, hak cipta kode, dan pelatihan (training) staf Klien hingga mandiri.

Pasal 8 — Pengendalian Ruang Lingkup (Scope Creep Prevention)
1. Segala permintaan penambahan fitur di luar dokumen spesifikasi teknis awal (RAB) dikategorikan sebagai Pekerjaan Tambahan (Change Request).
2. Pekerjaan tambahan wajib dituangkan dalam Adendum Kontrak tertulis beserta penyesuaian biaya dan waktu penyelesaian.
3. Tim Engineer dilarang mengeksekusi fitur baru atas instruksi lisan Klien sebelum mendapat konfirmasi tertulis dari Admin/Superadmin.

================================================================================

BAB IV: STANDAR OPERASIONAL PEMELIHARAAN & DUKUNGAN (MAINTENANCE & SLA)

Pasal 9 — Klasifikasi Paket Pemeliharaan
Perusahaan menyediakan 2 tingkatan paket pemeliharaan berkala:
1. Paket Perawatan Dasar (Basic Care - Rp 1.500.000 / $99 per bulan):
   - Pemeliharaan infrastruktur, instalasi & perpanjangan SSL, dan monitoring uptime server.
   - Pencadangan (backup) data database mingguan.
   - Perbaikan galat teknis/bug ringan yang mengganggu kelancaran operasional harian.
   - Jaminan SLA penanganan darurat maksimal 1x24 jam di hari kerja.
2. Paket Perawatan Prioritas (Priority Care - Rp 3.500.000 / $225 per bulan):
   - Seluruh fasilitas Paket Perawatan Dasar.
   - Pencadangan (backup) data database harian.
   - Fasilitas penambahan fitur ringan maksimal 2 (dua) kali per bulan.
   - Jalur komunikasi prioritas via grup WhatsApp khusus langsung bersama Lead Engineer.

Pasal 10 — Batasan Teknis Biaya Server & Fitur Ringan
1. Biaya Sewa Infrastruktur Cloud Terpisah:
   - Biaya paket pemeliharaan adalah murni jasa pemeliharaan, monitoring, dan penanganan teknis (retaining fee).
   - Biaya sewa server pihak ketiga (seperti VPS Niagahoster, Biznet, CloudKilat, DigitalOcean, AWS, GCP) adalah biaya terpisah dan dibayarkan secara transparan sesuai kebutuhan kapasitas Klien.
2. Definisi Ruang Lingkup Fitur Ringan (Minor Feature):
   - Yang dimaksud penambahan fitur ringan pada Priority Care adalah pekerjaan dengan estimasi waktu maksimal 2 hingga 3 jam kerja teknis per permintaan.
   - Contoh fitur ringan yang diakomodasi: penyesuaian/penambahan field formulir baru, perbaikan format cetak dokumen PDF/kuitansi, pembaruan filter laporan tabel, atau penggantian teks/banner pengumuman.
   - Pembuatan modul baru, penambahan alur bisnis baru, atau integrasi API pihak ketiga baru TIDAK TERMASUK fitur ringan dan harus dihitung melalui RAB Adendum.

Pasal 11 — Prosedur Penanganan Tiket Darurat SLA 24 Jam
1. Klien melaporkan kendala melalui tombol tiket di portal.meldir.id atau grup WhatsApp darurat.
2. Status tiket otomatis tercatat pada office.meldir.id dengan stempel waktu (timestamp).
3. Admin/Lead Engineer wajib memberikan konfirmasi awal penanganan dalam waktu maksimal 2 (dua) jam kerja.
4. Proses investigasi, isolasi bug, dan perbaikan wajib diselesaikan atau diberikan solusi sementara (workaround) maksimal dalam tempo 24 jam kerja.
5. Jika kendala disebabkan oleh gangguan pihak ketiga (third-party downtime seperti ISP atau bencana data center), Perusahaan wajib menyertakan bukti status resmi dari penyedia layanan tersebut.

================================================================================

BAB V: TATA KELOLA TRANSAKSI KEUANGAN, INVOICING, & PAYMENT GATEWAY

Pasal 12 — Saluran Pembayaran Resmi
1. Pembayaran yang diakui sah oleh Perusahaan hanya melalui saluran berikut:
   - Rekening Bank Resmi Perusahaan: Rekening giro/tabungan atas nama resmi PT. Melayani Digital Raya.
   - Payment Gateway Otomatis Terverifikasi: Saluran pembayaran otomatis (QRIS, Virtual Account, Kartu Kredit) yang tertaut secara resmi pada webapp Perusahaan melalui penyedia berizin Bank Indonesia (misal: Xendit / Midtrans).
2. Perusahaan TIDAK PERNAH membenarkan pembayaran ke rekening pribadi karyawan, engineer, atau pihak ketiga mana pun. Pembayaran ke rekening pribadi dianggap tidak sah dan Perusahaan tidak bertanggung jawab atas kerugian yang ditimbulkan.

Pasal 13 — Penerbitan Invoice & Faktur Pajak
1. Setiap transaksi otomatis menghasilkan tagihan digital (Digital Invoice) dengan nomor unik berkode tahun dan bulan berjalan.
2. Klien yang membutuhkan Faktur Pajak wajib melampirkan salinan NPWP Perusahaan / NIK dan status PKP pada profil portal.meldir.id.
3. Invoice memiliki masa jatuh tempo (due date) 7 (tujuh) hari kalender sejak diterbitkan.
4. Keterlambatan pembayaran paket pemeliharaan melebihi 14 hari berakibat pada penonaktifan sementara (temporary suspension) dukungan pemeliharaan hingga tunggakan diselesaikan.

Pasal 14 — Skema Termin Pembayaran Proyek Kustom
Kecuali disepakati lain dalam SPK tertulis, skema standar termin proyek adalah:
- Termin I (Uang Muka / DP): Sebesar 40% dibayarkan sebelum pekerjaan desain & arsitektur dimulai.
- Termin II (Milestone Fungsional): Sebesar 30% dibayarkan setelah fungsionalitas inti selesai dan siap diuji coba di server staging.
- Termin III (Pelunasan & Serah Terima): Sebesar 30% dibayarkan setelah pengujian UAT disetujui, sebelum penyerahan kredensial produksi dan hak cipta source code.

Pasal 15 — Prosedur Honorarium & Bagi Hasil Engineer
1. Nilai honorarium bagi Software Engineer ditentukan berdasarkan kesepakatan per proyek atau per milestone.
2. Khusus proyek nirlaba/sosial yayasan, dapat berlaku skema honorarium khusus atau pro-bono yang disepakati sukarela oleh Engineer.
3. Pembayaran honorarium dicairkan melalui transfer bank setelah Engineer menyelesaikan deliverables dan disetujui (signed off) oleh Lead Engineer / Superadmin.

================================================================================

BAB VI: KEAMANAN INFORMASI, PRIVASI DATA, & KERAHASIAAN (NDA)

Pasal 16 — Perjanjian Kerahasiaan (Non-Disclosure Agreement - NDA)
1. Seluruh Superadmin, Admin, dan Software Engineer terikat secara otomatis oleh klausul kerahasiaan Perusahaan sejak hari pertama bertugas.
2. Informasi rahasia mencakup:
   - Kode sumber (source code) dan arsitektur sistem Klien.
   - Data transaksi, omzet, basis data pelanggan Klien, dan strategi bisnis Klien.
   - Kunci enkripsi, token API, dan konfigurasi server internal PT. MELDIR.
3. Kewajiban menjaga kerahasiaan ini tetap berlaku mengikat tanpa batas waktu, bahkan setelah masa kerja atau kerja sama proyek berakhir.

Pasal 17 — Protokol Privasi Data & Isolasi Kredensial
1. Data bisnis Klien yang tersimpan dalam basis data produksi tidak boleh diunduh, disalin, atau didistribusikan ke komputer lokal pribadi tanpa enkripsi dan izin tertulis Superadmin.
2. Lingkungan pengujian (development/staging) wajib menggunakan data tiruan (dummy data) yang disamarkan (sanitized), dilarang menggunakan data sensitif riil milik Klien.
3. Kredensial server Klien yang disimpan di vault office.meldir.id hanya boleh diakses saat melakukan tindakan darurat perbaikan sistem.

================================================================================

BAB VII: KODE ETIK KERJA, INTEGRITAS, & LARANGAN BENTURAN KEPENTINGAN

Pasal 18 — Larangan Transaksi Gelap (Anti-Side Channeling)
1. Seluruh karyawan staf dan Software Engineer dilarang keras menawarkan jasa pribadi secara langsung (moonlighting/side-channeling) kepada Klien yang diperkenalkan oleh atau bertransaksi dengan PT. Melayani Digital Raya.
2. Segala bentuk penawaran proyek tambahan dari Klien wajib dilaporkan kepada Admin/Superadmin untuk diproses melalui kontrak resmi Perusahaan.
3. Pelanggaran terhadap pasal ini dikenakan sanksi pemutusan hubungan kerja seketika dan tuntutan ganti rugi perdata atas potensi pendapatan Perusahaan yang hilang.

Pasal 19 — Klausul Larangan Bersaing (Non-Compete)
1. Tenaga ahli dan karyawan dilarang mendirikan, mengelola, atau bekerja pada entitas usaha sejenis yang secara langsung menjadi kompetitor aktif PT. Melayani Digital Raya selama masa kerja aktif.
2. Dilarang memanfaatkan aset intelektual, pustaka kode internal (proprietary boilerplates), atau template rancangan milik Perusahaan untuk kepentingan komersial pihak ketiga di luar naungan PT. MELDIR.

Pasal 20 — Etika Komunikasi Profesional
1. Segala bentuk interaksi dengan Klien wajib dilakukan dengan sopan, transparan, dan profesional melalui kanal resmi (WhatsApp bisnis PT, email domain @meldir.id, atau fitur tiket portal).
2. Dilarang memberikan janji lisan di luar kapasitas teknis atau di luar klausul yang tercantum dalam kontrak resmi.

================================================================================

BAB VIII: SANKSI PELANGGARAN & MEKANISME RESOLUSI

Pasal 21 — Tingkatan Sanksi Internal
Setiap pelanggaran terhadap ketentuan dalam Peraturan Internal dan SOP ini akan dikenakan tindakan pendisiplinan berjenjang:
1. Surat Peringatan Pertama (SP 1): Untuk kelalaian administratif ringan, keterlambatan respons SLA tanpa alasan sah, atau kelalaian dokumentasi.
2. Surat Peringatan Kedua (SP 2): Untuk pengulangan pelanggaran SP 1 dalam kurun waktu 3 (tiga) bulan atau kelalaian teknis yang mengakibatkan downtime sistem Klien.
3. Surat Peringatan Ketiga / Terakhir (SP 3) & Penonaktifan: Untuk pelanggaran berat, pengabaian instruksi manajemen, atau pembocoran informasi non-kritis.
4. Pemutusan Hubungan Kerja (PHK) & Tindakan Hukum: Berlaku seketika tanpa peringatan untuk tindak pidana penggelapan uang, pencurian aset source code, pembocoran basis data Klien (data breach), atau transaksi gelap (side-channeling).

Pasal 22 — Tuntutan Ganti Rugi
Perusahaan berhak menuntut ganti rugi materiil dan imateriil kepada oknum yang tindakannya terbukti mengakibatkan kerugian finansial langsung bagi PT. Melayani Digital Raya atau tuntutan hukum dari pihak Klien.

Pasal 23 — Domisili Hukum & Penyelesaian Sengketa
1. Segala perselisihan yang timbul terkait penafsiran atau pelaksanaan SOP ini diutamakan diselesaikan secara musyawarah mufakat.
2. Apabila mufakat tidak tercapai dalam waktu 30 (tiga puluh) hari kalender, para pihak sepakat memilih domisili hukum yang tetap di Kantor Kepaniteraan Pengadilan Negeri tempat kedudukan hukum PT. Melayani Digital Raya.

================================================================================

BAB IX: KETENTUAN PERALIHAN & PENGESAHAN

Pasal 24 — Pemberlakuan & Pembaruan
1. Dokumen Peraturan Internal dan Standar Operasional Prosedur ini sah dan mulai berlaku terhitung sejak tanggal ditetapkan oleh Direksi.
2. Peraturan ini dapat ditinjau kembali dan disesuaikan secara berkala oleh Direksi sesuai dengan perkembangan teknologi dan regulasi perundang-undangan Republik Indonesia.

================================================================================

LEMBAR PENGESAHAN DIREKSI
Ditetapkan di: Surabaya, Indonesia  
Pada tanggal: 7 September 2026  

Atas Nama Manajemen & Dewan Direksi  
PT. MELAYANI DIGITAL RAYA (MELDIR)



______________________________________  
Direktur Utama / Founder  
PT. Melayani Digital Raya  
NIB / SK Kemenkumham: AHU-A104016.AH.01.30.Tahun 2026
