package logging

import "log/slog"

// ParseLogLevel converts a string representation of a log level to its corresponding slog.Level constant.
//
// Parameters:
//   - level: The string representation of the log level.
//
// Returns:
//   - The equivalent slog.Level.
func ParseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
	default:
	}

	return slog.LevelError
}
