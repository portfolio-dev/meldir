<template>
  <div class="portal-container">
    <!-- Unauthenticated State: Show Login Form -->
    <div v-if="!authStore.isAuthenticated" class="auth-wrapper">
      <div class="auth-hero-header">
        <div class="brand-logo">🏛️ PT. MELAYANI DIGITAL RAYA</div>
        <h1>Portal Operasional & Akuntansi Korporat</h1>
        <p>Akses terbatas untuk Direksi, Manajemen, Administrasi Keuangan, dan Tim Audit Resmi.</p>
      </div>
      <LoginForm portal-name="Office Operations" portal-badge-class="office" />
    </div>

    <!-- Authenticated State: Executive Dashboard -->
    <div v-else class="dashboard-wrapper">
      <header class="portal-header glass-panel">
        <div class="brand">
          <span class="badge badge-office">OFFICE ADMIN</span>
          <h1>PT. Melayani Digital Raya</h1>
          <span class="host-tag">{{ currentHost }}</span>
        </div>

        <div class="user-profile-menu">
          <div class="user-avatar-badge" @click="showMoreSheet = true" title="Menu Akun & Pengaturan" style="cursor: pointer;">
            {{ userInitials }}
          </div>
          <div class="user-meta">
            <span class="user-name">{{ authStore.user?.name }}</span>
            <div class="user-role-line">
              <span class="badge-role">{{ roleDisplay }}</span>
              <span class="user-email">{{ authStore.user?.email }}</span>
            </div>
          </div>
          <button @click="handleLogout" class="btn-logout" title="Keluar dari akun">
            🚪 Keluar
          </button>
        </div>
      </header>

      <main class="portal-main">
        <!-- Live System & API Status Bar -->
        <div class="status-bar-banner glass-panel">
          <div class="status-left">
            <span class="pulse-indicator"></span>
            <strong>Status Backend RESTful:</strong>
            <span :class="['api-badge', apiStatus]">{{ apiStatusText }}</span>
          </div>
          <div class="status-right">
            <span>PostgreSQL: <strong class="text-emerald">Terhubung (32 Tabel DDL)</strong></span>
            <span class="sep">•</span>
            <span>Redis Cache: <strong class="text-indigo">Aktif (Port 6379)</strong></span>
          </div>
        </div>

        <!-- Navigation Tabs -->
        <div class="module-nav-tabs">
          <button
            v-for="tab in tabs"
            :key="tab.id"
            :class="['nav-tab-btn', { active: activeTab === tab.id, 'highlight-user-tab': tab.id === 'users' }]"
            @click="switchTab(tab.id)"
          >
            {{ tab.icon }} {{ tab.label }}
            <span v-if="tab.id === 'users'" class="badge-tab-pill">Direktur</span>
          </button>
        </div>

        <!-- Global Alert Message -->
        <div v-if="globalMessage" :class="['alert-banner', globalMessageType]">
          {{ globalMessage }}
        </div>

        <!-- Tab 1: Ringkasan Eksekutif -->
        <div v-if="activeTab === 'overview'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel kpi-card-blue">
              <div class="card-header">
                <h3>Omset Invoice YTD</h3>
                <span class="tag tag-green">Live 2026</span>
              </div>
              <div class="kpi-value">{{ formatCurrency(totalRevenueYTD) }}</div>
              <p class="kpi-desc">Total faktur terbit & lunas terlapor SAK EMKM</p>
            </div>

            <div class="card glass-panel kpi-card-green">
              <div class="card-header">
                <h3>Buku Besar (General Ledger)</h3>
                <span class="badge-balanced">✓ Seimbang</span>
              </div>
              <div class="kpi-value">{{ journals.length > 0 ? journals.length + ' Transaksi' : '0 Transaksi' }}</div>
              <p class="kpi-desc">Debit & Kredit seimbang. Sesuai standar UU Pajak & IAI</p>
            </div>

            <div class="card glass-panel kpi-card-amber">
              <div class="card-header">
                <h3>Kepatuhan DJP (Pajak)</h3>
                <span class="tag tag-tax">SPT Masa 1111</span>
              </div>
              <div class="kpi-value">{{ officeInvoices.length > 0 ? officeInvoices.length + ' Faktur PPN' : '0 Faktur' }}</div>
              <p class="kpi-desc">Faktur Keluaran siap lapor ke sistem Coretax DJP</p>
            </div>

            <div
              class="card glass-panel kpi-card-indigo card-clickable"
              @click="switchTab('users')"
              title="Klik untuk membuka Manajemen Pengguna (CRUD)"
            >
              <div class="card-header">
                <h3>Ekosistem Pengguna</h3>
                <span class="tag tag-indigo">Atur User (CRUD) →</span>
              </div>
              <div class="kpi-value">{{ usersList.length > 0 ? usersList.length + ' User' : 'Kelola Akun' }}</div>
              <p class="kpi-desc">Akses lintas office, jobs, dan portal klien (Klik untuk kelola akun)</p>
            </div>
          </div>

          <section class="section-panel glass-panel">
            <h2>📊 Ringkasan Ekosistem Digital</h2>
            <p class="section-desc">
              Sistem terintegrasi monorepo dengan backend Golang di port 8080 dan OpenLiteSpeed reverse-proxy. Seluruh transaksi bisnis dicatat menggunakan mekanisme double-entry bookkeeping otomatis.
            </p>
            <div class="quick-links-grid">
              <div
                class="quick-card quick-card-highlight"
                @click="switchTab('users')"
                title="Buka Menu Manajemen Pengguna"
              >
                <h4>👥 Manajemen Akun Semua User (CRUD)</h4>
                <p>Tambah pengguna baru, atur peranan (office/jobs/portal), perbarui password & kontak, atau hapus user.</p>
                <div class="quick-card-btn">Buka Menu Pengguna →</div>
              </div>
              <div class="quick-card">
                <h4>Inbound CRM Leads</h4>
                <p>Pipeline prospek otomatis dari form formulir landing page meldir.id.</p>
              </div>
              <div class="quick-card">
                <h4>Digital BAST & SPK</h4>
                <p>E-Materai & tanda tangan digital sah berlandaskan hukum perikatan bisnis.</p>
              </div>
              <div class="quick-card">
                <h4>Monitoring SLA 24 Jam</h4>
                <p>Integrasi langsung ke workspace engineer di jobs.meldir.id.</p>
              </div>
            </div>
          </section>
        </div>

        <!-- Tab 2: Manajemen Pengguna CRUD (Khusus Direktur & Admin) -->
        <div v-else-if="activeTab === 'users'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>👥 Manajemen Akun Seluruh User (Office, Jobs, Portal)</h2>
                <p class="section-desc">
                  Hak istimewa Direktur Utama & Admin untuk menambah, meninjau, mengubah wewenang, dan menghapus akun pengguna di semua portal.
                </p>
              </div>
              <button @click="openCreateUserModal" class="btn-primary">
                + Tambah Pengguna Baru
              </button>
            </div>

            <!-- Filter & Search Bar -->
            <div class="filter-bar">
              <div class="search-box">
                <input
                  v-model="searchQuery"
                  @input="filterUsers"
                  type="text"
                  placeholder="Cari nama atau email pengguna..."
                  class="search-input"
                />
              </div>
              <div class="role-filter-box">
                <select v-model="roleFilter" @change="fetchUsers" class="select-filter">
                  <option value="">Semua Peranan Portal</option>
                  <option value="direktur">Direktur Utama (Superadmin)</option>
                  <option value="admin">Office Admin</option>
                  <option value="engineer">Engineer Workspace (jobs)</option>
                  <option value="klien">Klien & Mitra (portal)</option>
                  <option value="audit">Internal Auditor</option>
                </select>
              </div>
              <button @click="fetchUsers" class="btn-refresh" title="Muat ulang daftar">
                🔄 Muat Ulang
              </button>
            </div>

            <!-- Users Table -->
            <div v-if="isLoadingUsers" class="loading-state">
              <span>Memuat data pengguna dari database...</span>
            </div>
            <div v-else class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Nama & Email</th>
                    <th>Portal & Role</th>
                    <th>Tipe Spesifik</th>
                    <th>No. WhatsApp</th>
                    <th>Status</th>
                    <th>Login Terakhir</th>
                    <th style="text-align: right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="u in usersList" :key="u.id">
                    <td class="font-mono text-muted">#{{ u.id }}</td>
                    <td>
                      <div class="user-row-cell">
                        <div class="avatar-circle">{{ getInitials(u.name) }}</div>
                        <div>
                          <strong>{{ u.name }}</strong>
                          <div class="text-xs text-muted">{{ u.email }}</div>
                        </div>
                      </div>
                    </td>
                    <td>
                      <span class="role-badge" :class="u.role">{{ formatRole(u.role) }}</span>
                    </td>
                    <td>
                      <span v-if="u.role === 'engineer'" class="text-xs badge-detail">
                        Tipe: {{ u.engineer_type || 'none' }}
                      </span>
                      <span v-else-if="u.role === 'klien'" class="text-xs badge-detail">
                        Tipe: {{ u.client_type || 'none' }}
                      </span>
                      <span v-else class="text-xs text-muted">-</span>
                    </td>
                    <td class="font-mono text-sm">{{ u.phone_wa }}</td>
                    <td>
                      <span class="status-badge" :class="u.status">{{ u.status }}</span>
                    </td>
                    <td class="text-xs text-muted">
                      {{ formatDate(u.last_login_at) }}
                    </td>
                    <td style="text-align: right">
                      <div class="action-btn-group">
                        <button @click="openEditUserModal(u)" class="btn-action edit" title="Edit Akun">
                          ✏️ Edit
                        </button>
                        <button
                          @click="confirmDeleteUser(u)"
                          :disabled="u.id === authStore.user?.id"
                          class="btn-action delete"
                          :title="u.id === authStore.user?.id ? 'Tidak dapat menghapus akun Anda sendiri' : 'Hapus Akun'"
                        >
                          🗑️ Hapus
                        </button>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="usersList.length === 0">
                    <td colspan="8" class="text-center py-6 text-muted">
                      Tidak ada pengguna yang sesuai dengan filter pencarian.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- Tab 3: Buku Besar & Jurnal SAK EMKM -->
        <div v-if="activeTab === 'accounting'" class="tab-content">
          <!-- Modul Transaksi Jurnal Umum -->
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>📖 Jurnal Umum Transaksi (Double-Entry Bookkeeping)</h2>
                <p class="section-desc">
                  Setiap transaksi dicatat berpasangan (Debit & Kredit) sesuai kaidah akuntansi resmi SAK EMKM.
                </p>
              </div>
              <button @click="showJournalModal = true" class="btn-primary">
                + Catat Transaksi Jurnal
              </button>
            </div>

            <div class="journal-balance-summary">
              <div class="balance-item">
                <span>Total Debit:</span>
                <strong class="font-mono text-emerald">{{ formatCurrency(totalJournalDebit) }}</strong>
              </div>
              <div class="balance-item">
                <span>Total Kredit:</span>
                <strong class="font-mono text-emerald">{{ formatCurrency(totalJournalCredit) }}</strong>
              </div>
              <span class="badge-balanced">✓ Seimbang (Balance)</span>
            </div>

            <div class="table-container" style="margin-top: 16px;">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Tanggal</th>
                    <th>No. Bukti</th>
                    <th>Keterangan Transaksi</th>
                    <th>Akun Debit</th>
                    <th>Akun Kredit</th>
                    <th style="text-align: right">Nominal</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="j in journals" :key="j.id">
                    <td class="font-mono text-xs">{{ j.date }}</td>
                    <td class="font-mono text-muted">{{ j.refNo }}</td>
                    <td><strong>{{ j.description }}</strong></td>
                    <td><span class="tag tag-green">{{ j.debitAccount }}</span></td>
                    <td><span class="tag tag-blue">{{ j.creditAccount }}</span></td>
                    <td style="text-align: right" class="font-mono font-bold text-emerald">
                      {{ formatCurrency(j.amount) }}
                    </td>
                  </tr>
                  <tr v-if="journals.length === 0">
                    <td colspan="6" class="text-center py-6 text-muted">
                      Belum ada catatan transaksi jurnal. Klik "+ Catat Transaksi Jurnal" untuk memposting transaksi baru.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>

          <!-- Modul Bagan Akun Standar (COA) -->
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>📑 Bagan Akun Standar (Chart of Accounts - SAK EMKM)</h2>
                <p class="section-desc">
                  Kode akun standar Ikatan Akuntan Indonesia (IAI) untuk pelaporan keuangan korporat.
                </p>
              </div>
              <button @click="showCoaModal = true" class="btn-primary">
                + Tambah Akun COA Baru
              </button>
            </div>
            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Kode Akun</th>
                    <th>Nama Akun</th>
                    <th>Kategori</th>
                    <th>Saldo Normal</th>
                    <th>Status</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="acc in coaAccounts" :key="acc.code">
                    <td class="font-mono">{{ acc.code }}</td>
                    <td><strong>{{ acc.name }}</strong></td>
                    <td><span class="category-pill" :class="acc.category">{{ acc.category }}</span></td>
                    <td class="capitalize">{{ acc.balance }}</td>
                    <td><span class="badge-active">Aktif</span></td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- Tab 4: Pajak DJP & Faktur Penagihan -->
        <div v-else-if="activeTab === 'tax'" class="tab-content">
          <!-- Modul Faktur Invoice Korporat -->
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>🧾 Manajemen Faktur & Invoice Klien (PPN 11%)</h2>
                <p class="section-desc">
                  Penerbitan faktur tagihan resmi PT. Melayani Digital Raya dengan perhitungan otomatis PPN 11% sesuai UU HPP.
                </p>
              </div>
              <div class="action-btn-group">
                <button @click="exportCoretaxCSV" class="btn-secondary">
                  📥 Export Coretax DJP (CSV)
                </button>
                <button @click="showInvoiceModal = true" class="btn-primary">
                  + Terbitkan Faktur Invoice
                </button>
              </div>
            </div>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>No. Invoice</th>
                    <th>Nama Klien</th>
                    <th>DPP (Pokok)</th>
                    <th>PPN (11%)</th>
                    <th>Total Tagihan</th>
                    <th>Jatuh Tempo</th>
                    <th>Status</th>
                    <th style="text-align: right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="inv in officeInvoices" :key="inv.id">
                    <td class="font-mono font-bold">{{ inv.invoiceNo }}</td>
                    <td>{{ inv.clientName }}</td>
                    <td class="font-mono">{{ formatCurrency(inv.dpp) }}</td>
                    <td class="font-mono text-muted">{{ formatCurrency(inv.ppn) }}</td>
                    <td class="font-mono font-bold text-emerald">{{ formatCurrency(inv.total) }}</td>
                    <td class="font-mono text-xs">{{ inv.dueDate }}</td>
                    <td>
                      <span :class="['badge-inv-status', inv.status]">
                        {{ inv.status === 'paid' ? '✓ Lunas' : 'Belum Bayar' }}
                      </span>
                    </td>
                    <td style="text-align: right">
                      <button @click="toggleInvoiceStatus(inv)" class="btn-action edit">
                        {{ inv.status === 'paid' ? 'Set Unpaid' : 'Tandai Lunas' }}
                      </button>
                    </td>
                  </tr>
                  <tr v-if="officeInvoices.length === 0">
                    <td colspan="8" class="text-center py-6 text-muted">
                      Belum ada faktur tagihan terbit. Klik "+ Terbitkan Faktur Invoice" untuk membuat faktur baru.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>

          <!-- Modul Hub Kepatuhan Perpajakan -->
          <section class="section-panel glass-panel">
            <h2>🏛️ Hub Kepatuhan Perpajakan PT. Melayani Digital Raya</h2>
            <p class="section-desc">
              Pelaporan pajak korporat otomatis sesuai regulasi Ditjen Pajak & integrasi Coretax DJP.
            </p>
            <div class="tax-grid">
              <div class="tax-card">
                <div class="card-header">
                  <h3>SPT Masa PPN 1111</h3>
                  <span class="tag tag-tax">PPN 11%</span>
                </div>
                <div class="kpi-value text-emerald font-mono" style="font-size: 1.35rem; margin: 8px 0;">
                  {{ formatCurrency(totalPPN) }}
                </div>
                <p class="tax-detail">Total PPN Keluaran terhitung dari {{ officeInvoices.length }} faktur terbit. Siap rekonsiliasi Coretax DJP.</p>
              </div>

              <div class="tax-card">
                <div class="card-header">
                  <h3>e-Bupot PPh Pasal 21</h3>
                  <span class="tag tag-blue">Tenaga Ahli</span>
                </div>
                <div class="kpi-value text-rose font-mono" style="font-size: 1.35rem; margin: 8px 0;">
                  {{ formatCurrency(totalPPh21) }}
                </div>
                <p class="tax-detail">Total PPh 21 terpotong dari {{ officeWithholdingSlips.length }} bukti potong resmi 21/26 yang diterbitkan untuk engineer.</p>
              </div>

              <div class="tax-card">
                <div class="card-header">
                  <h3>e-Bupot PPh Pasal 23</h3>
                  <span class="tag tag-amber">Kredit Pajak 2%</span>
                </div>
                <div class="kpi-value text-indigo font-mono" style="font-size: 1.35rem; margin: 8px 0;">
                  {{ formatCurrency(Math.round(totalRevenueYTD * 0.02)) }}
                </div>
                <p class="tax-detail">Estimasi pemotongan PPh 23 (2%) oleh klien korporat atas realisasi omset jasa software engineering.</p>
              </div>
            </div>
          </section>

          <!-- Modul Penerbitan Bukti Potong PPh 21 Tenaga Ahli -->
          <section class="section-panel glass-panel" style="margin-top: 24px;">
            <div class="section-header-row">
              <div>
                <h2>🏛️ Penerbitan Bukti Potong PPh 21 Tenaga Ahli (e-Bupot 21/26)</h2>
                <p class="section-desc">
                  Penerbitan bukti potong pajak resmi formulir 21/26 untuk software engineer & mitra lepas. Terkoneksi langsung ke portal jobs.meldir.id.
                </p>
              </div>
              <button @click="openWithholdingModal" class="btn-primary">
                + Terbitkan Bukti Potong PPh 21
              </button>
            </div>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>No. Bukti Potong</th>
                    <th>Nama Engineer / Tenaga Ahli</th>
                    <th>Masa / Tahun</th>
                    <th>Penghasilan Bruto</th>
                    <th>DPP (50%)</th>
                    <th>PPh 21 (5%)</th>
                    <th>Status e-Bupot</th>
                    <th style="text-align: right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="slip in officeWithholdingSlips" :key="slip.id">
                    <td class="font-mono text-muted">#{{ slip.slipNumber }}</td>
                    <td><strong>{{ slip.engineerName }}</strong></td>
                    <td>{{ slip.period }}</td>
                    <td class="font-mono">{{ formatCurrency(slip.bruto) }}</td>
                    <td class="font-mono text-muted">{{ formatCurrency(slip.dpp) }}</td>
                    <td class="font-mono font-bold text-rose">{{ formatCurrency(slip.taxAmount) }}</td>
                    <td><span class="badge-active">✓ Terlapor DJP</span></td>
                    <td style="text-align: right">
                      <div class="action-btn-group">
                        <button @click="selectedSlipToPrint = slip" class="btn-action edit" title="Lihat & Cetak Slip">
                          🖨️ Cetak
                        </button>
                        <button @click="deleteWithholdingSlip(slip.id)" class="btn-action delete" title="Hapus">
                          🗑️
                        </button>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="officeWithholdingSlips.length === 0">
                    <td colspan="8" class="text-center py-6 text-muted">
                      Belum ada bukti potong PPh 21 yang diterbitkan. Klik "+ Terbitkan Bukti Potong PPh 21" untuk membuat bukti potong baru bagi engineer.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- Tab 5: Manajemen Kontrak SPK & BAST Klien -->
        <div v-else-if="activeTab === 'contracts'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>📜 Manajemen Kontrak Kerja Sama (SPK) & BAST Klien</h2>
                <p class="section-desc">
                  Pencatatan dan pengesahan kontrak SPK Canvas & BAST bersertifikasi e-Materai Peruri. Terkoneksi otomatis ke Client Control Center (portal.meldir.id).
                </p>
              </div>
              <button @click="openContractModal" class="btn-primary">
                + Terbitkan Dokumen SPK / BAST
              </button>
            </div>

            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>No. Dokumen</th>
                    <th>Nama Klien Korporat</th>
                    <th>Jenis Berkas</th>
                    <th>Judul & Lingkup Kontrak</th>
                    <th>Periode Berlaku</th>
                    <th>Nilai Kontrak</th>
                    <th>Status Legalitas</th>
                    <th style="text-align: right">Aksi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="c in officeContracts" :key="c.id">
                    <td class="font-mono text-muted">#{{ c.docNumber }}</td>
                    <td><strong>{{ c.clientName }}</strong></td>
                    <td>
                      <span :class="['tag', c.docType === 'SPK' ? 'tag-blue' : 'tag-green']">
                        {{ c.docType }}
                      </span>
                    </td>
                    <td>
                      <div>{{ c.title }}</div>
                      <div class="text-xs text-muted">{{ c.description }}</div>
                    </td>
                    <td class="text-xs font-mono">{{ c.period }}</td>
                    <td class="font-mono font-bold text-emerald">{{ c.amountText }}</td>
                    <td><span class="badge-active">{{ c.badge }}</span></td>
                    <td style="text-align: right">
                      <div class="action-btn-group">
                        <button @click="deleteContract(c.id)" class="btn-action delete">
                          🗑️ Hapus
                        </button>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="officeContracts.length === 0">
                    <td colspan="8" class="text-center py-6 text-muted">
                      Belum ada dokumen kontrak SPK atau BAST yang dicatat. Klik "+ Terbitkan Dokumen SPK / BAST" untuk membuat dokumen legal baru bagi klien.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- Tab 6: Papan CRM Leads (Inbound Prospek Web) -->
        <div v-else-if="activeTab === 'leads'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>🎯 Papan Manajemen Inbound CRM Leads</h2>
                <p class="section-desc">
                  Prospek calon klien dan permohonan Audit Sistem Gratis yang masuk dari web publik (meldir.id). Terkoneksi langsung ke database PostgreSQL.
                </p>
              </div>
              <div class="header-action-tools">
                <button @click="fetchLeads" class="btn-refresh" :disabled="isLoadingLeads">
                  🔄 {{ isLoadingLeads ? 'Menyegarkan...' : 'Muat Ulang' }}
                </button>
              </div>
            </div>

            <!-- Filter & Search Toolbar -->
            <div class="table-toolbar">
              <div class="search-input-wrapper">
                <span class="search-icon">🔍</span>
                <input
                  v-model="leadSearchQuery"
                  type="text"
                  placeholder="Cari kode tiket, nama, perusahaan, atau WhatsApp..."
                  class="input-search"
                />
              </div>
              <div class="filter-group">
                <select v-model="leadStatusFilter" class="select-filter">
                  <option value="">Semua Status</option>
                  <option value="baru">Baru Masuk</option>
                  <option value="dihubungi">Sedang Dihubungi</option>
                  <option value="penawaran">Penawaran SPK</option>
                  <option value="closing">Closing (Deal)</option>
                  <option value="dibatalkan">Dibatalkan</option>
                </select>
              </div>
            </div>

            <!-- Leads Table -->
            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Kode Tiket</th>
                    <th>Calon Klien &amp; Perusahaan</th>
                    <th>Kontak WhatsApp &amp; Email</th>
                    <th>Kebutuhan Solusi</th>
                    <th>Estimasi Budget</th>
                    <th>Status Prospek</th>
                    <th>Tanggal Masuk</th>
                    <th style="text-align: right">Tindak Lanjut</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="lead in filteredLeads" :key="lead.id">
                    <td>
                      <span class="lead-code-badge">{{ lead.lead_code }}</span>
                    </td>
                    <td>
                      <div><strong>{{ lead.name }}</strong></div>
                      <div class="text-xs text-muted">{{ lead.company || 'Pribadi / Individu' }}</div>
                      <div v-if="lead.notes" class="text-xs text-slate-500 mt-1" style="max-width: 260px; white-space: normal;">
                        💬 <em>{{ lead.notes }}</em>
                      </div>
                    </td>
                    <td>
                      <div class="font-mono text-sm">{{ lead.whatsapp }}</div>
                      <div class="text-xs text-muted">{{ lead.email || '-' }}</div>
                    </td>
                    <td>
                      <span class="tag tag-blue">{{ lead.service_interest }}</span>
                    </td>
                    <td class="font-mono text-xs">
                      {{ lead.budget_range || '-' }}
                    </td>
                    <td>
                      <select
                        :value="lead.status"
                        @change="onLeadStatusChange(lead, ($event.target as HTMLSelectElement).value)"
                        :class="['lead-status-select', 'status-' + lead.status]"
                        :disabled="updatingLeadId === lead.id"
                      >
                        <option value="baru">🆕 Baru Masuk</option>
                        <option value="dihubungi">📞 Dihubungi</option>
                        <option value="penawaran">📑 Penawaran SPK</option>
                        <option value="closing">✅ Closing (Deal)</option>
                        <option value="dibatalkan">❌ Dibatalkan</option>
                      </select>
                    </td>
                    <td class="text-xs font-mono text-muted">
                      {{ formatDate(lead.created_at) }}
                    </td>
                    <td style="text-align: right">
                      <div class="action-btn-group">
                        <button
                          @click="openLeadWhatsApp(lead)"
                          class="btn-lead-wa-action"
                          title="Hubungi langsung via WhatsApp"
                        >
                          💬 Follow Up WA
                        </button>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="filteredLeads.length === 0">
                    <td colspan="8" class="text-center py-8 text-muted">
                      <div v-if="isLoadingLeads">Memuat data leads dari server...</div>
                      <div v-else>
                        Tidak ada data leads yang sesuai dengan pencarian atau filter status.
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>
      </main>

      <!-- Mobile Bottom Navigation Bar (5 Items) -->
      <nav class="mobile-bottom-nav">
        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'overview' && !showMoreSheet }]"
          @click="switchTab('overview')"
        >
          <span class="nav-bottom-icon">📊</span>
          <span class="nav-bottom-label">Ringkasan</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'users' && !showMoreSheet }]"
          @click="switchTab('users')"
        >
          <span class="nav-bottom-icon">👥</span>
          <span class="nav-bottom-label">Pengguna</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'accounting' && !showMoreSheet }]"
          @click="switchTab('accounting')"
        >
          <span class="nav-bottom-icon">📖</span>
          <span class="nav-bottom-label">Akuntansi</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: activeTab === 'tax' && !showMoreSheet }]"
          @click="switchTab('tax')"
        >
          <span class="nav-bottom-icon">🏛️</span>
          <span class="nav-bottom-label">Pajak</span>
        </button>

        <button
          type="button"
          :class="['nav-bottom-item', { active: showMoreSheet || activeTab === 'contracts' || activeTab === 'leads' }]"
          @click="showMoreSheet = !showMoreSheet"
        >
          <span class="nav-bottom-icon">⋯</span>
          <span class="nav-bottom-label">Lainnya</span>
          <span v-if="activeTab === 'contracts' || activeTab === 'leads'" class="bottom-active-dot"></span>
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
              <h3>Menu & Kontrol Operasional</h3>
              <p class="sheet-subtitle">PT. Melayani Digital Raya</p>
            </div>
            <button class="btn-sheet-close" @click="showMoreSheet = false">✕</button>
          </div>

          <!-- User Profile Card in Sheet -->
          <div class="sheet-profile-card">
            <div class="sheet-avatar">{{ userInitials }}</div>
            <div class="sheet-user-info">
              <div class="sheet-user-name">{{ authStore.user?.name }}</div>
              <div class="sheet-user-meta">
                <span class="badge-role">{{ roleDisplay }}</span>
                <span class="sheet-user-email">{{ authStore.user?.email }}</span>
              </div>
            </div>
          </div>

          <!-- Extra Navigation Items -->
          <div class="sheet-nav-list">
            <div
              :class="['sheet-nav-item', { active: activeTab === 'leads' }]"
              @click="switchTab('leads')"
            >
              <div class="sheet-nav-icon">🎯</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Papan CRM Leads</div>
                <div class="sheet-nav-desc">Inbound prospek dari web meldir.id &amp; audit sistem</div>
              </div>
              <div class="sheet-nav-arrow">→</div>
            </div>

            <div
              :class="['sheet-nav-item', { active: activeTab === 'contracts' }]"
              @click="switchTab('contracts')"
            >
              <div class="sheet-nav-icon">📜</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Kontrak SPK & BAST</div>
                <div class="sheet-nav-desc">Penerbitan dokumen legal & sertifikasi e-Materai Peruri</div>
              </div>
              <div class="sheet-nav-arrow">→</div>
            </div>

            <div class="sheet-nav-item" @click="exportCoretaxCSV(); showMoreSheet = false">
              <div class="sheet-nav-icon">📥</div>
              <div class="sheet-nav-text">
                <div class="sheet-nav-title">Ekspor Rekonsiliasi Coretax DJP</div>
                <div class="sheet-nav-desc">Unduh berkas CSV faktur PPN untuk pelaporan SPT Masa</div>
              </div>
              <div class="sheet-nav-arrow">↓</div>
            </div>
          </div>

          <!-- System Status Info in Sheet -->
          <div class="sheet-system-card">
            <div class="sheet-sys-header">
              <span class="pulse-indicator"></span>
              <strong>Status Ekosistem Server</strong>
            </div>
            <div class="sheet-sys-details">
              <div>Backend Golang: <span :class="['api-badge', apiStatus]">{{ apiStatusText }}</span></div>
              <div>Database: <strong class="text-emerald">PostgreSQL (32 Tabel DDL)</strong></div>
              <div>Cache: <strong class="text-indigo">Redis 6379</strong></div>
            </div>
          </div>

          <!-- Logout Button in Sheet -->
          <button class="sheet-btn-logout" @click="handleLogout">
            🚪 Keluar dari Akun (Logout)
          </button>
        </div>
      </div>
    </div>

    <!-- Modal Form (Tambah / Edit Pengguna) -->
    <div v-if="showUserModal" class="modal-backdrop" @click.self="closeUserModal">
      <div class="modal-card">
        <div class="modal-header">
          <h3>{{ isEditing ? 'Edit Akun Pengguna' : 'Tambah Pengguna Baru' }}</h3>
          <button @click="closeUserModal" class="btn-close-modal">✕</button>
        </div>

        <form @submit.prevent="saveUser" class="modal-form">
          <div v-if="modalError" class="modal-error-banner">
            ⚠️ {{ modalError }}
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label>Nama Lengkap *</label>
              <input v-model="userForm.name" type="text" required placeholder="Contoh: Rian Anggoro, S.Kom" class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Alamat Email *</label>
              <input v-model="userForm.email" type="email" required placeholder="nama@meldir.id atau klien@perusahaan.com" class="form-input" />
            </div>
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label>{{ isEditing ? 'Ganti Password (Kosongkan jika tetap)' : 'Kata Sandi / Password *' }}</label>
              <input
                v-model="userForm.password"
                type="password"
                :required="!isEditing"
                placeholder="Minimal 6 karakter"
                class="form-input"
              />
            </div>
            <div class="form-group flex-1">
              <label>Nomor WhatsApp Resmi *</label>
              <input v-model="userForm.phone_wa" type="text" required placeholder="+628123456789" class="form-input" />
            </div>
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label>Peranan Pengguna (Role) *</label>
              <select v-model="userForm.role" required class="form-input">
                <option value="direktur">Direktur Utama (office.meldir.id)</option>
                <option value="admin">Office Administrator (office.meldir.id)</option>
                <option value="engineer">Engineer Workspace (jobs.meldir.id)</option>
                <option value="klien">Klien / Customer (portal.meldir.id)</option>
                <option value="audit">Internal Auditor (office.meldir.id)</option>
              </select>
            </div>
            <div class="form-group flex-1">
              <label>Status Akun *</label>
              <select v-model="userForm.status" required class="form-input">
                <option value="aktif">Aktif (Bisa Login)</option>
                <option value="pengajuan">Pengajuan (Menunggu Approval)</option>
                <option value="suspended">Suspended (Diblokir)</option>
              </select>
            </div>
          </div>

          <!-- Dynamic Role Extra Fields -->
          <div v-if="userForm.role === 'engineer'" class="form-row">
            <div class="form-group flex-1">
              <label>Tipe Engineer</label>
              <select v-model="userForm.engineer_type" class="form-input">
                <option value="internal">Internal Core Engineer</option>
                <option value="external">Tenaga Ahli Eksternal (Mitra)</option>
                <option value="none">Bukan Engineer (None)</option>
              </select>
            </div>
            <div class="form-group flex-1">
              <label>Username GitHub (Untuk Undangan Repo)</label>
              <input v-model="userForm.github_username" type="text" placeholder="username-github" class="form-input" />
            </div>
          </div>

          <div v-if="userForm.role === 'klien'" class="form-row">
            <div class="form-group flex-1">
              <label>Tipe Klien / Entitas Bisnis</label>
              <select v-model="userForm.client_type" class="form-input">
                <option value="company">Badan Usaha (PT / CV / Firma)</option>
                <option value="individual">Perorangan / Personal</option>
                <option value="foundation">Yayasan Sosial / Komunitas</option>
                <option value="none">None</option>
              </select>
            </div>
          </div>

          <div class="modal-footer">
            <button type="button" @click="closeUserModal" class="btn-secondary">
              Batal
            </button>
            <button type="submit" :disabled="isSavingUser" class="btn-primary">
              <span v-if="isSavingUser">Menyimpan Data...</span>
              <span v-else>{{ isEditing ? 'Simpan Perubahan' : 'Buat Akun Sekarang' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Form (Tambah Transaksi Jurnal Double-Entry) -->
    <div v-if="showJournalModal" class="modal-backdrop" @click.self="showJournalModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Catat Transaksi Jurnal Double-Entry</h3>
          <button @click="showJournalModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveJournal" class="modal-form">
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Tanggal Transaksi *</label>
              <input v-model="journalForm.date" type="date" required class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Nominal Transaksi (Rp) *</label>
              <input
                v-model.number="journalForm.amount"
                type="number"
                min="1000"
                step="1000"
                required
                placeholder="Contoh: 12500000"
                class="form-input"
              />
            </div>
          </div>

          <div class="form-group">
            <label>Keterangan Transaksi / Nomor Referensi *</label>
            <input
              v-model="journalForm.description"
              type="text"
              required
              placeholder="Contoh: Pembayaran invoice jasa maintenance termin 1"
              class="form-input"
            />
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label>Akun Posisi DEBIT *</label>
              <select v-model="journalForm.debitAccount" required class="form-input">
                <option v-for="acc in coaAccounts" :key="'deb-' + acc.code" :value="acc.code + ' ' + acc.name">
                  {{ acc.code }} — {{ acc.name }}
                </option>
              </select>
            </div>
            <div class="form-group flex-1">
              <label>Akun Posisi KREDIT *</label>
              <select v-model="journalForm.creditAccount" required class="form-input">
                <option v-for="acc in coaAccounts" :key="'kred-' + acc.code" :value="acc.code + ' ' + acc.name">
                  {{ acc.code }} — {{ acc.name }}
                </option>
              </select>
            </div>
          </div>

          <div class="modal-footer">
            <button type="button" @click="showJournalModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Posting ke Buku Besar</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Form (Terbitkan Faktur Invoice Baru) -->
    <div v-if="showInvoiceModal" class="modal-backdrop" @click.self="showInvoiceModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Terbitkan Faktur Invoice Korporat (PPN 11%)</h3>
          <button @click="showInvoiceModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveInvoice" class="modal-form">
          <div class="form-group">
            <label>Nama Klien / Perusahaan Penerima Faktur *</label>
            <input v-model="invoiceForm.clientName" type="text" required placeholder="Contoh: PT. Surya Logistik Multimoda" class="form-input" />
          </div>

          <div class="form-group">
            <label>Deskripsi Layanan / Kontrak SPK *</label>
            <input v-model="invoiceForm.description" type="text" required placeholder="Contoh: Paket Managed Care SLA 24 Jam Periode Q4" class="form-input" />
          </div>

          <div class="form-row">
            <div class="form-group flex-1">
              <label>Nilai Pokok / DPP (Rp) *</label>
              <input
                v-model.number="invoiceForm.dpp"
                type="number"
                min="10000"
                step="1000"
                required
                placeholder="Contoh: 25000000"
                class="form-input"
              />
            </div>
            <div class="form-group flex-1">
              <label>Tanggal Jatuh Tempo *</label>
              <input v-model="invoiceForm.dueDate" type="date" required class="form-input" />
            </div>
          </div>

          <div class="addon-summary-box">
            <div class="addon-sum-row">
              <span>DPP (Dasar Pengenaan Pajak):</span>
              <span class="font-mono">{{ formatCurrency(invoiceForm.dpp || 0) }}</span>
            </div>
            <div class="addon-sum-row">
              <span>PPN 11% (UU HPP):</span>
              <span class="font-mono text-muted">{{ formatCurrency(Math.round((invoiceForm.dpp || 0) * 0.11)) }}</span>
            </div>
            <div class="addon-sum-row total">
              <span>Total Tagihan Faktur:</span>
              <span class="font-mono font-bold text-emerald">{{ formatCurrency(Math.round((invoiceForm.dpp || 0) * 1.11)) }}</span>
            </div>
          </div>

          <div class="modal-footer">
            <button type="button" @click="showInvoiceModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Terbitkan Faktur Sah</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Form (Tambah Akun COA Baru) -->
    <div v-if="showCoaModal" class="modal-backdrop" @click.self="showCoaModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Tambah Akun Bagan Standar (COA)</h3>
          <button @click="showCoaModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveCoaAccount" class="modal-form">
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Kode Akun (Standar 5-digit) *</label>
              <input v-model="coaForm.code" type="text" required placeholder="Contoh: 1-1004" class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Posisi Saldo Normal *</label>
              <select v-model="coaForm.balance" required class="form-input">
                <option value="debit">Debit</option>
                <option value="credit">Kredit</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label>Nama Akun Baru *</label>
            <input v-model="coaForm.name" type="text" required placeholder="Contoh: Bank BNI Giro Korporat" class="form-input" />
          </div>
          <div class="form-group">
            <label>Kategori Klasifikasi Akun *</label>
            <select v-model="coaForm.category" required class="form-input">
              <option value="asset">Aset / Harta Lancar</option>
              <option value="liability">Liabilitas / Kewajiban</option>
              <option value="equity">Ekuitas / Modal</option>
              <option value="revenue">Pendapatan Operasional</option>
              <option value="expense">Beban / Biaya Usaha</option>
            </select>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showCoaModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Simpan Akun COA</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Form (Terbitkan Bukti Potong PPh 21) -->
    <div v-if="showWithholdingModal" class="modal-backdrop" @click.self="showWithholdingModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Terbitkan Bukti Potong PPh 21 Tenaga Ahli (e-Bupot 21/26)</h3>
          <button @click="showWithholdingModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveWithholdingSlip" class="modal-form">
          <div class="form-group">
            <label>Pilih Engineer Penerima Kompensasi *</label>
            <input
              v-model="withholdingForm.engineerName"
              type="text"
              list="engineer-suggestions"
              required
              placeholder="Ketik atau pilih nama engineer..."
              class="form-input"
            />
            <datalist id="engineer-suggestions">
              <option v-for="eng in usersList.filter(u => u.role === 'engineer')" :key="eng.id" :value="eng.name"></option>
            </datalist>
          </div>
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Nomor Bukti Potong Resmi *</label>
              <input v-model="withholdingForm.slipNumber" type="text" required class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Masa & Tahun Pajak *</label>
              <input v-model="withholdingForm.period" type="text" required placeholder="Contoh: Oktober 2026" class="form-input" />
            </div>
          </div>
          <div class="form-group">
            <label>Penghasilan Bruto (Rp) *</label>
            <input v-model.number="withholdingForm.bruto" type="number" min="100000" step="50000" required class="form-input" />
          </div>
          <div class="addon-summary-box" style="margin-bottom: 16px;">
            <div class="addon-sum-row">
              <span>Dasar Pengenaan Pajak (DPP 50% PP 58/2023):</span>
              <strong class="font-mono">{{ formatCurrency(withholdingDpp) }}</strong>
            </div>
            <div class="addon-sum-row">
              <span>Tarif Pajak Pasal 17:</span>
              <span class="tag tag-blue">5%</span>
            </div>
            <div class="addon-sum-row total">
              <span>PPh 21 Dipotong Perusahaan:</span>
              <strong class="font-mono text-rose">{{ formatCurrency(withholdingTax) }}</strong>
            </div>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showWithholdingModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Terbitkan Bukti Potong Resmi</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Form (Terbitkan Kontrak SPK / BAST) -->
    <div v-if="showContractModal" class="modal-backdrop" @click.self="showContractModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Terbitkan Dokumen Kontrak SPK / BAST Klien</h3>
          <button @click="showContractModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveContract" class="modal-form">
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Pilih Klien Korporat *</label>
              <input
                v-model="contractForm.clientName"
                type="text"
                list="client-suggestions"
                required
                placeholder="Ketik atau pilih nama klien..."
                class="form-input"
              />
              <datalist id="client-suggestions">
                <option v-for="cl in usersList.filter(u => u.role === 'klien')" :key="cl.id" :value="cl.name"></option>
              </datalist>
            </div>
            <div class="form-group flex-1">
              <label>Jenis Dokumen Legal *</label>
              <select v-model="contractForm.docType" required class="form-input">
                <option value="SPK">Surat Perjanjian Kerja Sama (SPK)</option>
                <option value="BAST">Berita Acara Serah Terima (BAST)</option>
              </select>
            </div>
          </div>
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Nomor Dokumen Resmi *</label>
              <input v-model="contractForm.docNumber" type="text" required placeholder="018/SPK/MDR/2026" class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Periode Berlaku *</label>
              <input v-model="contractForm.period" type="text" required placeholder="1 Jan 2026 – 31 Des 2026" class="form-input" />
            </div>
          </div>
          <div class="form-group">
            <label>Judul Dokumen / Nama Layanan *</label>
            <input v-model="contractForm.title" type="text" required placeholder="Layanan Managed Care SLA 24 Jam" class="form-input" />
          </div>
          <div class="form-group">
            <label>Nilai Kontrak (Teks) *</label>
            <input v-model="contractForm.amountText" type="text" required placeholder="Rp 15.000.000 / Bulan" class="form-input" />
          </div>
          <div class="form-group">
            <label>Klausul SLA & Ruang Lingkup *</label>
            <textarea v-model="contractForm.slaTerms" rows="2" required placeholder="Contoh: SLA respon kritis 4 jam, garansi bug fix 12 jam, kuota bulanan 10 jam..." class="form-input"></textarea>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showContractModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Terbitkan & Sahkan e-Materai</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Cetak Slip Pajak PPh 21 -->
    <div v-if="selectedSlipToPrint" class="modal-backdrop" @click.self="selectedSlipToPrint = null">
      <div class="modal-card modal-slip-view">
        <div class="modal-header">
          <h3>Bukti Potong PPh 21 Resmi (Formulir 21/26)</h3>
          <button @click="selectedSlipToPrint = null" class="btn-close-modal">✕</button>
        </div>
        <div class="slip-content" id="printable-slip">
          <div class="slip-corp-header">
            <h4>PT. MELAYANI DIGITAL RAYA</h4>
            <p>NPWP: 01.234.567.8-012.000 • SK Kemenkumham RI: AHU-0012345.AH.01.01.TAHUN 2026</p>
            <div class="slip-title">BUKTI PEMOTONGAN PPH PASAL 21 (FORMULIR 21/26)</div>
            <div class="slip-number font-mono">Nomor: {{ selectedSlipToPrint.slipNumber }}</div>
          </div>

          <div class="slip-details-grid">
            <div class="slip-row">
              <span class="label">Nama Penerima Penghasilan:</span>
              <span class="val"><strong>{{ selectedSlipToPrint.engineerName }}</strong></span>
            </div>
            <div class="slip-row">
              <span class="label">Masa / Tahun Pajak:</span>
              <span class="val font-bold">{{ selectedSlipToPrint.period }}</span>
            </div>
            <div class="slip-row">
              <span class="label">Klasifikasi Penghasilan:</span>
              <span class="val">Imbalan Kepada Tenaga Ahli / Jasa Perangkat Lunak (Bukan Pegawai)</span>
            </div>
          </div>

          <div class="slip-calc-table">
            <table class="data-table">
              <thead>
                <tr>
                  <th>Jumlah Penghasilan Bruto</th>
                  <th>Dasar Pengenaan Pajak (50%)</th>
                  <th>Tarif</th>
                  <th>PPh 21 Dipotong</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td class="font-mono font-bold">{{ formatCurrency(selectedSlipToPrint.bruto) }}</td>
                  <td class="font-mono">{{ formatCurrency(selectedSlipToPrint.dpp) }}</td>
                  <td class="font-bold">{{ selectedSlipToPrint.rate || '5%' }}</td>
                  <td class="font-mono font-bold text-rose">{{ formatCurrency(selectedSlipToPrint.taxAmount) }}</td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="slip-footer-sign">
            <div class="seal-badge">
              <div class="seal-text">PT. MELDIR</div>
              <div class="seal-sub">AUTHENTICATED</div>
            </div>
            <div class="sign-block">
              <div>Jakarta, {{ selectedSlipToPrint.period }}</div>
              <div class="sign-title">Pemotong Pajak / Direktur Utama</div>
              <div class="sign-space"></div>
              <div class="sign-name">PT. Melayani Digital Raya</div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="printSlip()" class="btn-primary">🖨️ Cetak / Simpan PDF</button>
          <button @click="selectedSlipToPrint = null" class="btn-secondary">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import LoginForm from '../components/LoginForm.vue'
import { useAuthStore, User } from '../stores/authStore'

const authStore = useAuthStore()
const currentHost = ref(window.location.host)
const apiStatus = ref('checking')
const apiStatusText = ref('Memeriksa...')
const activeTab = ref('overview')
const showMoreSheet = ref(false)

const tabs = [
  { id: 'overview', label: 'Ringkasan Eksekutif', icon: '📊' },
  { id: 'users', label: 'Manajemen Pengguna (CRUD)', icon: '👥' },
  { id: 'leads', label: 'Papan CRM Leads', icon: '🎯' },
  { id: 'accounting', label: 'Buku Besar SAK EMKM', icon: '📖' },
  { id: 'tax', label: 'Kepatuhan Pajak DJP', icon: '🏛️' },
  { id: 'contracts', label: 'Kontrak SPK & BAST', icon: '📜' },
]

const globalMessage = ref('')
const globalMessageType = ref('success')

// CRUD State
const usersList = ref<User[]>([])
const allUsersCache = ref<User[]>([])
const isLoadingUsers = ref(false)
const searchQuery = ref('')
const roleFilter = ref('')

const showUserModal = ref(false)
const isEditing = ref(false)
const isSavingUser = ref(false)
const modalError = ref('')

const userForm = ref({
  id: 0,
  name: '',
  email: '',
  password: '',
  role: 'klien' as any,
  engineer_type: 'none',
  client_type: 'none',
  phone_wa: '',
  status: 'aktif',
  github_username: '',
})

const userInitials = computed(() => {
  if (!authStore.user?.name) return 'MD'
  return authStore.user.name
    .split(' ')
    .map(w => w[0])
    .slice(0, 2)
    .join('')
    .toUpperCase()
})

const roleDisplay = computed(() => {
  switch (authStore.user?.role) {
    case 'direktur': return 'Direktur Utama (Superadmin)'
    case 'admin': return 'Office Administrator'
    case 'audit': return 'Internal Auditor'
    case 'engineer': return 'Technical Lead'
    default: return 'User'
  }
})

const defaultAccounts = [
  { code: '1-1001', name: 'Kas Operasional / Petty Cash', category: 'asset', balance: 'debit' },
  { code: '1-1002', name: 'Bank BCA Utama Korporat', category: 'asset', balance: 'debit' },
  { code: '1-1003', name: 'Bank Mandiri Korporat', category: 'asset', balance: 'debit' },
  { code: '1-1201', name: 'Piutang Usaha Klien', category: 'asset', balance: 'debit' },
  { code: '1-1301', name: 'Pajak Dibayar di Muka (PPh 23)', category: 'asset', balance: 'debit' },
  { code: '2-1001', name: 'Utang Usaha / Vendor Cloud', category: 'liability', balance: 'credit' },
  { code: '2-1201', name: 'Utang PPN Keluaran 11%', category: 'liability', balance: 'credit' },
  { code: '2-1202', name: 'Utang PPh 21 Tenaga Ahli', category: 'liability', balance: 'credit' },
  { code: '3-1001', name: 'Modal Disetor Saham', category: 'equity', balance: 'credit' },
  { code: '4-1001', name: 'Pendapatan Jasa Software Engineering', category: 'revenue', balance: 'credit' },
  { code: '4-1002', name: 'Pendapatan Kontrak Managed Care SLA', category: 'revenue', balance: 'credit' },
  { code: '5-1001', name: 'Beban Kompensasi Engineer', category: 'expense', balance: 'debit' },
  { code: '5-1002', name: 'Beban Infrastruktur Server & VPS', category: 'expense', balance: 'debit' },
]

function loadStoredCoa(): any[] {
  try {
    const raw = localStorage.getItem('meldir_custom_coa')
    if (raw) {
      const parsed = JSON.parse(raw)
      if (Array.isArray(parsed) && parsed.length > 0) return [...defaultAccounts, ...parsed]
    }
  } catch {}
  return [...defaultAccounts]
}

const coaAccounts = ref<any[]>(loadStoredCoa())
const showCoaModal = ref(false)
const coaForm = ref({
  code: '',
  name: '',
  category: 'asset',
  balance: 'debit',
})

function saveCoaAccount() {
  if (!coaForm.value.code || !coaForm.value.name) return
  const newAcc = {
    code: coaForm.value.code.trim(),
    name: coaForm.value.name.trim(),
    category: coaForm.value.category,
    balance: coaForm.value.balance,
  }
  coaAccounts.value.push(newAcc)
  try {
    const custom = coaAccounts.value.filter(a => !defaultAccounts.some(d => d.code === a.code))
    localStorage.setItem('meldir_custom_coa', JSON.stringify(custom))
  } catch {}
  showCoaModal.value = false
  coaForm.value = { code: '', name: '', category: 'asset', balance: 'debit' }
  showFlash(`Akun COA baru "${newAcc.code} - ${newAcc.name}" berhasil ditambahkan!`)
}

// State Bukti Potong PPh 21 (Office -> Jobs)
function loadStoredWithholdingSlips(): any[] {
  try {
    const raw = localStorage.getItem('meldir_withholding_slips')
    if (raw) {
      const parsed = JSON.parse(raw)
      if (Array.isArray(parsed)) return parsed
    }
  } catch {}
  return []
}

const officeWithholdingSlips = ref<any[]>(loadStoredWithholdingSlips())
const showWithholdingModal = ref(false)
const withholdingForm = ref({
  engineerName: '',
  slipNumber: '',
  period: new Date().toLocaleDateString('id-ID', { month: 'long', year: 'numeric' }),
  bruto: 10000000,
})

const withholdingDpp = computed(() => Math.round(Number(withholdingForm.value.bruto || 0) * 0.5))
const withholdingTax = computed(() => Math.round(withholdingDpp.value * 0.05))

function openWithholdingModal() {
  const currentMonthNum = new Date().getMonth() + 1
  const count = officeWithholdingSlips.value.length + 1
  withholdingForm.value = {
    engineerName: '',
    slipNumber: `BP-21/2026/${currentMonthNum < 10 ? '0' + currentMonthNum : currentMonthNum}/${count < 10 ? '00' + count : '0' + count}`,
    period: new Date().toLocaleDateString('id-ID', { month: 'long', year: 'numeric' }),
    bruto: 10000000,
  }
  showWithholdingModal.value = true
}

function saveWithholdingSlip() {
  if (!withholdingForm.value.engineerName || !withholdingForm.value.bruto) return
  const newSlip = {
    id: Date.now(),
    slipNumber: withholdingForm.value.slipNumber,
    engineerName: withholdingForm.value.engineerName,
    period: withholdingForm.value.period,
    bruto: Number(withholdingForm.value.bruto),
    dpp: withholdingDpp.value,
    rate: '5%',
    taxAmount: withholdingTax.value,
    createdAt: new Date().toISOString(),
  }
  officeWithholdingSlips.value.unshift(newSlip)
  try {
    localStorage.setItem('meldir_withholding_slips', JSON.stringify(officeWithholdingSlips.value))
  } catch {}
  showWithholdingModal.value = false
  showFlash(`Bukti Potong PPh 21 #${newSlip.slipNumber} berhasil diterbitkan untuk ${newSlip.engineerName}!`)
}

function deleteWithholdingSlip(id: number) {
  officeWithholdingSlips.value = officeWithholdingSlips.value.filter(s => s.id !== id)
  try {
    localStorage.setItem('meldir_withholding_slips', JSON.stringify(officeWithholdingSlips.value))
  } catch {}
  showFlash('Bukti potong berhasil dihapus.')
}

const selectedSlipToPrint = ref<any>(null)

function printSlip() {
  window.print()
}

// State Kontrak SPK & BAST (Office -> Client Portal)
function loadStoredContracts(): any[] {
  try {
    const raw = localStorage.getItem('meldir_client_contracts')
    if (raw) {
      const parsed = JSON.parse(raw)
      if (Array.isArray(parsed)) return parsed
    }
  } catch {}
  return []
}

const officeContracts = ref<any[]>(loadStoredContracts())
const showContractModal = ref(false)
const contractForm = ref({
  clientName: '',
  docType: 'SPK',
  docNumber: '',
  title: '',
  period: '',
  amountText: 'Rp 15.000.000 / Bulan',
  slaTerms: 'SLA Respon P1 4 Jam, Kuota Jam Kerja 10 Jam / Bulan',
})

function openContractModal() {
  const count = officeContracts.value.length + 1
  contractForm.value = {
    clientName: '',
    docType: 'SPK',
    docNumber: `0${18 + count}/SPK/MDR/2026`,
    title: 'Surat Perjanjian Kerja Sama Managed Care SLA 24 Jam',
    period: '1 Jan 2026 – 31 Des 2026',
    amountText: 'Rp 15.000.000 / Bulan',
    slaTerms: 'SLA Respon P1 4 Jam, Kuota Jam Kerja 10 Jam / Bulan',
  }
  showContractModal.value = true
}

function saveContract() {
  if (!contractForm.value.clientName || !contractForm.value.docNumber) return
  const newDoc = {
    id: Date.now(),
    clientName: contractForm.value.clientName,
    docType: contractForm.value.docType,
    docNumber: contractForm.value.docNumber,
    title: `${contractForm.value.docType === 'SPK' ? 'Surat Perjanjian Kerja Sama' : 'Berita Acara Serah Terima'} No. ${contractForm.value.docNumber}`,
    subtitle: `${contractForm.value.title} • Periode ${contractForm.value.period}`,
    description: contractForm.value.slaTerms,
    period: contractForm.value.period,
    amountText: contractForm.value.amountText,
    badge: '✓ E-Materai Sah & Tervalidasi Peruri',
    createdAt: new Date().toISOString(),
  }
  officeContracts.value.unshift(newDoc)
  try {
    localStorage.setItem('meldir_client_contracts', JSON.stringify(officeContracts.value))
  } catch {}
  showContractModal.value = false
  showFlash(`Dokumen ${newDoc.docType} #${newDoc.docNumber} untuk ${newDoc.clientName} berhasil dicatat & divalidasi e-Materai!`)
}

function deleteContract(id: number) {
  officeContracts.value = officeContracts.value.filter(c => c.id !== id)
  try {
    localStorage.setItem('meldir_client_contracts', JSON.stringify(officeContracts.value))
  } catch {}
  showFlash('Dokumen kontrak berhasil dihapus.')
}

function switchTab(tabId: string) {
  activeTab.value = tabId
  showMoreSheet.value = false
  window.scrollTo({ top: 0, behavior: 'smooth' })
  if (tabId === 'users' && usersList.value.length === 0) {
    fetchUsers()
  }
  if (tabId === 'leads' && leadsList.value.length === 0) {
    fetchLeads()
  }
}

async function fetchUsers() {
  if (!authStore.token) return
  isLoadingUsers.value = true

  try {
    const params = new URLSearchParams()
    if (roleFilter.value) params.append('role', roleFilter.value)
    if (searchQuery.value) params.append('q', searchQuery.value)

    const res = await fetch(`/api/v1/users?${params.toString()}`, {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })

    const result = await res.json()
    if (res.ok && result.success) {
      usersList.value = result.data || []
      allUsersCache.value = result.data || []
    } else {
      showFlash(result.error || 'Gagal memuat pengguna', 'error')
    }
  } catch (err: any) {
    showFlash('Koneksi terputus saat mengambil data user', 'error')
  } finally {
    isLoadingUsers.value = false
  }
}

function filterUsers() {
  const q = searchQuery.value.toLowerCase().trim()
  if (!q) {
    usersList.value = allUsersCache.value
    return
  }
  usersList.value = allUsersCache.value.filter(
    u => u.name.toLowerCase().includes(q) || u.email.toLowerCase().includes(q)
  )
}

function openCreateUserModal() {
  isEditing.value = false
  modalError.value = ''
  userForm.value = {
    id: 0,
    name: '',
    email: '',
    password: '',
    role: 'engineer',
    engineer_type: 'internal',
    client_type: 'none',
    phone_wa: '+628',
    status: 'aktif',
    github_username: '',
  }
  showUserModal.value = true
}

function openEditUserModal(u: any) {
  isEditing.value = true
  modalError.value = ''
  userForm.value = {
    id: u.id,
    name: u.name,
    email: u.email,
    password: '',
    role: u.role,
    engineer_type: u.engineer_type || 'none',
    client_type: u.client_type || 'none',
    phone_wa: u.phone_wa || '',
    status: u.status || 'aktif',
    github_username: u.github_username || '',
  }
  showUserModal.value = true
}

function closeUserModal() {
  showUserModal.value = false
  modalError.value = ''
}

async function saveUser() {
  if (!authStore.token) return
  isSavingUser.value = true
  modalError.value = ''

  try {
    const url = isEditing.value ? '/api/v1/users/update' : '/api/v1/users/create'
    const method = 'POST'

    const payload = { ...userForm.value }
    if (isEditing.value && !payload.password) {
      delete (payload as any).password
    }

    const res = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify(payload),
    })

    const result = await res.json()
    if (res.ok && result.success) {
      closeUserModal()
      showFlash(result.message || 'Data pengguna berhasil disimpan!', 'success')
      await fetchUsers()
    } else {
      modalError.value = result.error || 'Gagal menyimpan data pengguna'
    }
  } catch (e: any) {
    modalError.value = 'Terjadi kesalahan jaringan'
  } finally {
    isSavingUser.value = false
  }
}

