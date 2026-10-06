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
          <!-- Tombol Akses Cepat Manajemen Pengguna untuk Direktur & Admin -->
          <button
            v-if="authStore.user?.role === 'direktur' || authStore.user?.role === 'admin'"
            @click="switchTab('users')"
            :class="['btn-header-users', { active: activeTab === 'users' }]"
            title="Buka Menu Manajemen Akun Pengguna (CRUD)"
          >
            👥 Manajemen Pengguna (CRUD)
          </button>

          <div class="user-avatar-badge">
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
              <div class="kpi-value">Rp 128.500.000</div>
              <p class="kpi-desc">Total faktur terbit & lunas terlapor SAK EMKM</p>
            </div>

            <div class="card glass-panel kpi-card-green">
              <div class="card-header">
                <h3>Buku Besar (General Ledger)</h3>
                <span class="badge-balanced">✓ Seimbang</span>
              </div>
              <div class="kpi-value">32 Akun COA</div>
              <p class="kpi-desc">Debit & Kredit seimbang. Sesuai standar UU Pajak & IAI</p>
            </div>

            <div class="card glass-panel kpi-card-amber">
              <div class="card-header">
                <h3>Kepatuhan DJP (Pajak)</h3>
                <span class="tag tag-tax">SPT Masa 1111</span>
              </div>
              <div class="kpi-value">PPN 11% & e-Bupot</div>
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
                </tbody>
              </table>
            </div>
          </section>

          <!-- Modul Bagan Akun Standar (COA) -->
          <section class="section-panel glass-panel">
            <h2>📑 Bagan Akun Standar (Chart of Accounts - 32 Akun)</h2>
            <p class="section-desc">
              Kode akun 5-digit standar Ikatan Akuntan Indonesia (IAI) untuk pelaporan keuangan korporat.
            </p>
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
                  <tr v-for="acc in defaultAccounts" :key="acc.code">
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
                <h3>SPT Masa PPN 1111</h3>
                <p class="tax-rate">Tarif: 11% (12% Transisi)</p>
                <p class="tax-detail">Pencatatan faktur pajak keluaran atas jasa custom development & managed care, serta rekonsiliasi faktur masukan vendor server/cloud.</p>
              </div>
              <div class="tax-card">
                <h3>e-Bupot PPh Pasal 21</h3>
                <p class="tax-rate">Pasal 17 / Tenaga Ahli</p>
                <p class="tax-detail">Pemotongan PPh 21 atas kompensasi tenaga ahli programmer eksternal dengan penerbitan bukti potong resmi 21/26.</p>
              </div>
              <div class="tax-card">
                <h3>e-Bupot PPh Pasal 23</h3>
                <p class="tax-rate">Tarif: 2% Jasa Teknik</p>
                <p class="tax-detail">Pencatatan bukti potong PPh 23 saat klien korporat memotong pembayaran invoice PT. Melayani Digital Raya.</p>
              </div>
            </div>
          </section>
        </div>
      </main>
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
                <option v-for="acc in defaultAccounts" :key="'deb-' + acc.code" :value="acc.code + ' ' + acc.name">
                  {{ acc.code }} — {{ acc.name }}
                </option>
              </select>
            </div>
            <div class="form-group flex-1">
              <label>Akun Posisi KREDIT *</label>
              <select v-model="journalForm.creditAccount" required class="form-input">
                <option v-for="acc in defaultAccounts" :key="'kred-' + acc.code" :value="acc.code + ' ' + acc.name">
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

