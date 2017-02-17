package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

// START 1 OMIT
type ModerationData struct {
	Text string `json:"text"`
}

type MTKReq struct {
	Data     ModerationData `json: "data"`
	Criteria []string       `json: "criteria"`
}

// END 1 OMIT
func main() {

	// START 2 OMIT
	input := MTKReq{
		Data:     ModerationData{Text: "El votante de Macri es 50% idiota y 50% hdp"},
		Criteria: []string{"hate_speech"},
	}

	body, err := json.Marshal(input)

	req, _ := http.NewRequest("POST", "https://api.sherloq.io/moderations", bytes.NewBuffer(body))
	req.SetBasicAuth("you need a valid token here", "")

	res, err := http.DefaultClient.Do(req)

	// END 2 OMIT
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	out, _ := ioutil.ReadAll(res.Body)
	var outIndent bytes.Buffer
	json.Indent(&outIndent, out, "=", "\t")

	outIndent.WriteTo(os.Stdout)
}
