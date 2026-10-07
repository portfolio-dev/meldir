<template>
  <div class="login-card glass-panel" :class="portalBadgeClass">
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
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '../stores/authStore'

defineProps<{
  portalName: string
  portalBadgeClass?: string
}>()

const authStore = useAuthStore()
const email = ref('')
const password = ref('')

async function handleLogin() {
  await authStore.login(email.value, password.value)
}
</script>

<style scoped>
.login-card {
  max-width: 440px;
  width: 100%;
  margin: 32px auto;
  padding: 36px 32px;
  background: #ffffff;
  border: 1px solid #cbd5e1;
  border-radius: 14px;
  box-shadow: 0 10px 25px -5px rgba(15, 23, 42, 0.08), 0 8px 10px -6px rgba(15, 23, 42, 0.04);
}
.login-card.office { border-top: 4px solid #0284c7; }
.login-card.jobs { border-top: 4px solid #4f46e5; }
.login-card.portal { border-top: 4px solid #059669; }

.login-header {
  text-align: center;
  margin-bottom: 24px;
}
.portal-badge {
  font-size: 0.72rem;
  font-weight: 700;
  padding: 4px 12px;
  border-radius: 9999px;
  display: inline-block;
  margin-bottom: 12px;
  letter-spacing: 0.06em;
  background: #0284c7;
  color: #ffffff;
}
.portal-badge.office {
  background: #0284c7;
}
.portal-badge.jobs {
  background: #4f46e5;
}
.portal-badge.portal {
  background: #059669;
}
.login-header h2 {
  font-size: 1.5rem;
  font-weight: 800;
  color: #0f172a;
  margin-bottom: 6px;
}
.subtitle {
  font-size: 0.85rem;
  color: #64748b;
}
.login-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}
.error-banner {
  background: #fef2f2;
  border: 1px solid #fecaca;
  color: #b91c1c;
  padding: 10px 14px;
  border-radius: 8px;
  font-size: 0.85rem;
  line-height: 1.4;
  font-weight: 500;
}
.form-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
  text-align: left;
}
.form-group label {
  font-size: 0.82rem;
  font-weight: 600;
  color: #334155;
}
.form-input {
  background: #f8fafc;
  border: 1px solid #cbd5e1;
  color: #0f172a;
  padding: 11px 14px;
  border-radius: 8px;
  font-family: inherit;
  font-size: 0.9rem;
  transition: border-color 0.2s, box-shadow 0.2s, background-color 0.2s;
}
.form-input:focus {
  outline: none;
  background: #ffffff;
  border-color: #0284c7;
  box-shadow: 0 0 0 3px rgba(2, 132, 199, 0.12);
}
.btn-submit {
  background: #0284c7;
  color: #ffffff;
  border: none;
  padding: 12px;
  border-radius: 8px;
  font-weight: 700;
  font-size: 0.95rem;
  cursor: pointer;
  transition: background 0.2s, transform 0.1s;
  margin-top: 6px;
}
.btn-submit:hover:not(:disabled) {
  background: #0369a1;
  transform: translateY(-1px);
}
.btn-submit:disabled {
  opacity: 0.65;
  cursor: not-allowed;
}

@media (max-width: 480px) {
  .login-card {
    padding: 24px 18px;
    margin: 16px auto;
    border-radius: 12px;
    max-width: 100%;
  }
  .login-header h2 {
    font-size: 1.25rem;
  }
  .subtitle {
    font-size: 0.8rem;
  }
}
</style>