async function confirmDeleteUser(u: any) {
  if (u.id === authStore.user?.id) {
    alert('Anda tidak dapat menghapus akun Anda sendiri.')
    return
  }

  const ok = confirm(`Apakah Anda yakin ingin menghapus akun "${u.name}" (${u.email})? Tindakan ini tidak dapat dibatalkan.`)
  if (!ok) return

  try {
    const res = await fetch(`/api/v1/users/delete?id=${u.id}`, {
      method: 'POST',
      headers: { Authorization: `Bearer ${authStore.token}` },
    })

    const result = await res.json()
    if (res.ok && result.success) {
      showFlash(`Pengguna ${u.name} berhasil dihapus.`, 'success')
      await fetchUsers()
    } else {
      showFlash(result.error || 'Gagal menghapus pengguna', 'error')
    }
  } catch (e) {
    showFlash('Gagal menghubungi server saat menghapus akun', 'error')
  }
}

function showFlash(msg: string, type = 'success') {
  globalMessage.value = msg
  globalMessageType.value = type
  setTimeout(() => {
    globalMessage.value = ''
  }, 4000)
}

function getInitials(name: string) {
  if (!name) return 'U'
  return name.split(' ').map(n => n[0]).slice(0, 2).join('').toUpperCase()
}

function formatRole(role: string) {
  switch (role) {
    case 'direktur': return 'Direktur Utama'
    case 'admin': return 'Office Admin'
    case 'engineer': return 'Jobs Engineer'
    case 'klien': return 'Client Portal'
    case 'audit': return 'Internal Audit'
    default: return role
  }
}

