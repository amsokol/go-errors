package errors

import (
	"fmt"
	"strings"
)

// Wrap wraps error and returns the string
// as a value that satisfies error.
func Wrap(err error, text string) error {
	return Error(strings.Join([]string{text, err.Error()}, "-> "))
}

// Wrapf wraps error, formats according to a format specifier and returns the string
// as a value that satisfies error.
func Wrapf(err error, format string, a ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, a...))
}
