package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		value, err := getData(client, "mykey")
		var counter int
		if err != nil {
			putData(client, "mykey", strconv.Itoa(1))
			counter = 1
		} else {
			counter, _ = strconv.Atoi(value)
			counter++
			putData(client, "mykey", strconv.Itoa(counter))
		}
		fmt.Fprintf(w, "counter = %d\n", counter)
		fmt.Fprintf(w, "Hello this is server 1, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8080", nil)
}

func putData(client *redis.Client, key string, value string) error {
	// Set
	err := client.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("redis.Client.Set Error:", err)
		return err
	}
	return nil
}

func getData(client *redis.Client, key string) (string, error) {
	// Set
	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println("redis.Client.Get Error:", err)
		return "", err
	}

	return val, nil
}
