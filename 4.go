package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org", "https://google.com", "https://github.com",
	}

	sem := make(chan struct{}, 3)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			resp, err := http.Get(u)
			if err != nil {
				fmt.Printf("%s -> Ошибка: %v\n", u, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("%s -> %s\n", u, resp.Status)
		}(url)
	}
	wg.Wait()
}
