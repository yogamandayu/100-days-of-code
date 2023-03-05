package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var x int

func main() {
	http.HandleFunc("/add", autoIncrement)
	http.HandleFunc("/reset", reset)
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

func autoIncrement(w http.ResponseWriter, r *http.Request) {
	var m sync.Mutex
	m.Lock()
	time.Sleep(100 * time.Millisecond)
	x++
	fmt.Println("Add value to", x)
	m.Unlock()
	w.WriteHeader(200)
	w.Write([]byte("Added!"))
	return
}

func reset(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Reset value")
	x = 0
	return
}

func status(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(strconv.Itoa(x)))
	return
}
