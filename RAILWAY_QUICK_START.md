# ⚡ Quick Start - Deploy ke Railway (5 Menit)

Panduan cepat deploy backend Go ke Railway.

---

## 🚀 Step-by-Step (5 Menit)

### **1. Push ke GitHub (1 menit)**

```bash
cd /Users/ghifaryahmada/best-portfolio-go
git add .
git commit -m "Prepare for Railway"
git push origin main
```

---

### **2. Deploy di Railway (4 menit)**

1. **Railway Dashboard**: https://railway.app
2. **New Project** → **Deploy from GitHub repo**
3. **Pilih repo**: `best-portfolio-go`
4. Railway auto-detect Go project

5. **Variables** tab → Add:
   ```
   DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
   DB_PORT=5432
   DB_USER=postgres.bpudjwunrwsgjyhqhssh
   DB_PASSWORD=tnqb2289
   DB_NAME=postgres
   DB_SSLMODE=require
   PORT=8080
   CORS_ORIGIN=https://*.vercel.app
   JWT_SECRET=[random-32-chars]
   JWT_TOKEN_TTL_MINUTES=720
   ADMIN_USERNAME=ghifary
   ADMIN_PASSWORD=[secure-password]
   ```

6. **Settings** → **Generate Domain** → Copy URL

7. Tunggu deploy selesai (2-5 menit)

---

### **3. Test API**

```bash
curl https://your-service.up.railway.app/health
curl https://your-service.up.railway.app/api/projects
```

---

### **4. Update Frontend**

**Vercel/Railway**:
```
API_BASE_URL=https://your-service.up.railway.app
NEXT_PUBLIC_API_BASE=https://your-service.up.railway.app
```

---

### **5. Update CORS**

1. Railway → **Variables**
2. Update `CORS_ORIGIN` dengan domain Vercel:
   ```
   CORS_ORIGIN=https://portfolio.vercel.app,https://dashboard.vercel.app
   ```

---

## ⚠️ Important Notes

1. **Service Sleep**: Free tier sleep setelah 5 menit idle
   - **Solusi**: Setup UptimeRobot untuk ping setiap 5 menit

2. **Auto-Deploy**: Railway auto-deploy setiap push ke GitHub

3. **Environment Variables**: Pastikan semua sudah di-set sebelum deploy

---

## ✅ Checklist

- [ ] Code pushed ke GitHub
- [ ] Railway project created
- [ ] Environment variables set
- [ ] Domain generated
- [ ] Deploy successful
- [ ] API test berhasil
- [ ] Frontend updated
- [ ] CORS configured

---

**Done!** 🎉

API URL: `https://your-service.up.railway.app`


