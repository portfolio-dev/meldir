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
              <span class="badge-role">{{ roleDisplay }}</span>
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
            <span>Anda sedang melihat tampilan workspace Jobs Engineer. Anda memiliki hak akses penuh ke seluruh portal.</span>
          </div>
          <a href="https://office.meldir.id" class="btn-director-action">
            🏛️ Kelola Akun & Manajemen di Office →
          </a>
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

        <!-- Flash Notice -->
        <div v-if="flashNotice" class="alert-banner success">
          {{ flashNotice }}
        </div>

        <!-- TAB 1: SLA TASKBOARD -->
        <div v-if="activeTab === 'tasks'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel kpi-card-indigo">
              <div class="card-header">
                <h3>Total Tiket Ditugaskan</h3>
                <span class="tag tag-indigo">Active Sprint</span>
              </div>
              <div class="kpi-value">{{ tickets.length }} Tiket</div>
              <p class="kpi-desc">Tiket SLA terdistribusi untuk engineer</p>
            </div>

            <div class="card glass-panel kpi-card-rose">
              <div class="card-header">
                <h3>Prioritas Kritis (P1)</h3>
                <span class="tag tag-rose">SLA 4 Jam</span>
              </div>
              <div class="kpi-value">{{ p1Count }} Tiket</div>
              <p class="kpi-desc">Wajib direspon & ditangani segera</p>
            </div>

            <div class="card glass-panel kpi-card-amber">
              <div class="card-header">
                <h3>Sedang Dikerjakan</h3>
                <span class="tag tag-amber">In Progress</span>
              </div>
              <div class="kpi-value">{{ inProgressCount }} Tiket</div>
              <p class="kpi-desc">Sedang dalam proses coding / investigasi</p>
            </div>

            <div class="card glass-panel kpi-card-green">
              <div class="card-header">
                <h3>Terselesaikan (Resolved)</h3>
                <span class="tag tag-green">Done</span>
              </div>
              <div class="kpi-value">{{ resolvedCount }} Tiket</div>
              <p class="kpi-desc">Bug fixed & verifikasi lolos QA</p>
            </div>
          </div>

          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>📋 Papan Kerja Tiket SLA & Bug Tracker</h2>
                <p class="section-desc">
                  Pantau dan perbarui progres perbaikan sesuai target jaminan Service Level Agreement (SLA).
                </p>
              </div>
              <button @click="showNewTicketModal = true" class="btn-primary">
                + Tambah Tiket / Bug Baru
              </button>
            </div>

            <!-- Filter Bar -->
            <div class="filter-bar">
              <div class="search-box">
                <input
                  v-model="ticketSearch"
                  type="text"
                  placeholder="Cari kode tiket, judul, atau klien..."
                  class="search-input"
                />
              </div>
              <div class="role-filter-box">
                <select v-model="ticketFilterPriority" class="select-filter">
                  <option value="">Semua Prioritas</option>
                  <option value="p1_critical">P1 — Kritis (Maks 4 Jam)</option>
                  <option value="p2_major">P2 — Mayor (Maks 12 Jam)</option>
                  <option value="p3_low">P3 — Minor / Request (Maks 48 Jam)</option>
                </select>
              </div>
              <div class="role-filter-box">
                <select v-model="ticketFilterStatus" class="select-filter">
                  <option value="">Semua Status</option>
                  <option value="open">Open (Baru)</option>
                  <option value="in_progress">In Progress (Dikerjakan)</option>
                  <option value="review">Review / Testing</option>
                  <option value="resolved">Resolved (Selesai)</option>
                </select>
              </div>
            </div>

            <!-- Tickets Table -->
            <div class="table-container">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Kode</th>
                    <th>Judul Kendala & Klien</th>
                    <th>Prioritas</th>
                    <th>Status</th>
                    <th>Deadline SLA</th>
                    <th>Jam Terpakai</th>
                    <th style="text-align: right">Aksi Progres</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="t in filteredTickets" :key="t.id">
                    <td class="font-mono text-muted">{{ t.code }}</td>
                    <td>
                      <div>
                        <strong>{{ t.title }}</strong>
                        <div class="text-xs text-muted">Klien: {{ t.clientName }} • {{ t.project }}</div>
                      </div>
                    </td>
                    <td>
                      <span :class="['priority-badge', t.priority]">
                        {{ formatPriority(t.priority) }}
                      </span>
                    </td>
                    <td>
                      <span :class="['status-badge', t.status]">
                        {{ formatStatus(t.status) }}
                      </span>
                    </td>
                    <td class="text-xs font-mono">
                      {{ t.deadline }}
                    </td>
                    <td class="font-mono">{{ t.hoursSpent }} Jam</td>
                    <td style="text-align: right">
                      <div class="action-btn-group">
                        <select
                          :value="t.status"
                          @change="changeTicketStatus(t.id, ($event.target as HTMLSelectElement).value)"
                          class="select-action-status"
                        >
                          <option value="open">Open</option>
                          <option value="in_progress">In Progress</option>
                          <option value="review">Review</option>
                          <option value="resolved">Resolved</option>
                        </select>
                        <button @click="openQuickLog(t)" class="btn-action edit" title="Catat Jam">
                          ⏱️ Log Jam
                        </button>
                      </div>
                    </td>
                  </tr>
                  <tr v-if="filteredTickets.length === 0">
                    <td colspan="7" class="text-center py-6 text-muted">
                      Tidak ada tiket yang sesuai dengan filter pencarian.
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- TAB 2: TIMESHEET LOGGER -->
        <div v-else-if="activeTab === 'timesheet'" class="tab-content">
          <div class="grid-overview">
            <div class="card glass-panel kpi-card-blue">
              <div class="card-header">
                <h3>Total Jam Logged</h3>
                <span class="tag tag-blue">Bulan Berjalan</span>
              </div>
              <div class="kpi-value">{{ totalLoggedHours }} Jam</div>
              <p class="kpi-desc">Akumulasi seluruh pengerjaan engineer</p>
            </div>

            <div class="card glass-panel kpi-card-green">
              <div class="card-header">
                <h3>Estimasi Kompensasi</h3>
                <span class="tag tag-green">Billing Rate</span>
              </div>
              <div class="kpi-value">{{ formatCurrency(totalLoggedHours * 175000) }}</div>
              <p class="kpi-desc">Tarif acuan Rp 175.000 / Jam pengerjaan</p>
            </div>

            <div class="card glass-panel kpi-card-indigo">
              <div class="card-header">
                <h3>Pemotong Saldo Add-On</h3>
                <span class="tag tag-indigo">Metered SLA</span>
              </div>
              <div class="kpi-value">12.5 Jam</div>
              <p class="kpi-desc">Otomatis dipotong dari saldo kuota klien</p>
            </div>
          </div>

          <div class="two-col-layout">
            <!-- Form Input Log Jam Kerja -->
            <section class="section-panel glass-panel">
              <h2>⏱️ Catat Waktu Pengerjaan (Timesheet)</h2>
              <p class="section-desc">
                Input jam kerja aktual untuk tiket maintenance atau fitur add-on klien.
              </p>

              <form @submit.prevent="submitTimesheet" class="form-grid">
                <div class="form-group">
                  <label>Pilih Proyek / Kontrak *</label>
                  <select v-model="timesheetForm.project" required class="form-input">
                    <option value="PT. Surya Logistik — Managed Care SLA">PT. Surya Logistik — Managed Care SLA</option>
                    <option value="CV. Sejahtera Abadi — Core Backend API">CV. Sejahtera Abadi — Core Backend API</option>
                    <option value="Yayasan Harapan Bangsa — Portal Siswa">Yayasan Harapan Bangsa — Portal Siswa</option>
                    <option value="Internal Meldir — Core Engine Maintenance">Internal Meldir — Core Engine Maintenance</option>
                  </select>
                </div>

                <div class="form-group">
                  <label>Tiket Terkait (Opsional)</label>
                  <select v-model="timesheetForm.ticketCode" class="form-input">
                    <option value="">Tanpa Tiket (Maintenance Umum)</option>
                    <option v-for="t in tickets" :key="t.id" :value="t.code">
                      {{ t.code }} — {{ t.title }}
                    </option>
                  </select>
                </div>

                <div class="form-row">
                  <div class="form-group flex-1">
                    <label>Tanggal Kerja *</label>
                    <input v-model="timesheetForm.date" type="date" required class="form-input" />
                  </div>
                  <div class="form-group flex-1">
                    <label>Durasi (Jam) *</label>
                    <input
                      v-model.number="timesheetForm.hours"
                      type="number"
                      step="0.5"
                      min="0.5"
                      max="16"
                      required
                      placeholder="Contoh: 2.5"
                      class="form-input"
                    />
                  </div>
                </div>

                <div class="form-group">
                  <label>Deskripsi Detail Pekerjaan Selesai *</label>
                  <textarea
                    v-model="timesheetForm.description"
                    rows="3"
                    required
                    placeholder="Jelaskan modul/bug yang diperbaiki, branch Git, dan hasil pengetesan..."
                    class="form-input"
                  ></textarea>
                </div>

                <button type="submit" class="btn-primary" style="width: 100%;">
                  💾 Simpan Log Jam Kerja
                </button>
              </form>
            </section>

            <!-- Tabel Riwayat Timesheet Terkini -->
            <section class="section-panel glass-panel">
              <h2>📜 Riwayat Jam Kerja Terkini</h2>
              <p class="section-desc">
                Daftar log jam kerja yang telah direkam dan siap diaudit manajemen.
              </p>

              <div class="table-container">
                <table class="data-table">
                  <thead>
                    <tr>
                      <th>Tanggal</th>
                      <th>Proyek & Tiket</th>
                      <th>Jam</th>
                      <th>Deskripsi</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr v-for="log in timesheets" :key="log.id">
                      <td class="font-mono text-xs">{{ log.date }}</td>
                      <td>
                        <strong>{{ log.project }}</strong>
                        <div v-if="log.ticketCode" class="text-xs text-muted">{{ log.ticketCode }}</div>
                      </td>
                      <td class="font-mono font-bold text-indigo">{{ log.hours }} Jam</td>
                      <td class="text-xs text-muted">{{ log.description }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </section>
          </div>
        </div>

        <!-- TAB 3: BUKTI POTONG PPH 21 -->
        <div v-else-if="activeTab === 'tax21'" class="tab-content">
          <section class="section-panel glass-panel">
            <div class="section-header-row">
              <div>
                <h2>🏛️ Bukti Pemotongan Pajak PPh 21 Tenaga Ahli</h2>
                <p class="section-desc">
                  Slip bukti potong resmi formulir 21/26 dari PT. Melayani Digital Raya yang dapat Anda gunakan sebagai kredit pajak pada SPT Tahunan Pribadi.
                </p>
              </div>
            </div>

            <div class="tax-info-banner">
              <div class="tax-info-icon">ℹ️</div>
              <div>
                <strong>Dasar Pengenaan Pajak (DPP) Tenaga Ahli Bukan Pegawai:</strong>
                <p>Sesuai PP 58/2023 & UU HPP: DPP = 50% × Penghasilan Bruto. Tarif progresif Pasal 17 dikenakan atas DPP tersebut.</p>
              </div>
            </div>

            <div class="table-container" style="margin-top: 20px;">
              <table class="data-table">
                <thead>
                  <tr>
                    <th>Nomor Bukti Potong</th>
                    <th>Masa / Tahun</th>
                    <th>Penghasilan Bruto</th>
                    <th>DPP (50%)</th>
                    <th>Tarif Pajak</th>
                    <th>PPh 21 Terpotong</th>
                    <th>Status DJP</th>
                    <th style="text-align: right">Slip Resmi</th>
                  </tr>
                </thead>
                <tbody>
                  <tr v-for="slip in withholdingSlips" :key="slip.id">
                    <td class="font-mono text-muted">#{{ slip.slipNumber }}</td>
                    <td><strong>{{ slip.period }}</strong></td>
                    <td class="font-mono">{{ formatCurrency(slip.bruto) }}</td>
                    <td class="font-mono text-muted">{{ formatCurrency(slip.dpp) }}</td>
                    <td><span class="tag tag-blue">{{ slip.rate }}</span></td>
                    <td class="font-mono font-bold text-rose">{{ formatCurrency(slip.taxAmount) }}</td>
                    <td><span class="badge-active">✓ Terlapor e-Bupot</span></td>
                    <td style="text-align: right">
                      <button @click="openSlipModal(slip)" class="btn-action edit">
                        🖨️ Cetak Slip
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </section>
        </div>

        <!-- TAB 4: GITHUB & REPOSITORI -->
        <div v-else-if="activeTab === 'repos'" class="tab-content">
          <section class="section-panel glass-panel">
            <h2>🐙 Akses Repositori GitHub & Standar Kerja</h2>
            <p class="section-desc">
              Kode sumber resmi PT. Melayani Digital Raya diakses melalui organisasi privat GitHub. Pastikan username GitHub Anda sudah terdaftar.
            </p>

            <div class="github-profile-card">
              <div class="gh-icon">🐙</div>
              <div class="gh-info">
                <h4>Username GitHub Terdaftar:</h4>
                <div class="font-mono text-indigo font-bold">{{ authStore.user?.github_username || 'Belum diatur (Hubungi Direktur/Admin)' }}</div>
                <p class="text-xs text-muted">Akses repositori otomatis diberikan ke username di atas.</p>
              </div>
            </div>

            <div class="repo-grid">
              <div v-for="repo in assignedRepos" :key="repo.name" class="repo-card glass-panel">
                <div class="repo-header">
                  <h3>📦 {{ repo.name }}</h3>
                  <span class="tag tag-indigo">{{ repo.visibility }}</span>
                </div>
                <p class="repo-desc">{{ repo.description }}</p>
                <div class="repo-meta font-mono text-xs">
                  <span>Branch Aktif: <strong>{{ repo.defaultBranch }}</strong></span>
                  <span>•</span>
                  <span>Stack: {{ repo.stack }}</span>
                </div>
                <div class="repo-actions">
                  <a :href="repo.url" target="_blank" class="btn-gh-link">
                    Buka Repositori ↗
                  </a>
                </div>
              </div>
            </div>
          </section>
        </div>
      </main>
    </div>

    <!-- Modal Buat Tiket Baru -->
    <div v-if="showNewTicketModal" class="modal-backdrop" @click.self="showNewTicketModal = false">
      <div class="modal-card">
        <div class="modal-header">
          <h3>+ Buat Tiket SLA / Kendala Baru</h3>
          <button @click="showNewTicketModal = false" class="btn-close-modal">✕</button>
        </div>
        <form @submit.prevent="saveNewTicket" class="modal-form">
          <div class="form-group">
            <label>Judul Kendala / Fitur *</label>
            <input v-model="newTicketForm.title" type="text" required placeholder="Contoh: Bug perhitungan diskon invoice checkout" class="form-input" />
          </div>
          <div class="form-row">
            <div class="form-group flex-1">
              <label>Klien / Proyek Terkait *</label>
              <input v-model="newTicketForm.clientName" type="text" required placeholder="Contoh: PT. Surya Logistik" class="form-input" />
            </div>
            <div class="form-group flex-1">
              <label>Prioritas SLA *</label>
              <select v-model="newTicketForm.priority" required class="form-input">
                <option value="p1_critical">P1 — Kritis (Maks 4 Jam)</option>
                <option value="p2_major">P2 — Mayor (Maks 12 Jam)</option>
                <option value="p3_low">P3 — Minor / Request (Maks 48 Jam)</option>
              </select>
            </div>
          </div>
          <div class="form-group">
            <label>Deskripsi Masalah / Langkah Reproduksi *</label>
            <textarea v-model="newTicketForm.description" rows="3" required placeholder="Jelaskan detail error log atau perilaku yang diharapkan..." class="form-input"></textarea>
          </div>
          <div class="modal-footer">
            <button type="button" @click="showNewTicketModal = false" class="btn-secondary">Batal</button>
            <button type="submit" class="btn-primary">Buat Tiket Sekarang</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Modal Cetak Slip Pajak PPh 21 -->
    <div v-if="selectedSlip" class="modal-backdrop" @click.self="selectedSlip = null">
      <div class="modal-card modal-slip-view">
        <div class="modal-header">
          <h3>Bukti Potong PPh 21 Resmi</h3>
          <button @click="selectedSlip = null" class="btn-close-modal">✕</button>
        </div>
        <div class="slip-content" id="printable-slip">
          <div class="slip-corp-header">
            <h4>PT. MELAYANI DIGITAL RAYA</h4>
            <p>NPWP: 01.234.567.8-012.000 • SK Kemenkumham RI: AHU-0012345.AH.01.01.TAHUN 2026</p>
            <div class="slip-title">BUKTI PEMOTONGAN PPH PASAL 21 (FORMULIR 21/26)</div>
            <div class="slip-number font-mono">Nomor: {{ selectedSlip.slipNumber }}</div>
          </div>

          <div class="slip-details-grid">
            <div class="slip-row">
              <span class="label">Nama Penerima Penghasilan:</span>
              <span class="val"><strong>{{ authStore.user?.name }}</strong></span>
            </div>
            <div class="slip-row">
              <span class="label">Email & Telepon:</span>
              <span class="val">{{ authStore.user?.email }} • {{ authStore.user?.phone_wa }}</span>
            </div>
            <div class="slip-row">
              <span class="label">Masa / Tahun Pajak:</span>
              <span class="val font-bold">{{ selectedSlip.period }}</span>
            </div>
            <div class="slip-row">
              <span class="label">Klasifikasi Penghasilan:</span>
              <span class="val">Imbalan Kepada Tenaga Ahli / Jasa Perangkat Lunak</span>
            </div>
          </div>

          <div class="slip-calc-table">
            <table class="data-table">
              <thead>
                <tr>
                  <th>Jumlah Penghasilan Bruto</th>
                  <th>Dasar Pengenaan Pajak (50%)</th>
                  <th>Tarif</th>
                  <th>PPh 21 Dipotong</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td class="font-mono font-bold">{{ formatCurrency(selectedSlip.bruto) }}</td>
                  <td class="font-mono">{{ formatCurrency(selectedSlip.dpp) }}</td>
                  <td class="font-bold">{{ selectedSlip.rate }}</td>
                  <td class="font-mono font-bold text-rose">{{ formatCurrency(selectedSlip.taxAmount) }}</td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="slip-footer-sign">
            <div class="seal-badge">
              <div class="seal-text">PT. MELDIR</div>
              <div class="seal-sub">AUTHENTICATED</div>
            </div>
            <div class="sign-block">
              <div>Jakarta, {{ selectedSlip.period }}</div>
              <div class="sign-title">Pemotong Pajak / Direktur Utama</div>
              <div class="sign-space"></div>
              <div class="sign-name">PT. Melayani Digital Raya</div>
            </div>
          </div>
        </div>
        <div class="modal-footer">
          <button @click="printSlip" class="btn-primary">🖨️ Cetak / Simpan PDF</button>
          <button @click="selectedSlip = null" class="btn-secondary">Tutup</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import LoginForm from '../components/LoginForm.vue'
import { useAuthStore } from '../stores/authStore'

const authStore = useAuthStore()
const currentHost = ref(window.location.host)
const activeTab = ref('tasks')
const flashNotice = ref('')

const tabs = [
  { id: 'tasks', label: 'SLA Taskboard (Tiket)', icon: '📋' },
  { id: 'timesheet', label: 'Timesheet Logger', icon: '⏱️' },
  { id: 'tax21', label: 'Bukti Potong PPh 21', icon: '📄' },
  { id: 'repos', label: 'Repositori GitHub', icon: '🐙' },
]

const userInitials = computed(() => {
  if (!authStore.user?.name) return 'EN'
  return authStore.user.name
    .split(' ')
    .map(w => w[0])
    .slice(0, 2)
    .join('')
    .toUpperCase()
})

const roleDisplay = computed(() => {
  switch (authStore.user?.role) {
    case 'engineer': return 'Technical Software Engineer'
    case 'direktur': return 'Direktur Utama (Superadmin)'
    case 'admin': return 'Office Administrator'
    default: return 'Engineer'
  }
})

// Data Tiket SLA
const tickets = ref([
  {
    id: 1,
    code: 'TKT-2026-081',
    title: 'Database connection pool timeout under peak traffic',
    clientName: 'PT. Surya Logistik',
    project: 'Logistics ERP Core',
    priority: 'p1_critical',
    status: 'in_progress',
    deadline: '4 Jam (Hari Ini 23:00 WIB)',
    hoursSpent: 2.5,
    description: 'Postgres connection pool maxed out at 25 connections. Need to tune max connections and add Redis cache.',
  },
  {
    id: 2,
    code: 'TKT-2026-082',
    title: 'Implementasi Webhook Notifikasi Pembayaran Invoice',
    clientName: 'CV. Sejahtera Abadi',
    project: 'B2B Client Portal',
    priority: 'p2_major',
    status: 'open',
    deadline: '12 Jam (Besok 08:00 WIB)',
    hoursSpent: 0,
    description: 'Integrasi callback payment gateway untuk update status invoice otomatis.',
  },
  {
    id: 3,
    code: 'TKT-2026-083',
    title: 'Penyesuaian teks template email BAST & e-Materai',
    clientName: 'Yayasan Harapan Bangsa',
    project: 'Scholarship Management',
    priority: 'p3_low',
    status: 'resolved',
    deadline: '48 Jam (Selesai)',
    hoursSpent: 1.5,
    description: 'Update format alamat resmi dan nomor surat izin operasional yayasan pada footer BAST.',
  },
])

const ticketSearch = ref('')
const ticketFilterPriority = ref('')
const ticketFilterStatus = ref('')

const filteredTickets = computed(() => {
  return tickets.value.filter(t => {
    const matchSearch = ticketSearch.value === '' ||
      t.title.toLowerCase().includes(ticketSearch.value.toLowerCase()) ||
      t.code.toLowerCase().includes(ticketSearch.value.toLowerCase()) ||
      t.clientName.toLowerCase().includes(ticketSearch.value.toLowerCase())
    const matchPriority = ticketFilterPriority.value === '' || t.priority === ticketFilterPriority.value
    const matchStatus = ticketFilterStatus.value === '' || t.status === ticketFilterStatus.value
    return matchSearch && matchPriority && matchStatus
  })
})

const p1Count = computed(() => tickets.value.filter(t => t.priority === 'p1_critical').length)
const inProgressCount = computed(() => tickets.value.filter(t => t.status === 'in_progress').length)
const resolvedCount = computed(() => tickets.value.filter(t => t.status === 'resolved').length)

async function fetchTickets() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/tickets', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data) && result.data.length > 0) {
      tickets.value = result.data.map((item: any) => ({
        id: item.id,
        code: item.ticket_code,
        title: item.title,
        clientName: item.client_name || 'Klien Korporat',
        project: item.project_name || 'Managed Care SLA',
        priority: item.priority,
        status: item.status,
        deadline: new Date(item.sla_deadline).toLocaleString('id-ID', { dateStyle: 'medium', timeStyle: 'short' }) + ' WIB',
        hoursSpent: item.hours_spent || 0,
        description: item.description,
      }))
    }
  } catch (err) {
    console.error('Gagal mengambil data tiket dari backend:', err)
  }
}

