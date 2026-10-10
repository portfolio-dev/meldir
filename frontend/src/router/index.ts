import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import OfficeDashboardView from '../views/OfficeDashboardView.vue'
import JobsDashboardView from '../views/JobsDashboardView.vue'
import ClientPortalView from '../views/ClientPortalView.vue'
import PortalHubView from '../views/PortalHubView.vue'

// Deteksi Subdomain Otomatis untuk Component Halaman Utama (/)
function getPortalComponent() {
  const host = window.location.hostname.toLowerCase()
  if (host.startsWith('office.')) return OfficeDashboardView
  if (host.startsWith('jobs.')) return JobsDashboardView
  if (host.startsWith('portal.')) return ClientPortalView
  return PortalHubView
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: getPortalComponent(),
  },
  {
    path: '/office',
    name: 'OfficeDashboard',
    component: OfficeDashboardView,
  },
  {
    path: '/jobs',
    name: 'JobsDashboard',
    component: JobsDashboardView,
  },
  {
    path: '/portal',
    name: 'ClientPortal',
    component: ClientPortalView,
  },
  {
    path: '/hub',
    name: 'PortalHub',
    component: PortalHubView,
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

// Auto-reload saat deploy baru menggantikan chunk lama sehingga pengguna tidak mengalami blank page
router.onError((error) => {
  const isChunkError = /Failed to fetch dynamically imported module|Importing a module script failed/i.test(
    error?.message || ''
  )
  if (isChunkError) {
    console.warn('[Vite Auto-Update] Chunk usang terdeteksi setelah deploy, merefresh halaman...')
    window.location.reload()
  }
})
