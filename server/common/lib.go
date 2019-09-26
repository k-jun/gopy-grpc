package common

import "os"

func GetEnv(key string, defaultVal string) string {
	val := defaultVal
	if os.Getenv(key) != "" {
		val = os.Getenv(key)
	}
	return val
}