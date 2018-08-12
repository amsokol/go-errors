# go-errors

[![GitHub release](https://img.shields.io/github/release/amsokol/go-errors.svg)](https://github.com/amsokol/go-errors/releases)
[![GitHub issues](https://img.shields.io/github/issues/amsokol/go-errors.svg)](https://github.com/amsokol/go-errors/issues)
[![GitHub issues closed](https://img.shields.io/github/issues-closed/amsokol/go-errors.svg)](https://github.com/amsokol/go-errors/issues)
[![Travis branch](https://img.shields.io/travis/amsokol/go-errors/master.svg)](https://travis-ci.org/amsokol/go-errors)
[![Coverage Status](https://coveralls.io/repos/github/amsokol/go-errors/badge.svg?branch=master)](https://coveralls.io/github/amsokol/go-errors?branch=master)
[![Go Report Card](https://goreportcard.com/badge/amsokol/go-errors)](http://goreportcard.com/report/amsokol/go-errors)
[![license](https://img.shields.io/github/license/amsokol/go-errors.svg)](https://github.com/amsokol/go-errors/blob/master/LICENSE)

This library contains simple error handling primitives.

---------------------------------

## Getting started / Prerequisites / Dependencies

```shell
go get -u github.com/amsokol/go-errors
```

See usage [example](https://github.com/amsokol/go-errors/blob/master/example_test.go):

```go
package errors

import (
    "fmt"
)

func Function(val1 int, val2 string, val3 bool) (string, error) {
    var err error

    // check input params
    Check(&err, val1 > 0, "val1 must be greater than 0")
    Check(&err, len(val2) > 0, "val2 can't be empty")
    Check(&err, val3, "val3 should be 'true'")
    if err != nil {
        return "", err
    }

    return fmt.Sprintf("val1=%d, val2='%s', val3=%t", val1, val2, val3), nil
}

func Example() {
    err := Error("error message")
    fmt.Println(err)
    // Output: error message

    err = Errorf("error message: %s", "reason")
    fmt.Println(err)
    // Output: error message: reason

    err = Wrap(err, "new error message")
    fmt.Println(err)
    // Output: new error message-> error message: reason

    err = Wrapf(err, "one more error message: %s", "another reason")
    fmt.Println(err)
    // Output: "one more error message: another reason-> new error message-> error message: reason

    // val1 is invalid
    _, err = Function(0, "text", true)
    if err != nil {
        fmt.Println(err)
        // Output: al1 must be greater than 0
    }

    // val2 is invalid
    _, err = Function(10, "", true)
    if err != nil {
        fmt.Println(err)
        // Output: val2 can't be empty
    }

    // val3 is invalid
    _, err = Function(10, "text", false)
    if err != nil {
        fmt.Println(err)
        // Output: val3 should be 'true'
    }

    // all values are valid
    res, err := Function(10, "text", true)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(res)
}
```

### Testing

```shell
go test ./...
```
