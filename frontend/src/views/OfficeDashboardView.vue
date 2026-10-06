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
            :class="['nav-tab-btn', { active: activeTab === tab.id }]"
            @click="activeTab = tab.id"
          >
            {{ tab.icon }} {{ tab.label }}
          </button>
        </div>

        <!-- Tab 1: Ringkasan Eksekutif -->
        <div v-if="activeTab === 'overview'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel">
              <div class="card-header">
                <h3>Omset Invoice YTD</h3>
                <span class="tag tag-green">Live 2026</span>
              </div>
              <div class="kpi-value">Rp 128.500.000</div>
              <p class="kpi-desc">Total faktur terbit & lunas terlapor SAK EMKM</p>
            </div>

            <div class="card glass-panel">
              <div class="card-header">
                <h3>Buku Besar (General Ledger)</h3>
                <span class="badge-balanced">✓ Seimbang</span>
              </div>
              <div class="kpi-value">32 Akun COA</div>
              <p class="kpi-desc">Debit & Kredit seimbang. Sesuai standar UU Pajak & IAI</p>
            </div>

            <div class="card glass-panel">
              <div class="card-header">
                <h3>Kepatuhan DJP (Pajak)</h3>
                <span class="tag tag-tax">SPT Masa 1111</span>
              </div>
              <div class="kpi-value">PPN 11% & e-Bupot</div>
              <p class="kpi-desc">Faktur Keluaran siap lapor ke sistem Coretax DJP</p>
            </div>
          </div>

          <section class="section-panel glass-panel">
            <h2>📊 Ringkasan Ekosistem Digital</h2>
            <p class="section-desc">
              Sistem terintegrasi monorepo dengan backend Golang di port 8080 dan NGINX/OpenLiteSpeed reverse-proxy. Seluruh transaksi bisnis dicatat menggunakan mekanisme double-entry bookkeeping otomatis.
            </p>
            <div class="quick-links-grid">
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

        <!-- Tab 2: Buku Besar SAK EMKM -->
        <div v-else-if="activeTab === 'accounting'" class="tab-content">
          <section class="section-panel glass-panel">
            <h2>📖 Bagan Akun Standar (Chart of Accounts - SAK EMKM)</h2>
            <p class="section-desc">
              Daftar akun buku besar 5-digit sesuai regulasi Standar Akuntansi Keuangan Entitas Mikro, Kecil, dan Menengah (SAK EMKM).
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

        <!-- Tab 3: Pajak DJP -->
        <div v-else-if="activeTab === 'tax'" class="tab-content">
          <section class="section-panel glass-panel">
            <h2>🏛️ Hub Kepatuhan Perpajakan PT. Melayani Digital Raya</h2>
            <p class="section-desc">
              Pelaporan pajak korporat otomatis sesuai tarif UU Harmonisasi Peraturan Perpajakan (HPP).
            </p>
            <div class="tax-grid">
              <div class="tax-card glass-panel">
                <h3>SPT Masa PPN 1111</h3>
                <p class="tax-rate">Tarif: 11% (12% Transisi)</p>
                <p class="tax-detail">Pencatatan faktur pajak keluaran atas jasa custom development & managed care, serta rekonsiliasi faktur masukan vendor server/cloud.</p>
              </div>
              <div class="tax-card glass-panel">
                <h3>e-Bupot PPh Pasal 21</h3>
                <p class="tax-rate">Pasal 17 / Bukan Pegawai</p>
                <p class="tax-detail">Pemotongan PPh 21 atas jasa tenaga ahli programmer/engineer eksternal dengan penerbitan bukti potong resmi 21/26.</p>
              </div>
              <div class="tax-card glass-panel">
                <h3>e-Bupot PPh Pasal 23</h3>
                <p class="tax-rate">Tarif: 2% atas Jasa Teknik</p>
                <p class="tax-detail">Pencatatan bukti potong PPh 23 saat klien korporat memotong pembayaran invoice PT. Melayani Digital Raya.</p>
              </div>
            </div>
          </section>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import LoginForm from '../components/LoginForm.vue'