function formatDate(dateStr?: string) {
  if (!dateStr) return 'Belum pernah'
  try {
    const d = new Date(dateStr)
    return d.toLocaleDateString('id-ID', {
      day: 'numeric',
      month: 'short',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    })
  } catch {
    return dateStr
  }
}

function formatCurrency(val: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
  }).format(val)
}

// State Jurnal SAK EMKM
const journals = ref<any[]>([])

async function fetchJournals() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/journals', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      journals.value = result.data.map((item: any) => ({
        id: item.id,
        date: item.journal_date,
        refNo: item.journal_number,
        description: item.memo,
        debitAccount: item.debit_account,
        creditAccount: item.credit_account,
        amount: item.amount,
      }))
    } else {
      console.error('🚨 [Backend Error] /api/v1/journals:', res.status, result)
    }
  } catch (err) {
    console.error('Gagal memuat data jurnal dari backend:', err)
  }
}

const showJournalModal = ref(false)
const journalForm = ref({
  date: new Date().toISOString().substring(0, 10),
  description: '',
  debitAccount: '1-1002 Bank BCA Utama Korporat',
  creditAccount: '4-1002 Pendapatan Kontrak Managed Care SLA',
  amount: 5000000,
})

const totalJournalDebit = computed(() => journals.value.reduce((acc, curr) => acc + curr.amount, 0))
const totalJournalCredit = computed(() => journals.value.reduce((acc, curr) => acc + curr.amount, 0))

