# 📝 Panduan Environment Variables

Panduan lengkap untuk environment variables yang perlu diubah di file `.env`.

---

## 🔧 Environment Variables yang Harus Diubah

### **1. Database Configuration (Supabase)**

Dari connection string:
```
postgresql://postgres.bpudjwunrwsgjyhqhssh:tnqb2289@aws-1-ap-northeast-1.pooler.supabase.com:5432/postgres
```

**Yang perlu diubah:**

```env
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require
```

**Penjelasan:**
- `DB_HOST`: Host database dari Supabase
- `DB_PORT`: Port database (5432 untuk direct, 6543 untuk pooler)
- `DB_USER`: Username database
- `DB_PASSWORD`: Password database (dari Supabase)
- `DB_NAME`: Nama database (biasanya `postgres`)
- `DB_SSLMODE`: `require` untuk production/Supabase

---

### **2. Server Configuration**

```env
PORT=8080
```

**Penjelasan:**
- Port untuk server API
- Default: `8080`
- Render akan override dengan env var `PORT` jika ada

---

### **3. CORS Configuration**

```env
CORS_ORIGIN=http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000,http://127.0.0.1:3001
```

**Yang perlu diubah untuk production:**

```env
CORS_ORIGIN=https://your-portfolio.vercel.app,https://your-dashboard.vercel.app
```

**Penjelasan:**
- Daftar domain yang diizinkan mengakses API
- Format: comma-separated, tanpa spasi
- Untuk development: localhost
- Untuk production: domain frontend yang sebenarnya

---

### **4. JWT Configuration**

```env
JWT_SECRET=change-me-to-secure-secret-key-in-production
JWT_TOKEN_TTL_MINUTES=720
```

**Yang perlu diubah:**

```env
JWT_SECRET=[generate-random-32-chars]
```

**Cara generate:**
```bash
openssl rand -base64 32
```

Atau online: https://generate-secret.vercel.app/32

**Penjelasan:**
- `JWT_SECRET`: Secret key untuk sign JWT tokens (WAJIB diganti!)
- `JWT_TOKEN_TTL_MINUTES`: Token expiry time (720 = 12 jam)

---

### **5. Admin Credentials**

```env
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=tnqb2289
```

**Yang perlu diubah:**

```env
ADMIN_PASSWORD=[your-secure-password]
```

**Penjelasan:**
- `ADMIN_USERNAME`: Username untuk login dashboard
- `ADMIN_PASSWORD`: Password untuk login (WAJIB diganti untuk production!)

---

## 📋 Template .env untuk Supabase

```env
# Database Configuration - Supabase
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require

# Server Configuration
PORT=8080

# CORS Configuration (Development)
CORS_ORIGIN=http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000,http://127.0.0.1:3001

# CORS Configuration (Production) - Uncomment dan ganti dengan domain sebenarnya
# CORS_ORIGIN=https://your-portfolio.vercel.app,https://your-dashboard.vercel.app

# JWT Configuration
JWT_SECRET=change-me-to-secure-secret-key-in-production
JWT_TOKEN_TTL_MINUTES=720

# Admin Credentials
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=tnqb2289
```

---

## 🔐 Security Checklist

Sebelum deploy ke production, pastikan:

- [ ] `JWT_SECRET` sudah diganti dengan random string (32+ chars)
- [ ] `ADMIN_PASSWORD` sudah diganti dari default
- [ ] `DB_PASSWORD` sudah benar (dari Supabase)
- [ ] `CORS_ORIGIN` sudah di-update dengan domain production
- [ ] `DB_SSLMODE=require` untuk production

---

## 🎯 Quick Setup untuk Supabase

Copy-paste ini ke `.env`:

```env
DB_HOST=aws-1-ap-northeast-1.pooler.supabase.com
DB_PORT=5432
DB_USER=postgres.bpudjwunrwsgjyhqhssh
DB_PASSWORD=tnqb2289
DB_NAME=postgres
DB_SSLMODE=require
PORT=8080
CORS_ORIGIN=http://localhost:3000,http://localhost:3001,http://127.0.0.1:3000,http://127.0.0.1:3001
JWT_SECRET=change-me-to-secure-secret-key-in-production
JWT_TOKEN_TTL_MINUTES=720
ADMIN_USERNAME=ghifary
ADMIN_PASSWORD=tnqb2289
```

---

## 📝 Untuk Render Deploy

Gunakan env vars yang sama, tapi:
1. `CORS_ORIGIN` ganti dengan domain Vercel
2. `JWT_SECRET` generate random baru
3. `ADMIN_PASSWORD` ganti dengan password kuat


