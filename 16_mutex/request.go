package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 30000; i++ {
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
			req = nil
		}()
	}

	time.Sleep(20 * time.Second)
}
