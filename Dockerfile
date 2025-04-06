# Build stage
FROM golang:1.22-alpine AS builder

# Set working directory
WORKDIR /app

# Copy source code
COPY . .

# Build the application
RUN go build -o backend .

# Final stage
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy only the built binary from builder stage
COPY --from=builder /app/backend .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./backend"]

# Builder stage: Contains Go compiler, build tools, source code  

# Final stage: Only contains the compiled binary 






