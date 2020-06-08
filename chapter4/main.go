package main

import (
	"fmt"
	"time"
)

func timer(duration int) {
	fmt.Println("start")

	// done := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)

		// done <- "result ok"
	}()

	<-time.After(time.Duration(duration) * time.Second)
	fmt.Println("timeout")
}

func main() {
	timer(2)
}
