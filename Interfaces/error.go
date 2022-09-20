package main

import (
	"fmt"
)

type errFoo struct {
	err  error
	path string
}

func (e errFoo) Error() string {
	return fmt.Sprintf("%s: %s", e.path, e.err)
}

func XYZ(a int) error {
	return nil
}

func main() {
	var err error = XYZ(1)

	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("OK!")
	}
}
