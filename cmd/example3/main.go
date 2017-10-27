package main

import (
	"log"
	"io"

	"github.com/amsokol/go-errors"
)

func f11() error {
	return errors.Wrapff(f12(), errors.Fields{"b": "bbbb"}, "cause from f11")
}

func f12() error {
	return errors.Wrapff(f13(), errors.Fields{"a": "aa\"aa"}, "cause from f12")
}

func f13() error {
	return io.EOF
}

func main() {
	log.Print(f11())
}
