package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name   string
	Status string
}

var (
	user = &User{
		Name:   "John",
		Status: "INACTIVE",
	}
)

func main() {

	r := GetRedisConn()

	handler := NewHandler(r)

	http.HandleFunc("/get", handler.Get)
	http.HandleFunc("/active", handler.SetActive)
	http.HandleFunc("/inactive", handler.SetInactive)
	http.HandleFunc("/invalidate", handler.Invalidate)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}

}

type Handler struct {
	redis *redis.Client
}

func NewHandler(r *redis.Client) *Handler {
	return &Handler{
		redis: r,
	}
}

func (h *Handler) SetActive(w http.ResponseWriter, r *http.Request) {
	user.Status = "ACTIVE"

	w.Write([]byte("Updated!"))
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) SetInactive(w http.ResponseWriter, r *http.Request) {
	user.Status = "INACTIVE"

	w.Write([]byte("Updated!"))
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) Invalidate(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	h.redis.Del(ctx, "user")
	w.Write([]byte("Redis Invalidate"))
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	v, err := h.redis.Get(ctx, "user").Result()
	if err == nil {
		w.Write([]byte(v))
		w.WriteHeader(http.StatusOK)
		return
	}

	b, _ := json.Marshal(user)
	h.redis.Set(ctx, "user", string(b), time.Minute*15)
	w.Write(b)
	w.WriteHeader(http.StatusOK)
	return
}

func GetRedisConn() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return conn
}
