PERATURAN INTERNAL & STANDAR OPERASIONAL PROSEDUR (SOP)
TATA KELOLA PERUSAHAAN & PLATFORM DIGITAL
PT. MELAYANI DIGITAL RAYA (MELDIR)
SK Kemenkum: AHU-A104016.AH.01.30.Tahun 2026 | NIB: 0709260111296

================================================================================

PANDUAN CETAK & LEGALITAS:
Dokumen ini merupakan instrumen peraturan internal resmi (Internal Corporate Bylaws & SOP) yang mengikat seluruh jajaran pimpinan, karyawan staf, tenaga ahli (engineer), serta menjadi rujukan operasional dalam melayani klien PT. Melayani Digital Raya. Dokumen ini dirancang untuk dicetak secara fisik, dibubuhi tanda tangan pengesahan Direksi, dan disimpan sebagai arsip tata kelola perusahaan (corporate governance handbook).

================================================================================

DAFTAR ISI

1. BAB I: KETENTUAN UMUM & RUANG LINGKUP
2. BAB II: STRUKTUR PERANAN & HAK AKSES PENGGUNA (ROLE-BASED ACCESS CONTROL)
3. BAB III: STANDAR OPERASIONAL PENGEMBANGAN SISTEM (CUSTOM DEVELOPMENT)
4. BAB IV: STANDAR OPERASIONAL PEMELIHARAAN & DUKUNGAN SISTEM (MAINTENANCE & SLA)
5. BAB V: TATA KELOLA KEUANGAN, AKUNTANSI BERPASANGAN, & PERPAJAKAN PERUSAHAAN
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

Pasal 8 — Pengendalian Ruang Lingkup (Scope Creep Prevention) & Skema Add-On
1. Segala permintaan penambahan fitur di luar dokumen spesifikasi teknis awal (RAB) dikategorikan sebagai Pekerjaan Tambahan (Change Request).
2. Pekerjaan tambahan dapat diselesaikan melalui dua mekanisme resmi:
   - Skema Add-On Terukur (Metered Hours): Klien dapat memesan paket jam kerja on-demand (Hourly Extension Rp 175.000/jam, Sprint Block 10 Jam Rp 1.500.000, atau Sprint Block 20 Jam Rp 2.800.000) langsung melalui portal.meldir.id.
   - Adendum Kontrak SPK: Untuk penambahan modul skala besar atau perombakan sistem struktural yang membutuhkan RAB dan timeline terpisah.
3. Tim Engineer dilarang mengeksekusi fitur baru atas instruksi lisan Klien sebelum mendapat tiket resmi atau konfirmasi tertulis dari Admin/Superadmin.
4. Seluruh pengurangan saldo jam kerja Add-On wajib didasarkan pada catatan log kerja (Timesheet) terverifikasi.

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

Pasal 10 — Batasan Teknis Biaya Server, Kuota Fitur Ringan, & Add-On
1. Biaya Sewa Infrastruktur Cloud Terpisah:
   - Biaya paket pemeliharaan adalah murni jasa pemeliharaan, monitoring, dan penanganan teknis (retaining fee).
   - Biaya sewa server pihak ketiga (seperti VPS Niagahoster, Biznet, CloudKilat, DigitalOcean, AWS, GCP) adalah biaya terpisah dan dibayarkan secara transparan sesuai kebutuhan kapasitas Klien.
2. Definisi & Batasan Kuota Fitur Ringan (Minor Feature):
   - Yang dimaksud penambahan fitur ringan pada Priority Care adalah pekerjaan dengan estimasi waktu maksimal 2 hingga 3 jam kerja teknis per permintaan (maksimal 2x per bulan).
   - Kuota fitur ringan otomatis di-reset menjadi 2 (dua) kali pada tanggal 1 setiap awal bulan (00:01 WIB) dan tidak dapat diakumulasi (*non-rollover*).
   - Permintaan penambahan fitur melebihi kuota 2x sebulan atau pekerjaan yang melebihi estimasi 3 jam kerja wajib menggunakan saldo paket Add-On Jam Kerja Klien.

