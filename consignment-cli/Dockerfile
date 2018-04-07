FROM golang:1.9.5 as builder
WORKDIR /go/src/github.com/infoslack/go-microservice/consignment-cli
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN GCO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk add --update ca-certificates
WORKDIR /app
COPY consignment.json .
COPY --from=builder /go/src/github.com/infoslack/go-microservice/consignment-cli/consignment-cli .
CMD ["./consignment-cli"]
