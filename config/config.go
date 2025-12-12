package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
	CORS     CORSConfig
	Auth     AuthConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type ServerConfig struct {
	Port string
}

type CORSConfig struct {
	AllowOrigins []string
}

// AuthConfig holds authentication-related configuration
type AuthConfig struct {
	JWTSecret       string
	TokenTTLMinutes int
	AdminUsername   string
	AdminPassword   string
}

func LoadConfig() *Config {
	return &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "postgres"),
			DBName:   getEnv("DB_NAME", "portfolio_db"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		Server: ServerConfig{
			Port: getEnv("PORT", "8080"),
		},
		CORS: CORSConfig{
			AllowOrigins: getCORSOrigins(),
		},
		Auth: AuthConfig{
			JWTSecret:       getEnv("JWT_SECRET", "change-me"),
			TokenTTLMinutes: getEnvAsInt("JWT_TOKEN_TTL_MINUTES", 720), // default 12 hours
			AdminUsername:   getEnv("ADMIN_USERNAME", "admin"),
			AdminPassword:   getEnv("ADMIN_PASSWORD", "admin123"),
		},
	}
}

func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getCORSOrigins returns allowed CORS origins
// Always includes http://localhost:3000, http://localhost:3001, and http://127.0.0.1:3000
// Also supports additional origins from CORS_ORIGIN env var (comma-separated)
func getCORSOrigins() []string {
	// Always include both localhost ports and 127.0.0.1 (browser treats them differently)
	origins := []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
		"https://dashboard-admin-portfolio.vercel.app",
		"https://portfolio-main-go.vercel.app",
	}

	// Add additional origins from CORS_ORIGIN if specified
	corsOrigin := getEnv("CORS_ORIGIN", "")
	if corsOrigin != "" {
		// Split by comma and trim spaces
		additionalOrigins := strings.Split(corsOrigin, ",")
		for _, origin := range additionalOrigins {
			trimmed := strings.TrimSpace(origin)
			if trimmed != "" {
				// Check if not already in the list
				exists := false
				for _, existing := range origins {
					if existing == trimmed {
						exists = true
						break
					}
				}
				if !exists {
					origins = append(origins, trimmed)
				}
			}
		}
	}

	return origins
}
