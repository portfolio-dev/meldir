const CACHE_NAME = 'meldir-v2026.10.10.3';
const STATIC_ASSETS = [
  './',
  'index.html',
  'service/',
  'support/',
  'privacy/',
  'terms/',
  'refund/',
  'style.css',
  'app.js',
  'manifest.json',
  'favicon.ico',
  'images/hero-bg.jpg',
  'images/service-hero-bg.jpg',
  'icons/icon-192.png',
  'icons/icon-512.png'
];

// Install Event - Pre-cache Static Assets & Segera Ambil Alih
self.addEventListener('install', (e) => {
  self.skipWaiting();
  e.waitUntil(
    caches.open(CACHE_NAME).then((cache) => {
      console.log('[Service Worker] Pre-caching static assets for', CACHE_NAME);
      return cache.addAll(STATIC_ASSETS);
    })
  );
});

// Activate Event - Hapus Semua Cache Lama Agar Bersih
self.addEventListener('activate', (e) => {
  e.waitUntil(
    caches.keys().then((keys) => {
      return Promise.all(
        keys.map((key) => {
          if (key !== CACHE_NAME) {
            console.log('[Service Worker] Purging outdated cache:', key);
            return caches.delete(key);
          }
        })
      );
    }).then(() => self.clients.claim())
  );
});

// Fetch Event - Smart Strategy:
// 1. Backend API: Network Only (tidak pernah di-cache)
// 2. Navigasi HTML & Kode Inti (CSS & JS): NETWORK-FIRST (selalu update terbaru saat online, cache hanya saat offline)
// 3. Media Gambar & Icon: CACHE-FIRST (ringan dan hemat bandwidth)
self.addEventListener('fetch', (e) => {
  const url = new URL(e.request.url);

  // Jangan sentuh API requests atau non-GET
  if (url.pathname.startsWith('/api') || e.request.method !== 'GET') {
    return;
  }

  // A. Navigasi / Dokumen HTML: NETWORK-FIRST agar setiap update langsung tampil
  if (e.request.mode === 'navigate' || (e.request.headers.get('accept') && e.request.headers.get('accept').includes('text/html'))) {
    e.respondWith(
      fetch(e.request)
        .then((networkResponse) => {
          if (networkResponse && networkResponse.status === 200) {
            const copy = networkResponse.clone();
            caches.open(CACHE_NAME).then((cache) => cache.put(e.request, copy));
          }
          return networkResponse;
        })
        .catch(() => {
          // Fallback offline
          return caches.match(e.request).then((cached) => cached || caches.match('index.html'));
        })
    );
    return;
  }

  // B. Style & Script (CSS & JS): NETWORK-FIRST agar perubahan UI & logika langsung aktif
  if (url.pathname.endsWith('.css') || url.pathname.endsWith('.js')) {
    e.respondWith(
      fetch(e.request)
        .then((networkResponse) => {
          if (networkResponse && networkResponse.status === 200) {
            const copy = networkResponse.clone();
            caches.open(CACHE_NAME).then((cache) => cache.put(e.request, copy));
          }
          return networkResponse;
        })
        .catch(() => caches.match(e.request))
    );
    return;
  }

  // C. Media Statis (Gambar, Icon, Font): CACHE-FIRST, fallback ke Network
  e.respondWith(
    caches.match(e.request).then((cachedResponse) => {
      if (cachedResponse) {
        return cachedResponse;
      }
      return fetch(e.request).then((networkResponse) => {
        if (networkResponse && networkResponse.status === 200 && networkResponse.type === 'basic') {
          const copy = networkResponse.clone();
          caches.open(CACHE_NAME).then((cache) => cache.put(e.request, copy));
        }
        return networkResponse;
      }).catch(() => {
        // Abaikan jika media offline gagal
      });
    })
  );
});
