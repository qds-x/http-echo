FROM golang:1.16-alpine AS build

WORKDIR /app

COPY . /app/

RUN go build ./cmd/http-echo

FROM alpine

COPY --from=build /app/http-echo /usr/local/bin

EXPOSE 8080
ENTRYPOINT ["http-echo"]