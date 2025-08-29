package core

import (
    log "log"
    os "os"
    sync "sync"
    godotenv "github.com/joho/godotenv"
)

var once sync.Once

// InitEnv loads .env file only once
func InitEnv() {
    once.Do(func() {
        if err := godotenv.Load(); err != nil {
            log.Println("⚠️ No .env file found, using system env vars")
        }
    })
}

// GetEnv returns the value of an environment variable, or fallback if not set
func EnvOptional(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func EnvMandatory(key string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    log.Fatalf("❌ Required environment variable %s is not set", key)
    return "" // unreachable
}
