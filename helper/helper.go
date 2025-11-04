package helper

import (
	"log"
)

// Functions starting with Must conventionally indicate that they will panic
// if an error occurs.
// Must wraps a function that returns an error, and panics if is not nil.
// It helps avoiding repetitive error checking boilerplate.
func Must(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
