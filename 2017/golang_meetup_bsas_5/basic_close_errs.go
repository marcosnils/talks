package main

import "net/http"

func main() {
	res, err := http.Get("https://www.frutacosmica.com")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close() // HL

}