// Modal Buat Tiket
const showNewTicketModal = ref(false)
const newTicketForm = ref({
  title: '',
  clientName: '',
  priority: 'p2_major',
  description: '',
})

async function saveNewTicket() {
  if (!newTicketForm.value.title) return

  if (authStore.token) {
    try {
      const res = await fetch('/api/v1/tickets/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          title: newTicketForm.value.title,
          description: newTicketForm.value.description || newTicketForm.value.title,
          priority: newTicketForm.value.priority,
          project_name: 'Managed Care SLA',
        }),
      })
      const result = await res.json()
      if (res.ok && result.success) {
        await fetchTickets()
        showNewTicketModal.value = false
        const code = result.data?.ticket_code || 'TKT-SLA'
        newTicketForm.value = { title: '', clientName: '', priority: 'p2_major', description: '' }
        showFlashMsg(`Tiket baru ${code} berhasil dibuat dan ditambahkan ke database!`)
        return
      }
    } catch (e) {
      console.error('Gagal membuat tiket di backend:', e)
    }
  }

  // Fallback local update
  const newId = tickets.value.length + 1
  const code = `TKT-2026-0${80 + newId}`
  tickets.value.unshift({
    id: newId,
    code,
    title: newTicketForm.value.title,
    clientName: newTicketForm.value.clientName || 'PT. Surya Logistik',
    project: 'Managed Care SLA',
    priority: newTicketForm.value.priority,
    status: 'open',
    deadline: newTicketForm.value.priority === 'p1_critical' ? '4 Jam' : '12 Jam',
    hoursSpent: 0,
    description: newTicketForm.value.description,
  })
  showNewTicketModal.value = false
  newTicketForm.value = { title: '', clientName: '', priority: 'p2_major', description: '' }
  showFlashMsg(`Tiket baru ${code} berhasil dibuat dan ditambahkan ke antrian!`)
}

