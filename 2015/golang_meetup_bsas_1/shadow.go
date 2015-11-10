package main

import (
	"errors"
	"fmt"
)

// START OMIT
func main() {
	list := []string{"a", "b", "c"}
	for {
		list, err := repeat(list)
		if err != nil {
			panic(err)
		}
		fmt.Println(list)
		break
	}
	fmt.Println(list)
}

func repeat(list []string) ([]string, error) {
	if len(list) == 0 {
		return nil, errors.New("Nothing to repeat!")
	}
	list = append(list, list...)
	return list, nil
}

// END OMIT
