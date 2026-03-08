FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gouxchecker ./cmd/main.go

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
RUN adduser -D appuser
WORKDIR /home/appuser/

COPY --from=builder /app/gouxchecker .

USER appuser

ENTRYPOINT ["./gouxchecker"]
CMD []
