package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"time"
)

type ModerationData struct {
	Text string `json:"text"`
}

type MTKReq struct {
	Data     ModerationData `json: "data"`
	Criteria []string       `json: "criteria"`
}

type Result struct {
}

var concurrency = 10
var c = make(chan int, concurrency)
var start = 4521105664

func main() {

	// START 1 OMIT
	transport := &http.Transport{ // HL
		Proxy: http.ProxyFromEnvironment, // HL
		DialContext: (&net.Dialer{ // HL
			Timeout:   30 * time.Second, // HL
			KeepAlive: 30 * time.Second, // HL
		}).DialContext, // HL
		MaxIdleConns:        100, // HL
		MaxIdleConnsPerHost: 10,  // HL
	} // HL
	// END 1 OMIT

	http.DefaultTransport = transport

	for i := 0; i < concurrency; i++ {
		go worker()
	}

	for i := start; i < start+5000; i++ {
		c <- i
	}
}

// START 2 OMIT
func worker() {
	for {
		question_id := <-c
		res, _ := http.Get("https://api.mercadolibre.com/questions/" + strconv.Itoa(question_id))
		if res.StatusCode == 200 {
			fmt.Printf("Question %d exists\n", question_id)
		}
		io.Copy(ioutil.Discard, res.Body) // HL
		res.Body.Close()
	}
}

// END 2 OMIT
