package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// START OMIT
type MELIResponse struct {
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Pictures []struct {
		Url string `json:"url"`
	} `json: "pictures"`
}

func main() {
	res, err := http.Get("https://api.mercadolibre.com/items/MLA649184315")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var resp MELIResponse
	_ = json.NewDecoder(res.Body).Decode(&resp) // HL

	fmt.Println(resp.Title)
	fmt.Println(resp.Price)
	fmt.Println(resp.Pictures)

}

// END OMIT
