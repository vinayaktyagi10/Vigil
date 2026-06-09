FROM golang:1.23-alpine AS builder

ARG SERVICE_NAME

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-binary ./cmd/${SERVICE_NAME}/main.go

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /app-binary /app-binary

CMD ["/app-binary"]
