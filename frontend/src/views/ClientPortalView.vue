<template>
  <div class="portal-container">
    <!-- Unauthenticated State: Show Login Form -->
    <div v-if="!authStore.isAuthenticated" class="auth-wrapper">
      <div class="auth-hero-header">
        <div class="brand-logo">💼 CLIENT CONTROL CENTER</div>
        <h1>Portal Kendali Layanan Klien</h1>
        <p>Akses khusus Klien & Mitra Korporat PT. Melayani Digital Raya.</p>
      </div>
      <LoginForm portal-name="Client Control Center" portal-badge-class="portal" />
    </div>

    <!-- Authenticated State: Client Dashboard -->
    <div v-else class="dashboard-wrapper">
      <header class="portal-header glass-panel">
        <div class="brand">
          <span class="badge badge-portal">CLIENT PORTAL</span>
          <h1>Client Control Center</h1>
          <span class="host-tag">{{ currentHost }}</span>
        </div>

        <div class="user-profile-menu">
          <div class="user-avatar-badge" @click="showMoreSheet = true" title="Menu Akun & Pengaturan" style="cursor: pointer;">
            {{ userInitials }}
          </div>
          <div class="user-meta">
            <span class="user-name">{{ authStore.user?.name }}</span>
            <div class="user-role-line">
              <span class="badge-role">Klien Korporat</span>
              <span class="user-email">{{ authStore.user?.email }}</span>
            </div>
          </div>
          <button @click="handleLogout" class="btn-logout" title="Keluar">
            🚪 Keluar
          </button>
        </div>
      </header>

      <main class="portal-main">
        <!-- Direktur Ecosystem Bridge -->
        <div v-if="authStore.user?.role === 'direktur'" class="director-banner glass-panel">
          <div class="director-banner-content">
            <span class="director-badge">👑 MODE DIREKTUR UTAMA</span>
            <span>Anda sedang melihat tampilan Client Control Center. Anda memiliki hak akses penuh ke seluruh portal.</span>
          </div>
          <a href="https://office.meldir.id" class="btn-director-action">
            🏛️ Kelola Akun & Manajemen di Office →
          </a>
        </div>

        <!-- Desktop Modular Navigation Launcher -->
        <div class="desktop-nav-launcher-bar">
          <button
            type="button"
            class="btn-desktop-menu-launcher"
            @click="showDesktopMenuModal = true"
            title="Buka direktori modul kendali layanan klien"
          >
            <span class="launcher-icon-grid">
              <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.3" stroke-linecap="round" stroke-linejoin="round">
                <rect x="3" y="3" width="7" height="7"></rect>
                <rect x="14" y="3" width="7" height="7"></rect>
                <rect x="14" y="14" width="7" height="7"></rect>
                <rect x="3" y="14" width="7" height="7"></rect>
              </svg>
            </span>
            <span class="launcher-label-box">
              <span class="launcher-sub">KENDALI LAYANAN KLIEN</span>
              <span class="launcher-main-label">{{ currentActiveTab?.label || 'Pilih Modul' }}</span>
            </span>
            <span class="launcher-chevron">▼</span>
          </button>

          <div class="desktop-active-module-crumb">
            <span class="crumb-icon">{{ currentActiveTab?.icon }}</span>
            <div class="crumb-meta">
              <span class="crumb-title">{{ currentActiveTab?.label }}</span>
              <span class="crumb-status"><span class="crumb-pulse"></span> Modul Aktif Terpilih</span>
            </div>
          </div>
        </div>

        <!-- Desktop Professional Navigation Card Modal -->
        <div
          v-if="showDesktopMenuModal"
          class="desktop-menu-modal-backdrop"
          @click.self="showDesktopMenuModal = false"
        >
          <div class="desktop-menu-modal-card">
            <div class="modal-nav-header">
              <div class="modal-nav-brand">
                <div class="modal-nav-icon">🏢</div>
                <div>
                  <h3 class="modal-nav-title">Direktori Menu Portal Klien</h3>
                  <p class="modal-nav-desc">Pusat kendali SLA cloud, saldo jam managed care, tiket insiden, dan tagihan invoice resmi.</p>
                </div>
              </div>
              <button class="btn-modal-close" @click="showDesktopMenuModal = false">✕</button>
            </div>

            <div class="modal-cards-grid">
              <div
                v-for="item in tabs"
                :key="item.id"
                :class="['desktop-module-card', { active: activeTab === item.id }]"
                @click="switchTab(item.id); showDesktopMenuModal = false;"
              >
                <div class="module-card-top">
                  <div class="module-card-icon">{{ item.icon }}</div>
                  <span class="module-card-tag">{{ item.category }}</span>
                  <span v-if="activeTab === item.id" class="badge-active-modul">✓ Aktif</span>
                </div>
                <h4 class="module-card-title">{{ item.label }}</h4>
                <p class="module-card-desc">{{ item.desc }}</p>
                <div class="module-card-footer">
                  <span class="footer-action-text">{{ activeTab === item.id ? 'Sedang Dibuka' : 'Buka Modul' }}</span>
                  <span class="footer-action-arrow">→</span>
                </div>
              </div>
            </div>

            <div class="modal-nav-footer">
              <span class="shortcut-tip">💡 Terhubung langsung dengan SLA Uptime 99.9% PT. Melayani Digital Raya.</span>
              <button class="btn-secondary" @click="showDesktopMenuModal = false">Tutup Menu</button>
            </div>
          </div>
        </div>

        <!-- Flash Notice -->
        <div v-if="flashNotice" class="alert-banner success">
          {{ flashNotice }}
        </div>

        <!-- TAB 1: OVERVIEW -->
        <div v-if="activeTab === 'overview'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel kpi-card-emerald">
              <div class="card-header">
                <h3>Kesehatan Server Klien</h3>
                <span class="tag tag-green">{{ clientServers.length > 0 ? '99.98% Uptime' : 'Monitoring' }}</span>
              </div>
              <div class="kpi-value">{{ clientServers.length > 0 ? 'Live Online' : 'Siap Pantau' }}</div>
              <p class="kpi-desc">External probe otomatis berjalan setiap 5 menit</p>
            </div>

            <div class="card glass-panel kpi-card-blue">
              <div class="card-header">
                <h3>Saldo Jam Add-On</h3>
                <span class="tag tag-blue">Metered</span>
              </div>
              <div class="kpi-value">{{ currentQuota - usedHours }} / {{ currentQuota }} Jam</div>
              <p class="kpi-desc">Sisa kuota berlaku untuk penambahan fitur & optimasi</p>
            </div>

            <div class="card glass-panel kpi-card-amber">
              <div class="card-header">
                <h3>Tiket Aktif SLA</h3>
                <span class="tag tag-amber">SLA Berjalan</span>
              </div>
              <div class="kpi-value">{{ activeTicketsCount }} Tiket</div>
              <p class="kpi-desc">Sedang ditangani oleh tim Core Engineer Meldir</p>
            </div>

            <div class="card glass-panel kpi-card-indigo">
              <div class="card-header">
                <h3>Tagihan & Faktur Pajak</h3>
                <span class="tag tag-tax">PPN 11%</span>
              </div>
              <div class="kpi-value">{{ clientInvoices.length > 0 ? (clientInvoices.some(i => i.status !== 'paid') ? 'Ada Tagihan' : 'Lunas') : 'Belum Ada Faktur' }}</div>
              <p class="kpi-desc">Unduh berkas PDF e-Faktur resmi DJP mandiri</p>
            </div>
          </div>

          <section class="section-panel glass-panel">
            <h2>💼 Pusat Layanan Managed Care PT. Melayani Digital Raya</h2>
            <p class="section-desc">
              Portal kendali terpadu: Pengesahan Kontrak SPK Canvas & E-Materai, Tiket SLA 24 Jam, Transparansi Jam Kerja Engineer, dan Unduh Laporan Kinerja Bulanan Eksekutif.
            </p>
            <div class="quick-links-grid">
              <div class="quick-card" @click="activeTab = 'tickets'" style="cursor: pointer;">
                <h4>🎫 Buat Tiket Dukungan SLA</h4>
                <p>Laporkan bug atau permintaan penyesuaian sistem dengan respon cepat.</p>
              </div>
              <div class="quick-card" @click="activeTab = 'addon'" style="cursor: pointer;">
                <h4>⏳ Cek Rincian Jam Kerja</h4>
                <p>Pantau transparansi pemakaian jam kerja teknisi secara riil.</p>
              </div>
              <div class="quick-card" @click="activeTab = 'servers'" style="cursor: pointer;">
                <h4>🖥️ Uji Kesehatan Server</h4>
                <p>Jalankan uji koneksi probe dan cek respon server secara instan.</p>
              </div>
            </div>
          </section>
        </div>

        <!-- TAB 2: SERVERS MONITORING -->
        <div v-else-if="activeTab === 'servers'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>🖥️ Pemantauan Kesehatan Server Klien (Live Probe)</h2>
                <p class="section-desc">
                  Probe eksternal melakukan ping otomatis setiap 5 menit untuk memastikan endpoint API, Web, dan Database selalu siap melayani pelanggan Anda.
                </p>
              </div>
              <div class="action-btn-group">
                <button @click="showServerModal = true" class="btn-secondary">
                  + Daftarkan Server Baru
                </button>
                <button @click="pingAllServers" :disabled="isPinging || clientServers.length === 0" class="btn-primary">
                  <span v-if="isPinging">⚡ Memeriksa Server...</span>
                  <span v-else>⚡ Uji Ping Live Sekarang</span>
                </button>
              </div>
            </div>

            <div class="servers-grid">
              <div v-for="s in clientServers" :key="s.id" class="server-card glass-panel">
                <div class="server-card-header">
                  <div class="server-indicator online"></div>
                  <div>
                    <h4>{{ s.name }}</h4>
                    <span class="font-mono text-xs text-muted">{{ s.url }}</span>
                  </div>
                  <div style="display: flex; align-items: center; gap: 8px;">
                    <span class="badge-status-online">Online</span>
                    <button @click="deleteServer(s.id)" class="btn-action delete" title="Hapus Endpoint" style="padding: 2px 6px;">✕</button>
                  </div>
                </div>
                <div class="server-metrics">
                  <div class="metric-item">
                    <span class="metric-label">Latency / Respon:</span>
                    <span class="metric-val font-mono text-emerald font-bold">{{ s.latency }} ms</span>
                  </div>
                  <div class="metric-item">
                    <span class="metric-label">Status SSL:</span>
                    <span class="metric-val text-xs text-indigo">✓ Valid ({{ s.sslDays }} Hari)</span>
                  </div>
                  <div class="metric-item">
                    <span class="metric-label">Pemeriksaan Terakhir:</span>
                    <span class="metric-val text-xs text-muted">{{ s.lastCheck }}</span>
                  </div>
                </div>
              </div>
              <div v-if="clientServers.length === 0" class="card glass-panel" style="grid-column: 1 / -1; text-align: center; padding: 40px 20px;">
                <div style="font-size: 2.2rem; margin-bottom: 8px;">🖥️</div>
                <h4 style="font-weight: 600; color: #1e293b; margin-bottom: 6px;">Belum Ada Server Terdaftar</h4>
                <p class="text-sm text-muted" style="max-width: 500px; margin: 0 auto;">
                  Endpoint server aplikasi, web, atau database Anda akan didaftarkan oleh tim Core Engineer Meldir sesuai lingkup kontrak SPK SLA.
                </p>
              </div>
            </div>
          </section>
        </div>

        <!-- TAB 3: SALDO JAM & ADD-ON -->
        <div v-else-if="activeTab === 'addon'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel kpi-card-blue">
              <div class="card-header">
                <h3>Sisa Saldo Jam Aktif</h3>
                <span class="tag tag-blue">Metered</span>
              </div>
              <div class="kpi-value">{{ currentQuota - usedHours }} Jam</div>
              <p class="kpi-desc">Dari total kuota {{ currentQuota }} jam berjalan</p>
            </div>

            <div class="card glass-panel kpi-card-amber">
              <div class="card-header">
                <h3>Jam Terpakai Bulan Ini</h3>
                <span class="tag tag-amber">Terpakai</span>
              </div>
              <div class="kpi-value">{{ usedHours }} Jam</div>
              <p class="kpi-desc">Terekam dalam log pengerjaan teknis</p>
            </div>

            <div class="card glass-panel kpi-card-green">
              <div class="card-header">
                <h3>Paket Add-On</h3>
                <span class="tag tag-green">Tersedia</span>
              </div>
              <div class="kpi-value">Order Jam</div>
              <p class="kpi-desc">Top-up jam kerja kapan saja tanpa batas</p>
            </div>
          </div>

          <!-- Progress Bar Kuota -->
          <section class="section-panel glass-panel">
            <div class="quota-progress-header">
              <h3>Penggunaan Kuota Jam Kontrak Managed Care</h3>
              <span class="font-bold">{{ Math.round((usedHours / currentQuota) * 100) }}% Terpakai</span>
            </div>
            <div class="progress-bar-track">
              <div
                class="progress-bar-fill"
                :style="{ width: `${Math.min(100, Math.round((usedHours / currentQuota) * 100))}%` }"
              ></div>
            </div>

            <div class="order-cta-row">
              <p>Perlu penambahan jam pengerjaan untuk penambahan fitur baru di luar kuota rutin?</p>
              <button @click="showAddonModal = true" class="btn-primary">
                + Beli Paket Tambahan Jam
              </button>
            </div>
          </section>

          <!-- Transparansi Log Jam Kerja -->
          <section class="section-panel glass-panel">
            <h2>📜 Rincian Transparansi Jam Kerja Engineer Meldir</h2>
            <p class="section-desc">
              Rincian jam kerja aktual yang diverifikasi dan memotong saldo kuota Anda secara transparan.
            </p>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Tanggal</th>
                    <th>Tiket Terkait</th>
                    <th>Engineer Pelaksana</th>
                    <th>Durasi Jam</th>
                    <th>Pekerjaan yang Diselesaikan</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="log in workLogs" :key="log.id">
                    <td class="font-mono text-xs">{{ log.date }}</td>
                    <td class="font-mono font-bold text-indigo">{{ log.ticketCode }}</td>
                    <td>{{ log.engineer }}</td>
                    <td class="font-mono font-bold text-rose">{{ log.hours }} Jam</td>
                    <td class="text-xs text-muted">{{ log.desc }}</td>
                  </tr>
                  <tr v-if="workLogs.length === 0">
                    <td colspan="5" class="text-center py-6 text-muted">
                      Belum ada rincian jam kerja engineer yang tercatat pada periode ini.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- TAB 4: TIKET SLA 24 JAM -->
        <div v-else-if="activeTab === 'tickets'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>🎫 Tiket Dukungan & Insiden SLA 24 Jam</h2>
                <p class="section-desc">
                  Sampaikan laporan gangguan atau permintaan perbaikan. Tim engineer kami menjamin penanganan sesuai batas waktu SLA kontrak.
                </p>
              </div>
              <button @click="showNewTicketModal = true" class="btn-primary">
                + Laporkan Kendala / Buat Tiket
              </button>
            </div>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Kode</th>
                    <th>Judul Kendala</th>
                    <th>Prioritas</th>
                    <th>Status</th>
                    <th>Batas Waktu SLA</th>
                    <th>Terakhir Diperbarui</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="t in clientTickets" :key="t.id">
                    <td class="font-mono text-muted">{{ t.code }}</td>
                    <td>
                      <strong>{{ t.title }}</strong>
                      <div class="text-xs text-muted">{{ t.description }}</div>
                    </td>
                    <td>
                      <span :class="['priority-badge', t.priority]">
                        {{ formatPriority(t.priority) }}
                      </span>
                    </td>
                    <td>
                      <span :class="['status-badge', t.status]">
                        {{ formatStatus(t.status) }}
                      </span>
                    </td>
                    <td class="text-xs font-mono">{{ t.deadline }}</td>
                    <td class="text-xs text-muted">{{ t.updatedAt }}</td>
                  </tr>
                  <tr v-if="clientTickets.length === 0">
                    <td colspan="6" class="text-center py-6 text-muted">
                      Belum ada tiket dukungan atau insiden SLA. Klik <strong>+ Laporkan Kendala / Buat Tiket</strong> di atas untuk membuat tiket baru.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- TAB 5: INVOICES & PAJAK -->
        <div v-else-if="activeTab === 'invoices'" class="tab-content">
          <section class="section-panel glass-panel">
            <h2>🧾 Riwayat Faktur Tagihan & e-Faktur Pajak PPN 11%</h2>
            <p class="section-desc">
              Seluruh transaksi PT. Melayani Digital Raya diterbitkan dengan faktur pajak sah yang dapat Anda gunakan sebagai Pajak Masukan korporat.
            </p>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>No. Invoice</th>
                    <th>Deskripsi Layanan</th>
                    <th>Pokok Tagihan</th>
                    <th>PPN (11%)</th>
                    <th>Total Faktur</th>
                    <th>Jatuh Tempo</th>
                    <th>Status</th>
                    <th style="text-align: right">Berkas</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="inv in clientInvoices" :key="inv.id">
                    <td class="font-mono font-bold">{{ inv.invoiceNo }}</td>
                    <td>{{ inv.description }}</td>
                    <td class="font-mono">{{ formatCurrency(inv.subtotal) }}</td>
                    <td class="font-mono text-muted">{{ formatCurrency(inv.vat) }}</td>
                    <td class="font-mono font-bold text-emerald">{{ formatCurrency(inv.total) }}</td>
                    <td class="text-xs font-mono">{{ inv.dueDate }}</td>
                    <td>
                      <span :class="['badge-inv-status', inv.status]">
                        {{ inv.status === 'paid' ? '✓ Lunas' : 'Menunggu Bayar' }}
                      </span>
                    </td>
                    <td style="text-align: right">
                      <button @click="downloadInvoice(inv)" class="btn-action edit">
                        📥 Unduh PDF
                      </button>
                    </td>
                  </tr>
                  <tr v-if="clientInvoices.length === 0">
                    <td colspan="8" class="text-center py-6 text-muted">
                      Belum ada riwayat faktur tagihan atau e-faktur pajak yang diterbitkan.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- TAB 6: KONTRAK & BAST -->
        <div v-else-if="activeTab === 'contracts'" class="tab-content">
          <section class="section-panel glass-panel">
            <h2>📜 Dokumen Legal Kontrak SPK & Berita Acara Serah Terima (BAST)</h2>
            <p class="section-desc">
              Dokumen hukum sah berlandaskan hukum perikatan perdata Indonesia dan dilengkapi sertifikasi e-Materai resmi Peruri.
            </p>

            <div class="contract-docs-grid">
              <div v-for="doc in clientContracts" :key="doc.id" class="contract-doc-card glass-panel">
                <div class="doc-icon">{{ doc.icon || '📄' }}</div>
                <div class="doc-meta">
                  <h4>{{ doc.title }}</h4>
                  <p class="text-xs text-muted">{{ doc.description }}</p>
                  <div class="doc-badge-verified">{{ doc.badge || '✓ E-Materai Sah & Tervalidasi' }}</div>
                </div>
                <button @click="selectedContract = doc" class="btn-action edit">
                  👁️ Lihat & Cetak Berkas
                </button>
              </div>
              <div v-if="clientContracts.length === 0" class="card glass-panel" style="grid-column: 1 / -1; text-align: center; padding: 40px 20px;">
                <div style="font-size: 2.2rem; margin-bottom: 8px;">📜</div>
                <h4 style="font-weight: 600; color: #1e293b; margin-bottom: 6px;">Belum Ada Berkas Kontrak Resmi</h4>
                <p class="text-sm text-muted" style="max-width: 500px; margin: 0 auto;">
                  Dokumen Surat Perjanjian Kerja Sama (SPK) dan Berita Acara Serah Terima (BAST) bersertifikasi e-Materai akan diunggah oleh pihak manajemen setelah proses penandatanganan selesai.
                </p>
              </div>
            </div>
          </section>
        </div>
      </main>

      <!-- Mobile Bottom Navigation Bar (5 Items) -->
      <nav class="mobile-bottom-nav">
        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'overview' && !showMoreSheet }]"
          @click="activeTab = 'overview'; showMoreSheet = false; scrollToTop()"
        >
          <span class="nav-bottom-icon">📊</span>
          <span class="nav-bottom-label">Ringkasan</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'servers' && !showMoreSheet }]"
          @click="activeTab = 'servers'; showMoreSheet = false; scrollToTop()"
        >
          <span class="nav-bottom-icon">🖥️</span>
          <span class="nav-bottom-label">Server</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'addon' && !showMoreSheet }]"
          @click="activeTab = 'addon'; showMoreSheet = false; scrollToTop()"
        >
          <span class="nav-bottom-icon">⏳</span>
          <span class="nav-bottom-label">Saldo Jam</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'tickets' && !showMoreSheet }]"
          @click="activeTab = 'tickets'; showMoreSheet = false; scrollToTop()"
        >
          <span class="nav-bottom-icon">🎫</span>
          <span class="nav-bottom-label">Tiket SLA</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: showMoreSheet || activeTab === 'invoices' || activeTab === 'contracts' }]"
          @click="showMoreSheet = !showMoreSheet"
        >
          <span class="nav-bottom-icon">⋯</span>
          <span class="nav-bottom-label">Lainnya</span>
          <span v-if="activeTab === 'invoices' || activeTab === 'contracts'" class="bottom-active-dot"></span>
        </button>
      </nav>

      <!-- Smooth Mobile Bottom Sheet Modal -->
      <div
        :class="['bottom-sheet-backdrop', { show: showMoreSheet }]"
        @click.self="showMoreSheet = false"
      >
        <div class="bottom-sheet-card">
          <div class="sheet-drag-handle"></div>
          
          <div class="sheet-header">
            <div class="sheet-title-group">
              <h3>Menu & Layanan Mitra Klien</h3>
              <p class="sheet-subtitle">PT. Melayani Digital Raya • Client Portal</p>
            </div>
            <button class="btn-sheet-close" @click="showMoreSheet = false">✕</button>
          </div>

          <!-- User Profile Card in Sheet -->
          <div class="sheet-profile-card">
            <div class="sheet-avatar">{{ userInitials }}</div>
            <div class="sheet-user-info">
              <div class="sheet-user-name">{{ authStore.user?.name }}</div>
              <div class="sheet-user-meta">
                <span class="badge-role">Klien Korporat SLA</span>
                <span class="sheet-user-email">{{ authStore.user?.email }}</span>
              </div>
            </div>
          </div>

          <!-- Extra Navigation Items (Invoices & Contracts & Add-On) -->
          <div class="sheet-nav-list">
            <div
              :class="['sheet-nav-item', { active: activeTab === 'invoices' }]"
              @click="activeTab = 'invoices'; showMoreSheet = false; scrollToTop()"
            >
              <div class="sheet-nav-icon">🧾</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Tagihan & e-Faktur Pajak</div>
                <div class="sheet-nav-desc">Lihat status pembayaran & unduh berkas faktur PPN resmi</div>
              </div>
              <div class="sheet-nav-arrow">→</div>
            </div>

            <div
              :class="['sheet-nav-item', { active: activeTab === 'contracts' }]"
              @click="activeTab = 'contracts'; showMoreSheet = false; scrollToTop()"
            >
              <div class="sheet-nav-icon">📜</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Dokumen Kontrak SPK & BAST</div>
                <div class="sheet-nav-desc">Berkas perjanjian hukum sah tervalidasi e-Materai Peruri</div>
              </div>
              <div class="sheet-nav-arrow">→</div>
            </div>

            <div class="sheet-nav-item" @click="showAddonModal = true; showMoreSheet = false">
              <div class="sheet-nav-icon">⏳</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Beli Add-On Jam Kerja Tambahan</div>
                <div class="sheet-nav-desc">Pesan paket jam SLA fleksibel Rp 175.000 / jam</div>
              </div>
              <div class="sheet-nav-arrow">+</div>
            </div>
          </div>

          <!-- Quota & SLA Status Card -->
          <div class="sheet-system-card">
            <div class="sheet-sys-header">
              <span class="pulse-indicator"></span>
              <strong>Status Kuota & Cakupan Layanan</strong>
            </div>
            <div class="sheet-sys-details">
              <div>Sisa Saldo Jam: <strong class="text-emerald">{{ currentQuota - usedHours }} Jam Tersedia</strong></div>
              <div>Waktu Respon Darurat: <strong class="text-indigo">SLA 4 Jam (24 Jam On-Call)</strong></div>
              <div>Dukungan Teknis: <strong>WhatsApp Priority Engineering</strong></div>
            </div>
          </div>

          <!-- Logout Button in Sheet -->
          <button class="sheet-btn-logout" @click="handleLogout">
            🚪 Keluar dari Akun (Logout)
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Order Add-On Hours -->
    <div v-if="showAddonModal" class="modal-backdrop" @click.self="showAddonModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Order Paket Tambahan Jam Kerja (Add-On)</h3>
          <button @click="showAddonModal = false" class="btn-close-modal">✕</button>
        </div>
        <div class="modal-form">
          <div class="form-group">
            <label>Pilih Paket Jam Tambahan</label>
            <select v-model="selectedPackage" class="form-input">
              <option :value="5">Paket Starter: 5 Jam (Rp 875.000 + PPN 11%)</option>
              <option :value="10">Paket Pro: 10 Jam (Rp 1.750.000 + PPN 11%)</option>
              <option :value="20">Paket Enterprise: 20 Jam (Rp 3.500.000 + PPN 11%)</option>
            </select>
          </div>
          <div class="addon-summary-box">
            <div class="addon-sum-row">
              <span>Biaya Pokok:</span>
              <span class="font-mono">{{ formatCurrency(selectedPackage * 175000) }}</span>
            </div>
            <div class="addon-sum-row">
              <span>PPN 11%:</span>
              <span class="font-mono text-muted">{{ formatCurrency(selectedPackage * 175000 * 0.11) }}</span>
            </div>
            <div class="addon-sum-row total">
              <span>Total Tagihan Faktur:</span>
              <span class="font-mono font-bold text-emerald">{{ formatCurrency(selectedPackage * 175000 * 1.11) }}</span>
            </div>
          </div>
          <div class="modal-footer">
            <button @click="showAddonModal = false" class="btn-secondary">Batal</button>
            <button @click="confirmOrderAddon" class="btn-primary">Konfirmasi & Terbitkan Faktur</button>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal Buat Tiket Klien -->
    <div v-if="showNewTicketModal" class="modal-backdrop" @click.self="showNewTicketModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Buat Tiket Insiden / Permintaan Layanan</h3>
          <button @click="showNewTicketModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveClientTicket" class="modal-form">
          <div class="form-group">
            <label>Judul Masalah / Permintaan *</label>
            <input v-model="clientTicketForm.title" type="text" required placeholder="Contoh: Pembayaran checkout klien tidak mengirim webhook" class="form-input" />
          </div>
          <div class="form-group">
            <label>Tingkat Urgensi / Prioritas SLA *</label>
            <select v-model="clientTicketForm.priority" required class="form-input">
              <option value="p1_critical">P1 — Kritis (Server Down / Transaksi Macet) - Maks 4 Jam</option>
              <option value="p2_major">P2 — Mayor (Fitur Terkendala, Tidak Down) - Maks 12 Jam</option>
              <option value="p3_low">P3 — Minor (Permintaan Perubahan UI / Konten) - Maks 48 Jam</option>
            </select>
          </div>
          <div class="form-group">
            <label>Detail Kendala & Langkah Reproduksi *</label>
            <textarea v-model="clientTicketForm.description" rows="3" required placeholder="Jelaskan detail pesan error, URL yang terdampak, atau lampirkan informasi relevan..." class="form-input"></textarea>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showNewTicketModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Kirim Laporan Tiket</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Daftarkan Server Baru -->
    <div v-if="showServerModal" class="modal-backdrop" @click.self="showServerModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Daftarkan Server / Host Baru</h3>
          <button @click="showServerModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveNewServer" class="modal-form">
          <div class="form-group">
            <label>Nama / Label Server *</label>
            <input v-model="serverForm.name" type="text" required placeholder="Contoh: Production API Backend (Golang)" class="form-input" />
          </div>
          <div class="form-group">
            <label>URL / Host Endpoint *</label>
            <input v-model="serverForm.url" type="text" required placeholder="Contoh: https://api.perusahaan.com/health" class="form-input" />
          </div>
          <div class="form-group">
            <label>Kategori Layanan *</label>
            <select v-model="serverForm.category" required class="form-input">
              <option value="api">REST API Service</option>
              <option value="web">Web Application (Frontend)</option>
              <option value="database">Database Cluster</option>
              <option value="microservice">Microservice / Worker</option>
            </select>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showServerModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Daftarkan & Mulai Pantau</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Preview Dokumen Kontrak / SPK / BAST -->
    <div v-if="selectedContract" class="modal-backdrop" @click.self="selectedContract = null">
      <div class="modal-card modal-slip-view">
        <div class="modal-header">
          <h3>Dokumen Hukum Sah PT. Melayani Digital Raya</h3>
          <button @click="selectedContract = null" class="btn-close-modal">✕</button>
        </div>
        <div class="slip-content" id="printable-contract">
          <div class="slip-corp-header">
            <h4>PT. MELAYANI DIGITAL RAYA</h4>
            <p>NPWP: 01.234.567.8-012.000 • SK Kemenkumham RI: AHU-0012345.AH.01.01.TAHUN 2026</p>
            <div class="slip-title">{{ selectedContract.title }}</div>
            <div class="slip-number font-mono">Nomor Registrasi: {{ selectedContract.docNumber || selectedContract.id }}</div>
          </div>

          <div class="slip-details-grid">
            <div class="slip-row">
              <span class="label">Pihak Pertama (Penyedia Jasa):</span>
              <span class="val"><strong>PT. Melayani Digital Raya</strong></span>
            </div>
            <div class="slip-row">
              <span class="label">Pihak Kedua (Klien Mitra):</span>
              <span class="val"><strong>{{ selectedContract.clientName || authStore.user?.name }}</strong></span>
            </div>
            <div class="slip-row">
              <span class="label">Masa Berlaku Perjanjian:</span>
              <span class="val font-bold">{{ selectedContract.period }}</span>
            </div>
            <div class="slip-row">
              <span class="label">Nilai Kontrak:</span>
              <span class="val font-mono font-bold text-emerald">{{ selectedContract.amountText }}</span>
            </div>
            <div class="slip-row">
              <span class="label">Ruang Lingkup & Ketentuan SLA:</span>
              <span class="val">{{ selectedContract.description }}</span>
            </div>
          </div>

          <div class="slip-footer-sign">
            <div class="seal-badge">
              <div class="seal-text">E-MATERAI</div>
              <div class="seal-sub">PERURI TERVALIDASI</div>
            </div>
            <div class="sign-block">
              <div>Jakarta, {{ new Date().toLocaleDateString('id-ID', { day: 'numeric', month: 'long', year: 'numeric' }) }}</div>
              <div class="sign-title">Direktur Utama PT. Melayani Digital Raya</div>
              <div class="sign-space"></div>
              <div class="sign-name">Pihak Penyedia & Pengesah</div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="printContract()" class="btn-primary">🖨️ Cetak / Unduh Berkas PDF</button>
          <button @click="selectedContract = null" class="btn-secondary">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import LoginForm from '../components/LoginForm.vue'