async function changeTicketStatus(id: number, newStatus: string) {
  const t = tickets.value.find(item => item.id === id)
  if (t) {
    t.status = newStatus
    showFlashMsg(`Status tiket ${t.code} diubah menjadi "${formatStatus(newStatus)}"`)

    if (authStore.token) {
      try {
        await fetch('/api/v1/tickets/status', {
          method: 'PATCH',
          headers: {
            'Content-Type': 'application/json',
            Authorization: `Bearer ${authStore.token}`,
          },
          body: JSON.stringify({ id, status: newStatus }),
        })
      } catch (e) {
        console.error('Gagal update status tiket di backend:', e)
      }
    }
  }
}

// Timesheet Data
const timesheets = ref([
  {
    id: 1,
    project: 'PT. Surya Logistik — Managed Care SLA',
    ticketCode: 'TKT-2026-081',
    date: '2026-10-06',
    hours: 2.5,
    description: 'Diagnostik PostgreSQL connection pool bottleneck dan konfigurasi Redis caching.',
  },
  {
    id: 2,
    project: 'Internal Meldir — Core Engine Maintenance',
    ticketCode: '',
    date: '2026-10-05',
    hours: 4.0,
    description: 'Refactoring RESTful handler Golang dan optimasi CORS OpenLiteSpeed.',
  },
  {
    id: 3,
    project: 'CV. Sejahtera Abadi — Core Backend API',
    ticketCode: 'TKT-2026-079',
    date: '2026-10-04',
    hours: 3.5,
    description: 'Implementasi middleware otentikasi JWT dan pembatasan peran akses user.',
  },
])

