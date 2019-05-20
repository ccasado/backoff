package main

import (
	"fmt"
	"net/http"
	"time"
)

// backoff algorithm that gradually increase the rate at which retries
// are performed, thus avoiding network congestion
func main() {

	cliente := &http.Client{
		Timeout: time.Second * 100,
	}

	retries := 0
	retry := true
	maxRetries := 100

	for retry && retries < maxRetries {
		time.Sleep(time.Duration(retries) * time.Millisecond)
		status, err := cliente.Get("https://www.google.com.br")
		if err != nil {
			retry = true
			retries = retries + 1
			fmt.Println("[main] Request NOTOK. Retries:", retries, err.Error())
		} else {
			retry = false
			fmt.Println("[main] Request OK")
			defer status.Body.Close()
		}
	}
}
