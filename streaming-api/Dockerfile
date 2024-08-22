FROM golang:1.20 AS builder

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./main.go

FROM alpine:latest

COPY --from=builder ./main .

EXPOSE 8000

CMD ["./main"]
