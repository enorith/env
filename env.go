package environment

import (
	"strconv"
	"syscall"
)

// GetInt get int value from system environment variable
func GetInt(key string, def ...int) int {
	v, found := syscall.Getenv(key)
	if value, err := strconv.Atoi(v); err == nil && found {
		return value
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// GetInt64 get int64 value from system environment variable
func GetInt64(key string, def ...int64) int64 {
	v, found := syscall.Getenv(key)
	if value, err := strconv.ParseInt(v, 10, 64); err == nil && found {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}

	return 0
}


// GetFloat64 get float64 value from system environment variable
func GetFloat64(key string, def ...float64) float64 {
	v, found := syscall.Getenv(key)
	if value, err := strconv.ParseFloat(v, 64); err == nil && found {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}

	return 0
}

// GetString get string value from system environment variable
func GetString(key string, def... string) string {
	if v, found := syscall.Getenv(key); found {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

// GetBoolean get boolean value from system environment variable
func GetBoolean(key string, def... bool) bool {
	if v, found := syscall.Getenv(key); found {
		if value, err := strconv.ParseBool(v); err == nil && found {
			return value
		}
	}

	if len(def) > 0 {
		return def[0]
	}

	return false
}
