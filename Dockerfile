FROM golang:1.27.2 AS BUILDER
WORKDIR /app
RUN curl -sL https://taskfile.dev/install.sh | sh
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN /app/bin/task build

FROM alpine:3.24.3
RUN apk update && apk add --no-cache ca-certificates
WORKDIR /
COPY --from=BUILDER /app/dist/my-app .
USER nobody
ENTRYPOINT ["/my-app"]

# FROM gcr.io/distroless/static-debian13
# WORKDIR /
# COPY --from=BUILDER /app/dist/my-app .
# ENTRYPOINT ["/my-app"]
