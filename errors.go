package errors

import (
	"fmt"
	"path/filepath"
	"runtime"
	"bytes"
	"strconv"
	"strings"
)

type Fields map[string]interface{}

type errorItem struct {
	fileName string
	fileLine int
	msg      string
}

type field struct {
	name  string
	value interface{}
}

type errors struct {
	fields []field
	stack  []errorItem
}

func wrap(err error, fields Fields, msg string, a ...interface{}) error {
	var errs *errors
	var item errorItem

	if err != nil {
		switch v := err.(type) {
		case *errors:
			errs = v
		}
	}

	if errs == nil {
		errs = &errors{stack: []errorItem{}, fields: make([]field, 0, 10)}
		if err != nil {
			errs.stack = append(errs.stack,
				errorItem{fileName: "", fileLine: 0, msg: err.Error()})
		}
	}

	_, fileName, fileLine, ok := runtime.Caller(2)
	if ok {
		_, fileName = filepath.Split(fileName)
		item.fileName = fileName
		item.fileLine = fileLine
	}
	item.msg = fmt.Sprintf(msg, a...)

	errs.stack = append(errs.stack, item)

	if fields != nil {
		for name, value := range fields {
			errs.fields = append(errs.fields, field{name: name, value: value})
		}
	}

	return errs
}
func (e *errors) Error() string {
	var buf [1024]byte
	msg := bytes.NewBuffer(buf[:0])

	for i := len(e.stack) - 1; i >= 0; i-- {
		err := e.stack[i]
		if len(err.fileName) > 0 && err.fileLine != 0 {
			msg.WriteString("[")
			msg.WriteString(err.fileName)
			msg.WriteString(":")
			msg.WriteString(strconv.Itoa(err.fileLine))
			msg.WriteString("] ")
			msg.WriteString(err.msg)
		} else {
			msg.WriteString(err.msg)
		}
		if i > 0 {
			msg.WriteString(": ")
		}
	}

	if len(e.fields) > 0 {
		msg.WriteString(";")
	}

	for _, f := range e.fields {
		msg.WriteString(" ")
		msg.WriteString(f.name)
		msg.WriteString("=\"")
		msg.WriteString(strings.Replace(fmt.Sprintf("%+v", f.value), "\"", "\\\"", -1))
		msg.WriteString("\"")
	}
	return msg.String()
}

// New returns error with caller
func New(msg string) error {
	return wrap(nil, Fields{}, msg)
}

// New returns error with caller
func Newf(msg string, a ...interface{}) error {
	return wrap(nil, Fields{}, msg, a...)
}

// New returns error with caller
func Newff(msg string, fields Fields, a ...interface{}) error {
	return wrap(nil, fields, msg, a...)
}

// Wrap returns error wrapped by caller info if 'err' is not nil
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return wrap(err, Fields{}, "")
}

// Wrapf returns error wrapped by message and caller info if 'err' is not nil
func Wrapf(err error, msg string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrap(err, Fields{}, msg, a...)
}

// Wrapf returns error wrapped by message and caller info if 'err' is not nil
func Wrapff(err error, fields Fields, msg string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	return wrap(err, fields, msg, a...)
}
