# Build stage
FROM --platform=linux/amd64 golang@sha256:68dfce93aabedced2731fb2f799ab7c4b7191131e76317a6a0293eb8ffc861d2 AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s \
    -X main.Version=$(git describe --tags --always --dirty) \
    -X main.BuildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o /app/qlik-sudoku-puzzle 

# Run tests
RUN go test ./... -v

# ############################
# # STEP 2: create image
# ############################
# alpine:3.22.1 using digest
FROM --platform=linux/amd64 alpine@sha256:eafc1edb577d2e9b458664a15f23ea1c370214193226069eb22921169fc7e43f AS production

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates tzdata

# Create non-root user
RUN addgroup -g 1000 appuser && \
    adduser -D -u 1000 -G appuser appuser

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/qlik-sudoku-puzzle .

# Change ownership
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Set environment variables
ENV APP_ENVIRONMENT=production \
    LOG_LEVEL=info \
    LOG_FORMAT=json

# Default command
ENTRYPOINT ["/app/qlik-sudoku-puzzle"]

# Default puzzle (can be overridden)
CMD ["-input=5,3,0,0,7,0,0,0,0,6,0,0,1,9,5,0,0,0,0,9,8,0,0,0,0,6,0,8,0,0,0,6,0,0,0,3,4,0,0,8,0,3,0,0,1,7,0,0,0,2,0,0,0,6,0,6,0,0,0,0,2,8,0,0,0,0,4,1,9,0,0,5,0,0,0,0,8,0,0,7,9"]