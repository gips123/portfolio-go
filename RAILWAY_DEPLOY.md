# 🚂 Panduan Deploy Backend Go ke Railway

Panduan lengkap step-by-step untuk deploy backend Go ke Railway.

---

## 📋 Prerequisites

- ✅ Akun GitHub (untuk connect repo)
- ✅ Code sudah di-push ke GitHub
- ✅ Database sudah setup (Supabase/Railway PostgreSQL)

---

## 🎯 Step 1: Push Code ke GitHub

Pastikan semua code sudah di-push:

```bash
cd /Users/ghifaryahmada/best-portfolio-go

# Check status
git status

# Add semua file
git add .

# Commit
git commit -m "Prepare for Railway deployment"

# Push ke GitHub
git push origin main
```

**Pastikan file ini sudah di-commit:**
- ✅ `railway.json` (konfigurasi Railway)
- ✅ `go.mod` dan `go.sum`
- ✅ Semua source code
- ✅ `Dockerfile` (jika menggunakan Docker)

---

## 🎯 Step 2: Daftar & Login Railway

1. **Daftar di Railway**: https://railway.app
2. **Sign up dengan GitHub** (recommended)
3. **Authorize Railway** untuk akses GitHub repos

---

## 🎯 Step 3: Create New Project

1. **Dashboard** → **New Project**
2. **Deploy from GitHub repo**
3. **Pilih repository**: `best-portfolio-go`
4. Railway akan auto-detect Go project

---

## 🎯 Step 4: Configure Service

### 4.1. Basic Settings

Railway akan auto-detect:
- **Language**: Go
- **Build Command**: Auto-detect dari `railway.json` atau `go build`
- **Start Command**: Auto-detect dari `railway.json` atau binary name

**Jika perlu manual configure:**
- **Build Command**: `go build -o portfolio-api`
- **Start Command**: `./portfolio-api`

### 4.2. Environment Variables

Klik **Variables** tab, tambahkan:

```
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require
PORT=8080
CORS_ORIGIN=https://your-portfolio.vercel.app,https://your-dashboard.vercel.app
JWT_SECRET=[generate-random-32-chars]
JWT_TOKEN_TTL_MINUTES=720
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=[your-secure-password]
```

**Cara generate JWT_SECRET:**
```bash
openssl rand -base64 32
```

### 4.3. Generate Public Domain

1. **Settings** → **Generate Domain**
2. Railway akan generate: `https://your-service.up.railway.app`
3. **Copy domain** ini untuk setup frontend

---

## 🎯 Step 5: Deploy

1. Railway akan **auto-deploy** setelah connect GitHub
2. **Monitor logs** di **Deployments** tab
3. Tunggu sampai status **Deployed** (biasanya 2-5 menit)

---

## 🎯 Step 6: Verifikasi Deploy

### 6.1. Check Logs

1. **Deployments** tab → Klik deployment terbaru
2. **View Logs** → Pastikan tidak ada error
3. Cek apakah server sudah start: `Server starting on port 8080`

### 6.2. Test API

```bash
# Test health endpoint
curl https://your-service.up.railway.app/health

# Test API endpoint
curl https://your-service.up.railway.app/api/projects
```

### 6.3. Update Frontend

Update environment variables di frontend (Vercel/Railway):

```
API_BASE_URL=https://your-service.up.railway.app
NEXT_PUBLIC_API_BASE=https://your-service.up.railway.app
```

---

## 🔧 Railway Configuration

### File `railway.json`

Railway akan menggunakan konfigurasi dari `railway.json` jika ada:

```json
{
  "$schema": "https://railway.app/railway.schema.json",
  "build": {
    "builder": "NIXPACKS",
    "buildCommand": "go build -o portfolio-api"
  },
  "deploy": {
    "startCommand": "./portfolio-api",
    "restartPolicyType": "ON_FAILURE",
    "restartPolicyMaxRetries": 10
  }
}
```

