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
	h := NewHandler()

	http.HandleFunc("/add", h.autoIncrement)
	http.HandleFunc("/reset", reset)
	http.HandleFunc("/status", status)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}

type Handler struct {
	mutex sync.Mutex
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) autoIncrement(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	h.mutex.Lock()
	x++
	fmt.Printf("Add value to %d \n", x)
	fmt.Printf("runtime : %v\n", time.Since(t))
	h.mutex.Unlock()
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
