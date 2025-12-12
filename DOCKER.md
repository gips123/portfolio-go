# Docker Setup Guide

Panduan untuk menjalankan Portfolio API menggunakan Docker dan Docker Compose.

## 📋 Prerequisites

- Docker (version 20.10+)
- Docker Compose (version 2.0+)

## 🚀 Quick Start

### 1. Development Mode

Jalankan aplikasi dengan Docker Compose:

```bash
# Build dan start semua services
docker-compose up -d

# View logs
docker-compose logs -f api

# Stop services
docker-compose down

# Stop dan hapus volumes (database data akan hilang)
docker-compose down -v
```

Aplikasi akan berjalan di:
- **API**: http://localhost:8080
- **PostgreSQL**: localhost:5432

### 2. Setup Database

**Option A: Menggunakan Setup Script (Recommended)**

```bash
# Jalankan script setup otomatis
./docker-setup.sh
```

Script ini akan:
- Membuat database jika belum ada
- Menjalankan migration
- Menjalankan seed data (optional)
- Fix sequences

**Option B: Manual Setup**

```bash
# Masuk ke container postgres
docker-compose exec postgres psql -U postgres -d portfolio

# Atau jalankan migration dari host (jika file migration ada)
docker-compose exec postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/01_create_tables.sql

# Jalankan seed data (optional)
docker-compose exec postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/seed.sql

# Fix sequences setelah seed
docker-compose exec postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/03_fix_sequences.sql
```

### 3. Environment Variables

Untuk development, environment variables sudah di-set di `docker-compose.yml`. Untuk production, gunakan file `.env`:

```bash
# Copy dan edit .env
cp .env.example .env

# Edit sesuai kebutuhan
nano .env
```

Kemudian jalankan dengan:

```bash
docker-compose --env-file .env up -d
```

### 4. Production Mode

Untuk production, gunakan `docker-compose.prod.yml`:

```bash
# Build
docker-compose -f docker-compose.yml -f docker-compose.prod.yml build

# Run
docker-compose -f docker-compose.yml -f docker-compose.prod.yml up -d
```

**Penting untuk production:**
- Ganti `JWT_SECRET` dengan secret key yang kuat
- Ganti `ADMIN_USERNAME` dan `ADMIN_PASSWORD`
- Set `CORS_ORIGIN` sesuai domain frontend
- Pertimbangkan menggunakan reverse proxy (nginx) di depan API
- Setup backup database secara berkala

## 🛠️ Useful Commands

### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f api
docker-compose logs -f postgres
```

### Restart Services

```bash
# Restart all
docker-compose restart

# Restart specific service
docker-compose restart api
```

### Execute Commands in Container

```bash
# Run migration
docker-compose exec api ./portfolio-api

# Access PostgreSQL
docker-compose exec postgres psql -U postgres -d portfolio

# Access API container shell
docker-compose exec api sh
```

### Rebuild After Code Changes

```bash
# Rebuild and restart
docker-compose up -d --build

# Rebuild without cache
docker-compose build --no-cache
```

### Database Backup

```bash
# Backup database
docker-compose exec postgres pg_dump -U postgres portfolio > backup.sql

# Restore database
docker-compose exec -T postgres psql -U postgres portfolio < backup.sql
```

## 📁 File Structure

```
.
├── Dockerfile                 # Multi-stage build untuk aplikasi
├── docker-compose.yml         # Development configuration
├── docker-compose.prod.yml    # Production overrides
├── .dockerignore              # Files excluded from Docker build
└── migrations/                # SQL migration files
    ├── 01_create_tables.sql
    ├── 03_fix_sequences.sql
    └── seed.sql
```

## 🔍 Troubleshooting

### Port Already in Use

Jika port 8080 atau 5432 sudah digunakan:

```bash
# Edit docker-compose.yml dan ubah port mapping
ports:
  - "8081:8080"  # API akan di localhost:8081
  - "5433:5432"  # PostgreSQL akan di localhost:5433
```

### Database Connection Error

Pastikan:
1. PostgreSQL container sudah running: `docker-compose ps`
2. Database sudah dibuat: `docker-compose exec postgres psql -U postgres -l`
3. Environment variables benar di `docker-compose.yml`

### Permission Denied

Jika ada permission issues:

```bash
# Fix permissions
sudo chown -R $USER:$USER .
```

## 🐳 Docker Image Details

- **Base Image**: `golang:1.21-alpine` (builder), `alpine:latest` (runtime)
- **Final Image Size**: ~15-20MB (optimized)
- **Health Check**: `/health` endpoint setiap 30 detik
- **Timezone**: Asia/Jakarta (bisa diubah di Dockerfile)

## 📝 Notes

- Database data akan persist di Docker volume `postgres_data`
- Migration files di `migrations/` akan otomatis di-copy ke container
- Untuk development, gunakan volume mounting untuk hot reload (tidak di-setup di docker-compose.yml default)

