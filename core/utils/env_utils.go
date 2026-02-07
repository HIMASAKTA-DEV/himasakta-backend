package utils

import "os"

func GetEnvWithFallback(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return value
}

func GetDomain() (string, bool) {
	mode, exists := os.LookupEnv("APP_MODE")
	if !exists {
		return "localhost", false
	}

	if mode == "development" {
		return "localhost", false
	} else {
		value, exists := os.LookupEnv("APP_DOMAIN")
		if !exists {
			return "localhost", false
		}
		return value, true
	}
}
