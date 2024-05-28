package log

import (
	"log/slog"
)

// Exported constants representing [slog.Level].
//
// - Trace for tracing program's execution.
//
// - Debug for providing contextual information in debugging phase.
//
// - Info for informing about general system operations.
//
// - Notice for conditions that are not errors but might need handling.
//
// - Warning for warning conditions.
//
// - Error for error conditions.
//
// - Emergency for system-unusable conditions.
const (
	Trace     = slog.Level(-8)
	Debug     = slog.LevelDebug
	Info      = slog.LevelInfo
	Notice    = slog.Level(2)
	Warning   = slog.LevelWarn
	Error     = slog.LevelError
	Emergency = slog.Level(12)
)