async function fetchTimesheets() {
  if (!authStore.token) return
  try {
    const res = await fetch('/api/v1/timesheets', {
      headers: { Authorization: `Bearer ${authStore.token}` },
    })
    const result = await res.json()
    if (res.ok && result.success && Array.isArray(result.data) && result.data.length > 0) {
      timesheets.value = result.data.map((item: any) => ({
        id: item.id,
        project: item.project_name || 'Managed Care SLA',
        ticketCode: item.ticket_code || '',
        date: item.log_date,
        hours: item.hours_spent,
        description: item.work_description,
      }))
    }
  } catch (err) {
    console.error('Gagal mengambil timesheet dari backend:', err)
  }
}

const timesheetForm = ref({
  project: 'PT. Surya Logistik — Managed Care SLA',
  ticketCode: '',
  date: new Date().toISOString().substring(0, 10),
  hours: 2,
  description: '',
})

const totalLoggedHours = computed(() => {
  return timesheets.value.reduce((acc, curr) => acc + curr.hours, 0)
})

async function submitTimesheet() {
  if (!timesheetForm.value.hours || !timesheetForm.value.description) return

  if (authStore.token) {
    try {
      const res = await fetch('/api/v1/timesheets/log', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify({
          project_name: timesheetForm.value.project,
          ticket_code: timesheetForm.value.ticketCode,
          hours_spent: Number(timesheetForm.value.hours),
          work_description: timesheetForm.value.description,
          log_date: timesheetForm.value.date,
        }),
      })
      const result = await res.json()
      if (res.ok && result.success) {
        await fetchTimesheets()
        if (timesheetForm.value.ticketCode) {
          await fetchTickets()
        }
        showFlashMsg(`Berhasil mencatat ${timesheetForm.value.hours} jam kerja ke database timesheet!`)
        timesheetForm.value.description = ''
        return
      }
    } catch (e) {
      console.error('Gagal submit timesheet ke backend:', e)
    }
  }

  // Fallback local update
  timesheets.value.unshift({
    id: Date.now(),
    project: timesheetForm.value.project,
    ticketCode: timesheetForm.value.ticketCode,
    date: timesheetForm.value.date,
    hours: Number(timesheetForm.value.hours),
    description: timesheetForm.value.description,
  })

  if (timesheetForm.value.ticketCode) {
    const t = tickets.value.find(item => item.code === timesheetForm.value.ticketCode)
    if (t) t.hoursSpent += Number(timesheetForm.value.hours)
  }

  showFlashMsg(`Berhasil mencatat ${timesheetForm.value.hours} jam kerja ke timesheet!`)
  timesheetForm.value.description = ''
}

