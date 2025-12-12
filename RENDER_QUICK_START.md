# ⚡ Quick Start - Deploy ke Render (5 Menit)

Panduan cepat deploy backend Go ke Render.

---

## 🚀 Step-by-Step (5 Menit)

### **1. Database Setup (2 menit)**

**Supabase** (Recommended):
1. https://supabase.com → New Project
2. Settings → Database → Copy Connection String
3. SQL Editor → Jalankan migration & seed

**Atau Render PostgreSQL**:
1. Render → New → PostgreSQL
2. Copy Internal Database URL

---

### **2. Deploy di Render (3 menit)**

1. **Render Dashboard**: https://render.com
2. **New +** → **Web Service**
3. **Connect GitHub** → Pilih repo `best-portfolio-go`
4. **Configure**:
   - Name: `portfolio-api`
   - Region: Singapore
   - Environment: `Docker` (atau `Go`)
   - Build Command: `go build -o portfolio-api`
   - Start Command: `./portfolio-api`

5. **Environment Variables**:
   ```
   DB_HOST=[dari-supabase]
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=[dari-supabase]
   DB_NAME=postgres
   DB_SSLMODE=require
   PORT=8080
   CORS_ORIGIN=https://*.vercel.app
   JWT_SECRET=[random-32-chars]
   JWT_TOKEN_TTL_MINUTES=720
   ADMIN_USERNAME=ghifary
   ADMIN_PASSWORD=[secure-password]
   ```

6. **Create Web Service** → Tunggu deploy selesai

---

### **3. Generate JWT Secret**

```bash
openssl rand -base64 32
```

Atau: https://generate-secret.vercel.app/32

---

### **4. Test API**

```bash
# Health check
curl https://your-service.onrender.com/health

# Test endpoint
curl https://your-service.onrender.com/api/projects
```

---

### **5. Update Frontend**

**Vercel/Railway**:
```
API_BASE_URL=https://your-service.onrender.com
NEXT_PUBLIC_API_BASE=https://your-service.onrender.com
```

---

## ⚠️ Important Notes

1. **Service Sleep**: Free tier sleep setelah 15 menit idle
   - **Solusi**: Setup UptimeRobot untuk ping setiap 5 menit
   - URL: https://uptimerobot.com

2. **Cold Start**: Request pertama setelah sleep akan lambat (~30 detik)
   - Normal untuk free tier

3. **Database SSL**: Pastikan `DB_SSLMODE=require` untuk production

4. **CORS**: Update `CORS_ORIGIN` dengan domain frontend yang sebenarnya

---

## 🔗 Quick Links

- **Render Dashboard**: https://dashboard.render.com
- **Supabase Dashboard**: https://supabase.com/dashboard
- **UptimeRobot**: https://uptimerobot.com

---

## ✅ Checklist

- [ ] Database setup & schema migrated
- [ ] Code pushed ke GitHub
- [ ] Render service created
- [ ] Environment variables set
- [ ] Service status: Live
- [ ] API test berhasil
- [ ] Frontend updated
- [ ] UptimeRobot setup (optional)

---

**Done!** 🎉

API URL: `https://your-service.onrender.com`


