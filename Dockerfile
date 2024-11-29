FROM golang:1.23 AS build-golang

WORKDIR /go/src/app
COPY . .
RUN go get -v
RUN go install github.com/air-verse/air@v1.52.3
EXPOSE 8080

CMD ["air"]
