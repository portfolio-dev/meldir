<template>
  <div class="login-card glass-panel">
    <div class="login-header">
      <span class="portal-badge" :class="portalBadgeClass">{{ portalName }}</span>
      <h2>Masuk ke Akun Anda</h2>
      <p class="subtitle">Platform Ekosistem Digital PT. Melayani Digital Raya</p>
    </div>

    <form @submit.prevent="handleLogin" class="login-form">
      <!-- Error Message -->
      <div v-if="authStore.errorMessage" class="error-banner">
        ⚠️ {{ authStore.errorMessage }}
      </div>

      <div class="form-group">
        <label for="email">Alamat Email Resmi</label>
        <input
          id="email"
          v-model="email"
          type="email"
          placeholder="nama@meldir.id atau email Anda"
          required
          class="form-input"
        />
      </div>

      <div class="form-group">
        <label for="password">Kata Sandi (Password)</label>
        <input
          id="password"
          v-model="password"
          type="password"
          placeholder="••••••••••••"
          required
          class="form-input"
        />
      </div>

      <button type="submit" :disabled="authStore.isLoading" class="btn-submit">
        <span v-if="authStore.isLoading">Memverifikasi Sesi...</span>
        <span v-else>Masuk ke Dashboard →</span>
      </button>

      <!-- Quick Fill Helper for Demo / Testing -->
      <div class="quick-fill-box">
        <button type="button" @click="fillDefaultSuperadmin" class="btn-quick-fill">
          ⚡ Isi Cepat Akun Direktur Utama (Default)
        </button>
      </div>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/authStore'

const props = defineProps<{
  portalName: string
  portalBadgeClass?: string
}>()

const authStore = useAuthStore()
const email = ref('')
const password = ref('')

async function handleLogin() {
  await authStore.login(email.value, password.value)
}

function fillDefaultSuperadmin() {
  email.value = 'direktur@meldir.id'
  password.value = 'MeldirAdmin2026!'
}
</script>

<style scoped>
.login-card {
  max-width: 440px;
  width: 100%;
  margin: 40px auto;
  padding: 36px 32px;
}
.login-header {
  text-align: center;
  margin-bottom: 28px;
}
.portal-badge {
  font-size: 0.7rem;
  font-weight: 800;
  padding: 3px 10px;
  border-radius: 9999px;
  display: inline-block;
  margin-bottom: 12px;
  letter-spacing: 0.05em;
  background: #0284c7;
  color: #fff;
}
.portal-badge.office {
  background: #0284c7;
}
.portal-badge.jobs {
  background: #6366f1;
}
.portal-badge.portal {
  background: #10b981;
}
.login-header h2 {
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--text-main);
  margin-bottom: 6px;
}
.subtitle {
  font-size: 0.85rem;
  color: var(--text-muted);
}
.login-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}
.error-banner {
  background: rgba(244, 63, 94, 0.15);
  border: 1px solid rgba(244, 63, 94, 0.4);
  color: #fb7185;
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 0.85rem;
  line-height: 1.4;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
  text-align: left;
}
.form-group label {
  font-size: 0.8rem;
  font-weight: 600;
  color: #cbd5e1;
}
.form-input {
  background: rgba(15, 23, 42, 0.6);
  border: 1px solid var(--border-color);
  color: var(--text-main);
  padding: 12px 14px;
  border-radius: 8px;
  font-family: inherit;
  font-size: 0.9rem;
  transition: border-color 0.2s, box-shadow 0.2s;
}
.form-input:focus {
  outline: none;
  border-color: #38bdf8;
  box-shadow: 0 0 0 3px rgba(56, 189, 248, 0.15);
}
.btn-submit {
  background: linear-gradient(135deg, #0284c7, #2563eb);
  color: #fff;
  border: none;
  padding: 12px;
  border-radius: 8px;
  font-weight: 700;
  font-size: 0.95rem;
  cursor: pointer;
  transition: opacity 0.2s, transform 0.1s;
  margin-top: 6px;
}
.btn-submit:hover:not(:disabled) {
  opacity: 0.95;
  transform: translateY(-1px);
}
.btn-submit:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}
.quick-fill-box {
  margin-top: 10px;
  text-align: center;
}
.btn-quick-fill {
  background: rgba(255, 255, 255, 0.05);
  border: 1px dashed rgba(255, 255, 255, 0.2);
  color: #94a3b8;
  padding: 8px 12px;
  border-radius: 6px;
  font-size: 0.75rem;
  cursor: pointer;
  transition: all 0.2s;
  width: 100%;
}
.btn-quick-fill:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #38bdf8;
  border-color: #38bdf8;
}
</style>
