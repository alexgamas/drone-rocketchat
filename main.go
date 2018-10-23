package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	POST           = "POST"
	baseURL        = "https://localhost:3000"
	loginURL       = "api/v1/login"
	postMessageURL = "api/v1/chat.postMessage"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginTokenResponse struct {
	UserID    string `json:"userId"`
	AuthToken string `json:"authToken"`
}

type LoginResponse struct {
	Status string             `json:"status"`
	Data   LoginTokenResponse `json:"data"`
}

type MessageRequest struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	Avatar string  `json:"avatar"`
}

func newRequest(method, url string, headers map[string]string, body interface{}) *http.Request {
	fmt.Println(url)

	fmt.Printf("Request Headers:\n")

	//body io.Reader
	payload, err := json.Marshal(body)

	fmt.Println("Payload", string(payload))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	for k, v := range headers {
		fmt.Printf(" + %s: %s\n", k, v)
		req.Header.Add(k, v)
	}

	if err != nil {
		fmt.Printf("%s", err)
		return nil
	}

	return req
}

func main() {

	debug := true

	loginRequestObj := LoginRequest{Username: "user", Password: "test"}

	headers := map[string]string{"Content-Type": "application/json"}

	url := fmt.Sprintf("%s/%s", baseURL, loginURL)

	loginRequest := newRequest(POST, url, headers, loginRequestObj)

	var client = http.DefaultClient

	response, err := client.Do(loginRequest)

	if err != nil {
		log.Println("FATAL: Error: ", err)
		os.Exit(1)
	}

	defer response.Body.Close()

	if debug {
		fmt.Println("Response Status:", response.Status)
	}

	if response.StatusCode == 200 {

		if debug {
			fmt.Printf("Response Headers:\n")
			for k, v := range response.Header {
				fmt.Printf(" + %s: %s\n", k, v)
			}
		}

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Println("FATAL: Error: ", err)
			os.Exit(1)
		}

		var loginResponse LoginResponse

		err = json.Unmarshal(body, &loginResponse)

		if err != nil {
			log.Println("FATAL: Error: ", err)
			os.Exit(1)
		}

		url := fmt.Sprintf("%s/%s", baseURL, postMessageURL)

		// headers := map[string]string{
		// 	"Content-Type": "application/json",
		// 	"X-User-Id":    loginResponse.Data.UserID,
		// 	"X-Auth-Token": loginResponse.Data.AuthToken}

		headers := map[string]string{
			"Content-Type": "application/json",
			"X-User-Id":    loginResponse.Data.UserID,
			"X-Auth-Token": loginResponse.Data.AuthToken}

		messageRequestObj := MessageRequest{Channel: "ci", Text: "Teste de mensagem", Avatar: "http://res.guggy.com/logo_128.png"}

		messageRequest := newRequest(POST, url, headers, messageRequestObj)

		var client = http.DefaultClient
		response, err := client.Do(messageRequest)

		if err != nil {
			log.Println("FATAL: Error: ", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		body2, err := ioutil.ReadAll(response.Body)

		if err != nil {
			log.Println("FATAL: Error: ", err)
			os.Exit(1)
		}

		log.Println("response.StatusCode: ", response.StatusCode)
		log.Println("response.StatusCode: ", string(body2))

	}

}
