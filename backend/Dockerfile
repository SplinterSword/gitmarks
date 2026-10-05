# syntax=docker/dockerfile:1

# --- Builder: compile static Go binary ---
FROM golang:1.27-alpine AS builder
WORKDIR /src

# Cache deps first (better layer reuse)
COPY go.mod go.sum ./
RUN go mod download

# Copy source and build
COPY cmd/ ./cmd/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux go build \
  -trimpath \
  -ldflags="-s -w" \
  -o /out/server ./cmd/server

# --- Runner: minimal non-root image ---
FROM gcr.io/distroless/static-debian12:nonroot
WORKDIR /app
COPY --from=builder /out/server /app/server

EXPOSE 34179
ENV PORT=34179
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
