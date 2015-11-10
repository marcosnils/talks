package main

import "fmt"

// START OMIT
type T struct {
	Foo string
	Bar int
}

func main() {
	t := T{"example", 123} // untagged literal
	fmt.Printf("t %+v\n", t)
}

// END OMIT
