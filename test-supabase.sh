#!/bin/bash

# Quick test script untuk Supabase connection

echo "🧪 Testing Supabase Database Connection..."
echo ""

# Set environment variables
export DB_HOST="aws-1-ap-northeast-1.pooler.supabase.com"
export DB_PORT="5432"
export DB_USER="postgres.bpudjwunrwsgjyhqhssh"
export DB_PASSWORD="tnqb2289"
export DB_NAME="postgres"
export DB_SSLMODE="require"
export PORT="8080"
export CORS_ORIGIN="http://localhost:3000,http://localhost:3001"
export JWT_SECRET="test-secret-for-local-testing"
export JWT_TOKEN_TTL_MINUTES="720"
export ADMIN_USERNAME="ghifary"
export ADMIN_PASSWORD="tnqb2289"

echo "📋 Environment Variables Set:"
echo "  DB_HOST: $DB_HOST"
echo "  DB_PORT: $DB_PORT"
echo "  DB_USER: $DB_USER"
echo "  DB_NAME: $DB_NAME"
echo "  DB_SSLMODE: $DB_SSLMODE"
echo ""

echo "🚀 Starting Go server..."
echo ""

# Run the application
go run main.go


