package errors

// Check function raises error if condition is false.
// If *err is not nil (means error has raised early) function does nothing and returns immediately.
// Otherwise if case of cond is false functions create Error and puts to *err.
func Check(err *error, cond bool, text string) {
	// do nothing if error has raised early
	if *err != nil {
		return
	}

	if !cond {
		*err = Error(text)
	}
}

// Checkf function raises error if condition is false.
// If *err is not nil (means error has raised early) function does nothing and returns immediately.
// Otherwise if case of cond is false functions create Error (format with parameters) and puts to *err.
func Checkf(err *error, cond bool, format string, a ...interface{}) {
	// do nothing if error has raised early
	if *err != nil {
		return
	}

	if !cond {
		*err = Errorf(format, a...)
	}
}