### Auto-Deploy

Railway akan **auto-deploy** setiap kali:
- Push ke branch yang di-monitor (default: `main`)
- Manual trigger dari dashboard

---

## 🔐 Environment Variables untuk Railway

Setelah deploy, pastikan semua env vars sudah di-set:

### Database (Supabase)
```
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require
```

### Server
```
PORT=8080
```

### CORS
```
CORS_ORIGIN=https://your-portfolio.vercel.app,https://your-dashboard.vercel.app
```

### JWT & Auth
```
JWT_SECRET=[random-32-chars]
JWT_TOKEN_TTL_MINUTES=720
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=[secure-password]
```

---

## 🆓 Free Tier Limits

- ✅ **$5 credit gratis/bulan**
- ✅ **512 MB RAM**
- ✅ **Auto-sleep** setelah 5 menit idle (free tier)
- ⚠️ **Bisa habis** jika traffic tinggi

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
- Railway menggunakan port dari `PORT` env var
- Pastikan `PORT=8080` di environment variables
- Railway akan auto-assign port jika tidak di-set

### Error: Database Connection Failed

**Masalah**: Tidak bisa connect ke database
**Solusi**:
- Pastikan `DB_SSLMODE=require` untuk Supabase
- Check database firewall/whitelist
- Pastikan semua env var sudah benar
- Test connection string dengan psql dulu

### Error: CORS Error

**Masalah**: Frontend tidak bisa akses API
**Solusi**:
- Pastikan domain frontend ada di `CORS_ORIGIN`
- Format: `https://domain1.com,https://domain2.com` (tanpa spasi)
- Redeploy setelah update env var

### Service Sleep (Free Tier)

**Masalah**: Service sleep setelah 5 menit idle
**Solusi**:
- Free tier akan sleep setelah idle
- Request pertama setelah sleep akan lambat (cold start)
- Gunakan UptimeRobot (gratis) untuk ping setiap 5 menit:
  - Setup: https://uptimerobot.com
  - Monitor URL: `https://your-service.up.railway.app/health`
  - Interval: 5 minutes

---

## 📊 Monitoring & Maintenance

### 1. View Logs

- **Deployments** tab → Klik deployment → **View Logs**
- Real-time logs untuk debugging
- Download logs untuk analisis

### 2. Update Environment Variables

- **Variables** tab → Edit/Add variables
- **Save Changes** → Auto-redeploy

### 3. Manual Deploy

- **Deployments** tab → **Redeploy**
- Atau push ke GitHub untuk auto-deploy

### 4. Setup Monitoring

- **UptimeRobot**: https://uptimerobot.com
- Monitor: `https://your-service.up.railway.app/health`
- Interval: 5 minutes (untuk prevent sleep)

---

## 📝 Checklist Deploy

- [ ] Code sudah di-push ke GitHub
- [ ] Railway account sudah dibuat
- [ ] Project sudah dibuat
- [ ] GitHub repo sudah terhubung
- [ ] Environment variables sudah di-set
- [ ] Public domain sudah di-generate
- [ ] Build berhasil
- [ ] Deploy berhasil
- [ ] Health endpoint bisa diakses
- [ ] API endpoints bisa diakses
- [ ] CORS sudah dikonfigurasi
- [ ] Frontend sudah di-update dengan API URL
- [ ] UptimeRobot sudah di-setup (optional)

---

## 🔗 Quick Links

- **Railway Dashboard**: https://railway.app/dashboard
- **Railway Docs**: https://docs.railway.app
- **Supabase Dashboard**: https://supabase.com/dashboard
- **UptimeRobot**: https://uptimerobot.com

---

## 🎉 Selesai!

Backend API sudah live di Railway! 

**URL API**: `https://your-service.up.railway.app`

**Next Steps**:
1. Update frontend dengan API URL baru
2. Test semua endpoints
3. Setup monitoring dengan UptimeRobot
4. Update CORS jika perlu


