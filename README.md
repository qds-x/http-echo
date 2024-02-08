# http-echo

Super simple echo request properties, headers etc.

## Actions
- build `go build -o out cmd/`
- build sans container, standalone executable `go build -o http-echo cmd/main.go`
- build container `docker build . -t http-echo:latest`
- push image `docker push wherever/http-echo:latest` 
- helm install `helm install http-echo deployments/helm/http-echo` 
  
## Run
```
./main
```
