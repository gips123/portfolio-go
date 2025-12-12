# 🔧 Fix CORS Error - Railway Backend

Panduan untuk memperbaiki error CORS di Railway backend.

---

## ❌ Masalah

Error yang muncul:
```
Access to fetch at 'https://portfolio-go-production-35c5.up.railway.app/api/projects/5/images' 
from origin 'https://portfolio-main-go.vercel.app' 
has been blocked by CORS policy
```

**Penyebab:**
- Backend di Railway belum mengizinkan origin `https://portfolio-main-go.vercel.app` di CORS

---

## ✅ Solusi: Update CORS di Railway

### **Step 1: Update Environment Variable di Railway**

1. **Railway Dashboard**: https://railway.app/dashboard
2. **Pilih service**: `portfolio-go-production` (atau nama service Anda)
3. **Variables** tab
4. **Edit `CORS_ORIGIN`** atau **Add** jika belum ada:

```
CORS_ORIGIN=https://portfolio-main-go.vercel.app,https://dashboard-admin-portfolio.vercel.app
```

**Format penting:**
- ✅ Comma-separated, **tanpa spasi**
- ✅ Harus include `https://` (bukan `http://`)
- ✅ Tidak ada trailing slash

**Contoh lengkap:**
```
CORS_ORIGIN=https://portfolio-main-go.vercel.app,https://dashboard-admin-portfolio.vercel.app,http://localhost:3000,http://localhost:3001
```

### **Step 2: Redeploy Backend**

Setelah update env var:
1. Railway akan **auto-redeploy** setelah save
2. Atau **Manual Redeploy**: Deployments → Redeploy
3. Tunggu sampai status **Deployed**

### **Step 3: Verifikasi**

Test dengan curl:
```bash
curl -H "Origin: https://portfolio-main-go.vercel.app" \
     -H "Access-Control-Request-Method: GET" \
     -H "Access-Control-Request-Headers: Content-Type" \
     -X OPTIONS \
     https://portfolio-go-production-35c5.up.railway.app/api/projects/5/images -v
```

Harus return header:
```
Access-Control-Allow-Origin: https://portfolio-main-go.vercel.app
```

---

## 📋 Environment Variables Lengkap untuk Railway

Pastikan semua env vars sudah di-set:

```
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require
PORT=8080
CORS_ORIGIN=https://portfolio-main-go.vercel.app,https://dashboard-admin-portfolio.vercel.app
JWT_SECRET=[your-jwt-secret]
JWT_TOKEN_TTL_MINUTES=720
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=[your-password]
```

---

## 🔍 Debugging

### Check CORS Configuration

1. **Railway** → **Variables** → Pastikan `CORS_ORIGIN` sudah benar
2. **Railway** → **Deployments** → **View Logs** → Cek apakah ada log:
   ```
   CORS Allowed Origins: [https://portfolio-main-go.vercel.app ...]
   ```

### Test CORS dengan Browser

1. Buka browser console di `https://portfolio-main-go.vercel.app`
2. Test fetch:
   ```javascript
   fetch('https://portfolio-go-production-35c5.up.railway.app/api/projects')
     .then(r => r.json())
     .then(console.log)
     .catch(console.error)
   ```
3. Pastikan tidak ada CORS error

---

## ✅ Checklist

- [ ] `CORS_ORIGIN` di Railway sudah include `https://portfolio-main-go.vercel.app`
- [ ] `CORS_ORIGIN` format benar (comma-separated, tanpa spasi)
- [ ] Backend sudah di-redeploy setelah update env var
- [ ] Test CORS berhasil
- [ ] Frontend bisa akses API tanpa error

---

## 🎯 Quick Fix

1. **Railway** → **Variables** → Edit `CORS_ORIGIN`:
   ```
   https://portfolio-main-go.vercel.app,https://dashboard-admin-portfolio.vercel.app
   ```
2. **Save** → Tunggu auto-redeploy
3. **Test** → Refresh frontend

---

**Done!** 🎉

