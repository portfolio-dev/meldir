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
          <div class="user-avatar-badge">
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
        <div class="grid-overview">
          <div class="card glass-panel">
            <div class="card-header">
              <h3>Kesehatan Server Klien</h3>
              <span class="tag tag-green">99.98% Uptime</span>
            </div>
            <div class="kpi-value">Live Online</div>
            <p class="kpi-desc">External probe otomatis berjalan setiap 5 menit</p>
          </div>

          <div class="card glass-panel">
            <div class="card-header">
              <h3>Saldo Jam Add-On</h3>
              <span class="tag tag-blue">Metered</span>
            </div>
            <div class="kpi-value">7.5 / 10 Jam</div>
            <p class="kpi-desc">Berlaku untuk penambahan fitur & optimasi</p>
          </div>

          <div class="card glass-panel">
            <div class="card-header">
              <h3>Tagihan & e-Faktur</h3>
              <span class="tag tag-tax">PPN 11%</span>
            </div>
            <div class="kpi-value">e-Faktur Sah</div>
            <p class="kpi-desc">Unduh berkas PDF e-Faktur resmi DJP mandiri</p>
          </div>
        </div>

        <section class="section-panel glass-panel">
          <h2>💼 Portal Layanan Klien (portal.meldir.id)</h2>
          <p class="section-desc">
            Layanan terpadu: Pengesahan Kontrak SPK Canvas & E-Materai, Tiket SLA 24 Jam, BAST Digital, dan Unduh Laporan Kinerja Bulanan Eksekutif.
          </p>
        </section>
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

const userInitials = computed(() => {
  if (!authStore.user?.name) return 'KL'
  return authStore.user.name
    .split(' ')
    .map(w => w[0])
    .slice(0, 2)
    .join('')
    .toUpperCase()
})

async function handleLogout() {
  await authStore.logout()
}

onMounted(async () => {
  if (authStore.token) {
    await authStore.fetchProfile()
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
  color: #10b981;
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
  margin-bottom: 24px;
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
.badge-portal {
  background: #10b981;
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
  background: linear-gradient(135deg, #10b981, #059669);
  color: #fff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 0.85rem;
  box-shadow: 0 0 12px rgba(16, 185, 129, 0.3);
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
  background: rgba(16, 185, 129, 0.15);
  color: #34d399;
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

/* Grid & Cards */
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
  font-size: 1.5rem;
  font-weight: 800;
  color: #f8fafc;
  margin-bottom: 6px;
}
.kpi-desc {
  font-size: 0.8rem;
  color: var(--text-muted);
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
.tag-blue {
  background: rgba(56, 189, 248, 0.2);
  color: #38bdf8;
}
.tag-tax {
  background: rgba(245, 158, 11, 0.15);
  color: #f59e0b;
}
.section-panel {
  padding: 24px;
}
.section-panel h2 {
  font-size: 1.25rem;
  margin-bottom: 8px;
}
.section-desc {
  color: var(--text-muted);
  font-size: 0.9rem;
  line-height: 1.5;
}
</style>
