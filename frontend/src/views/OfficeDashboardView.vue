<template>
  <div class="portal-container">
    <header class="portal-header glass-panel">
      <div class="brand">
        <span class="badge badge-office">OFFICE ADMIN</span>
        <h1>PT. Melayani Digital Raya</h1>
      </div>
      <div class="user-info">
        <span class="status-indicator online"></span>
        <span class="host-info">{{ currentHost }}</span>
      </div>
    </header>

    <main class="portal-main">
      <div class="grid-overview">
        <!-- Card 1: Executive KPI -->
        <div class="card glass-panel">
          <div class="card-header">
            <h3>Ringkasan Operasional</h3>
            <span class="tag">Live System</span>
          </div>
          <div class="kpi-value">Rp 128.500.000</div>
          <p class="kpi-desc">Total Omset Invoice Terverifikasi (YTD 2026)</p>
        </div>

        <!-- Card 2: Accounting & Double-Entry -->
        <div class="card glass-panel">
          <div class="card-header">
            <h3>Buku Besar & Akuntansi</h3>
            <span class="badge-balanced">✓ Balanced</span>
          </div>
          <div class="kpi-value">32 Akun COA</div>
          <p class="kpi-desc">Standar SAK EMKM 5-Digit (Aset, Kewajiban, Ekuitas)</p>
        </div>

        <!-- Card 3: Tax Compliance Hub -->
        <div class="card glass-panel">
          <div class="card-header">
            <h3>Kepatuhan Perpajakan</h3>
            <span class="tag tag-tax">PPN & PPh</span>
          </div>
          <div class="kpi-value">SPT Masa 1111</div>
          <p class="kpi-desc">PPN 11% Keluaran & Masukan | e-Bupot PPh 21/23</p>
        </div>
      </div>

      <!-- Quick Action / Status Board -->
      <section class="section-panel glass-panel">
        <h2>🏛️ Pusat Kendali Korporat (office.meldir.id)</h2>
        <p class="section-desc">
          Backend Golang API terhubung di <code>/api/health</code>. Seluruh modul Inbound Leads CRM, Kontrak SPK, BAST Digital, Jurnal Umum, dan Perpajakan aktif.
        </p>
        <div class="api-status-box">
          <strong>Status Koneksi API:</strong> 
          <span :class="['api-badge', apiStatus]">{{ apiStatusText }}</span>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const currentHost = ref(window.location.host)
const apiStatus = ref('checking')
const apiStatusText = ref('Memeriksa...')

onMounted(async () => {
  try {
    const res = await fetch('/api/health')
    if (res.ok) {
      const data = await res.json()
      apiStatus.value = 'online'
      apiStatusText.value = `Online (${data.service})`
    } else {
      apiStatus.value = 'error'
      apiStatusText.value = `Respon HTTP ${res.status}`
    }
  } catch (e) {
    apiStatus.value = 'offline'
    apiStatusText.value = 'Tidak Terhubung (Gunakan /api/health)'
  }
})
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
  letter-spacing: 0.05em;
}
.badge-office {
  background: #0284c7;
  color: #fff;
}
.badge-balanced {
  background: rgba(16, 185, 129, 0.2);
  color: #10b981;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 0.75rem;
}
.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  font-family: var(--font-mono);
  font-size: 0.85rem;
  color: var(--text-muted);
}
.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
}
.status-indicator.online {
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
  font-size: 1.75rem;
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
  margin-bottom: 16px;
  line-height: 1.5;
}
.section-desc code {
  background: rgba(255, 255, 255, 0.1);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: var(--font-mono);
  color: #38bdf8;
}
.api-status-box {
  display: flex;
  align-items: center;
  gap: 12px;
  font-size: 0.9rem;
}
.api-badge {
  padding: 4px 10px;
  border-radius: 6px;
  font-weight: 600;
  font-size: 0.8rem;
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
</style>
