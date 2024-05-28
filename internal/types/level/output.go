package level

import (
    "errors"
    "log/slog"
    "strings"
)

// Type string that implements Cobra's Type interface for valid string enumeration values.
type Type string

const (
    Trace     Type = "trace"
    Debug     Type = "debug"
    Info      Type = "info"
    Notice    Type = "notice"
    Warning   Type = "warning"
    Error     Type = "error"
    Emergency Type = "emergency"
)

// String is used both by fmt.Print and by Cobra in help text
func (o *Type) String() string {
    return string(*o)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (o *Type) Set(v string) error {
    switch strings.ToLower(v) {
    case "trace", "debug", "info", "notice", "warning", "error", "emergency":
        *o = Type(v)

        return nil
    default:
        return errors.New("must be one of \"trace\", \"debug\", \"info\", \"notice\", \"warning\", \"error\", \"emergency\"")
    }
}

// Type is only used in help text
func (o *Type) Type() string {
    return "[\"trace\"|\"debug\"|\"info\"|\"notice\"|\"warning\"|\"error\"|\"emergency\"]"
}

// Level - Exported constants representing [slog.Level].
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
func (o *Type) Level() slog.Level {
    const Default = slog.LevelError

    if o == nil {
        return Default
    }

    switch *o {
    case Trace:
        return slog.Level(-8)

    case Debug:
        return slog.LevelDebug

    case Info:
        return slog.LevelInfo

    case Notice:
        return slog.Level(2)

    case Warning:
        return slog.LevelWarn

    case Error:
        return slog.LevelError

    case Emergency:
        return slog.Level(12)
    default:
        return Default
    }
}
