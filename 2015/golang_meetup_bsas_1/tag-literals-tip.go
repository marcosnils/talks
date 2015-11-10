package main

import "fmt"

// START OMIT
type T struct {
	Foo string
	Bar int
}

func main() {
	t := T{Foo: "example", Bar: 123} // untagged literal
	fmt.Printf("t %+v\n", t)
}

// END OMIT
