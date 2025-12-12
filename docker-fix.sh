#!/bin/bash

# Script untuk fix Docker I/O error
# Error: input/output error saat pull image

echo "🔧 Docker I/O Error Fix Script"
echo "==============================="
echo ""

echo "1. Cleaning Docker system..."
docker system prune -f

echo ""
echo "2. Cleaning build cache..."
docker builder prune -f

echo ""
echo "3. Checking Docker disk usage..."
docker system df

echo ""
echo "4. Restarting Docker Desktop..."
echo "   ⚠️  Silakan restart Docker Desktop secara manual:"
echo "   - Klik icon Docker di menu bar"
echo "   - Pilih 'Quit Docker Desktop'"
echo "   - Buka Docker Desktop lagi"
echo ""

echo "5. Alternative: Reset Docker Desktop (jika masih error)"
echo "   - Docker Desktop > Settings > Troubleshoot > Clean / Purge data"
echo ""

echo "✅ Cleanup selesai!"
echo ""
echo "💡 Tips:"
echo "   - Pastikan disk space cukup (minimal 5GB free)"
echo "   - Coba pull image lagi: docker pull golang:1.21-alpine"
echo "   - Atau langsung build: docker-compose build"

