package main

import (
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.com") // HL

	if err != nil {
		panic(err)
	}
	fmt.Printf("Done %d\n", res.StatusCode)

	res, err = http.Get("https://www.facebook.com") // HL

	if err != nil {
		panic(err)
	}
	fmt.Printf("Done %d\n", res.StatusCode)
}
