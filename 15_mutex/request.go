package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			req, err := http.NewRequest(http.MethodGet, "http://localhost:8000/add", nil)
			if err != nil {
				fmt.Println(err.Error())
			}
			client := http.Client{}
			_, err = client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
