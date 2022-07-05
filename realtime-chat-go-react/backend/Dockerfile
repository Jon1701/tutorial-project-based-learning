FROM golang:1.18.3-alpine3.16
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod download
RUN go build -o main ./...
CMD ["/app/main"]