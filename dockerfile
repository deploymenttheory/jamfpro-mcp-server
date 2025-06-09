# Build stage
FROM golang:1.21-alpine AS builder

# Install git and ca-certificates (needed for go modules and HTTPS)
RUN apk add --no-cache git ca-certificates

# Set the working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o jamfpro-mcp-server ./cmd/jamfpro-mcp-server

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create a non-root user
RUN addgroup -g 1001 jamfpro && \
    adduser -D -s /bin/sh -u 1001 -G jamfpro jamfpro

# Set the working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/jamfpro-mcp-server .

# Copy any additional files needed
COPY --from=builder /app/jamfpro-mcp-server-config.json* ./

# Change ownership to non-root user
RUN chown -R jamfpro:jamfpro /app

# Switch to non-root user
USER jamfpro

# Set the entrypoint
ENTRYPOINT ["./jamfpro-mcp-server"]

# Default command (can be overridden)
CMD ["stdio"]

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD ./jamfpro-mcp-server --help > /dev/null || exit 1

# Labels
LABEL org.opencontainers.image.title="Jamf Pro MCP Server"
LABEL org.opencontainers.image.description="A Model Context Protocol server for Jamf Pro APIs"
LABEL org.opencontainers.image.vendor="Your Organization"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/deploymenttheory/jamfpro-mcp-server"