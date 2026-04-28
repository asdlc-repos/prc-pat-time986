# time-service

A lightweight Go HTTP service that returns the current UTC time in RFC3339 format.

## Endpoints

| Method | Path      | Description                             |
|--------|-----------|-----------------------------------------|
| GET    | `/time`   | Returns current UTC time as RFC3339 JSON |
| GET    | `/health` | Liveness check — returns `{"status":"ok"}` |

### Example responses

```bash
# /time
curl http://localhost:9090/time
{"now":"2024-01-15T14:30:45Z"}

# /health
curl http://localhost:9090/health
{"status":"ok"}
```

## Configuration

| Environment Variable | Default | Description         |
|----------------------|---------|---------------------|
| `PORT`               | `9090`  | HTTP listening port |

## Build & run

```bash
# Build binary
cd time-service
go build -o time-service .

# Run
./time-service

# Override port
PORT=8080 ./time-service
```

## Docker

```bash
# Build image
docker build -t time-service .

# Run container
docker run -p 9090:9090 time-service

# Override port inside container
docker run -e PORT=8080 -p 8080:8080 time-service
```

The Dockerfile uses a multi-stage build: Go source is compiled with `golang:1.21-alpine` and the resulting static binary is copied into a `scratch` final image for a minimal footprint.
