package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func foo() error {
	err := errors.New("foo")
	return err
}

func bar() error {
	err := foo()
	return errors.Wrap(err, "bar")
}

func baz() error {
	err := bar()
	return errors.Wrap(err, "baz")
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func main() {
	err, ok := baz().(stackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}
	st := err.StackTrace()
	fmt.Printf("%+v", st[:]) // top two frames
}
