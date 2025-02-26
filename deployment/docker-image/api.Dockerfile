FROM golang:1.23.2-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download && go mod verify

COPY backend/cmd ./cmd
COPY backend/data ./data
COPY backend/internal ./internal

RUN go mod tidy

RUN go build -o main cmd/main.go

RUN chmod +x main

CMD ["./main"]