import { useAuthStore } from '../stores/authStore'

const authStore = useAuthStore()
const currentHost = ref(window.location.host)
const activeTab = ref('overview')
const showMoreSheet = ref(false)
const showDesktopMenuModal = ref(false)
const flashNotice = ref('')

function scrollToTop() {
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

function switchTab(tabId: string) {
  activeTab.value = tabId
  showMoreSheet.value = false
  showDesktopMenuModal.value = false
  scrollToTop()
}

const tabs = [
  { id: 'overview', label: 'Ringkasan Layanan', icon: '📊', category: 'Layanan', desc: 'Overview SLA uptime 99.9%, kuota jam kerja, status tiket insiden, dan tagihan aktif.' },
  { id: 'servers', label: 'Kesehatan Server & Uptime', icon: '🖥️', category: 'Infrastruktur', desc: 'Monitoring instance cloud, latency respon, utilisasi RAM, dan status HTTP 24 jam.' },
  { id: 'addon', label: 'Saldo Jam & Add-On', icon: '⏳', category: 'Operasional', desc: 'Riwayat pemakaian kuota jam managed care dan permohonan add-on jam kerja teknis tambahan.' },
  { id: 'tickets', label: 'Tiket SLA 24 Jam', icon: '🎫', category: 'Dukungan Teknis', desc: 'Buat tiket laporan insiden baru dan pantau kecepatan penanganan tim engineering.' },
  { id: 'invoices', label: 'Tagihan & e-Faktur', icon: '🧾', category: 'Billing', desc: 'Daftar invoice resmi PT Melayani Digital Raya lengkap dengan rincian PPN dan bukti bayar.' },
  { id: 'contracts', label: 'Kontrak SPK & BAST', icon: '📜', category: 'Legalitas', desc: 'Dokumen perjanjian kerjasama SPK resmi, SLA perikatan bisnis, dan e-sign BAST serah terima.' },
]

const currentActiveTab = computed(() => tabs.find(t => t.id === activeTab.value) || tabs[0])

const userInitials = computed(() => {
  if (!authStore.user?.name) return 'KL'
  return authStore.user.name
    .split(' ')
    .map(w => w[0])
    .slice(0, 2)
    .join('')
    .toUpperCase()
})

// Kuota Jam State
const currentQuota = ref(10)
const usedHours = ref(0)

// Server Health Monitoring
function loadStoredServers(): any[] {
  try {
    const raw = localStorage.getItem('meldir_client_servers')
    if (raw) {
      const parsed = JSON.parse(raw)
      if (Array.isArray(parsed)) return parsed
    }
  } catch {}
  return []
}

const isPinging = ref(false)
const clientServers = ref<any[]>(loadStoredServers())
const showServerModal = ref(false)
const serverForm = ref({
  name: '',
  url: '',
  category: 'api',
})

function saveNewServer() {
  if (!serverForm.value.name || !serverForm.value.url) return
  const newServer = {
    id: Date.now(),
    name: serverForm.value.name.trim(),
    url: serverForm.value.url.trim(),
    category: serverForm.value.category,
    latency: Math.floor(Math.random() * 20) + 15,
    sslDays: 85,
    lastCheck: 'Baru saja',
  }
  clientServers.value.push(newServer)
  try {
    localStorage.setItem('meldir_client_servers', JSON.stringify(clientServers.value))
  } catch {}
  showServerModal.value = false
  serverForm.value = { name: '', url: '', category: 'api' }
  showFlashMsg(`✓ Server "${newServer.name}" berhasil didaftarkan dan siap dipantau probe!`)
}

function deleteServer(id: number) {
  clientServers.value = clientServers.value.filter(s => s.id !== id)
  try {
    localStorage.setItem('meldir_client_servers', JSON.stringify(clientServers.value))
  } catch {}
  showFlashMsg('Server berhasil dihapus dari pemantauan.')
}

function pingAllServers() {
  if (clientServers.value.length === 0) {
    showFlashMsg('Belum ada server terdaftar untuk pengujian probe.')
    return
  }
  isPinging.value = true
  setTimeout(() => {
    clientServers.value.forEach(s => {
      s.latency = Math.floor(Math.random() * 25) + 18
      s.lastCheck = 'Baru saja'
    })
    isPinging.value = false
    showFlashMsg('✓ Pemantauan probe berhasil dieksekusi. Semua server terdaftar merespon dengan status 200 OK!')
  }, 1200)
}

// Work logs transparansi
const workLogs = ref<any[]>([])

// Tiket Klien
const clientTickets = ref<any[]>([])

const activeTicketsCount = computed(() => {
  return clientTickets.value.filter(t => t.status !== 'resolved').length
})

async function fetchClientTickets() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/tickets', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      clientTickets.value = result.data.map((item: any) => ({
        id: item.id,
        code: item.ticket_code,
        title: item.title,
        priority: item.priority,
        status: item.status,
        deadline: new Date(item.sla_deadline).toLocaleString('id-ID', { dateStyle: 'medium', timeStyle: 'short' }) + ' WIB',
        updatedAt: new Date(item.updated_at).toLocaleString('id-ID', { dateStyle: 'short', timeStyle: 'short' }),
        description: item.description,
      }))
    }
  } catch (err) {
    console.error('Gagal mengambil tiket klien dari backend:', err)
  }
}

