package constants

import (
	_ "embed"
	"sync/atomic"
)

var executable atomic.Value

// Name is a function that returns the name of the executable.
func Name() (cli string) {
	var valid bool
	if cli, valid = executable.Load().(string); !(valid) {
		panic("unable to typecast atomic.Value \"executable\" to string")
	}

	return
}

func init() {
	executable.Store("ethr-cli")
}
