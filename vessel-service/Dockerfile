FROM golang:1.9.5 as builder
WORKDIR /go/src/github.com/infoslack/go-microservice/vessel-service
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go

FROM alpine:latest
RUN apk add --update ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/infoslack/go-microservice/vessel-service/vessel-service .
CMD ["./vessel-service"]
