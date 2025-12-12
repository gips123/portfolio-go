# 🚀 Panduan Deploy Backend Go ke Render (Free Tier)

Panduan lengkap step-by-step untuk deploy backend Go ke Render secara gratis.

---

## 📋 Prerequisites

- ✅ Akun GitHub (untuk connect repo)
- ✅ Code sudah di-push ke GitHub
- ✅ Database sudah setup (Supabase/Railway/atau Render PostgreSQL)

---

## 🎯 Step 1: Siapkan Database

### Opsi A: Gunakan Supabase (Recommended - Gratis)

1. **Daftar di Supabase**: https://supabase.com
2. **Create New Project**:
   - Project name: `portfolio-db`
   - Database password: (simpan dengan aman!)
   - Region: pilih yang terdekat (Singapore recommended)
3. **Dapatkan Connection String**:
   - Settings → Database → Connection String
   - Copy "URI" connection string
   - Format: `postgresql://postgres:[PASSWORD]@[HOST]:5432/postgres`
4. **Setup Schema**:
   - SQL Editor → Jalankan migration `20251212150000_create_tables.sql`
   - SQL Editor → Jalankan seed data `20251212152735_new-migration.sql`

### Opsi B: Gunakan Render PostgreSQL (Free 90 hari)

1. **Render Dashboard** → New → PostgreSQL
2. **Settings**:
   - Name: `portfolio-db`
   - Database: `portfolio`
   - User: `postgres`
   - Region: Singapore (atau terdekat)
   - Plan: Free (90 hari gratis)
3. **Copy Connection String** dari Internal Database URL
4. **Setup Schema** (sama seperti Supabase)

---

## 🎯 Step 2: Push Code ke GitHub

Pastikan semua code sudah di-push:

```bash
cd /Users/ghifaryahmada/best-portfolio-go

# Check status
git status

# Add semua file
git add .

# Commit
git commit -m "Prepare for Render deployment"

# Push ke GitHub
git push origin main
```

---

## 🎯 Step 3: Deploy di Render

### 3.1. Daftar & Login

1. **Daftar di Render**: https://render.com
2. **Sign up dengan GitHub** (recommended)
3. **Authorize Render** untuk akses GitHub repos

### 3.2. Create New Web Service

1. **Dashboard** → **New +** → **Web Service**
2. **Connect Repository**:
   - Pilih repo: `best-portfolio-go`
   - Klik **Connect**

### 3.3. Configure Service

**Basic Settings**:
- **Name**: `portfolio-api` (atau nama lain)
- **Region**: Singapore (atau terdekat)
- **Branch**: `main` (atau branch yang digunakan)
- **Root Directory**: (kosongkan, atau `./` jika ada subfolder)

**Build & Deploy**:
- **Environment**: `Docker` (atau `Go`)
- **Build Command**: 
  ```
  go build -o portfolio-api
  ```
- **Start Command**: 
  ```
  ./portfolio-api
  ```

**Jika menggunakan Go (bukan Docker)**:
- **Build Command**: `go build -o portfolio-api`
- **Start Command**: `./portfolio-api`

**Jika menggunakan Docker**:
- Render akan auto-detect Dockerfile
- Pastikan Dockerfile sudah ada di root

### 3.4. Environment Variables

Klik **Environment** tab, tambahkan:

```
DB_HOST=[database-host]
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=[database-password]
DB_NAME=postgres
DB_SSLMODE=require
PORT=8080
CORS_ORIGIN=https://your-portfolio.vercel.app,https://your-dashboard.vercel.app
JWT_SECRET=[generate-random-32-chars]
JWT_TOKEN_TTL_MINUTES=720
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=[your-secure-password]
```

**Cara mendapatkan nilai**:

1. **DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME**:
   - Dari Supabase: Settings → Database → Connection String
   - Parse connection string: `postgresql://postgres:PASSWORD@HOST:5432/postgres`
   - DB_HOST = bagian setelah `@` dan sebelum `:`
   - DB_PASSWORD = bagian setelah `postgresql://postgres:` dan sebelum `@`

2. **JWT_SECRET**:
   ```bash
   # Generate random secret
   openssl rand -base64 32
   ```
   Atau online: https://generate-secret.vercel.app/32

3. **CORS_ORIGIN**:
   - Tambahkan domain frontend yang akan mengakses API
   - Format: `https://domain1.com,https://domain2.com`
   - Untuk development: `http://localhost:3000,http://localhost:3001`

### 3.5. Advanced Settings (Optional)

- **Auto-Deploy**: `Yes` (deploy otomatis saat push ke GitHub)
- **Health Check Path**: `/health` (jika ada)
- **Dockerfile Path**: `Dockerfile` (jika menggunakan Docker)

### 3.6. Deploy

1. Klik **Create Web Service**
2. Render akan mulai build dan deploy
3. Tunggu sampai status **Live** (biasanya 2-5 menit)

---

## 🎯 Step 4: Setup Custom Domain (Optional)

1. **Settings** → **Custom Domains**
2. **Add Custom Domain**:
   - Masukkan domain: `api.yourdomain.com`
   - Follow instructions untuk setup DNS
3. **SSL Certificate**: Auto-generated oleh Render

---

## 🎯 Step 5: Verifikasi Deploy

