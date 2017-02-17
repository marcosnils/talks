package main

import (
	"fmt"
	"net/http"
	"strconv"
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

// START OMIT
var concurrency = 10
var c = make(chan int, concurrency)
var start = 4521105664

func main() {

	for i := 0; i < concurrency; i++ {
		go worker()
	}
	for i := start; i < start+5000; i++ {
		c <- i
	}
}

func worker() {
	for {
		question_id := <-c
		res, _ := http.Get("https://api.mercadolibre.com/questions/" + strconv.Itoa(question_id))
		if res.StatusCode == 200 {
			fmt.Printf("Question %d exists\n", question_id)
		}
		res.Body.Close()
	}
}

//END OMIT
