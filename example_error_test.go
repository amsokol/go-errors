package errors

import (
	"fmt"
)

func ExampleError() {
	err := Error("error message")
	fmt.Println(err)
	// Output: error message
}

func ExampleErrorf() {
	err := Errorf("error message: %s", "reason")
	fmt.Println(err)
	// Output: error message: reason
}
