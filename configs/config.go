package configs

import (
    "log"
    "os"
)

func GetEnv(key, fallback string) string {
    val := os.Getenv(key)
    if val == "" {
        log.Printf("Warning: %s not set, using fallback", key)
        return fallback
    }
    return val
}
