# -------- Stage 1: Build --------
    FROM golang:1.23-alpine AS builder

    # Install git (needed for go mod in some cases)
    RUN apk add --no-cache git
    
    WORKDIR /app
    
    # Leverage Docker cache
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copy source code
    COPY . .
    
    # Build static binary
    RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o nrb-mvs-email-service .
    
    # -------- Stage 2: Run --------
    FROM alpine:latest
    
    # Install CA certs
    RUN apk --no-cache add ca-certificates
    
    WORKDIR /root/
    
    # Copy the built binary
    COPY --from=builder /app/nrb-mvs-email-service .
    
    EXPOSE 9015
    
    CMD ["./nrb-mvs-nrb-mvs-email-service"]
    