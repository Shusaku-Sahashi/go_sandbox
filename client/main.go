package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	timeout := 2 * time.Second

	url := "http://localhost:8080"

	logfile, err := os.OpenFile("info.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logfile)
	log.SetPrefix("[info]:")
	log.SetFlags(log.Llongfile | log.Ltime)

	var result string
	var counter int

loop:
	for {
		res := sendRequest(url, timeout, counter)

		select {
		case result = <-res:
			break loop
		case <-time.After(timeout):
			if counter <= 3 {
				counter++
				log.Printf("retry %v times.", counter)
			} else {
				log.Println("timeout error")
				break
			}
		}
	}

	fmt.Print(result)
}

func sendRequest(url string, timeout time.Duration, count int) <-chan string {
	client := http.Client{
		Timeout: timeout,
	}

	resChan := make(chan string)
	timeoutError := make(chan struct{})

	log.Printf("%v: start request", count)

	go func() {
		log.Print("gorutine request")
		defer close(resChan)

		log.Print("request prepare")
		var err error
		res, err := client.Get(url)
		if err != nil {
			log.Printf("%v: HTTP Request Error: %v", count, err)
			timeoutError <- struct{}{}
			return
		}
		defer res.Body.Close()

		var buf bytes.Buffer
		io.Copy(&buf, res.Body)
		if res.StatusCode != 200 {
			// bodyを文字絵列で取得するには、ioutilやbyteBufferを使用する必要あり。
			log.Printf("%v: HTTP status: %v, error: %v", count, res.StatusCode, buf.String())
			return
		}

		resChan <- buf.String()
	}()

	log.Printf("%v: start select", count)

	select {
	case <-timeoutError:
		log.Printf("%v: timeout", count)
	}

	log.Printf("%v: send request", count)
	return resChan
}
