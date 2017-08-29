package errors

import (
	"fmt"
	"runtime"
)

type errorItem struct {
	fileName string
	fileLine int
	msg      string
}

type errors struct {
	stack []errorItem
}

func wrap(err error, msg string, a ...interface{}) error {
	var errs *errors
	var item errorItem

	_, fileName, fileLine, ok := runtime.Caller(2)
	if ok {
		item.fileName = fileName
		item.fileLine = fileLine
	}
	item.msg = fmt.Sprintf(msg, a...)

	if err != nil {
		switch v := err.(type) {
		case *errors:
			errs = v
		}
	}

	if err == nil {
		errs = &errors{stack: []errorItem{}}
	}

	errs.stack = append(errs.stack, item)

	return errs
}

func (e *errors) Error() string {
	var msg string
	for _, item := range e.stack {
		if len(item.fileName) > 0 && item.fileLine != 0 {
			msg = fmt.Sprintf("[%s:%d] %s\n", item.fileName, item.fileLine, item.msg) + msg
		} else {
			msg = fmt.Sprintf("%s\n", item.msg) + msg
		}
	}
	return msg
}

// New returns error with caller
func New(msg string, a ...interface{}) error {
	return wrap(nil, msg, a...)
}

// Wrap returns error wrapped by caller info if 'err' is not nil
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return wrap(err, "")
}

// Wrapf returns error wrapped by message and caller info if 'err' is not nil
func Wrapf(err error, msg string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrap(err, msg, a...)
}
