# Build stage
FROM golang:1.26.5-alpine AS builder

WORKDIR /app

# Download dependencies first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o ssh-portfolio .


# Runtime stage
FROM alpine:3.22

WORKDIR /app

# Create a non-root user
RUN adduser -D portfolio

# Copy compiled application
COPY --from=builder /app/ssh-portfolio .

# Run as non-root user
USER portfolio

# Start the portfolio
ENTRYPOINT ["./ssh-portfolio"]