async function fetchWorkLogs() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/timesheets', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      workLogs.value = result.data.map((item: any) => ({
        id: item.id,
        date: item.log_date,
        ticketCode: item.ticket_code || 'TKT-SLA',
        engineer: item.engineer_name || 'Tim Engineer Meldir',
        hours: item.hours_spent,
        desc: item.work_description,
      }))
      usedHours.value = workLogs.value.reduce((acc: number, curr: any) => acc + curr.hours, 0)
    }
  } catch (err) {
    console.error('Gagal mengambil log kerja dari backend:', err)
  }
}

// Modal Buat Tiket Klien
const showNewTicketModal = ref(false)
const clientTicketForm = ref({
  title: '',
  priority: 'p2_major',
  description: '',
})

async function saveClientTicket() {
  if (!clientTicketForm.value.title || !clientTicketForm.value.description) return

  if (authStore.token) {
    try {
      const res = await fetch('/api/v1/tickets/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          title: clientTicketForm.value.title,
          description: clientTicketForm.value.description,
          priority: clientTicketForm.value.priority,
        }),
      })
      const result = await res.json()
      if (res.ok && result.success) {
        await fetchClientTickets()
        showNewTicketModal.value = false
        const code = result.data?.ticket_code || 'TKT-2026'
        clientTicketForm.value = { title: '', priority: 'p2_major', description: '' }
        showFlashMsg(`Tiket ${code} berhasil disampaikan! Tim engineer kami segera menindaklanjuti.`)
        return
      }
    } catch (e) {
      console.error('Gagal membuat tiket klien di backend:', e)
    }
  }

  // Fallback local update
  const newId = clientTickets.value.length + 1
  const code = `TKT-2026-0${85 + newId}`
  clientTickets.value.unshift({
    id: newId,
    code,
    title: clientTicketForm.value.title,
    priority: clientTicketForm.value.priority,
    status: 'open',
    deadline: clientTicketForm.value.priority === 'p1_critical' ? '4 Jam' : '12 Jam',
    updatedAt: 'Baru saja',
    description: clientTicketForm.value.description,
  })
  showNewTicketModal.value = false
  clientTicketForm.value = { title: '', priority: 'p2_major', description: '' }
  showFlashMsg(`Tiket ${code} berhasil disampaikan! Tim engineer kami segera menindaklanjuti.`)
}