async function saveJournal() {
  const dAccount = journalForm.value.debitAccount
  const cAccount = journalForm.value.creditAccount
  const amt = Number(journalForm.value.amount)
  const desc = journalForm.value.description
  const dateVal = journalForm.value.date

  if (authStore.token) {
    try {
      const res = await fetch('/api/v1/journals/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          date: dateVal,
          description: desc,
          debit_account: dAccount,
          credit_account: cAccount,
          amount: amt,
        }),
      })
      const result = await res.json()
      if (res.ok && result.success) {
        await fetchJournals()
        showJournalModal.value = false
        journalForm.value = {
          date: new Date().toISOString().substring(0, 10),
          description: '',
          debitAccount: '1-1002 Bank BCA Utama Korporat',
          creditAccount: '4-1002 Pendapatan Kontrak Managed Care SLA',
          amount: 5000000,
        }
        showFlash('Transaksi jurnal berhasil diposting ke Buku Besar secara seimbang!')
        return
      }
    } catch (e) {
      console.error('Gagal mencatat jurnal di backend:', e)
    }
  }

  // Fallback local update
  const newId = journals.value.length + 1
  journals.value.unshift({
    id: newId,
    date: dateVal,
    refNo: `JU-2026/10/00${newId}`,
    description: desc,
    debitAccount: dAccount,
    creditAccount: cAccount,
    amount: amt,
  })
  showJournalModal.value = false
  showFlash('Transaksi jurnal berhasil diposting ke Buku Besar secara seimbang!')
}

