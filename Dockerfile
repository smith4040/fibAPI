# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Add Maintainer Info
LABEL maintainer="Chris Smith <smith4040@gmail.com>"

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Run test
RUN go test ./...

# Build the application
RUN go build -o fib-api .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/fib-api .

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /dist/fib-api /

# Command to run the executable
ENTRYPOINT ["/fib-api"]