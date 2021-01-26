package httputils

import "os"

// GetEnv ...
func GetEnv() string {
	if _, err := os.Stat(".env.test"); os.IsNotExist(err) {
		return ""
	}
	return ".env.test"
}
