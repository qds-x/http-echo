package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func all(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Host: %v\n", req.Host)
	fmt.Fprintf(w, "URL: %v\n", req.URL)
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "Body:\n%s", b)
}

func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/", all)

    http.ListenAndServe(":8080", nil)
}