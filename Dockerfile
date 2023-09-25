# Start with the Go 1.21 image
FROM golang:1.21

# Set the PATH to include $GOBIN
ENV PATH="/go/bin:${PATH}"

# Set the working directory inside the container
WORKDIR /app

# Copy the rest of the project files to the working directory
COPY . .

# Install protoc
RUN apt-get update && apt-get install -y protobuf-compiler

# Install grpc-gateway
RUN make grpc-gateway

# Build the Go application
RUN make waas-proxy-server

# Expose port 8081
EXPOSE 8091

# Run the proxy-server binary
CMD ["./waas-proxy-server"]
