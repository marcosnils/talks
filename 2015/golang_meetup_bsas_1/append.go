package main

import "fmt"

func main() {
	a := []byte("foo")
	b := append(a, []byte("bar")...)
	c := append(a, []byte("baz")...)

	fmt.Println(string(a), string(b), string(c))
}