Pasal 11 — Prosedur Penanganan Tiket Darurat SLA 24 Jam & Laporan Bulanan
1. Klien melaporkan kendala melalui tombol tiket di portal.meldir.id atau grup WhatsApp darurat.
2. Status tiket otomatis tercatat pada office.meldir.id dengan stempel waktu (timestamp).
3. Admin/Lead Engineer wajib memberikan konfirmasi awal penanganan dalam waktu maksimal 2 (dua) jam kerja.
4. Proses investigasi, isolasi bug, dan perbaikan wajib diselesaikan atau diberikan solusi sementara (workaround) maksimal dalam tempo 24 jam kerja.
5. Jika kendala disebabkan oleh gangguan pihak ketiga (third-party downtime seperti ISP atau bencana data center), Perusahaan wajib menyertakan bukti status resmi dari penyedia layanan tersebut.
6. Laporan Kinerja Bulanan Eksekutif (Monthly Executive Report):
   - Pada tanggal 1 setiap bulan, sistem otomatis menggenerasi berkas PDF Laporan Kinerja Bulanan per proyek klien yang memuat persentase uptime server, riwayat pencadangan data, pembaruan keamanan, dan log tiket yang diselesaikan.
   - Berkas laporan tersedia di portal.meldir.id sebagai bukti transparansi operasional dan pertanggungjawaban legal PT. MELDIR.

================================================================================

BAB V: TATA KELOLA KEUANGAN, AKUNTANSI BERPASANGAN, & PERPAJAKAN PERUSAHAAN

Pasal 12 — Saluran Pembayaran Resmi
1. Pembayaran yang diakui sah oleh Perusahaan hanya melalui saluran berikut:
   - Rekening Bank Resmi Perusahaan: Rekening giro/tabungan atas nama resmi PT. Melayani Digital Raya.
   - Payment Gateway Otomatis Terverifikasi: Saluran pembayaran otomatis (QRIS, Virtual Account, Kartu Kredit) yang tertaut secara resmi pada webapp Perusahaan melalui penyedia berizin Bank Indonesia (misal: Xendit / Midtrans).
2. Perusahaan TIDAK PERNAH membenarkan pembayaran ke rekening pribadi karyawan, engineer, atau pihak ketiga mana pun. Pembayaran ke rekening pribadi dianggap tidak sah dan Perusahaan tidak bertanggung jawab atas kerugian yang ditimbulkan.

Pasal 13 — Penerbitan Invoice, e-Faktur Pajak (PPN 11%), & Jatuh Tempo
1. Setiap transaksi otomatis menghasilkan tagihan digital (Digital Invoice) dengan nomor unik berkode tahun dan bulan berjalan.
2. Klien yang membutuhkan Faktur Pajak wajib melampirkan salinan NPWP Perusahaan / NIK dan status PKP pada profil portal.meldir.id.
3. Penerbitan e-Faktur Pajak:
   - Admin Office wajib mengunggah salinan berkas PDF e-Faktur Pajak resmi DJP dan mencatat Nomor Seri Faktur Pajak (NSFP) pada sistem invoice setelah pembayaran diverifikasi.
   - Klien dapat mengunduh berkas e-Faktur resmi langsung dari portal.meldir.id untuk keperluan pelaporan SPT Masa PPN perusahaan Klien.
4. Invoice memiliki masa jatuh tempo (due date) 7 (tujuh) hari kalender sejak diterbitkan.
5. Keterlambatan pembayaran paket pemeliharaan melebihi 14 hari berakibat pada penonaktifan sementara (temporary suspension) dukungan pemeliharaan hingga tunggakan diselesaikan.

