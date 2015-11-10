package main

import "fmt"

type Foo struct {
	bar string
}

func main() {
	list := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}

	list2 := make([]*Foo, len(list))
	for i, value := range list {
		list2[i] = &value
	}

	fmt.Println(list[0], list[1], list[2])
	fmt.Println(list2[0], list2[1], list2[2])
}
