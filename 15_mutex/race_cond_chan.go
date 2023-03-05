package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	var wg sync.WaitGroup
	c := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)i
		increment(&wg, c)
	const}

	wg.Wait()
	fmt.Println(x)
}

func increment(wg *sync.WaitGroup, c chan bool) {
	c <- true
	x++
	<-c
	wg.Done()
}
