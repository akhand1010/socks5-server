FROM golang:1.13-alpine as builder

WORKDIR /go/src/github.com/vlakam/socks5-server
RUN apk add --no-cache git
COPY . .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s' -o ./server

FROM scratch
COPY --from=builder /go/src/github.com/vlakam/socks5-server/server /
ENTRYPOINT ["/server"]