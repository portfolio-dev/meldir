import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

export interface User {
  id: number
  name: string
  email: string
  role: 'direktur' | 'admin' | 'engineer' | 'klien' | 'audit'
  phone_wa: string
  status: string
  last_login_at?: string
  avatar_url?: string
}

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('meldir_token'))
  const user = ref<User | null>(
    localStorage.getItem('meldir_user')
      ? JSON.parse(localStorage.getItem('meldir_user')!)
      : null
  )
  const isLoading = ref(false)
  const errorMessage = ref<string | null>(null)

  const isAuthenticated = computed(() => !!token.value && !!user.value)
  const role = computed(() => user.value?.role)

  async function login(email: string, password: string): Promise<boolean> {
    isLoading.value = true
    errorMessage.value = null

    try {
      const response = await fetch('/api/v1/auth/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      })

      const result = await response.json()

      if (!response.ok || !result.success) {
        throw new Error(result.error || 'Login gagal, periksa email dan password')
      }

      token.value = result.data.token
      user.value = result.data.user

      localStorage.setItem('meldir_token', result.data.token)
      localStorage.setItem('meldir_user', JSON.stringify(result.data.user))

      return true
    } catch (err: any) {
      errorMessage.value = err.message || 'Terjadi kesalahan jaringan'
      return false
    } finally {
      isLoading.value = false
    }
  }

  async function logout() {
    if (token.value) {
      try {
        await fetch('/api/v1/auth/logout', {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${token.value}`,
          },
        })
      } catch (e) {
        // Abaikan error jaringan saat logout
      }
    }

    token.value = null
    user.value = null
    localStorage.removeItem('meldir_token')
    localStorage.removeItem('meldir_user')
  }

  async function fetchProfile() {
    if (!token.value) return

    try {
      const res = await fetch('/api/v1/user/profile', {
        headers: {
          Authorization: `Bearer ${token.value}`,
        },
      })
      if (res.ok) {
        const result = await res.json()
        if (result.success && result.data) {
          user.value = result.data
          localStorage.setItem('meldir_user', JSON.stringify(result.data))
        }
      } else if (res.status === 401) {
        await logout()
      }
    } catch (e) {
      // Abaikan error background fetch
    }
  }

  return {
    token,
    user,
    role,
    isAuthenticated,
    isLoading,
    errorMessage,
    login,
    logout,
    fetchProfile,
  }
})
