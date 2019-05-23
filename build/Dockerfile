## BUilder
FROM  golang:1.11 as builder
LABEL maintainer="Luis Aguilar <luisaguilar@edlio.com>"
RUN mkdir -p /go/src/github.com/luisaguilar2910/go-server
WORKDIR /go/src/github.com/luisaguilar2910/go-server
COPY . .
RUN go get -d -v ./cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main ./cmd
RUN ls /go/bin/main

## Image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/bin/main .
EXPOSE 8080
CMD ["./main"]