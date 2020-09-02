package environment

import (
	"strconv"
	"syscall"
)

// GetInt get int value from system environment variable
func GetInt(key string, def int) int {
	v, found := syscall.Getenv(key)
	if value, err := strconv.Atoi(v); err == nil && found {
		return value
	}

	return def
}

// GetString get string value from system environment variable
func GetString(key string, def string) string {
	if v, found := syscall.Getenv(key); found {
		return v
	}

	return def
}

// GetBoolean get boolean value from system environment variable
func GetBoolean(key string, def bool) bool {
	if v, found := syscall.Getenv(key); found {
		if value, err := strconv.ParseBool(v); err == nil && found {
			return value
		}
	}

	return def
}
