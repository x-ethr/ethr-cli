package log

import (
    "strings"
    "sync/atomic"
)

// defaultLogLevel represents the default log level.
//
// It is an atomic value that stores the current log level.
var defaultLogLevel atomic.Value

// Default sets the default log level to the specified level, and will
// instantiate a global default logger.
//
//   - See Global for additional details.
func Default(level string) {
    defaultLogLevel.Store(strings.ToUpper(level))

    Global()
}

// Level gets the value of the default log level
//
// This function will panic if it fails to convert the log level to string.
func Level() (v string) {
    var valid bool
    if v, valid = defaultLogLevel.Load().(string); !(valid) {
        panic("unable to cast atomic default-log-level to string")
    }

    return
}

func init() {
    defaultLogLevel.Store("ERROR")
}
