package server 

import (
    "fmt"
    "net/http"
	"io/ioutil"
    "sort"
)

func Hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func All(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Host: %v\n", req.Host)
	fmt.Fprintf(w, "URL: %v\n\n", req.URL)

	var headerNames []string
	for name := range req.Header {
		headerNames = append(headerNames, name)
	}

    sort.Strings(headerNames)
    fmt.Fprintf(w, "Headers:\n")
	for _, name := range headerNames {
        headers := req.Header[name]
        for _, h := range headers {
            fmt.Fprintf(w, "    %v: %v\n", name, h)
        }
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "\nBody:\n%s", b)
}
