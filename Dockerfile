FROM golang:1.25.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -a -o YOUR_APP_NAME main.go

FROM alpine:3.22.2
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /
COPY --from=builder /app/YOUR_APP_NAME .
USER nobody
ENTRYPOINT ["/YOUR_APP_NAME"]