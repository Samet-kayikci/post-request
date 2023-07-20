package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	url := "http://localhost:8080/connect"

	// generate JSON data
	data := Message{
		Text: "connection request",
	}

	// convert byte
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error json conversion:", err)
		return
	}

	// generate HTTP POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("error request:", err)
		return
	}
	defer resp.Body.Close()

	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("response read error:", err)
		return
	}

	// read response from server
	fmt.Println("response:", string(bodyData))
	fmt.Println("Status code:", resp.Status)
}
