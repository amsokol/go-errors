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
