# ---------- builder stage ----------
FROM golang:1.21-alpine AS builder

WORKDIR /build

# Copy dependency files first for layer-cache efficiency.
COPY go.mod ./

# Download modules (none in this case — stdlib only).
RUN go mod download

# Copy source.
COPY . .

# Build a fully static binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o time-service .

# ---------- final stage ----------
FROM scratch

# Copy CA certificates so HTTPS calls can be made if needed in future.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the static binary.
COPY --from=builder /build/time-service /time-service

EXPOSE 9090

ENTRYPOINT ["/time-service"]
