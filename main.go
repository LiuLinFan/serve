package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		for {
			fmt.Println(time.Now())
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe(":8080", nil)
}
