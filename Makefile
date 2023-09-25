# The -race flag detects unsynchronised read/writes to the same memory and reports them as a failure.
GO_TEST_FLAGS=-race

# Run the proxy server
.PHONY: run
run: waas-proxy-server
	./waas-proxy-server

.PHONY: docker/build
docker/build:
	docker build -t waas-proxy-server .

.PHONY: docker/run
docker/run: docker/build
	docker run -p 8091:8091 waas-proxy-server

# Build the proxy server
waas-proxy-server: longrunning
	go build -o waas-proxy-server ./cmd/proxyserver/*.go

# Run unit tests
.PHONY: test
test:
	go test ${GO_TEST_FLAGS} ./...

# Run protocol service proxy unit tests
.PHONY: protocols/test
protocols/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/protocols/

# Run pool service proxy unit tests
.PHONY: pools/test
pools/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/pools/

# Run mpc transaction service proxy unit tests
.PHONY: mpctransactions/test
mpctransactions/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/mpctransactions/

# Run mpc key service proxy unit tests
.PHONY: mpckeys/test
mpckeys/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/mpckeys/

# Run mpc wallet service proxy unit tests
.PHONY: mpcwallets/test
mpcwallets/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/mpcwallets/

# Run blockchain service proxy unit tests
.PHONY: blockchain/test
blockchain/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/blockchain/

# Run operations service proxy unit tests
.PHONY: operations/test
operations/test:
	go test ${GO_TEST_FLAGS} ./internal/proxyserver/operations/

# Install grpc-gateway
.PHONY: grpc-gateway
grpc-gateway:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.0 \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.0

# Generate longrunning.pb.gw.go from longrunning.proto
longrunning: googleapis
	mkdir -p ./gen/go ;
	protoc -I .googleapis --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt standalone=true \
    .googleapis/google/longrunning/operations.proto;

# Download longrunning.proto
googleapis:
	@if [ ! -d .googleapis ]; then \
		echo "Downloading Google APIs..." ; \
		mkdir -p .proto ; \
		git clone --depth 1 --single-branch --no-tags https://github.com/googleapis/googleapis.git .googleapis ; \
	else \
		echo "Updating Google APIs..." ; \
		cd .googleapis && git pull -q ; \
	fi
