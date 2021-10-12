package main

import (
	"log"
	"net/http"
	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

func main() {
	initCache()
	
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}