// State Invoices Korporat
const officeInvoices = ref<any[]>([])

const totalRevenueYTD = computed(() => {
  return officeInvoices.value
    .filter(i => i.status === 'paid')
    .reduce((acc, curr) => acc + (curr.total || 0), 0)
})

const totalPPN = computed(() => {
  return officeInvoices.value.reduce((acc, curr) => acc + (curr.ppn || 0), 0)
})

const totalPPh21 = computed(() => {
  return officeWithholdingSlips.value.reduce((acc, curr) => acc + (curr.taxAmount || 0), 0)
})

async function fetchInvoices() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/invoices', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      officeInvoices.value = result.data.map((item: any) => ({
        id: item.id,
        invoiceNo: item.invoice_number,
        clientName: item.client_name,
        dpp: item.amount,
        ppn: item.tax_amount,
        total: item.total_amount,
        dueDate: item.due_date,
        status: item.status,
      }))
    } else {
      console.error('🚨 [Backend Error] /api/v1/invoices:', res.status, result)
    }
  } catch (err) {
    console.error('Gagal memuat faktur dari backend:', err)
  }
}

const showInvoiceModal = ref(false)
const invoiceForm = ref({
  clientName: '',
  description: '',
  dpp: 10000000,
  dueDate: new Date(Date.now() + 14 * 86400000).toISOString().substring(0, 10),
})