Pasal 14 — Skema Termin Pembayaran & Berita Acara Serah Terima (BAST) Digital
1. Kecuali disepakati lain dalam SPK tertulis, skema standar termin proyek adalah:
   - Termin I (Uang Muka / DP): Sebesar 40% dibayarkan sebelum pekerjaan desain & arsitektur dimulai.
   - Termin II (Milestone Fungsional): Sebesar 30% dibayarkan setelah fungsionalitas inti selesai dan siap diuji coba di server staging.
   - Termin III (Pelunasan & Serah Terima): Sebesar 30% dibayarkan setelah pengujian UAT disetujui.
2. Prasyarat Sah Pelunasan & Penyerahan Source Code:
   - Penandatanganan Berita Acara Serah Terima (BAST) secara digital (TTD canvas & upload e-Materai) merupakan prasyarat mutlak sebelum penerbitan invoice pelunasan Termin III.
   - Penyerahan kredensial produksi dan pengalihan repositori kode sumber (source code) hanya dilakukan setelah BAST ditandatangani dan Termin III terlunasi 100%.

Pasal 15 — Prosedur Honorarium Engineer & Kewajiban Timesheet
1. Nilai honorarium bagi Software Engineer ditentukan berdasarkan kesepakatan per proyek, per milestone, atau per jam kerja Add-On.
2. Kewajiban Pengisian Timesheet:
   - Setiap Software Engineer wajib mencatat jam kerja riil dan deskripsi pekerjaan secara detail pada fitur Timesheet di jobs.meldir.id untuk setiap tiket atau tugas yang dikerjakan.
   - Log timesheet ini menjadi dasar perhitungan pencairan honorarium teknis dan pemotongan saldo jam kerja Klien.
3. Khusus proyek nirlaba/sosial yayasan, dapat berlaku skema honorarium khusus atau pro-bono yang disepakati sukarela oleh Engineer.
4. Pembayaran honorarium dicairkan melalui transfer bank setelah Engineer menyelesaikan deliverables dan disetujui (signed off) oleh Lead Engineer / Superadmin.

Pasal 16 — Standar Pembukuan Akuntansi Berpasangan (Double-Entry Bookkeeping & SAK EMKM)
1. Perusahaan menyelenggarakan pembukuan akuntansi digital berpasangan (*Double-Entry Bookkeeping*) berpedoman pada Standar Akuntansi Keuangan Entitas Mikro, Kecil, dan Menengah (SAK EMKM) secara transparan dan akuntabel.
2. Seluruh transaksi finansial wajib diklasifikasikan menggunakan Bagan Akun Standar 5-Digit (Standard Chart of Accounts / COA):
   - Kelompok 10000: Aset (Kas Kecil, Rekening Bank Mandiri & BCA, Piutang Usaha, Aset Tetap, Uang Muka Pajak).
   - Kelompok 20000: Kewajiban (Utang Usaha, Utang Gaji/Honor Engineer, Utang Pajak PPN & PPh, Pendapatan Diterima di Muka).
   - Kelompok 30000: Ekuitas (Modal Disetor Pendiri, Saldo Laba Ditahan, Laba Tahun Berjalan).
   - Kelompok 40000: Pendapatan Operasional (Jasa Managed Care, Custom Development, Add-On Hours, Cloud Scaling).
   - Kelompok 50000: Beban Pokok Pendapatan / HPP (Honorarium Langsung Engineer, Biaya Staging & Third-Party API).
   - Kelompok 60000: Beban Operasional & Umum (Petty Cash Kantor, Legalitas, E-Materai, Pemasaran, Admin Bank, Penyusutan).
3. Integritas Keseimbangan Jurnal (*Balance Integrity*): Setiap transaksi jurnal wajib memiliki nilai total Debit yang sama persis dengan total Kredit (`Debit == Credit`). Sistem menolak pembukuan entri jurnal yang tidak seimbang.
4. Mesin Penjurnalan Otomatis (*Auto-Journaling Engine*):
   - Saat status invoice diverifikasi menjadi `paid`, sistem otomatis menerbitkan jurnal kas masuk, piutang, pendapatan, dan pengakuan PPN keluaran.
   - Saat honorarium engineer dicairkan, sistem otomatis mendebit beban honor, mengkredit utang PPh 21, dan mengkredit rekening bank PT.
   - Saat beban operasional disetujui, sistem otomatis mendebit akun beban terkait dan mengkredit akun kas/bank.