### 5.1. Check Logs

1. **Logs** tab di Render dashboard
2. Pastikan tidak ada error
3. Cek apakah server sudah start: `Server starting on port 8080`

### 5.2. Test API

```bash
# Test health endpoint
curl https://your-service.onrender.com/health

# Test API endpoint
curl https://your-service.onrender.com/api/projects
```

### 5.3. Update Frontend

Update environment variables di frontend:

**Vercel/Railway Frontend**:
```
API_BASE_URL=https://your-service.onrender.com
NEXT_PUBLIC_API_BASE=https://your-service.onrender.com
```

---

## 🔧 Troubleshooting

### Error: Build Failed

**Masalah**: Build command error
**Solusi**:
- Pastikan `go.mod` dan `go.sum` sudah commit
- Check build logs untuk detail error
- Pastikan Go version compatible (1.21+)

### Error: Port Already in Use

**Masalah**: Port conflict
**Solusi**:
- Render menggunakan port dari `PORT` env var
- Pastikan `PORT=8080` di environment variables
- Atau gunakan `PORT=10000` (Render default)

### Error: Database Connection Failed

**Masalah**: Tidak bisa connect ke database
**Solusi**:
- Pastikan `DB_SSLMODE=require` untuk production
- Check database firewall/whitelist
- Supabase: Settings → Database → Connection Pooling (gunakan connection pool URL)
- Pastikan semua env var sudah benar

### Error: CORS Error

**Masalah**: Frontend tidak bisa akses API
**Solusi**:
- Pastikan domain frontend ada di `CORS_ORIGIN`
- Format: `https://domain1.com,https://domain2.com` (tanpa spasi)
- Restart service setelah update env var

### Service Sleep (Free Tier)

**Masalah**: Service sleep setelah 15 menit idle
**Solusi**:
- Free tier akan sleep setelah 15 menit tidak ada request
- Request pertama setelah sleep akan lambat (cold start ~30 detik)
- Gunakan UptimeRobot (gratis) untuk ping setiap 5 menit:
  - Setup: https://uptimerobot.com
  - Monitor URL: `https://your-service.onrender.com/health`
  - Interval: 5 minutes

---

## 📊 Monitoring & Maintenance

### 1. Setup Uptime Monitoring

1. **Daftar di UptimeRobot**: https://uptimerobot.com (gratis)
2. **Add New Monitor**:
   - Monitor Type: HTTP(s)
   - Friendly Name: Portfolio API
   - URL: `https://your-service.onrender.com/health`
   - Monitoring Interval: 5 minutes
3. **Save** → Service akan selalu "awake"

### 2. View Logs

- **Render Dashboard** → Service → **Logs** tab
- Real-time logs untuk debugging
- Download logs untuk analisis

### 3. Update Environment Variables

- **Settings** → **Environment**
- Edit/Add variables
- **Save Changes** → Auto-redeploy

### 4. Manual Deploy

- **Manual Deploy** → **Deploy latest commit**
- Atau push ke GitHub untuk auto-deploy

---

## 🆓 Free Tier Limits

- ✅ **750 hours/month** (cukup untuk 1 service 24/7)
- ✅ **512 MB RAM**
- ✅ **0.1 CPU**
- ⚠️ **Sleep setelah 15 menit idle** (gunakan UptimeRobot)
- ⚠️ **Cold start ~30 detik** setelah sleep

---

## 💡 Tips Optimasi

1. **Gunakan Health Check**:
   - Setup endpoint `/health` untuk monitoring
   - UptimeRobot akan ping endpoint ini

2. **Optimize Build Time**:
   - Gunakan `.dockerignore` untuk exclude file tidak perlu
   - Cache dependencies jika mungkin

3. **Database Connection Pooling**:
   - Supabase: Gunakan Connection Pooling URL
   - Render PostgreSQL: Gunakan Internal Database URL

4. **Environment Variables**:
   - Jangan commit `.env` file
   - Gunakan Render Environment Variables

5. **Logs**:
   - Monitor logs secara berkala
   - Setup alerts jika ada error

---

## 📝 Checklist Deploy

- [ ] Database sudah setup (Supabase/Render)
- [ ] Database schema sudah di-migrate
- [ ] Seed data sudah di-insert
- [ ] Code sudah di-push ke GitHub
- [ ] Render account sudah dibuat
- [ ] Web Service sudah dibuat
- [ ] Environment variables sudah di-set
- [ ] Build berhasil
- [ ] Service status: Live
- [ ] Health endpoint bisa diakses
- [ ] API endpoints bisa diakses
- [ ] CORS sudah dikonfigurasi
- [ ] Frontend sudah di-update dengan API URL
- [ ] UptimeRobot sudah di-setup (optional)

---

## 🔗 Quick Links

- **Render Dashboard**: https://dashboard.render.com
- **Render Docs**: https://render.com/docs
- **Supabase Dashboard**: https://supabase.com/dashboard
- **UptimeRobot**: https://uptimerobot.com

---

## 🎉 Selesai!

Backend API sudah live di Render! 

**URL API**: `https://your-service.onrender.com`

**Next Steps**:
1. Update frontend dengan API URL baru
2. Test semua endpoints
3. Setup monitoring dengan UptimeRobot
4. Update CORS jika perlu