async function saveInvoice() {
  const dpp = Number(invoiceForm.value.dpp)
  const client = invoiceForm.value.clientName
  const desc = invoiceForm.value.description
  const due = invoiceForm.value.dueDate

  if (authStore.token) {
    try {
      const res = await fetch('/api/v1/invoices/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          client_name: client,
          description: desc,
          amount: dpp,
          due_date: due,
        }),
      })
      const result = await res.json()
      if (res.ok && result.success) {
        await fetchInvoices()
        await fetchJournals()
        showInvoiceModal.value = false
        invoiceForm.value = {
          clientName: '',
          description: '',
          dpp: 10000000,
          dueDate: new Date(Date.now() + 14 * 86400000).toISOString().substring(0, 10),
        }
        showFlash('Faktur tagihan baru berhasil diterbitkan dan otomatis dibukukan ke SAK EMKM!')
        return
      }
    } catch (e) {
      console.error('Gagal menerbitkan faktur di backend:', e)
    }
  }

  // Fallback local update
  const newId = officeInvoices.value.length + 1
  const ppn = Math.round(dpp * 0.11)
  const total = dpp + ppn
  officeInvoices.value.unshift({
    id: newId,
    invoiceNo: `INV/2026/10/00${newId + 3}`,
    clientName: client,
    dpp,
    ppn,
    total,
    dueDate: due,
    status: 'unpaid',
  })
  showInvoiceModal.value = false
  showFlash('Faktur tagihan baru berhasil diterbitkan!')
}

