package main

import (
    "net/http"
	"qds-x.io/http-echo/internal/server"
)

func main() {

    http.HandleFunc("/hello", server.Hello)
    http.HandleFunc("/", server.All)

    http.ListenAndServe(":8080", nil)
}
