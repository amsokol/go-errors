package errors

import (
	"fmt"
)

func ExampleWrap() {
	err := Error("error message 1")
	err = Wrap(err, "error message 2")
	fmt.Println(err)
	// Output: error message 2-> error message 1
}

func ExampleWrapf() {
	err := Error("error message 1")
	err = Wrapf(err, "error message 2: %s", "reason")
	fmt.Println(err)
	// Output: error message 2: reason-> error message 1
}