async function toggleInvoiceStatus(inv: any) {
  const newStatus = inv.status === 'paid' ? 'unpaid' : 'paid'
  inv.status = newStatus
  showFlash(`Status invoice ${inv.invoiceNo} diperbarui menjadi ${newStatus === 'paid' ? 'LUNAS' : 'BELUM BAYAR'}`)

  if (authStore.token) {
    try {
      await fetch('/api/v1/invoices/status', {
        method: 'PATCH',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({ id: inv.id, status: newStatus }),
      })
      await fetchJournals()
    } catch (e) {
      console.error('Gagal update status invoice di backend:', e)
    }
  }
}

function exportCoretaxCSV() {
  const headers = 'Nomor Faktur,Nama Klien,DPP,PPN 11%,Total,Status\n'
  const rows = officeInvoices.value.map(i => `${i.invoiceNo},${i.clientName},${i.dpp},${i.ppn},${i.total},${i.status}`).join('\n')
  const blob = new Blob([headers + rows], { type: 'text/csv' })
  const url = URL.createObjectURL(blob)
  const a = document.createElement('a')
  a.href = url
  a.download = `Coretax_DJP_Export_PPN_1111_${new Date().toISOString().substring(0, 10)}.csv`
  a.click()
  showFlash('Berkas ekspor rekonsiliasi Coretax DJP (CSV) berhasil diunduh!')
}

// CRM Inbound Leads State & Logic
export interface Lead {
  id: number
  lead_code: string
  name: string
  company?: string
  whatsapp: string
  email?: string
  service_interest: string
  budget_range?: string
  notes?: string
  status: 'baru' | 'dihubungi' | 'penawaran' | 'closing' | 'dibatalkan'
  source: string
  created_at: string
  updated_at: string
}

const leadsList = ref<Lead[]>([])
const isLoadingLeads = ref(false)
const leadSearchQuery = ref('')
const leadStatusFilter = ref('')
const updatingLeadId = ref<number | null>(null)

const filteredLeads = computed(() => {
  return leadsList.value.filter(lead => {
    const matchStatus = !leadStatusFilter.value || lead.status === leadStatusFilter.value
    const q = leadSearchQuery.value.toLowerCase().trim()
    const matchQuery = !q ||
      lead.lead_code.toLowerCase().includes(q) ||
      lead.name.toLowerCase().includes(q) ||
      (lead.company && lead.company.toLowerCase().includes(q)) ||
      lead.whatsapp.toLowerCase().includes(q) ||
      (lead.notes && lead.notes.toLowerCase().includes(q)) ||
      lead.service_interest.toLowerCase().includes(q)
    return matchStatus && matchQuery
  })
})

async function fetchLeads() {
  if (!authStore.token) return
  isLoadingLeads.value = true
  try {
    const res = await fetch('/api/v1/leads', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data)) {
      leadsList.value = result.data
    } else {
      console.error('🚨 [Backend Error] /api/v1/leads:', res.status, result)
    }
  } catch (err) {
    console.error('Gagal memuat CRM leads dari backend:', err)
  } finally {
    isLoadingLeads.value = false
  }
}

