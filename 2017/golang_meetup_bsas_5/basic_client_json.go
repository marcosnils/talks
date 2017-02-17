package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type MELIResponse struct {
	Title    string
	Price    float32
	Pictures []struct{ Url string }
}

func main() {
	// START OMIT

	var netTransport = &http.Transport{ // HL
		Dial: (&net.Dialer{ // HL
			Timeout: 5 * time.Second, // HL
		}).Dial, // HL
		TLSHandshakeTimeout: 5 * time.Second, // HL
	} // HL
	var netClient = &http.Client{ // HL
		Timeout:   time.Second * 10, // HL
		Transport: netTransport,     // HL
	} // HL

	res, err := netClient.Get("https://api.mercadolibre.com/items/MLA649184315")

	// END OMIT
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
