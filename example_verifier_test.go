package errors

import (
	"fmt"
	"math/rand"
)

func ExampleCheck() {
	// get fake invalid value
	r := rand.New(rand.NewSource(15))
	val := r.Intn(99)

	// check value
	var err error
	Check(&err, val > 100, "rand value must be more than 100")
	fmt.Println(err)
	// Output: rand value must be more than 100
}

func ExampleCheckf() {
	// get fake invalid value
	r := rand.New(rand.NewSource(15))
	val := r.Intn(99)

	// check value
	var err error
	Checkf(&err, val > 100, "rand value must be more than %d", 100)
	fmt.Println(err)
	// Output: rand value must be more than 100
}
