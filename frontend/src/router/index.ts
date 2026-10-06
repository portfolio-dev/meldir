import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import OfficeDashboardView from '../views/OfficeDashboardView.vue'
import JobsDashboardView from '../views/JobsDashboardView.vue'
import ClientPortalView from '../views/ClientPortalView.vue'
import PortalHubView from '../views/PortalHubView.vue'

// Deteksi Subdomain Otomatis
function getInitialPortalRoute(): string {
  const host = window.location.hostname.toLowerCase()
  if (host.startsWith('office.')) return '/office'
  if (host.startsWith('jobs.')) return '/jobs'
  if (host.startsWith('portal.')) return '/portal'
  return '/'
}

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: () => {
      const host = window.location.hostname.toLowerCase()
      if (host.startsWith('office.')) return OfficeDashboardView
      if (host.startsWith('jobs.')) return JobsDashboardView
      if (host.startsWith('portal.')) return ClientPortalView
      return PortalHubView
    },
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
    path: '/:pathMatch(.*)*',
    redirect: '/',
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})
