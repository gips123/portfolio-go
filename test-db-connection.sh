#!/bin/bash

# Script untuk test koneksi database ke Supabase

echo "🔍 Testing Database Connection to Supabase..."
echo "=============================================="
echo ""

# Set environment variables dari connection string
export DB_HOST="aws-1-ap-northeast-1.pooler.supabase.com"
export DB_PORT="5432"
export DB_USER="postgres.bpudjwunrwsgjyhqhssh"
export DB_PASSWORD="tnqb2289"
export DB_NAME="postgres"
export DB_SSLMODE="require"
export PORT="8080"

echo "📋 Connection Details:"
echo "  Host: $DB_HOST"
echo "  Port: $DB_PORT"
echo "  User: $DB_USER"
echo "  Database: $DB_NAME"
echo "  SSL Mode: $DB_SSLMODE"
echo ""

# Test dengan psql jika tersedia
if command -v psql &> /dev/null; then
    echo "🧪 Testing with psql..."
    PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -p "$DB_PORT" -U "$DB_USER" -d "$DB_NAME" -c "SELECT version();" 2>&1
    if [ $? -eq 0 ]; then
        echo "✅ psql connection successful!"
    else
        echo "❌ psql connection failed!"
    fi
    echo ""
fi

# Test dengan Go application
echo "🧪 Testing with Go application..."
echo "Starting server..."

# Run in background
go run main.go > /tmp/portfolio-api.log 2>&1 &
SERVER_PID=$!

# Wait for server to start
sleep 5

# Check if server is running
if ps -p $SERVER_PID > /dev/null; then
    echo "✅ Server started (PID: $SERVER_PID)"
    
    # Test health endpoint
    echo ""
    echo "🧪 Testing /health endpoint..."
    curl -s http://localhost:8080/health || echo "❌ Health check failed"
    
    # Test API endpoint
    echo ""
    echo "🧪 Testing /api/projects endpoint..."
    curl -s http://localhost:8080/api/projects | head -c 200
    echo ""
    
    # Stop server
    echo ""
    echo "🛑 Stopping server..."
    kill $SERVER_PID 2>/dev/null
    wait $SERVER_PID 2>/dev/null
    echo "✅ Server stopped"
else
    echo "❌ Server failed to start"
    echo "📋 Logs:"
    cat /tmp/portfolio-api.log
fi

echo ""
echo "📋 Full server logs:"
cat /tmp/portfolio-api.log