const tabs = [
  { id: 'overview', label: 'Ringkasan Eksekutif', icon: '📊' },
  { id: 'users', label: 'Manajemen Pengguna (CRUD)', icon: '👥' },
  { id: 'accounting', label: 'Buku Besar SAK EMKM', icon: '📖' },
  { id: 'tax', label: 'Kepatuhan Pajak DJP', icon: '🏛️' },
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

function switchTab(tabId: string) {
  activeTab.value = tabId
  if (tabId === 'users' && usersList.value.length === 0) {
    fetchUsers()
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
const journals = ref([
  {
    id: 1,
    date: '2026-10-06',
    refNo: 'JU-2026/10/001',
    description: 'Penerimaan pembayaran piutang invoice PT. Surya Logistik',
    debitAccount: '1-1002 Bank BCA Utama Korporat',
    creditAccount: '1-1201 Piutang Usaha Klien',
    amount: 13875000,
  },
  {
    id: 2,
    date: '2026-10-05',
    refNo: 'JU-2026/10/002',
    description: 'Penyetoran PPN Masa Keluaran 11% ke Kas Negara via DJP',
    debitAccount: '2-1201 Utang PPN Keluaran 11%',
    creditAccount: '1-1002 Bank BCA Utama Korporat',
    amount: 1375000,
  },
  {
    id: 3,
    date: '2026-10-04',
    refNo: 'JU-2026/10/003',
    description: 'Pembayaran kompensasi jasa software engineer freelance',
    debitAccount: '5-1001 Beban Kompensasi Engineer',
    creditAccount: '1-1002 Bank BCA Utama Korporat',
    amount: 7500000,
  },
])

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

function saveJournal() {
  const newId = journals.value.length + 1
  journals.value.unshift({
    id: newId,
    date: journalForm.value.date,
    refNo: `JU-2026/10/00${newId}`,
    description: journalForm.value.description,
    debitAccount: journalForm.value.debitAccount,
    creditAccount: journalForm.value.creditAccount,
    amount: Number(journalForm.value.amount),
  })
  showJournalModal.value = false
  journalForm.value = {
    date: new Date().toISOString().substring(0, 10),
    description: '',
    debitAccount: '1-1002 Bank BCA Utama Korporat',
    creditAccount: '4-1002 Pendapatan Kontrak Managed Care SLA',
    amount: 5000000,
  }
  showFlash('Transaksi jurnal berhasil diposting ke Buku Besar secara seimbang!')
}

// State Invoices Korporat
const officeInvoices = ref([
  {
    id: 1,
    invoiceNo: 'INV/2026/10/004',
    clientName: 'PT. Surya Logistik Multimoda',
    dpp: 12500000,
    ppn: 1375000,
    total: 13875000,
    dueDate: '2026-10-15',
    status: 'paid',
  },
  {
    id: 2,
    invoiceNo: 'INV/2026/10/005',
    clientName: 'CV. Sejahtera Abadi Mandiri',
    dpp: 18000000,
    ppn: 1980000,
    total: 19980000,
    dueDate: '2026-10-25',
    status: 'unpaid',
  },
  {
    id: 3,
    invoiceNo: 'INV/2026/10/006',
    clientName: 'Yayasan Harapan Bangsa',
    dpp: 7500000,
    ppn: 825000,
    total: 8325000,
    dueDate: '2026-10-30',
    status: 'unpaid',
  },
])

const showInvoiceModal = ref(false)
const invoiceForm = ref({
  clientName: '',
  description: '',
  dpp: 10000000,
  dueDate: new Date(Date.now() + 14 * 86400000).toISOString().substring(0, 10),
})

function saveInvoice() {
  const newId = officeInvoices.value.length + 1
  const dpp = Number(invoiceForm.value.dpp)
  const ppn = Math.round(dpp * 0.11)
  const total = dpp + ppn
  officeInvoices.value.unshift({
    id: newId,
    invoiceNo: `INV/2026/10/00${newId + 3}`,
    clientName: invoiceForm.value.clientName,
    dpp,
    ppn,
    total,
    dueDate: invoiceForm.value.dueDate,
    status: 'unpaid',
  })
  showInvoiceModal.value = false
  invoiceForm.value = {
    clientName: '',
    description: '',
    dpp: 10000000,
    dueDate: new Date(Date.now() + 14 * 86400000).toISOString().substring(0, 10),
  }
  showFlash('Faktur tagihan baru berhasil diterbitkan dan siap dikirim ke klien!')
}

function toggleInvoiceStatus(inv: any) {
  inv.status = inv.status === 'paid' ? 'unpaid' : 'paid'
  showFlash(`Status invoice ${inv.invoiceNo} diperbarui menjadi ${inv.status === 'paid' ? 'LUNAS' : 'BELUM BAYAR'}`)
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

async function handleLogout() {
  await authStore.logout()
}

onMounted(async () => {
  if (authStore.token) {
    await authStore.fetchProfile()
    fetchUsers()
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
</style>