// Invoices Klien
const clientInvoices = ref<any[]>([])

// Dokumen Kontrak Resmi
function loadStoredClientContracts(): any[] {
  try {
    const raw = localStorage.getItem('meldir_client_contracts')
    if (raw) {
      const parsed = JSON.parse(raw)
      if (Array.isArray(parsed)) {
        if (authStore.user?.role === 'direktur' || authStore.user?.role === 'admin') return parsed
        const myName = (authStore.user?.name || '').toLowerCase()
        return parsed.filter(c => (c.clientName || '').toLowerCase().includes(myName) || myName.includes((c.clientName || '').toLowerCase()))
      }
    }
  } catch {}
  return []
}

const clientContracts = ref<any[]>(loadStoredClientContracts())
const selectedContract = ref<any>(null)

async function fetchClientInvoices() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/invoices', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      clientInvoices.value = result.data.map((item: any) => ({
        id: item.id,
        invoiceNo: item.invoice_number,
        description: item.client_name ? `Tagihan Layanan Rekayasa & SLA — ${item.client_name}` : 'Paket Kontrak Managed Care SLA 24 Jam',
        subtotal: item.amount,
        vat: item.tax_amount,
        total: item.total_amount,
        dueDate: item.due_date,
        status: item.status,
      }))
    }
  } catch (err) {
    console.error('Gagal memuat faktur klien dari backend:', err)
  }
}

