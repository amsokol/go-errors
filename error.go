// Package errors provides primitives for simple error handling primitives.
package errors

import (
	go_errors "errors"
	"fmt"
)

// Error returns an error that formats as the given text.
func Error(text string) error {
	return go_errors.New(text)
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}
