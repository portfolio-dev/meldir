<template>
  <div class="portal-container">
    <header class="portal-header glass-panel">
      <div class="brand">
        <span class="badge badge-jobs">ENGINEER WORKSPACE</span>
        <h1>Jobs Developer Hub</h1>
      </div>
      <div class="user-info">
        <span class="status-indicator online"></span>
        <span class="host-info">{{ currentHost }}</span>
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
          Terhubung ke Golang API di <code>/api/health</code>. Akses kode sumber melalui undangan repositori resmi GitHub PT. Melayani Digital Raya.
        </p>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
const currentHost = ref(window.location.host)
</script>

<style scoped>
.portal-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 24px;
}
.portal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  margin-bottom: 24px;
}
.brand h1 {
  font-size: 1.25rem;
  font-weight: 700;
  margin-top: 4px;
}
.badge {
  font-size: 0.7rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;
}
.badge-jobs {
  background: #6366f1;
  color: #fff;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: var(--font-mono);
  font-size: 0.85rem;
  color: var(--text-muted);
}
.status-indicator.online {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #6366f1;
  box-shadow: 0 0 8px #6366f1;
}
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
.tag-blue {
  background: rgba(99, 102, 241, 0.2);
  color: #818cf8;
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
.section-desc code {
  background: rgba(255, 255, 255, 0.1);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: var(--font-mono);
  color: #818cf8;
}
</style>
