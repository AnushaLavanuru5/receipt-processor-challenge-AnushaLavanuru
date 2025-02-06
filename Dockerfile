FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy && go mod verify && go mod download

COPY . .

RUN go test ./...
RUN go build -o receipt-processor ./main.go

EXPOSE 8080

CMD ["./receipt-processor"]