import { useAuthStore } from '../stores/authStore'

const authStore = useAuthStore()
const currentHost = ref(window.location.host)
const apiStatus = ref('checking')
const apiStatusText = ref('Memeriksa...')
const activeTab = ref('overview')

const tabs = [
  { id: 'overview', label: 'Ringkasan Eksekutif', icon: '📊' },
  { id: 'accounting', label: 'Buku Besar SAK EMKM', icon: '📖' },
  { id: 'tax', label: 'Kepatuhan Pajak DJP', icon: '🏛️' },
]

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

async function handleLogout() {
  await authStore.logout()
}

onMounted(async () => {
  // Verifikasi token profile di background
  if (authStore.token) {
    await authStore.fetchProfile()
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
  color: #38bdf8;
  margin-bottom: 12px;
}
.auth-hero-header h1 {
  font-size: 1.85rem;
  font-weight: 800;
  margin-bottom: 8px;
  color: #f8fafc;
}
.auth-hero-header p {
  font-size: 0.9rem;
  color: var(--text-muted);
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
}
.brand {
  display: flex;
  align-items: center;
  gap: 12px;
}
.brand h1 {
  font-size: 1.25rem;
  font-weight: 700;
  margin: 0;
}
.badge {
  font-size: 0.7rem;
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
  letter-spacing: 0.05em;
}
.badge-office {
  background: #0284c7;
  color: #fff;
}
.host-tag {
  font-size: 0.75rem;
  font-family: var(--font-mono);
  background: rgba(255, 255, 255, 0.06);
  padding: 2px 8px;
  border-radius: 4px;
  color: var(--text-muted);
}

/* User Profile Header Menu */
.user-profile-menu {
  display: flex;
  align-items: center;
  gap: 14px;
}
.user-avatar-badge {
  width: 38px;
  height: 38px;
  border-radius: 50%;
  background: linear-gradient(135deg, #0284c7, #38bdf8);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 0.85rem;
  box-shadow: 0 0 12px rgba(56, 189, 248, 0.3);
}
.user-meta {
  display: flex;
  flex-direction: column;
  text-align: left;
}
.user-name {
  font-size: 0.9rem;
  font-weight: 700;
  color: #f8fafc;
}
.user-role-line {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.75rem;
}
.badge-role {
  background: rgba(14, 165, 233, 0.15);
  color: #38bdf8;
  padding: 1px 6px;
  border-radius: 4px;
  font-weight: 600;
}
.user-email {
  color: var(--text-muted);
}
.btn-logout {
  background: rgba(244, 63, 94, 0.12);
  border: 1px solid rgba(244, 63, 94, 0.3);
  color: #fb7185;
  padding: 6px 12px;
  border-radius: 6px;
  font-size: 0.8rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
}
.btn-logout:hover {
  background: rgba(244, 63, 94, 0.25);
  color: #fff;
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
}
.status-left {
  display: flex;
  align-items: center;
  gap: 8px;
}
.pulse-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #10b981;
  box-shadow: 0 0 8px #10b981;
}
.api-badge {
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
  font-size: 0.75rem;
}
.api-badge.online {
  background: rgba(16, 185, 129, 0.2);
  color: #34d399;
}
.api-badge.offline, .api-badge.error {
  background: rgba(244, 63, 94, 0.2);
  color: #fb7185;
}
.api-badge.checking {
  background: rgba(245, 158, 11, 0.2);
  color: #fbbf24;
}
.status-right {
  display: flex;
  align-items: center;
  gap: 8px;
  color: var(--text-muted);
}
.text-emerald { color: #34d399; }
.text-indigo { color: #818cf8; }
.sep { color: rgba(255, 255, 255, 0.2); }

/* Navigation Tabs */
.module-nav-tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
  border-bottom: 1px solid var(--border-color);
  padding-bottom: 8px;
  overflow-x: auto;
}
.nav-tab-btn {
  background: transparent;
  border: 1px solid transparent;
  color: var(--text-muted);
  padding: 8px 16px;
  border-radius: 6px;
  font-size: 0.85rem;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}
.nav-tab-btn:hover {
  color: #f8fafc;
  background: rgba(255, 255, 255, 0.05);
}
.nav-tab-btn.active {
  background: rgba(2, 132, 199, 0.2);
  border-color: #0284c7;
  color: #38bdf8;
}

/* Grid KPI */
.grid-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}
.card {
  padding: 20px;
}
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 12px;
}
.card-header h3 {
  font-size: 0.95rem;
  color: var(--text-muted);
}
.kpi-value {
  font-size: 1.75rem;
  font-weight: 800;
  color: #f8fafc;
  margin-bottom: 6px;
}
.kpi-desc {
  font-size: 0.8rem;
  color: var(--text-muted);
}
.badge-balanced {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
  font-weight: 600;
}
.tag {
  font-size: 0.7rem;
  background: rgba(255, 255, 255, 0.08);
  padding: 2px 6px;
  border-radius: 4px;
}
.tag-green {
  background: rgba(16, 185, 129, 0.2);
  color: #34d399;
}
.tag-tax {
  background: rgba(245, 158, 11, 0.15);
  color: #f59e0b;
}

/* Section Panels */
.section-panel {
  padding: 24px;
  margin-bottom: 24px;
}
.section-panel h2 {
  font-size: 1.25rem;
  margin-bottom: 8px;
}
.section-desc {
  color: var(--text-muted);
  font-size: 0.9rem;
  margin-bottom: 20px;
  line-height: 1.5;
}

/* Quick Links Grid */
.quick-links-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
  gap: 16px;
}
.quick-card {
  background: rgba(15, 23, 42, 0.5);
  border: 1px solid var(--border-color);
  padding: 18px;
  border-radius: 8px;
}
.quick-card h4 {
  font-size: 0.95rem;
  color: #38bdf8;
  margin-bottom: 6px;
}
.quick-card p {
  font-size: 0.8rem;
  color: var(--text-muted);
  line-height: 1.4;
}

