<template>
  <div class="hub-container">
    <div class="hub-card glass-panel">
      <div class="hub-header">
        <span class="badge">PT. MELAYANI DIGITAL RAYA</span>
        <h1>Multi-Portal Ecosystem</h1>
        <p>Pilih portal domain yang ingin Anda buka:</p>
      </div>

      <div class="portal-grid">
        <a href="https://office.meldir.id" class="portal-item glass-panel">
          <div class="portal-tag tag-office">OFFICE</div>
          <h3>office.meldir.id</h3>
          <p>Dashboard Corporate Operations, Akuntansi SAK EMKM & Pajak</p>
        </a>

        <a href="https://jobs.meldir.id" class="portal-item glass-panel">
          <div class="portal-tag tag-jobs">JOBS</div>
          <h3>jobs.meldir.id</h3>
          <p>Engineer Workspace, SLA Tasks, Timesheet & e-Bupot PPh 21</p>
        </a>

        <a href="https://portal.meldir.id" class="portal-item glass-panel">
          <div class="portal-tag tag-portal">PORTAL</div>
          <h3>portal.meldir.id</h3>
          <p>Client Control Center, Invoicing, Add-On Metered & BAST</p>
        </a>
      </div>

      <div class="hub-footer">
        <a href="https://meldir.id" class="back-link">← Kembali ke Landing Page meldir.id</a>
      </div>
    </div>
  </div>
</template>

<style scoped>
.hub-container {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 24px;
}
.hub-card {
  max-width: 900px;
  width: 100%;
  padding: 40px;
}
.hub-header {
  text-align: center;
  margin-bottom: 36px;
}
.badge {
  font-size: 0.75rem;
  font-weight: 700;
  letter-spacing: 0.1em;
  background: rgba(255, 255, 255, 0.1);
  padding: 4px 12px;
  border-radius: 9999px;
  display: inline-block;
  margin-bottom: 12px;
}
.hub-header h1 {
  font-size: 2rem;
  font-weight: 800;
  margin-bottom: 8px;
}
.hub-header p {
  color: var(--text-muted);
}
.portal-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}
.portal-item {
  display: block;
  padding: 24px;
  text-decoration: none;
  color: inherit;
  transition: transform 0.2s, border-color 0.2s, background 0.2s;
}
.portal-item:hover {
  transform: translateY(-4px);
  border-color: #38bdf8;
  background: var(--bg-card-hover);
}
.portal-tag {
  font-size: 0.7rem;
  font-weight: 800;
  padding: 2px 8px;
  border-radius: 4px;
  display: inline-block;
  margin-bottom: 12px;
}
.tag-office { background: #0284c7; color: #fff; }
.tag-jobs { background: #6366f1; color: #fff; }
.tag-portal { background: #10b981; color: #fff; }
.portal-item h3 {
  font-size: 1.1rem;
  margin-bottom: 8px;
  color: #f8fafc;
}
.portal-item p {
  font-size: 0.85rem;
  color: var(--text-muted);
  line-height: 1.4;
}
.hub-footer {
  text-align: center;
}
.back-link {
  color: #38bdf8;
  text-decoration: none;
  font-size: 0.9rem;
}
.back-link:hover {
  text-decoration: underline;
}
</style>
