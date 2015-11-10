package main

import (
	"fmt"
	"time"
)

func main() {

	// START OMIT
L:
	for {
		select {
		case <-time.After(time.Second):
			fmt.Println("hello")
		default:
			break L
		}
	}
	// END OMIT

	fmt.Println("ending")
}