Pasal 17 — Tata Kelola Kas Kecil (Petty Cash) & Beban Operasional Perusahaan
1. Kas Kecil (*Petty Cash*) dikelola untuk membiayai kebutuhan operasional harian kantor, pembelian e-Materai, biaya administrasi, dan perlengkapan mendesak.
2. Bukti Digital Wajib: Setiap pengeluaran kas kecil atau beban operasional wajib dicatat di `office.meldir.id` dengan melampirkan berkas digital struk/nota/kuitansi asli (format WebP atau PDF). Pengeluaran tanpa bukti digital dilarang dibukukan.
3. Batasan Wewenang Persetujuan (*Approval Limits*):
   - Pengeluaran sampai dengan Rp 2.000.000 dapat disetujui oleh Admin Keuangan.
   - Pengeluaran di atas Rp 2.000.000 wajib disetujui secara digital oleh Direktur Utama.
4. Klasifikasi Fiskal: Setiap entri beban wajib diklasifikasikan statusnya:
   - *Fiscal Deductible*: Biaya operasional murni yang dapat dikurangkan dari penghasilan bruto sesuai Pasal 6 UU PPh.
   - *Non-Deductible*: Biaya yang tidak dapat dikurangkan secara pajak (misal: natura tanpa fasilitas, sumbangan non-resmi, sanksi denda pajak) dan wajib dikoreksi positif pada akhir tahun.

Pasal 18 — Kepatuhan Pajak Pertambahan Nilai (PPN 11%) & SPT Masa 1111
1. Pemungutan PPN Keluaran: Setiap penyerahan Jasa Kena Pajak (JKP) kepada Klien wajib dikenakan PPN sebesar 11% (atau tarif resmi yang berlaku) dari Dasar Pengenaan Pajak (DPP).
2. Faktur Pajak Masukan: Bukti faktur pajak atas belanja kebutuhan perusahaan dari vendor resmi (misal: penyedia cloud server, ISP, atau konsultan rekanan ber-PKP) wajib direkam ke sistem sebagai PPN Masukan yang dapat dikreditkan.
3. Rekonsiliasi SPT Masa PPN 1111:
   - Sistem secara otomatis menghitung selisih antara PPN Keluaran dan PPN Masukan setiap akhir masa pajak.
   - Apabila terdapat PPN Kurang Bayar, Perusahaan wajib menyetorkan kekurangan tersebut ke kas negara sebelum akhir bulan berikutnya dan menginput kode NTPN resmi ke dalam modul perpajakan.

Pasal 19 — Pemotongan & Pemungutan PPh (PPh 21 Engineer & PPh 23 Klien/Vendor)
1. PPh Pasal 21 atas Honorarium Software Engineer:
   - Setiap pembayaran honorarium kepada Software Engineer dipotong PPh 21 sesuai ketentuan perpajakan yang berlaku (Tarif Efektif Rata-rata / TER atau Tarif Pasal 17 UU PPh atas Tenaga Ahli).
   - Perusahaan wajib menerbitkan Bukti Potong resmi (e-Bupot 21) dan menyediakannya di portal `jobs.meldir.id` agar dapat diunduh oleh Engineer untuk keperluan pelaporan SPT Tahunan Orang Pribadi.
2. PPh Pasal 23 yang Dipotong oleh Klien (Kredit Pajak PT):
   - Klien korporat/instansi yang memotong PPh Pasal 23 sebesar 2% atas invoice PT. Melayani Digital Raya wajib menyerahkan Bukti Potong PPh 23 resmi.
   - Bukti potong tersebut dicatat ke sistem sebagai **Aset Uang Muka Pajak (Akun 11050)** dan menjadi kredit pajak pengurang PPh Badan terutang pada SPT Tahunan 1771.
