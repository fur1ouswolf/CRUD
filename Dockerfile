# Build
FROM golang:latest AS build
WORKDIR /build

# Install dependencies
COPY go.* .
RUN go mod download

# Build the binary
ENV CGO_ENABLED=0
ENV GIN_MODE=release
RUN --mount=target=. \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /app ./cmd/app/main.go

# Create environment
COPY .env /

# Expose application port
EXPOSE ${APP_PORT}

# Run the binary
ENTRYPOINT ["/app", "--config=/app.env"]