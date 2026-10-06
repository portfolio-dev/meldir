<template>
  <div class="portal-container">
    <header class="portal-header glass-panel">
      <div class="brand">
        <span class="badge badge-portal">CLIENT CONTROL CENTER</span>
        <h1>Portal Klien meldir.id</h1>
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
.badge-portal {
  background: #10b981;
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
  background: #10b981;
  box-shadow: 0 0 8px #10b981;
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
