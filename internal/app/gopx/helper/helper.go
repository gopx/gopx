package helper

import (
	"fmt"
	"os"
)

// MustGetEnv returns the env value and panics when env variable not found.
func MustGetEnv(name string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	panic(fmt.Sprintf("Env variable not found: %s", name))
}
