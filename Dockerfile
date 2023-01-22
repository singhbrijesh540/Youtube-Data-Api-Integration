FROM golang:1.18.2

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download


COPY . ./

RUN go run main.go -b 0.0.0.0

EXPOSE 8081