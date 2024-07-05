# Stage 1: Use Ubuntu to install Wasmtime
FROM ubuntu:latest AS wasmtime-installer

# Install necessary tools
RUN apt-get update && \
    apt-get install -y \
    curl \
    gnupg \
    xz-utils \ 
    && rm -rf /var/lib/apt/lists/*

# Install Wasmtime
RUN curl https://wasmtime.dev/install.sh -sSf | bash

# Stage 2: Use the official Go image and copy Wasmtime from the previous stage
FROM golang:latest

WORKDIR /app

# Copy the Go modules files
COPY . .

# Copy Wasmtime from the first stage
COPY --from=wasmtime-installer /root/.wasmtime /root/.wasmtime

# Set the PATH environment variable
ENV PATH="/root/.wasmtime/bin:${PATH}"


# Example command to show Wasmtime installation
CMD ["go", "run", "main.go"]

EXPOSE 8080