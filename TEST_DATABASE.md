# 🧪 Test Database Connection ke Supabase

Panduan untuk test koneksi database ke Supabase sebelum deploy.

---

## 📋 Connection Details

**Connection String:**
```
postgresql://postgres.bpudjwunrwsgjyhqhssh:tnqb2289@aws-1-ap-northeast-1.pooler.supabase.com:5432/postgres
```

**Parsed:**
- Host: `aws-1-ap-northeast-1.pooler.supabase.com`
- Port: `5432`
- User: `postgres.bpudjwunrwsgjyhqhssh`
- Password: `tnqb2289`
- Database: `postgres`
- SSL Mode: `require`

---

## 🧪 Method 1: Test dengan Script (Recommended)

```bash
cd /Users/ghifaryahmada/best-portfolio-go
./test-db-connection.sh
```

Script ini akan:
1. Test koneksi dengan psql (jika tersedia)
2. Start Go server dengan env vars
3. Test `/health` endpoint
4. Test `/api/projects` endpoint
5. Show logs

---

## 🧪 Method 2: Test Manual dengan Go

### 2.1. Setup Environment Variables

```bash
export DB_HOST="aws-1-ap-northeast-1.pooler.supabase.com"
export DB_PORT="5432"
export DB_USER="postgres.bpudjwunrwsgjyhqhssh"
export DB_PASSWORD="tnqb2289"
export DB_NAME="postgres"
export DB_SSLMODE="require"
export PORT="8080"
export CORS_ORIGIN="http://localhost:3000,http://localhost:3001"
export JWT_SECRET="test-secret"
export JWT_TOKEN_TTL_MINUTES="720"
export ADMIN_USERNAME="ghifary"
export ADMIN_PASSWORD="tnqb2289"
```

### 2.2. Run Application

```bash
cd /Users/ghifaryahmada/best-portfolio-go
go run main.go
```

### 2.3. Test Endpoints

Di terminal lain:

```bash
# Health check
curl http://localhost:8080/health

# Test projects endpoint
curl http://localhost:8080/api/projects

# Test about endpoint
curl http://localhost:8080/api/about

# Test skills endpoint
curl http://localhost:8080/api/skills

# Test contact endpoint
curl http://localhost:8080/api/contact
```

---

## 🧪 Method 3: Test dengan .env File

### 3.1. Copy .env File

```bash
cp .env.supabase .env
```

### 3.2. Run Application

```bash
go run main.go
```

Aplikasi akan otomatis load `.env` file.

---

## 🧪 Method 4: Test dengan psql (Jika Installed)

```bash
psql "postgresql://postgres.bpudjwunrwsgjyhqhssh:tnqb2289@aws-1-ap-northeast-1.pooler.supabase.com:5432/postgres?sslmode=require"

# Atau dengan environment variable
PGPASSWORD=tnqb2289 psql -h aws-1-ap-northeast-1.pooler.supabase.com -p 5432 -U postgres.bpudjwunrwsgjyhqhssh -d postgres
```

Setelah connect, test query:

```sql
-- Test connection
SELECT version();

-- Check tables
\dt

-- Test query
SELECT COUNT(*) FROM projects;
```

---

## ✅ Expected Results

### Success Indicators:

1. **Server Start**:
   ```
   Database connected successfully
   CORS Allowed Origins: [...]
   Server starting on port 8080
   ```

2. **Health Endpoint**:
   ```json
   {
     "status": "ok",
     "message": "Portfolio API is running"
   }
   ```

3. **API Endpoints**:
   - `/api/projects` → Returns array of projects
   - `/api/about` → Returns array of about cards
   - `/api/skills` → Returns array of skill categories
   - `/api/contact` → Returns contact data

---

## ❌ Troubleshooting

### Error: "connection refused" atau "timeout"

**Kemungkinan**:
- Firewall blocking
- SSL mode salah
- Host/port salah

**Solusi**:
- Pastikan `DB_SSLMODE=require`
- Check connection string
- Test dengan psql dulu

### Error: "authentication failed"

**Kemungkinan**:
- Password salah
- User salah

**Solusi**:
- Double check password: `tnqb2289`
- Check user: `postgres.bpudjwunrwsgjyhqhssh`

### Error: "relation does not exist"

**Kemungkinan**:
- Tabel belum dibuat
- Migration belum dijalankan

**Solusi**:
- Jalankan migration di Supabase SQL Editor
- File: `supabase/migrations/20251212150000_create_tables.sql`

### Error: "no rows in result set"

**Kemungkinan**:
- Tabel kosong
- Seed data belum dijalankan

**Solusi**:
- Jalankan seed data di Supabase SQL Editor
- File: `supabase/migrations/20251212152735_new-migration.sql`

---

## 📝 Environment Variables untuk Render

Setelah test berhasil, gunakan env vars ini di Render:

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

---

## 🎯 Quick Test Command

```bash
# One-liner test
cd /Users/ghifaryahmada/best-portfolio-go && \
export DB_HOST="aws-1-ap-northeast-1.pooler.supabase.com" && \
export DB_PORT="5432" && \
export DB_USER="postgres.bpudjwunrwsgjyhqhssh" && \
export DB_PASSWORD="tnqb2289" && \
export DB_NAME="postgres" && \
export DB_SSLMODE="require" && \
export PORT="8080" && \
go run main.go
```

Lalu di terminal lain:
```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/projects
```


