FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . /app/

RUN go build -o http-echo ./cmd/main.go

FROM alpine

COPY --from=build /app/http-echo /usr/local/bin
COPY --from=build /app/config/config.yml /etc/http-echo/config.yml

EXPOSE 8080 8443
ENTRYPOINT ["http-echo"]