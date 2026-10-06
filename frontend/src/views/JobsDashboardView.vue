<template>
  <div class="portal-container">
    <!-- Unauthenticated State: Show Login Form -->
    <div v-if="!authStore.isAuthenticated" class="auth-wrapper">
      <div class="auth-hero-header">
        <div class="brand-logo">⚙️ JOBS DEVELOPER HUB</div>
        <h1>Workspace Rekayasa Perangkat Lunak</h1>
        <p>Akses khusus Tim Engineer PT. Melayani Digital Raya (Internal Core & Tenaga Ahli Eksternal).</p>
      </div>
      <LoginForm portal-name="Engineer Workspace" portal-badge-class="jobs" />
    </div>

    <!-- Authenticated State: Engineer Workspace -->
    <div v-else class="dashboard-wrapper">
      <header class="portal-header glass-panel">
        <div class="brand">
          <span class="badge badge-jobs">ENGINEER HUB</span>
          <h1>Jobs Developer Workspace</h1>
          <span class="host-tag">{{ currentHost }}</span>
        </div>

        <div class="user-profile-menu">
          <div class="user-avatar-badge">
            {{ userInitials }}
          </div>
          <div class="user-meta">
            <span class="user-name">{{ authStore.user?.name }}</span>
            <div class="user-role-line">
              <span class="badge-role">{{ authStore.user?.role }}</span>
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
              <h3>SLA Taskboard</h3>
              <span class="tag">P1 / P2 / P3</span>
            </div>
            <div class="kpi-value">Active Tasks</div>
            <p class="kpi-desc">Tiket bug & sprint maintenance real-time</p>
          </div>

          <div class="card glass-panel">
            <div class="card-header">
              <h3>Timesheet Metered</h3>
              <span class="tag tag-blue">Log Jam</span>
            </div>
            <div class="kpi-value">Input Waktu</div>
            <p class="kpi-desc">Pencatatan jam kerja pemotong saldo Add-On</p>
          </div>

          <div class="card glass-panel">
            <div class="card-header">
              <h3>Pajak PPh 21</h3>
              <span class="tag tag-tax">e-Bupot 21</span>
            </div>
            <div class="kpi-value">Bukti Potong</div>
            <p class="kpi-desc">Unduh slip e-Bupot 21 resmi untuk SPT Pribadi</p>
          </div>
        </div>

        <section class="section-panel glass-panel">
          <h2>⚙️ Workspace Engineer (jobs.meldir.id)</h2>
          <p class="section-desc">
            Sesi Anda terverifikasi oleh Golang Backend API. Akses kode sumber melalui undangan repositori resmi GitHub PT. Melayani Digital Raya.
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
  if (!authStore.user?.name) return 'EN'
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
  color: #4f46e5;
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
  margin-bottom: 24px;
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
.badge-jobs {
  background: #4f46e5;
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
  background: linear-gradient(135deg, #4f46e5, #7c3aed);
  color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: 800;
  font-size: 0.9rem;
  box-shadow: 0 2px 8px rgba(79, 70, 229, 0.25);
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
  background: #e0e7ff;
  color: #4338ca;
  padding: 1px 6px;
  border-radius: 4px;
  font-weight: 700;
  text-transform: uppercase;
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

/* Grid & Cards */
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
.tag {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.tag-blue {
  background: #e0e7ff;
  color: #3730a3;
}
.tag-tax {
  background: #fef3c7;
  color: #92400e;
}
.section-panel {
  padding: 28px;
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
  line-height: 1.6;
}
</style>
