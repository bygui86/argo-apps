package envvars

import (
	"os"
	"strconv"

	"github.com/bygui86/go-metrics/logging"
)

// GetStringEnv -
func GetStringEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	logging.Log.Infoln("[UTILS] Environment variable", key, "not found, using fallback", fallback)
	return fallback
}

// GetIntEnv -
func GetIntEnv(key string, fallback int) int {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.Atoi(strValue)
		if err != nil {
			logging.Log.Errorln("[UTILS] Error reading the environment variable", key, ", not an int! Using fallback", fallback)
			return fallback
		}
		return value
	}
	logging.Log.Infoln("[UTILS] Environment variable", key, "not found, using fallback", fallback)
	return fallback
}

// GetBoolEnv -
func GetBoolEnv(key string, fallback bool) bool {
	if strValue, ok := os.LookupEnv(key); ok {
		value, err := strconv.ParseBool(strValue)
		if err != nil {
			logging.Log.Errorln("[UTILS] Error reading the environment variable", key, ", not an boolean! Using fallback", fallback)
			return fallback
		}
		return value
	}
	logging.Log.Infoln("[UTILS] Environment variable", key, "not found, using fallback", fallback)
	return fallback
}
