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
  background-color: var(--bg-primary);
}
.hub-card {
  max-width: 900px;
  width: 100%;
  padding: 44px;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 16px;
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.05);
}
.hub-header {
  text-align: center;
  margin-bottom: 36px;
}
.badge {
  font-size: 0.75rem;
  font-weight: 800;
  letter-spacing: 0.08em;
  background: #f1f5f9;
  border: 1px solid #e2e8f0;
  color: #334155;
  padding: 4px 14px;
  border-radius: 9999px;
  display: inline-block;
  margin-bottom: 12px;
}
.hub-header h1 {
  font-size: 2.1rem;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 8px;
}
.hub-header p {
  color: #64748b;
  font-size: 1rem;
}
.portal-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(240px, 1fr));
  gap: 20px;
  margin-bottom: 32px;
}
.portal-item {
  display: block;
  padding: 26px;
  text-decoration: none;
  color: inherit;
  background: #ffffff;
  border: 1px solid #e2e8f0;
  border-radius: 12px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
  transition: transform 0.2s, border-color 0.2s, box-shadow 0.2s;
}
.portal-item:hover {
  transform: translateY(-3px);
  border-color: #0284c7;
  box-shadow: 0 8px 16px -4px rgba(2, 132, 199, 0.12);
  background: #f8fafc;
}
.portal-tag {
  font-size: 0.72rem;
  font-weight: 800;
  padding: 3px 10px;
  border-radius: 4px;
  display: inline-block;
  margin-bottom: 14px;
  letter-spacing: 0.05em;
}
.tag-office { background: #0284c7; color: #ffffff; }
.tag-jobs { background: #4f46e5; color: #ffffff; }
.tag-portal { background: #059669; color: #ffffff; }
.portal-item h3 {
  font-size: 1.15rem;
  font-weight: 700;
  margin-bottom: 8px;
  color: #0f172a;
}
.portal-item p {
  font-size: 0.88rem;
  color: #64748b;
  line-height: 1.5;
}
.hub-footer {
  text-align: center;
}
.back-link {
  color: #0284c7;
  text-decoration: none;
  font-size: 0.92rem;
  font-weight: 600;
}
.back-link:hover {
  text-decoration: underline;
}
</style>
