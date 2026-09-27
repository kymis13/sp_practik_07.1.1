package main

import (
	"fmt"
	"sync"
	"time"
)

func count(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Print(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go count(&wg)
	wg.Wait()
}
