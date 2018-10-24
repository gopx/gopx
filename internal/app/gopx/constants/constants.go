package constants

// The required environment variable names
const (
	EnvNameHost        = "HOST"
	EnvNamePort        = "PORT"
	EnvNameDatabaseURL = "DATABASE_URL"
)

// Request timeout constants for the server.
const (
	ServerReadTimeout  = 15
	ServerWriteTimeout = 15
	ServerIdleTimeout  = 15
)