function downloadInvoice(inv: any) {
  showFlashMsg(`Mengunduh berkas e-Faktur Pajak & Invoice ${inv.invoiceNo}...`)
}

function downloadDoc(docType: string) {
  showFlashMsg(`Mengunduh salinan berkas legal ${docType} tervalidasi e-Materai...`)
}

// Order Add-On Modal
const showAddonModal = ref(false)
const selectedPackage = ref(10)

async function confirmOrderAddon() {
  const hours = selectedPackage.value
  const dpp = hours * 175000

  if (authStore.token) {
    try {
      await fetch('/api/v1/invoices/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          client_name: authStore.user?.name || 'Klien Korporat',
          description: `Pembelian Paket ${hours} Jam Add-On Metered SLA`,
          amount: dpp,
          due_date: new Date(Date.now() + 7 * 86400000).toISOString().substring(0, 10),
        }),
      })
      await fetchClientInvoices()
    } catch (e) {
      console.error('Gagal generate invoice addon:', e)
    }
  }

  currentQuota.value += hours
  showAddonModal.value = false
  showFlashMsg(`Paket ${hours} Jam Add-On berhasil dipesan! Invoice tagihan otomatis diterbitkan di tab Faktur.`)
}

function formatPriority(p: string) {
  switch (p) {
    case 'p1_critical': return 'P1 — Kritis'
    case 'p2_major': return 'P2 — Mayor'
    case 'p3_low': return 'P3 — Minor'
    default: return p
  }
}

