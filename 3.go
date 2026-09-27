package main

import (
	"fmt"
	"time"
)

func main() {
	var requestCount = 15
	request := make(chan int, requestCount)

	for i := 1; i <= requestCount; i++ {
		request <- i
	}
	close(request)

	tick := time.Tick(200 * time.Millisecond)

	for requestPerSecond := range request {
		<-tick

		fmt.Printf("[%s] обработан запрос №%d\n", time.Now(), requestPerSecond)
	}
}
