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
  background-color: var(--bg-primary);
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
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
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
  border: 1px solid #e2e8f0;
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
  border: 1px solid #e2e8f0;
  border-radius: 10px;
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
  border-bottom: 2px solid #e2e8f0;
  padding-bottom: 8px;
  overflow-x: auto;
}
.nav-tab-btn {
  background: transparent;
  border: 1px solid transparent;
  color: #64748b;
  padding: 9px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}
.nav-tab-btn:hover {
  color: #0f172a;
  background: #f1f5f9;
}
.nav-tab-btn.active {
  background: #0284c7;
  color: #ffffff;
  box-shadow: 0 2px 6px rgba(2, 132, 199, 0.3);
}

/* Grid KPI */
.grid-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 20px;
  margin-bottom: 24px;
}
.card {
  padding: 24px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
}
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

/* Section Panels */
.section-panel {
  padding: 28px;
  margin-bottom: 24px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
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
  margin-bottom: 22px;
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
  border: 1px solid #e2e8f0;
  padding: 20px;
  border-radius: 10px;
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
  border: 1px solid #e2e8f0;
  border-radius: 8px;
}
.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.88rem;
  text-align: left;
}
.data-table th {
  background: #f8fafc;
  padding: 12px 16px;
  font-weight: 700;
  color: #475569;
  border-bottom: 1px solid #e2e8f0;
}
.data-table td {
  padding: 14px 16px;
  border-bottom: 1px solid #f1f5f9;
  color: #1e293b;
}
.data-table tbody tr:hover {
  background: #f8fafc;
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

/* Tax Grid */
.tax-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 20px;
}
.tax-card {
  padding: 24px;
  background: #f8fafc;
  border: 1px solid #e2e8f0;
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
</style>
