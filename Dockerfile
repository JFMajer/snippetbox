# Build stage
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.* ./

# Download Go modules
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o snippetbox ./cmd/web


# Final stage
FROM scratch

WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/snippetbox .

# Copy static files and templates (adjust paths according to your project structure)
COPY --from=builder /app/ui ./ui

# Copy /etc/passwd to the container
COPY --from=builder /etc/passwd /etc/passwd

# Run as a non-root user
USER 1001

# Expose the port the app runs on
EXPOSE 4000

# Run the binary
CMD ["./snippetbox"]