3. PPh Pasal 23 atas Pengeluaran Jasa Pihak Ketiga:
   - Apabila Perusahaan menggunakan jasa sewa atau jasa teknik dari pihak ketiga, Perusahaan bertindak sebagai pemotong PPh Pasal 23 sebesar 2%, menyetorkannya ke kas negara dengan kode NTPN, dan memberikan bukti potong kepada vendor terkait.

Pasal 20 — Rekonsiliasi Fiskal & SPT Tahunan PPh Badan 1771
1. Pada akhir tahun buku (31 Desember), Perusahaan menyusun Laporan Keuangan Komersial terpadu yang memuat Laba Bersih Komersial.
2. Rekonsiliasi Fiskal Otomatis:
   - Sistem melakukan penyesuaian koreksi fiskal positif atas seluruh beban non-deductible yang tercatat sepanjang tahun buku.
   - Menghasilkan nilai Penghasilan Kena Pajak (PKP) resmi untuk formulir SPT 1771.
3. Penerapan Fasilitas Pasal 31E UU HPP:
   - Selama peredaran bruto (omset) Perusahaan tidak melebihi Rp 4.800.000.000 (empat koma delapan milyar rupiah), Perusahaan berhak memperoleh fasilitas pengurangan tarif sebesar 50% dari tarif normal PPh Badan 22%, sehingga tarif efektif PPh Badan yang berlaku adalah **11%**.
4. Pelunasan PPh Pasal 29:
   - PPh Badan terutang dikurangi akumulasi kredit pajak PPh 23 (bukti potong dari klien) dan setoran angsuran PPh 25 bulanan.
   - Sisa PPh Kurang Bayar (PPh Pasal 29) wajib disetor ke kas negara sebelum penyampaian SPT Tahunan PPh Badan pada bulan April.
5. Hak Akses Auditor & Konsultan Pajak: Konsultan pajak terdaftar atau auditor eksternal dapat diberikan akun dengan peranan `audit` (Read-Only) untuk memvalidasi kebenaran mutasi buku besar, jurnal, dan kepatuhan faktur tanpa risiko manipulasi data.

================================================================================

BAB VI: KEAMANAN INFORMASI, PRIVASI DATA, & KERAHASIAAN (NDA)

Pasal 21 — Perjanjian Kerahasiaan (Non-Disclosure Agreement - NDA)
1. Seluruh Superadmin, Admin, dan Software Engineer terikat secara otomatis oleh klausul kerahasiaan Perusahaan sejak hari pertama bertugas.
2. Informasi rahasia mencakup:
   - Kode sumber (source code) dan arsitektur sistem Klien.
   - Data transaksi, omzet, basis data pelanggan Klien, dan strategi bisnis Klien.
   - Kunci enkripsi, token API, dan konfigurasi server internal PT. MELDIR.
3. Kewajiban menjaga kerahasiaan ini tetap berlaku mengikat tanpa batas waktu, bahkan setelah masa kerja atau kerja sama proyek berakhir.

Pasal 22 — Protokol Privasi Data & Isolasi Kredensial
1. Data bisnis Klien yang tersimpan dalam basis data produksi tidak boleh diunduh, disalin, atau didistribusikan ke komputer lokal pribadi tanpa enkripsi dan izin tertulis Superadmin.
2. Lingkungan pengujian (development/staging) wajib menggunakan data tiruan (dummy data) yang disamarkan (sanitized), dilarang menggunakan data sensitif riil milik Klien.
3. Kredensial server Klien yang disimpan di vault office.meldir.id hanya boleh diakses saat melakukan tindakan darurat perbaikan sistem.

================================================================================

BAB VII: KODE ETIK KERJA, INTEGRITAS, & LARANGAN BENTURAN KEPENTINGAN

