FROM golang:1.19-alpine

WORKDIR /go/src/github.com/tanapon395/sa-66-example
COPY backend .
RUN apk add build-base
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]