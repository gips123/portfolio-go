# Contributing Guide

Terima kasih atas minat Anda untuk berkontribusi pada Portfolio Backend API!

## 📋 Development Setup

1. Fork repository ini
2. Clone fork Anda:
   ```bash
   git clone <your-fork-url>
   cd best-portfolio-go
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Setup environment:
   ```bash
   cp .env.example .env
   # Edit .env dengan konfigurasi database Anda
   ```

5. Setup database (lihat `docs/SETUP_DATABASE.md`)

6. Run application:
   ```bash
   go run main.go
   ```

## 🔧 Code Style

- Gunakan `gofmt` untuk formatting
- Ikuti Go best practices
- Tambahkan comments untuk public functions
- Gunakan meaningful variable names

## 📝 Commit Messages

Gunakan format yang jelas:
- `feat: Add new endpoint for X`
- `fix: Fix database connection issue`
- `docs: Update README`
- `refactor: Clean up handlers`

## 🧪 Testing

Sebelum submit PR:
- Test semua endpoints dengan Postman
- Pastikan tidak ada error
- Cek response format sesuai spesifikasi

## 📦 Pull Request Process

1. Buat branch baru untuk feature/fix
2. Commit perubahan Anda
3. Push ke fork Anda
4. Buat Pull Request dengan deskripsi yang jelas

---

**Thank you for contributing! 🎉**

