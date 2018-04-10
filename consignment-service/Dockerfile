FROM golang:1.9.5 as builder
WORKDIR /go/src/github.com/infoslack/go-microservice/consignment-service
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
RUN apk add --update ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/infoslack/go-microservice/consignment-service/consignment-service .
CMD ["./consignment-service"]
