package environment

import (
    log "log"
    os "os"

    godotenv "github.com/joho/godotenv"
)


// InitEnv loads .env file only once
func Init() {
	if err := godotenv.Load(); err != nil {
        log.Println("⚠️ No .env file found, using system env vars")
    }
}

func Mandatory(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    log.Fatalf("❌ Required environment variable %s is not set", key)
    return "" // unreachable
}

// GetEnv returns the value of an environment variable, or fallback if not set
func Optional(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}
