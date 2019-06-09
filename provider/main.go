package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/", index)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	rand.Seed(time.Now().UnixNano())

	// 1 / 2の確率で20秒処理にかかる。
	time.Sleep(3 * time.Second)

	fmt.Fprint(res, "hello world.")
}
