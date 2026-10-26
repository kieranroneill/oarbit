package authentication

import "time"

const (
	Path                string        = "/auth"
	SessionTTLInSeconds time.Duration = 604800 * time.Second // 7 days (in seconds)
	StatusEndpoint      string        = "/status"
)