function openQuickLog(t: any) {
  activeTab.value = 'timesheet'
  timesheetForm.value.project = t.clientName + ' — Managed Care SLA'
  timesheetForm.value.ticketCode = t.code
}

// Pajak PPh 21 Slips
const withholdingSlips = ref([
  {
    id: 1,
    slipNumber: 'BP-21/2026/09/014',
    period: 'September 2026',
    bruto: 15000000,
    dpp: 7500000,
    rate: '5%',
    taxAmount: 375000,
  },
  {
    id: 2,
    slipNumber: 'BP-21/2026/08/011',
    period: 'Agustus 2026',
    bruto: 12500000,
    dpp: 6250000,
    rate: '5%',
    taxAmount: 312500,
  },
  {
    id: 3,
    slipNumber: 'BP-21/2026/07/008',
    period: 'Juli 2026',
    bruto: 18000000,
    dpp: 9000000,
    rate: '5%',
    taxAmount: 450000,
  },
])

const selectedSlip = ref<any>(null)
function openSlipModal(slip: any) {
  selectedSlip.value = slip
}
function printSlip() {
  window.print()
}

// GitHub Repositories
const assignedRepos = ref([
  {
    name: 'portfolio-dev/meldir',
    visibility: 'Privat',
    description: 'Repositori monorepo inti: Golang RESTful Backend & Vue 3 Multi-Portal Frontend.',
    defaultBranch: 'main',
    stack: 'Go 1.22, Vue 3, Pinia, Vite, PostgreSQL',
    url: 'https://github.com/portfolio-dev/meldir',
  },
  {
    name: 'meldir-enterprise/client-mobile',
    visibility: 'Privat',
    description: 'Aplikasi klien native/cross-platform untuk monitoring SLA dan e-Materai SPK.',
    defaultBranch: 'main',
    stack: 'TypeScript, Flutter / Vue Mobile',
    url: 'https://github.com/portfolio-dev/meldir',
  },
])

