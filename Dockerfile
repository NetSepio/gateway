# Stage 1: Build the Go application
FROM golang:bookworm as builder
WORKDIR /app

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application source code
COPY . .

# Build the Go application
RUN go build -o gateway .

# Stage 2: Prepare the Aptos CLI
FROM ubuntu:22.04 as aptos_builder
WORKDIR /app

# Install dependencies and download the Aptos CLI
RUN apt update -y && \
    apt install -y wget unzip && \
    wget https://github.com/aptos-labs/aptos-core/releases/download/aptos-cli-v2.3.0/aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && \
    unzip aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && \
    rm -rf aptos-cli-2.3.0-Ubuntu-22.04-x86_64.zip && \
    apt remove -y wget unzip

# Stage 3: Final image
FROM ubuntu:22.04
WORKDIR /app

# Copy the built Go application and Aptos CLI from previous stages
COPY --from=builder /app/gateway .
COPY --from=aptos_builder /app/aptos .

# Copy the startup script
COPY ./docker-start.sh .

# Set environment variables
ARG version
ENV VERSION=$version

# Set the default command
CMD [ "bash", "docker-start.sh" ]