Pasal 23 — Larangan Transaksi Gelap (Anti-Side Channeling)
1. Seluruh karyawan staf dan Software Engineer dilarang keras menawarkan jasa pribadi secara langsung (moonlighting/side-channeling) kepada Klien yang diperkenalkan oleh atau bertransaksi dengan PT. Melayani Digital Raya.
2. Segala bentuk penawaran proyek tambahan dari Klien wajib dilaporkan kepada Admin/Superadmin untuk diproses melalui kontrak resmi Perusahaan.
3. Pelanggaran terhadap pasal ini dikenakan sanksi pemutusan hubungan kerja seketika dan tuntutan ganti rugi perdata atas potensi pendapatan Perusahaan yang hilang.

Pasal 24 — Klausul Larangan Bersaing (Non-Compete)
1. Tenaga ahli dan karyawan dilarang mendirikan, mengelola, atau bekerja pada entitas usaha sejenis yang secara langsung menjadi kompetitor aktif PT. Melayani Digital Raya selama masa kerja aktif.
2. Dilarang memanfaatkan aset intelektual, pustaka kode internal (proprietary boilerplates), atau template rancangan milik Perusahaan untuk kepentingan komersial pihak ketiga di luar naungan PT. MELDIR.

Pasal 25 — Etika Komunikasi Profesional
1. Segala bentuk interaksi dengan Klien wajib dilakukan dengan sopan, transparan, dan profesional melalui kanal resmi (WhatsApp bisnis PT, email domain @meldir.id, atau fitur tiket portal).
2. Dilarang memberikan janji lisan di luar kapasitas teknis atau di luar klausul yang tercantum dalam kontrak resmi.

================================================================================

BAB VIII: SANKSI PELANGGARAN & MEKANISME RESOLUSI

Pasal 26 — Tingkatan Sanksi Internal
Setiap pelanggaran terhadap ketentuan dalam Peraturan Internal dan SOP ini akan dikenakan tindakan pendisiplinan berjenjang:
1. Surat Peringatan Pertama (SP 1): Untuk kelalaian administratif ringan, keterlambatan respons SLA tanpa alasan sah, atau kelalaian dokumentasi.
2. Surat Peringatan Kedua (SP 2): Untuk pengulangan pelanggaran SP 1 dalam kurun waktu 3 (tiga) bulan atau kelalaian teknis yang mengakibatkan downtime sistem Klien.
3. Surat Peringatan Ketiga / Terakhir (SP 3) & Penonaktifan: Untuk pelanggaran berat, pengabaian instruksi manajemen, atau pembocoran informasi non-kritis.
4. Pemutusan Hubungan Kerja (PHK) & Tindakan Hukum: Berlaku seketika tanpa peringatan untuk tindak pidana penggelapan uang, pencurian aset source code, pembocoran basis data Klien (data breach), atau transaksi gelap (side-channeling).

Pasal 27 — Tuntutan Ganti Rugi
Perusahaan berhak menuntut ganti rugi materiil dan imateriil kepada oknum yang tindakannya terbukti mengakibatkan kerugian finansial langsung bagi PT. Melayani Digital Raya atau tuntutan hukum dari pihak Klien.

Pasal 28 — Domisili Hukum & Penyelesaian Sengketa
1. Segala perselisihan yang timbul terkait penafsiran atau pelaksanaan SOP ini diutamakan diselesaikan secara musyawarah mufakat.
2. Apabila mufakat tidak tercapai dalam waktu 30 (tiga puluh) hari kalender, para pihak sepakat memilih domisili hukum yang tetap di Kantor Kepaniteraan Pengadilan Negeri tempat kedudukan hukum PT. Melayani Digital Raya.

================================================================================

BAB IX: KETENTUAN PERALIHAN & PENGESAHAN

Pasal 29 — Pemberlakuan & Pembaruan
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
NIB: 0709260111296 | SK Kemenkum: AHU-A104016.AH.01.30.Tahun 2026