function formatPriority(p: string) {
  switch (p) {
    case 'p1_critical': return 'P1 — Kritis'
    case 'p2_major': return 'P2 — Mayor'
    case 'p3_low': return 'P3 — Minor'
    default: return p
  }
}

function formatStatus(s: string) {
  switch (s) {
    case 'open': return 'Open'
    case 'in_progress': return 'In Progress'
    case 'review': return 'Review'
    case 'resolved': return 'Resolved'
    default: return s
  }
}

function formatCurrency(val: number) {
  return new Intl.NumberFormat('id-ID', {
    style: 'currency',
    currency: 'IDR',
    maximumFractionDigits: 0,
  }).format(val)
}

function showFlashMsg(msg: string) {
  flashNotice.value = msg
  setTimeout(() => {
    flashNotice.value = ''
  }, 4000)
}

async function handleLogout() {
  await authStore.logout()
}

onMounted(async () => {
  if (authStore.token) {
    await authStore.fetchProfile()
    await Promise.all([fetchTickets(), fetchTimesheets()])
  }
})
</script>

<style scoped>
.portal-container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 24px;
  min-height: 100vh;
  background-color: #edf2f7;
}

/* Auth Unauthenticated */
.auth-wrapper {
  max-width: 600px;
  margin: 40px auto;
  text-align: center;
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

/* Header */
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
}
.badge-jobs {
  background: #4f46e5;
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

/* User Menu */
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
}

