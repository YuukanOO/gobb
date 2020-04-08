// Package logging provides a simple Logger interface and a default logger which
// used the standard Go library to print stuff out.
package logging

type (
	// F type used to add logs data when using a logger.
	// I find it more convenient since we can format it the way we need.
	F map[string]interface{}

	// Logger interface used to log stuff.
	Logger interface {
		// Info logs an info message using this logger.
		Info(string, ...F)
		// Error logs an error message using this logger.
		Error(string, ...F)
		// Debug logs a debug message we should only be visible during development.
		Debug(string, ...F)
	}
)
