package main

import (
	"encoding/json"
	"fmt"
)

type MELIResponse struct {
	Title    string
	Price    float32
	Pictures []struct{ Url string }
}

func main() {
	res, err := netClient.Get("https://api.mercadolibre.com/items/MLA649184315")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var resp MELIResponse
	_ = json.NewDecoder(res.Body).Decode(&resp)

	fmt.Println(resp.Title)
	fmt.Println(resp.Price)
	fmt.Println(resp.Pictures)

}
