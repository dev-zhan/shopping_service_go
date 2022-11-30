FROM golang:latest

RUN go version
ENV GOPATH=/

COPY . / ./

RUN go mod download

RUN go build -o shopping_service_go ./cmd/main.go
CMD ["./shopping_service_go"]