
FROM golang:1.24-alpine AS builder

WORKDIR /src/Server-Side

COPY Server-Side/go.mod Server-Side/go.sum ./
COPY Server-Side/ ./

RUN go mod download


COPY Server-Side ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .


FROM alpine:latest
RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /server .
COPY .env .env
COPY honors-client-side/build ./static

EXPOSE 8080

ENTRYPOINT ["./server"]