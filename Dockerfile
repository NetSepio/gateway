# Stage 1: Prepare the Go environment
FROM golang:bookworm AS builder
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o gateway cmd/main.go cmd/server.go

# Stage 2: Prepare the Aptos CLI
FROM ubuntu:22.04 AS aptos_builder
WORKDIR /app

# Install dependencies and download the Aptos CLI
RUN apt update -y && \
    apt install -y wget unzip && \
    wget https://github.com/aptos-labs/aptos-core/releases/download/aptos-cli-v2.3.0/aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && \
    unzip aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && \
    rm -rf /app/aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && apt remove -y wget unzip

# Stage 3: Final image
FROM ubuntu:22.04
WORKDIR /app

# ✅ Install CA certificates to fix TLS issues
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the source code and Aptos CLI from previous stages
COPY --from=builder /app/gateway .
COPY --from=aptos_builder /app/aptos .
COPY ./docker-start.sh .

# Set environment variables
ARG version
ENV VERSION=$version

# Set the default command to run the Go application
CMD [ "bash", "docker-start.sh" ]
