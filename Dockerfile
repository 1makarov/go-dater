FROM golang:1.17 AS builder

COPY . /github.com/1makarov/go-dater/server
WORKDIR /github.com/1makarov/go-dater/server

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/1makarov/go-dater/server/.bin/app .

CMD ["./app"]