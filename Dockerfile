FROM golang:1.20

WORKDIR /app

COPY . .

RUN go build ./cmd/rick/main.go

EXPOSE 8080

ENTRYPOINT ["/app/main"]
