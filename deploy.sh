#!/bin/bash
set -e

# ==============================================================================
# MELDIR CI/CD PRODUCTION DEPLOYMENT SCRIPT
# Server: Hostinger VPS - AlmaLinux 9 + CyberPanel + OpenLiteSpeed
# ==============================================================================

REPO_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
BACKEND_BIN_DEST="/home/meldir.id/backend/main"
STORAGE_DIR="/home/meldir.id/storage"
PUBLIC_HTML_MAIN="/home/meldir.id/public_html"
PUBLIC_HTML_OFFICE="/home/office.meldir.id/public_html"
PUBLIC_HTML_JOBS="/home/jobs.meldir.id/public_html"
PUBLIC_HTML_PORTAL="/home/portal.meldir.id/public_html"

echo "=================================================="
echo "🚀 [1/5] Pulling Latest Commits from GitHub (main)"
echo "=================================================="
cd "$REPO_DIR"
git fetch origin main
git reset --hard origin/main

echo "=================================================="
echo "⚙️ [2/5] Building Golang Backend Service"
echo "=================================================="
mkdir -p /home/meldir.id/backend
cd "$REPO_DIR/backend"

# Ensure dependencies are clean & compile production binary
go mod tidy
go build -ldflags="-s -w" -o "$BACKEND_BIN_DEST" ./cmd/api/main.go
chmod +x "$BACKEND_BIN_DEST"
chown nobody:nobody "$BACKEND_BIN_DEST" 2>/dev/null || true

echo "🔄 Restarting meldir-backend.service..."
systemctl restart meldir-backend
sleep 2

if systemctl is-active --quiet meldir-backend; then
    echo "✅ Backend service is ACTIVE and running on :8080"
else
    echo "❌ ERROR: Backend service failed to start! Check logs:"
    journalctl -u meldir-backend -n 30 --no-pager
    exit 1
fi

echo "=================================================="
echo "🎨 [3/5] Building Vue 3 Frontend (Vite)"
echo "=================================================="
cd "$REPO_DIR/frontend"
npm install --prefer-offline --no-audit
npm run build

echo "=================================================="
echo "📦 [4/5] Deploying Web Roots & Static Assets"
echo "=================================================="
# 1. Main Landing Website (meldir.id)
mkdir -p "$PUBLIC_HTML_MAIN"
rsync -av --delete \
    --exclude 'backend' \
    --exclude 'frontend' \
    --exclude 'docs' \
    --exclude '.git' \
    --exclude '.github' \
    --exclude 'deploy.sh' \
    "$REPO_DIR/" "$PUBLIC_HTML_MAIN/"

# 2. Subdomain Multi-Portal Webapps (office, jobs, portal)
for SUB_DIR in "$PUBLIC_HTML_OFFICE" "$PUBLIC_HTML_JOBS" "$PUBLIC_HTML_PORTAL"; do
    mkdir -p "$SUB_DIR"
    rsync -av --delete "$REPO_DIR/frontend/dist/" "$SUB_DIR/"
done

# Pastikan kepemilikan file & hak akses OpenLiteSpeed (nobody:nobody & 755)
chown -R nobody:nobody "$PUBLIC_HTML_MAIN" "$PUBLIC_HTML_OFFICE" "$PUBLIC_HTML_JOBS" "$PUBLIC_HTML_PORTAL" "$STORAGE_DIR" 2>/dev/null || true
chmod -R 755 "$PUBLIC_HTML_MAIN" "$PUBLIC_HTML_OFFICE" "$PUBLIC_HTML_JOBS" "$PUBLIC_HTML_PORTAL" 2>/dev/null || true

echo "=================================================="
echo "🔄 [5/5] Purging Server Cache & Reloading Web Server"
echo "=================================================="
# Bersihkan cache internal OpenLiteSpeed & CyberPanel agar tampilan teranyar langsung aktif
echo "🧹 Membersihkan direktori cache OpenLiteSpeed..."
rm -rf /tmp/lshttpd/swap/* 2>/dev/null || true
rm -rf /tmp/lshttpd/bak_swap/* 2>/dev/null || true
rm -rf /usr/local/lsws/cachedata/* 2>/dev/null || true

if [ -f "/usr/local/lsws/bin/lswsctrl" ]; then
    /usr/local/lsws/bin/lswsctrl restart || true
else
    systemctl restart lsws || true
fi

echo "=================================================="
echo "🎉 DEPLOYMENT SELESAI DENGAN SUKSES! SISTEM LIVE."
echo "=================================================="