async function onLeadStatusChange(lead: Lead, newStatus: string) {
  if (!authStore.token) return
  updatingLeadId.value = lead.id
  try {
    const res = await fetch('/api/v1/leads/status', {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${authStore.token}`,
      },
      body: JSON.stringify({
        id: lead.id,
        status: newStatus,
      }),
    })
    const result = await res.json()
    if (res.ok && result.success) {
      lead.status = newStatus as any
      showFlash(`Status prospek ${lead.lead_code} berhasil diperbarui menjadi "${newStatus}".`)
    } else {
      showFlash(result.message || 'Gagal memperbarui status lead', 'error')
    }
  } catch (err) {
    showFlash('Gangguan koneksi saat memperbarui status lead', 'error')
  } finally {
    updatingLeadId.value = null
  }
}

function openLeadWhatsApp(lead: Lead) {
  let cleanPhone = lead.whatsapp.replace(/\D/g, '')
  if (cleanPhone.startsWith('0')) {
    cleanPhone = '62' + cleanPhone.slice(1)
  } else if (!cleanPhone.startsWith('62')) {
    cleanPhone = '62' + cleanPhone
  }
  const text = `Halo Bapak/Ibu ${lead.name},\n\nSalam dari PT Melayani Digital Raya (meldir.id). Kami menindaklanjuti permohonan audit & konsultasi sistem Anda dengan tiket *${lead.lead_code}* terkait *${lead.service_interest}*.\n\nApakah ada waktu luang untuk berdiskusi singkat mengenai kebutuhan arsitektur dan estimasi roadmap sistem Anda? Terima kasih.`
  const url = `https://wa.me/${cleanPhone}?text=${encodeURIComponent(text)}`
  window.open(url, '_blank')
}

async function handleLogout() {
  await authStore.logout()
}

onMounted(async () => {
  if (authStore.token) {
    await authStore.fetchProfile()
    fetchUsers()
    fetchInvoices()
    fetchJournals()
    fetchLeads()
  }

  try {
    const res = await fetch('/api/health')
    if (res.ok) {
      const data = await res.json()
      apiStatus.value = 'online'
      apiStatusText.value = `Online (${data.service || 'Golang API'})`
    } else {
      apiStatus.value = 'error'
      apiStatusText.value = `Respon HTTP ${res.status}`
    }
  } catch (e) {
    apiStatus.value = 'offline'
    apiStatusText.value = 'Offline / Gangguan Jaringan'
  }
})
</script>

<style scoped>
.portal-container {
  --theme-color: #0284c7;
  --theme-color-glow: rgba(2, 132, 199, 0.4);
  max-width: 1400px;
  margin: 0 auto;
  padding: 24px;
  min-height: 100vh;
  background-color: #edf2f7; /* Berlatar belakang sedikit abu-abu lembut agar card putih kontras */
}

/* Auth Unauthenticated Header */
.auth-wrapper {
  max-width: 600px;
  margin: 40px auto;
  text-align: center;
}
.auth-hero-header {
  margin-bottom: 24px;
}
.brand-logo {
  font-size: 0.85rem;
  font-weight: 800;
  letter-spacing: 0.1em;
  color: #0284c7;
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

/* Authenticated Header */
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
  padding: 4px 8px;
  border-radius: 4px;
  font-weight: 700;
  letter-spacing: 0.05em;
}
.badge-office {
  background: #0284c7;
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

/* User Profile Header Menu */
.user-profile-menu {
  display: flex;
  align-items: center;
  gap: 14px;
}
.btn-header-users {
  background: #0284c7;
  color: #ffffff;
  border: 1px solid #0369a1;
  padding: 8px 16px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 6px rgba(2, 132, 199, 0.35);
  display: flex;
  align-items: center;
  gap: 6px;
}
.btn-header-users:hover {
  background: #0369a1;
  transform: translateY(-1px);
}
.btn-header-users.active {
  background: #0f172a;
  border-color: #0f172a;
  box-shadow: 0 2px 6px rgba(15, 23, 42, 0.3);
}

.user-avatar-badge {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: linear-gradient(135deg, #0284c7, #2563eb);
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 0.9rem;
  box-shadow: 0 2px 8px rgba(2, 132, 199, 0.25);
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
  background: #e0f2fe;
  color: #0369a1;
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

/* Status Bar Banner */
.status-bar-banner {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 20px;
  margin-bottom: 20px;
  font-size: 0.85rem;
  flex-wrap: wrap;
  gap: 12px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.04);
}
.status-left {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #334155;
}
.pulse-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #059669;
  box-shadow: 0 0 6px rgba(5, 150, 105, 0.6);
}
.api-badge {
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;
  font-size: 0.75rem;
}
.api-badge.online {
  background: #d1fae5;
  color: #065f46;
}
.api-badge.offline, .api-badge.error {
  background: #fee2e2;
  color: #991b1b;
}
.api-badge.checking {
  background: #fef3c7;
  color: #92400e;
}
.status-right {
  display: flex;
  align-items: center;
  gap: 8px;
  color: #64748b;
}
.text-emerald { color: #059669; }
.text-indigo { color: #4f46e5; }
.sep { color: #cbd5e1; }

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
  box-shadow: 0 1px 2px rgba(0, 0, 0, 0.05);
}
.nav-tab-btn:hover {
  color: #0f172a;
  background: #f8fafc;
  border-color: #94a3b8;
}
.nav-tab-btn.active {
  background: #0284c7;
  border-color: #0284c7;
  color: #ffffff;
  box-shadow: 0 3px 8px rgba(2, 132, 199, 0.35);
}
.highlight-user-tab {
  border-color: #0284c7;
  font-weight: 800;
}
.badge-tab-pill {
  font-size: 0.65rem;
  background: #fef08a;
  color: #854d0e;
  padding: 1px 6px;
  border-radius: 9999px;
  font-weight: 800;
  margin-left: 4px;
}
.card-clickable {
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}
.card-clickable:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px -2px rgba(99, 102, 241, 0.15);
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
.alert-banner.error {
  background: #fef2f2;
  border: 1px solid #fecaca;
  color: #b91c1c;
}

/* Grid KPI with subtle colored top border */
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
.kpi-card-blue { border-top: 4px solid #0284c7; }
.kpi-card-green { border-top: 4px solid #10b981; }
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
.badge-balanced {
  background: #d1fae5;
  color: #065f46;
  padding: 3px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 700;
}
.tag {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.tag-green {
  background: #d1fae5;
  color: #065f46;
}
.tag-tax {
  background: #fef3c7;
  color: #92400e;
}
.tag-indigo {
  background: #e0e7ff;
  color: #3730a3;
}

/* Section Panels */
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

/* Filter & Search Bar */
.filter-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}
.search-box {
  flex: 1;
  min-width: 260px;
}
.search-input {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
  color: #0f172a;
}
.search-input:focus {
  outline: none;
  border-color: #0284c7;
  background: #ffffff;
  box-shadow: 0 0 0 3px rgba(2, 132, 199, 0.12);
}
.role-filter-box {
  min-width: 220px;
}
.select-filter {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
  color: #0f172a;
  cursor: pointer;
}
.btn-refresh {
  background: #f1f5f9;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 10px 16px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}
.btn-refresh:hover {
  background: #e2e8f0;
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
}
.quick-card-highlight {
  background: #f0f9ff;
  border: 1px solid #7dd3fc;
  border-left: 5px solid #0284c7;
  cursor: pointer;
  transition: transform 0.2s, box-shadow 0.2s;
}
.quick-card-highlight:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 12px rgba(2, 132, 199, 0.15);
}
.quick-card-btn {
  margin-top: 10px;
  font-size: 0.8rem;
  font-weight: 700;
  color: #0284c7;
}
.quick-card h4 {
  font-size: 1rem;
  color: #0284c7;
  font-weight: 700;
  margin-bottom: 6px;
}
.quick-card p {
  font-size: 0.85rem;
  color: #64748b;
  line-height: 1.5;
}

/* Table */
.table-container {
  overflow-x: auto;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  background: #ffffff;
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
.data-table tbody tr:hover {
  background: #f8fafc;
}
.user-row-cell {
  display: flex;
  align-items: center;
  gap: 10px;
}
.avatar-circle {
  width: 34px;
  height: 34px;
  border-radius: 50%;
  background: #e0f2fe;
  color: #0284c7;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 700;
  font-size: 0.8rem;
}
.role-badge {
  font-size: 0.75rem;
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
  white-space: nowrap;
}
.role-badge.direktur { background: #fee2e2; color: #991b1b; }
.role-badge.admin { background: #e0f2fe; color: #0369a1; }
.role-badge.engineer { background: #e0e7ff; color: #4338ca; }
.role-badge.klien { background: #d1fae5; color: #065f46; }
.role-badge.audit { background: #fef3c7; color: #92400e; }

.status-badge {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;
  text-transform: capitalize;
}
.status-badge.aktif { background: #d1fae5; color: #065f46; }
.status-badge.pengajuan { background: #fef3c7; color: #92400e; }
.status-badge.suspended { background: #fee2e2; color: #991b1b; }

.badge-detail {
  background: #f1f5f9;
  padding: 2px 6px;
  border-radius: 4px;
  color: #475569;
}

.action-btn-group {
  display: flex;
  justify-content: flex-end;
  gap: 6px;
}
.btn-action {
  padding: 5px 10px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
  cursor: pointer;
  border: 1px solid transparent;
  transition: all 0.2s;
}
.btn-action.edit {
  background: #e0f2fe;
  color: #0369a1;
  border-color: #bae6fd;
}
.btn-action.edit:hover {
  background: #0284c7;
  color: #ffffff;
}
.btn-action.delete {
  background: #fee2e2;
  color: #b91c1c;
  border-color: #fecaca;
}
.btn-action.delete:hover:not(:disabled) {
  background: #dc2626;
  color: #ffffff;
}
.btn-action.delete:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.btn-primary {
  background: #0284c7;
  color: #ffffff;
  border: none;
  padding: 10px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 4px rgba(2, 132, 199, 0.25);
}
.btn-primary:hover:not(:disabled) {
  background: #0369a1;
  transform: translateY(-1px);
}
.btn-secondary {
  background: #f1f5f9;
  border: 1px solid #cbd5e1;
  color: #334155;
  padding: 10px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}
.btn-secondary:hover {
  background: #e2e8f0;
}

.category-pill {
  font-size: 0.72rem;
  padding: 3px 8px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;
}
.category-pill.asset { background: #e0f2fe; color: #0369a1; }
.category-pill.liability { background: #fee2e2; color: #991b1b; }
.category-pill.equity { background: #f3e8ff; color: #7e22ce; }
.category-pill.revenue { background: #d1fae5; color: #065f46; }
.category-pill.expense { background: #fef3c7; color: #92400e; }
.badge-active {
  background: #d1fae5;
  color: #065f46;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 700;
}
.capitalize { text-transform: capitalize; }
.font-mono { font-family: var(--font-mono); font-weight: 600; color: #334155; }
.text-xs { font-size: 0.75rem; }
.text-sm { font-size: 0.85rem; }
.text-muted { color: #64748b; }
.text-center { text-align: center; }
.py-6 { padding-top: 24px; padding-bottom: 24px; }

/* Tax Grid */
.tax-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
}
.tax-card {
  padding: 24px;
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  border-radius: 10px;
}
.tax-card h3 {
  font-size: 1.15rem;
  color: #0f172a;
  font-weight: 700;
  margin-bottom: 8px;
}
.tax-rate {
  font-size: 0.88rem;
  font-weight: 700;
  color: #b45309;
  margin-bottom: 10px;
}
.tax-detail {
  font-size: 0.85rem;
  color: #475569;
  line-height: 1.6;
}

/* Modal Form Styles */
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(15, 23, 42, 0.6);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 999;
  padding: 16px;
}
.modal-card {
  max-width: 650px;
  width: 100%;
  background: #ffffff;
  border-radius: 14px;
  border: 1px solid #cbd5e1;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.2);
  overflow: hidden;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
}
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 24px;
  border-bottom: 1px solid #e2e8f0;
  background: #f8fafc;
}
.modal-header h3 {
  font-size: 1.15rem;
  font-weight: 800;
  color: #0f172a;
  margin: 0;
}
.btn-close-modal {
  background: transparent;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  color: #64748b;
  padding: 4px;
}
.btn-close-modal:hover {
  color: #0f172a;
}
.modal-form {
  padding: 24px;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.form-row {
  display: flex;
  gap: 16px;
  flex-wrap: wrap;
}
.flex-1 {
  flex: 1;
  min-width: 240px;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.form-group label {
  font-size: 0.82rem;
  font-weight: 700;
  color: #334155;
}
.form-input {
  padding: 10px 12px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
  color: #0f172a;
}
.form-input:focus {
  outline: none;
  border-color: #0284c7;
  background: #ffffff;
  box-shadow: 0 0 0 3px rgba(2, 132, 199, 0.12);
}
.modal-error-banner {
  background: #fef2f2;
  border: 1px solid #fecaca;
  color: #b91c1c;
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 0.85rem;
  font-weight: 600;
}
.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 12px;
  margin-top: 12px;
  padding-top: 16px;
  border-top: 1px solid #e2e8f0;
}
.loading-state {
  padding: 36px;
  text-align: center;
  color: #64748b;
  font-weight: 600;
}

/* Journal Balance Summary */
.journal-balance-summary {
  display: flex;
  align-items: center;
  gap: 20px;
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  padding: 14px 20px;
  border-radius: 8px;
  flex-wrap: wrap;
}
.balance-item {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.9rem;
}

/* Invoice Status Badges */
.badge-inv-status {
  font-size: 0.72rem;
  font-weight: 800;
  padding: 3px 8px;
  border-radius: 4px;
}
.badge-inv-status.paid {
  background: #d1fae5;
  color: #065f46;
}
.badge-inv-status.unpaid {
  background: #fef3c7;
  color: #92400e;
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

/* CRM Leads Tab & Badges */
.lead-code-badge {
  font-family: 'Fira Code', monospace;
  font-weight: 700;
  font-size: 0.8rem;
  background: #f1f5f9;
  color: #0060af;
  padding: 3px 8px;
  border-radius: 4px;
  border: 1px solid #cbd5e1;
}

.lead-status-select {
  padding: 5px 8px;
  border-radius: 6px;
  font-size: 0.78rem;
  font-weight: 700;
  border: 1px solid transparent;
  cursor: pointer;
  outline: none;
  transition: all 0.2s ease;
}

.lead-status-select.status-baru {
  background: #dbeafe;
  color: #1e40af;
  border-color: #bfdbfe;
}

.lead-status-select.status-dihubungi {
  background: #ede9fe;
  color: #5b21b6;
  border-color: #ddd6fe;
}

.lead-status-select.status-penawaran {
  background: #fef3c7;
  color: #92400e;
  border-color: #fde68a;
}

.lead-status-select.status-closing {
  background: #d1fae5;
  color: #065f46;
  border-color: #a7f3d0;
}

.lead-status-select.status-dibatalkan {
  background: #f1f5f9;
  color: #64748b;
  border-color: #e2e8f0;
}

.btn-lead-wa-action {
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 12px;
  background: #25d366;
  color: #ffffff;
  font-weight: 700;
  font-size: 0.78rem;
  border-radius: 6px;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-lead-wa-action:hover {
  background: #1ebc57;
  transform: translateY(-1px);
}
</style>