/* Table */
.table-container {
  overflow-x: auto;
}
.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.85rem;
  text-align: left;
}
.data-table th {
  background: rgba(15, 23, 42, 0.8);
  padding: 12px 14px;
  font-weight: 600;
  color: var(--text-muted);
  border-bottom: 1px solid var(--border-color);
}
.data-table td {
  padding: 12px 14px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}
.category-pill {
  font-size: 0.7rem;
  padding: 2px 8px;
  border-radius: 4px;
  text-transform: uppercase;
  font-weight: 700;
}
.category-pill.asset { background: rgba(56, 189, 248, 0.15); color: #38bdf8; }
.category-pill.liability { background: rgba(244, 63, 94, 0.15); color: #fb7185; }
.category-pill.equity { background: rgba(168, 85, 247, 0.15); color: #c084fc; }
.category-pill.revenue { background: rgba(16, 185, 129, 0.15); color: #34d399; }
.category-pill.expense { background: rgba(245, 158, 11, 0.15); color: #fbbf24; }
.badge-active {
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
  padding: 2px 6px;
  border-radius: 4px;
  font-size: 0.75rem;
}
.capitalize { text-transform: capitalize; }

/* Tax Grid */
.tax-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
}
.tax-card {
  padding: 20px;
}
.tax-card h3 {
  font-size: 1.1rem;
  color: #f8fafc;
  margin-bottom: 8px;
}
.tax-rate {
  font-size: 0.85rem;
  font-weight: 700;
  color: #f59e0b;
  margin-bottom: 10px;
}
.tax-detail {
  font-size: 0.8rem;
  color: var(--text-muted);
  line-height: 1.5;
}
</style>
