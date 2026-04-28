# Overview

A lightweight HTTP service that returns the current UTC time in a standard format when requested.

# Personas

- Developer — calls the service to get consistent UTC timestamps.

# Features

- A developer sends a GET request to /time and receives the current UTC time in RFC3339 format.
- The response is JSON with a single "now" field containing the timestamp.
- The service runs as a standalone HTTP server using Go's standard library.
- The service listens on a configurable port.
- Health checks can verify the service is running by calling the endpoint.