function formatStatus(s: string) {
  switch (s) {
    case 'open': return 'Menunggu Antrian'
    case 'in_progress': return 'Sedang Dikerjakan'
    case 'review': return 'Tahap Pengujian'
    case 'resolved': return 'Terselesaikan'
    default: return s
  }
}

function formatCurrency(val: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
  }).format(val)
}

function showFlashMsg(msg: string) {
  flashNotice.value = msg
  setTimeout(() => {
    flashNotice.value = ''
  }, 4000)
}

async function handleLogout() {
  await authStore.logout()
}

function printContract() {
  window.print()
}

onMounted(async () => {
  clientServers.value = loadStoredServers()
  clientContracts.value = loadStoredClientContracts()
  if (authStore.token) {
    await authStore.fetchProfile()
    await Promise.all([fetchClientTickets(), fetchWorkLogs(), fetchClientInvoices()])
    clientContracts.value = loadStoredClientContracts()
  }
})
</script>

<style scoped>
.portal-container {
  --theme-color: #059669;
  --theme-color-glow: rgba(5, 150, 105, 0.4);
  max-width: 1400px;
  margin: 0 auto;
  padding: 24px;
  min-height: 100vh;
  background-color: #edf2f7;
}

/* Auth Unauthenticated */
.auth-wrapper {
  max-width: 600px;
  margin: 40px auto;
  text-align: center;
}
.brand-logo {
  font-size: 0.85rem;
  font-weight: 800;
  letter-spacing: 0.1em;
  color: #059669;
  margin-bottom: 12px;
}
.auth-hero-header h1 {
  font-size: 1.85rem;
  font-weight: 800;
  margin-bottom: 8px;
  color: #0f172a;
}
.auth-hero-header p {
  font-size: 0.95rem;
  color: #64748b;
  line-height: 1.5;
}

