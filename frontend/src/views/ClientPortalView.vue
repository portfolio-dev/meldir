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

        <div class="grid-overview">
          <div class="card glass-panel kpi-card-emerald">
            <div class="card-header">
              <h3>Kesehatan Server Klien</h3>
              <span class="tag tag-green">99.98% Uptime</span>
            </div>
            <div class="kpi-value">Live Online</div>
            <p class="kpi-desc">External probe otomatis berjalan setiap 5 menit</p>
          </div>

          <div class="card glass-panel kpi-card-blue">
            <div class="card-header">
              <h3>Saldo Jam Add-On</h3>
              <span class="tag tag-blue">Metered</span>
            </div>
            <div class="kpi-value">7.5 / 10 Jam</div>
            <p class="kpi-desc">Berlaku untuk penambahan fitur & optimasi</p>
          </div>

          <div class="card glass-panel kpi-card-amber">
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
  background-color: #edf2f7; /* Background abu-abu lembut agar card putih kontras */
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
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
  letter-spacing: 0.05em;
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
  white-space: nowrap;
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
  box-shadow: 0 2px 4px rgba(2, 132, 199, 0.25);
}
.btn-director-action:hover {
  background: #0369a1;
  transform: translateY(-1px);
}

/* Grid & Cards with Contrasting Surfaces */
.grid-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
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
.tag-blue {
  background: #e0f2fe;
  color: #0369a1;
}
.tag-tax {
  background: #fef3c7;
  color: #92400e;
}
.section-panel {
  padding: 28px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.04);
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
  line-height: 1.6;
}
</style>
