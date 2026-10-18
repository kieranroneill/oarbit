package os

import "os"

// GetEnvWithDefault Retrieves the environment variable specified by the key or returns the default if the env var is empty or
// unset.
//
// Parameters:
//   - key: The key of the env var to get.
//   - _default: The default value to return if the env var is empty or unset.
//
// Returns:
//   - The env var or the default value if empty or unset.
func GetEnvWithDefault(key string, _default string) string {
	value := os.Getenv(key)

	if value == "" {
		return _default
	}

	return value
}
