# http-echo

Super simple echo request properties, headers etc.

## Build
- sans container `go build cmd/main.go`
- container `docker build . -t http-echo:latest`
- push image `docker push wherever/http-echo:latest` 
- helm install `helm install http-echo deployments/helm/http-echo` 
  
## Run
```
./main
```
