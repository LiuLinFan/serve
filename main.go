package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("hello")
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe(":8080")
}
