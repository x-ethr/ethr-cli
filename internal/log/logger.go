// Package log provides a logging framework with level support and pluggable
// export destinations.
package log

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// Global sets up the global slog configuration according to LOG_LEVEL environment variable.
// It manages log levels and creates a handler with options for log level,
// customization of attributes and output to stdout as JSON.
// Here is the list of log levels with their priorities:
//
// TRACE < DEBUG < INFO < NOTICE < WARNING < ERROR < EMERGENCY
//
// The default log level is INFO.
func Global() {
	var level slog.Level

	// Setting the log level based on LOG_LEVEL environment variable
	switch Level() {
	case "TRACE":
		level = Trace
	case "DEBUG":
		level = Debug
	case "INFO":
		level = Info
	case "NOTICE":
		level = Notice
	case "WARNING":
		level = Warning
	case "ERROR":
		level = Error
	case "EMERGENCY":
		level = Emergency
	default:
		// Default log level is INFO
		level = Error
	}

	// New handler to output the logs as JSON
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				// Remove time from output
				return slog.Attr{}
			case slog.LevelKey:
				// Customize the name of the level key and the output string, including
				// custom level values.
				level := a.Value.Any().(slog.Level)

				// Renaming the log levels based on their priority
				switch {
				case level <= Debug:
					a.Value = slog.StringValue("DEBUG")
				case level <= Info:
					a.Value = slog.StringValue("INFO")
				case level <= Notice:
					a.Value = slog.StringValue("NOTICE")
				case level <= Warning:
					a.Value = slog.StringValue("WARNING")
				case level <= Error:
					a.Value = slog.StringValue("ERROR")
				case level <= Emergency:
					a.Value = slog.StringValue("EMERGENCY")
				default:
					a.Value = slog.StringValue("ERROR")
				}

			case slog.SourceKey:
				a.Key = "$"

				value := a.Value.String()[2 : len(a.Value.String())-1]
				partials := strings.Split(value, " ")
				value = strings.Join(partials[1:], ":")

				a.Value = slog.StringValue(fmt.Sprintf("file://%s", value))
			}

			return a
		},
	})

	// Set the default logger
	slog.SetDefault(slog.New(handler))
}
