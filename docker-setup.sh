#!/bin/bash

# Docker Setup Script untuk Portfolio API
# Script ini membantu setup database setelah container berjalan

set -e

echo "🚀 Portfolio API Docker Setup"
echo "=============================="
echo ""

# Check if containers are running
if ! docker-compose ps | grep -q "portfolio-postgres.*Up"; then
    echo "❌ Error: PostgreSQL container tidak berjalan"
    echo "Jalankan: docker-compose up -d"
    exit 1
fi

echo "📦 Checking database connection..."
sleep 2

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
until docker-compose exec -T postgres pg_isready -U postgres > /dev/null 2>&1; do
    echo "   Waiting..."
    sleep 1
done

echo "✅ PostgreSQL is ready!"
echo ""

# Check if database exists
DB_EXISTS=$(docker-compose exec -T postgres psql -U postgres -tAc "SELECT 1 FROM pg_database WHERE datname='portfolio'" || echo "")

if [ "$DB_EXISTS" != "1" ]; then
    echo "📝 Creating database..."
    docker-compose exec -T postgres psql -U postgres -c "CREATE DATABASE portfolio;"
    echo "✅ Database created!"
else
    echo "ℹ️  Database 'portfolio' already exists"
fi

echo ""

# Check if migrations directory exists
if [ ! -d "migrations" ]; then
    echo "⚠️  Warning: migrations directory tidak ditemukan"
    echo "   Silakan buat file migration SQL terlebih dahulu"
    exit 0
fi

# Check for migration files
if [ -f "migrations/01_create_tables.sql" ]; then
    echo "📋 Running migrations..."
    docker-compose exec -T postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/01_create_tables.sql 2>/dev/null || \
    docker-compose exec -T postgres psql -U postgres -d portfolio < migrations/01_create_tables.sql
    echo "✅ Migrations completed!"
else
    echo "⚠️  Warning: migrations/01_create_tables.sql tidak ditemukan"
    echo "   Silakan buat file migration terlebih dahulu"
fi

echo ""

# Ask about seed data
read -p "🌱 Apakah ingin menjalankan seed data? (y/n): " -n 1 -r
echo ""
if [[ $REPLY =~ ^[Yy]$ ]]; then
    if [ -f "migrations/seed.sql" ]; then
        echo "📦 Running seed data..."
        docker-compose exec -T postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/seed.sql 2>/dev/null || \
        docker-compose exec -T postgres psql -U postgres -d portfolio < migrations/seed.sql
        
        if [ -f "migrations/03_fix_sequences.sql" ]; then
            echo "🔧 Fixing sequences..."
            docker-compose exec -T postgres psql -U postgres -d portfolio -f /docker-entrypoint-initdb.d/03_fix_sequences.sql 2>/dev/null || \
            docker-compose exec -T postgres psql -U postgres -d portfolio < migrations/03_fix_sequences.sql
        fi
        
        echo "✅ Seed data completed!"
    else
        echo "⚠️  Warning: migrations/seed.sql tidak ditemukan"
    fi
fi

echo ""
echo "🎉 Setup selesai!"
echo ""
echo "📊 Status:"
docker-compose ps
echo ""
echo "🔗 API tersedia di: http://localhost:8080"
echo "📚 Database: localhost:5432 (user: postgres, password: postgres, db: portfolio)"
echo ""
echo "💡 Tips:"
echo "   - View logs: docker-compose logs -f api"
echo "   - Stop: docker-compose down"
echo "   - Restart: docker-compose restart"