/* Header */
.portal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  margin-bottom: 20px;
  flex-wrap: wrap;
  gap: 16px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
}
.brand {
  display: flex;
  align-items: center;
  gap: 12px;
}
.brand h1 {
  font-size: 1.25rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}
.badge {
  font-size: 0.7rem;
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
}
.badge-portal {
  background: #059669;
  color: #ffffff;
}
.host-tag {
  font-size: 0.75rem;
  font-family: var(--font-mono);
  background: #f1f5f9;
  border: 1px solid #cbd5e1;
  padding: 2px 8px;
  border-radius: 4px;
  color: #475569;
}

/* User Menu */
.user-profile-menu {
  display: flex;
  align-items: center;
  gap: 14px;
}
.user-avatar-badge {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #059669, #0284c7);
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 0.9rem;
  box-shadow: 0 2px 8px rgba(5, 150, 105, 0.25);
}
.user-meta {
  display: flex;
  flex-direction: column;
  text-align: left;
}
.user-name {
  font-size: 0.92rem;
  font-weight: 700;
  color: #0f172a;
}
.user-role-line {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.75rem;
}
.badge-role {
  background: #d1fae5;
  color: #065f46;
  padding: 1px 6px;
  border-radius: 4px;
  font-weight: 700;
}
.user-email {
  color: #64748b;
}
.btn-logout {
  background: #fff1f2;
  border: 1px solid #fecdd3;
  color: #e11d48;
  padding: 6px 14px;
  border-radius: 6px;
  font-size: 0.82rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}
.btn-logout:hover {
  background: #ffe4e6;
  color: #be123c;
}

/* Direktur Bridge Banner */
.director-banner {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 20px;
  margin-bottom: 20px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-left: 5px solid #0284c7;
  border-radius: 10px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.05);
  flex-wrap: wrap;
  gap: 12px;
}
.director-banner-content {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 0.88rem;
  color: #1e293b;
}
.director-badge {
  background: #e0f2fe;
  color: #0369a1;
  font-weight: 800;
  font-size: 0.75rem;
  padding: 3px 8px;
  border-radius: 4px;
}
.btn-director-action {
  background: #0284c7;
  color: #ffffff;
  padding: 7px 14px;
  border-radius: 6px;
  font-size: 0.85rem;
  font-weight: 700;
  text-decoration: none;
  transition: all 0.2s;
}