/* Navigation Tabs */
.module-nav-tabs {
  display: flex;
  gap: 8px;
  margin-bottom: 24px;
  border-bottom: 2px solid #cbd5e1;
  padding-bottom: 8px;
  overflow-x: auto;
}
.nav-tab-btn {
  background: #ffffff;
  border: 1px solid #cbd5e1;
  color: #475569;
  padding: 9px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}
.nav-tab-btn:hover {
  background: #f8fafc;
  color: #0f172a;
}
.nav-tab-btn.active {
  background: #4f46e5;
  border-color: #4f46e5;
  color: #ffffff;
  box-shadow: 0 3px 8px rgba(79, 70, 229, 0.35);
}

/* Alert Banner */
.alert-banner {
  padding: 12px 18px;
  border-radius: 8px;
  font-size: 0.9rem;
  margin-bottom: 20px;
  font-weight: 600;
}
.alert-banner.success {
  background: #ecfdf5;
  border: 1px solid #a7f3d0;
  color: #047857;
}

/* Grid & Cards */
.grid-overview {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(260px, 1fr));
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
.kpi-card-indigo { border-top: 4px solid #4f46e5; }
.kpi-card-blue { border-top: 4px solid #0284c7; }
.kpi-card-amber { border-top: 4px solid #f59e0b; }
.kpi-card-rose { border-top: 4px solid #e11d48; }
.kpi-card-green { border-top: 4px solid #10b981; }

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

/* Tags & Badges */
.tag {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 600;
}
.tag-indigo { background: #e0e7ff; color: #3730a3; }
.tag-blue { background: #e0f2fe; color: #0369a1; }
.tag-rose { background: #ffe4e6; color: #9f1239; }
.tag-amber { background: #fef3c7; color: #92400e; }
.tag-green { background: #d1fae5; color: #065f46; }

.priority-badge {
  font-size: 0.75rem;
  padding: 3px 8px;
  border-radius: 4px;
  font-weight: 700;
}
.priority-badge.p1_critical { background: #fee2e2; color: #991b1b; }
.priority-badge.p2_major { background: #fef3c7; color: #92400e; }
.priority-badge.p3_low { background: #e0f2fe; color: #0369a1; }

.status-badge {
  font-size: 0.72rem;
  padding: 2px 8px;
  border-radius: 4px;
  font-weight: 700;
  text-transform: capitalize;
}
.status-badge.open { background: #f1f5f9; color: #475569; }
.status-badge.in_progress { background: #fef3c7; color: #92400e; }
.status-badge.review { background: #e0e7ff; color: #4338ca; }
.status-badge.resolved { background: #d1fae5; color: #065f46; }

/* Sections */
.section-panel {
  padding: 28px;
  margin-bottom: 24px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.04);
}
.section-header-row {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  flex-wrap: wrap;
  gap: 16px;
  margin-bottom: 16px;
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
  margin-bottom: 20px;
  line-height: 1.6;
}

/* Two Column Layout */
.two-col-layout {
  display: grid;
  grid-template-columns: 1fr 1.2fr;
  gap: 24px;
}
@media (max-width: 900px) {
  .two-col-layout { grid-template-columns: 1fr; }
}

/* Filter Bar */
.filter-bar {
  display: flex;
  gap: 12px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}
.search-box { flex: 1; min-width: 240px; }
.search-input {
  width: 100%;
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
}
.select-filter {
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
}

/* Table */
.table-container {
  overflow-x: auto;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
}
.data-table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.88rem;
  text-align: left;
}
.data-table th {
  background: #f1f5f9;
  padding: 12px 16px;
  font-weight: 700;
  color: #334155;
  border-bottom: 2px solid #cbd5e1;
}
.data-table td {
  padding: 14px 16px;
  border-bottom: 1px solid #f1f5f9;
  color: #1e293b;
}
.data-table tbody tr:hover { background: #f8fafc; }

.select-action-status {
  padding: 4px 8px;
  border-radius: 6px;
  border: 1px solid #cbd5e1;
  font-size: 0.8rem;
  background: #ffffff;
}

/* Buttons */
.btn-primary {
  background: #4f46e5;
  color: #ffffff;
  padding: 9px 18px;
  border-radius: 8px;
  font-size: 0.88rem;
  font-weight: 700;
  border: none;
  cursor: pointer;
  transition: all 0.2s;
  box-shadow: 0 2px 6px rgba(79, 70, 229, 0.3);
}
.btn-primary:hover {
  background: #4338ca;
}
.btn-secondary {
  background: #f1f5f9;
  color: #334155;
  border: 1px solid #cbd5e1;
  padding: 9px 16px;
  border-radius: 8px;
  font-weight: 600;
  cursor: pointer;
}
.action-btn-group {
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  align-items: center;
}
.btn-action.edit {
  background: #e0e7ff;
  color: #4338ca;
  border: 1px solid #c7d2fe;
  padding: 5px 10px;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 600;
  cursor: pointer;
}

/* Form */
.form-grid {
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}
.form-group label {
  font-size: 0.82rem;
  font-weight: 700;
  color: #334155;
}
.form-input {
  padding: 10px 14px;
  border: 1px solid #cbd5e1;
  border-radius: 8px;
  font-size: 0.88rem;
  background: #f8fafc;
}
.form-row {
  display: flex;
  gap: 12px;
}
.flex-1 { flex: 1; }

/* Tax Info Banner */
.tax-info-banner {
  display: flex;
  align-items: flex-start;
  gap: 14px;
  background: #eff6ff;
  border: 1px solid #bfdbfe;
  padding: 16px;
  border-radius: 8px;
  color: #1e3a8a;
  font-size: 0.88rem;
  line-height: 1.5;
}
.tax-info-icon { font-size: 1.4rem; }

/* GitHub Profile */
.github-profile-card {
  display: flex;
  align-items: center;
  gap: 16px;
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  padding: 20px;
  border-radius: 10px;
  margin-bottom: 24px;
}
.gh-icon { font-size: 2.2rem; }
.gh-info h4 { margin-bottom: 4px; font-size: 0.95rem; }
.repo-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(320px, 1fr));
  gap: 20px;
}
.repo-card {
  padding: 20px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 12px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.05);
}
.repo-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}
.repo-header h3 { font-size: 1.05rem; color: #0f172a; }
.repo-desc { font-size: 0.85rem; color: #64748b; margin-bottom: 12px; }
.repo-meta { color: #475569; margin-bottom: 16px; display: flex; gap: 8px; flex-wrap: wrap; }
.btn-gh-link {
  display: inline-block;
  background: #1e293b;
  color: #ffffff;
  padding: 6px 14px;
  border-radius: 6px;
  text-decoration: none;
  font-size: 0.82rem;
  font-weight: 600;
}

/* Modals */
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(15, 23, 42, 0.6);
  backdrop-filter: blur(4px);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 100;
  padding: 20px;
}
.modal-card {
  background: #ffffff;
  border-radius: 14px;
  max-width: 560px;
  width: 100%;
  border: 1px solid #cbd5e1;
  box-shadow: 0 20px 25px -5px rgba(0, 0, 0, 0.15);
  overflow: hidden;
}
.modal-slip-view { max-width: 700px; }
.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 18px 24px;
  border-bottom: 1px solid #e2e8f0;
}
.btn-close-modal {
  background: transparent;
  border: none;
  font-size: 1.2rem;
  cursor: pointer;
  color: #64748b;
}
.modal-form {
  padding: 24px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}
.modal-footer {
  padding: 16px 24px;
  background: #f8fafc;
  border-top: 1px solid #e2e8f0;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Slip Print View */
.slip-content {
  padding: 28px;
  background: #ffffff;
}
.slip-corp-header {
  text-align: center;
  border-bottom: 2px solid #0f172a;
  padding-bottom: 14px;
  margin-bottom: 20px;
}
.slip-corp-header h4 { font-size: 1.2rem; font-weight: 800; color: #0f172a; }
.slip-corp-header p { font-size: 0.75rem; color: #64748b; margin-bottom: 10px; }
.slip-title { font-size: 0.95rem; font-weight: 800; color: #1e293b; letter-spacing: 0.05em; }
.slip-number { font-size: 0.8rem; color: #475569; }

.slip-details-grid {
  margin-bottom: 20px;
  font-size: 0.88rem;
  display: flex;
  flex-direction: column;
  gap: 8px;
}
.slip-row { display: flex; justify-content: space-between; border-bottom: 1px dashed #e2e8f0; padding-bottom: 4px; }
.slip-row .label { color: #64748b; }
.slip-calc-table { margin-bottom: 24px; }
.slip-footer-sign {
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
  margin-top: 24px;
}
.seal-badge {
  border: 3px double #0284c7;
  padding: 8px 14px;
  border-radius: 8px;
  text-align: center;
  color: #0284c7;
  font-weight: 800;
}
.seal-text { font-size: 0.9rem; letter-spacing: 0.1em; }
.seal-sub { font-size: 0.65rem; }
.sign-block { text-align: right; font-size: 0.85rem; }
.sign-space { height: 48px; }
.sign-name { font-weight: 800; text-decoration: underline; }

.font-mono { font-family: var(--font-mono); }
.font-bold { font-weight: 700; }
.text-rose { color: #e11d48; }
.text-indigo { color: #4f46e5; }
.text-xs { font-size: 0.75rem; }
.text-muted { color: #64748b; }
.text-center { text-align: center; }
.py-6 { padding-top: 24px; padding-bottom: 24px; }
.badge-active { background: #d1fae5; color: #065f46; padding: 2px 8px; border-radius: 4px; font-weight: 700; font-size: 0.75rem; }
</style>
