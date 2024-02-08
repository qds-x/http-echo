package main

import (
    "net/http"
	"qds-x.io/http-echo/internal/server"
    "gopkg.in/yaml.v2"
    "fmt"
    "log"
    "io/ioutil"
)

// Field names must be capitalized for visibility to external packages e.g. yaml
type Config struct {
    Server struct {
        Tls struct {
            Enabled bool `yaml:"enabled"`
            Port int `yaml:"port"`
            Key string `yaml:"key"`
            Cert string `yaml:"cert"`

        } `yaml:"tls"`
        Port int `yaml:"http_port"`
    } `yaml:"server"`
}

func main() {
    var conf Config
    data, err := ioutil.ReadFile("/etc/http-echo/config.yml")
    if err != nil {
        log.Fatalf("error %v", err)
    }

    fmt.Printf("Loaded config:\n%s\n", string(data))
    err = yaml.Unmarshal(data, &conf)
    if err != nil {
        log.Fatalf("error: %v", err)
    }

    http.HandleFunc("/hello", server.Hello)
    http.HandleFunc("/", server.All)

    if conf.Server.Tls.Enabled {
        var port=fmt.Sprintf(":%d", conf.Server.Tls.Port)
        log.Printf("Listening on port %s", port)
        // certificate must be full chain including root CA
        err := http.ListenAndServeTLS(port, conf.Server.Tls.Cert, conf.Server.Tls.Key, nil)
        if err != nil {
            log.Fatalf("error: %v", err)
        }
    } else {
        var port=fmt.Sprintf(":%d", conf.Server.Port)
        log.Printf("Listening on port %s", port)
        http.ListenAndServe(port, nil)
    }
}