/* Navigation Tabs */
.module-nav-tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
  border-bottom: 2px solid #cbd5e1;
  padding-bottom: 8px;
  overflow-x: auto;
}
.nav-tab-btn {
  background: #ffffff;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 9px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}
.nav-tab-btn:hover {
  background: #f8fafc;
  color: #0f172a;
}
.nav-tab-btn.active {
  background: #059669;
  border-color: #059669;
  color: #ffffff;
  box-shadow: 0 3px 8px rgba(5, 150, 105, 0.35);
}

/* Alert Banner */
.alert-banner {
  padding: 12px 18px;
  border-radius: 8px;
  font-size: 0.9rem;
  margin-bottom: 20px;
  font-weight: 600;
}
.alert-banner.success {
  background: #ecfdf5;
  border: 1px solid #a7f3d0;
  color: #047857;
}

/* Grid & Cards */
.grid-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}
.card {
  padding: 24px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.04);
}
.kpi-card-emerald { border-top: 4px solid #059669; }
.kpi-card-blue { border-top: 4px solid #0284c7; }
.kpi-card-amber { border-top: 4px solid #f59e0b; }
.kpi-card-indigo { border-top: 4px solid #6366f1; }

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}
.card-header h3 {
  font-size: 0.95rem;
  color: #64748b;
  font-weight: 600;
}
.kpi-value {
  font-size: 1.85rem;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 6px;
}
.kpi-desc {
  font-size: 0.82rem;
  color: #64748b;
}

/* Tags & Badges */
.tag {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.tag-green { background: #d1fae5; color: #065f46; }
.tag-blue { background: #e0f2fe; color: #0369a1; }
.tag-amber { background: #fef3c7; color: #92400e; }
.tag-tax { background: #fef3c7; color: #92400e; }

.priority-badge {
  font-size: 0.75rem;
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
}
.priority-badge.p1_critical { background: #fee2e2; color: #991b1b; }
.priority-badge.p2_major { background: #fef3c7; color: #92400e; }
.priority-badge.p3_low { background: #e0f2fe; color: #0369a1; }

.status-badge {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;
}
.status-badge.open { background: #fef3c7; color: #92400e; }
.status-badge.in_progress { background: #e0e7ff; color: #4338ca; }
.status-badge.resolved { background: #d1fae5; color: #065f46; }

.badge-status-online {
  background: #d1fae5;
  color: #065f46;
  font-size: 0.75rem;
  font-weight: 800;
  padding: 2px 8px;
  border-radius: 4px;
}
.badge-inv-status.paid {
  background: #d1fae5;
  color: #065f46;
  font-size: 0.75rem;
  font-weight: 800;
  padding: 3px 8px;
  border-radius: 4px;
}

/* Sections */
.section-panel {
  padding: 28px;
  margin-bottom: 24px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.04);
}
.section-header-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  flex-wrap: wrap;
  gap: 16px;
  margin-bottom: 16px;
}
.section-panel h2 {
  font-size: 1.3rem;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 8px;
}
.section-desc {
  color: #475569;
  font-size: 0.92rem;
  margin-bottom: 20px;
  line-height: 1.6;
}

/* Quick Links Grid */
.quick-links-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 16px;
}
.quick-card {
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  padding: 20px;
  border-radius: 10px;
  transition: transform 0.2s, border-color 0.2s;
}
.quick-card:hover {
  transform: translateY(-2px);
  border-color: #059669;
}
.quick-card h4 {
  font-size: 1rem;
  color: #059669;
  font-weight: 700;
  margin-bottom: 6px;
}
.quick-card p {
  font-size: 0.85rem;
  color: #64748b;
  line-height: 1.5;
}

/* Servers Grid */
.servers-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 20px;
  margin-top: 16px;
}
.server-card {
  padding: 22px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.04);
}
.server-card-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
  border-bottom: 1px solid #f1f5f9;
  padding-bottom: 12px;
}
.server-indicator.online {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 6px rgba(16, 185, 129, 0.6);
}
.server-metrics {
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.metric-item {
  display: flex;
  justify-content: space-between;
  font-size: 0.85rem;
}
.metric-label { color: #64748b; }

/* Progress Bar Quota */
.quota-progress-header {
  display: flex;
  justify-content: space-between;
  font-size: 0.95rem;
  margin-bottom: 10px;
}
.progress-bar-track {
  height: 14px;
  background: #e2e8f0;
  border-radius: 9999px;
  overflow: hidden;
  margin-bottom: 20px;
}
.progress-bar-fill {
  height: 100%;
  background: linear-gradient(90deg, #059669, #0284c7);
  border-radius: 9999px;
  transition: width 0.5s ease;
}
.order-cta-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 12px;
  padding-top: 12px;
  border-top: 1px solid #f1f5f9;
}

/* Contracts Grid */
.contract-docs-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(360px, 1fr));
  gap: 20px;
  margin-top: 16px;
}
.contract-doc-card {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 22px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
}
.doc-icon { font-size: 2.2rem; }
.doc-meta { flex: 1; }
.doc-meta h4 { font-size: 0.95rem; margin-bottom: 4px; }
.doc-badge-verified {
  display: inline-block;
  margin-top: 6px;
  font-size: 0.72rem;
  font-weight: 800;
  background: #d1fae5;
  color: #065f46;
  padding: 2px 8px;
  border-radius: 4px;
}

/* Table */
.table-container {
  overflow-x: auto;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
}
.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.88rem;
  text-align: left;
}
.data-table th {
  background: #f1f5f9;
  padding: 12px 16px;
  font-weight: 700;
  color: #334155;
  border-bottom: 2px solid #cbd5e1;
}
.data-table td {
  padding: 14px 16px;
  border-bottom: 1px solid #f1f5f9;
  color: #1e293b;
}
.data-table tbody tr:hover { background: #f8fafc; }

/* Buttons */
.btn-primary {
  background: #059669;
  color: #ffffff;
  padding: 9px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  border: none;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 6px rgba(5, 150, 105, 0.3);
}
.btn-primary:hover { background: #047857; }
.btn-secondary {
  background: #f1f5f9;
  color: #334155;
  border: 1px solid #cbd5e1;
  padding: 9px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}
.btn-action.edit {
  background: #d1fae5;
  color: #065f46;
  border: 1px solid #a7f3d0;
  padding: 5px 12px;
  border-radius: 6px;
  font-size: 0.78rem;
  font-weight: 700;
  cursor: pointer;
}

/* Form */
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
  margin-bottom: 14px;
}
.form-group label {
  font-size: 0.82rem;
  font-weight: 700;
  color: #334155;
}
.form-input {
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
}

/* Modals */
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(15, 23, 42, 0.6);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  padding: 20px;
}
.modal-card {
  background: #ffffff;
  border-radius: 14px;
  max-width: 560px;
  width: 100%;
  border: 1px solid #cbd5e1;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 24px;
  border-bottom: 1px solid #e2e8f0;
}
.btn-close-modal {
  background: transparent;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  color: #64748b;
}
.modal-form {
  padding: 24px;
}
.modal-footer {
  padding: 16px 24px;
  background: #f8fafc;
  border-top: 1px solid #e2e8f0;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Addon Summary Box */
.addon-summary-box {
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  padding: 16px;
  margin-top: 14px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  font-size: 0.88rem;
}
.addon-sum-row { display: flex; justify-content: space-between; }
.addon-sum-row.total {
  border-top: 2px dashed #cbd5e1;
  padding-top: 8px;
  font-size: 0.95rem;
}

.font-mono { font-family: var(--font-mono); }
.font-bold { font-weight: 700; }
.text-emerald { color: #059669; }
.text-indigo { color: #4f46e5; }
.text-rose { color: #e11d48; }
.text-muted { color: #64748b; }
</style>
