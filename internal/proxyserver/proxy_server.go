package proxyserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"time"

	longrunninggw "github.com/coinbase/waas-proxy-server/gen/go/google/longrunning"
	"github.com/coinbase/waas-proxy-server/internal/blockchain"
	"github.com/coinbase/waas-proxy-server/internal/mpckeys"
	"github.com/coinbase/waas-proxy-server/internal/mpctransactions"
	"github.com/coinbase/waas-proxy-server/internal/mpcwallets"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	"github.com/coinbase/waas-proxy-server/internal/pools"
	"github.com/coinbase/waas-proxy-server/internal/protocols"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-client-library-go/auth"
	"github.com/coinbase/waas-client-library-go/clients"
	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	mpckeyspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	mpctransactionspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

const (
	// defaultGRPCAddress is the default address for the gRPC server.
	defaultGRPCAddress = "localhost:8080"

	// defaultHTTPAddress is the default address for the HTTP server.
	defaultHTTPAddress = "localhost:8081"

	// ReadHeaderTimeout is a timeout set to ensure that the HTTP server
	// is not vulnerable to Slowloris attacks that exhaust open connections
	// by holding them indefinitely and periodically sending partial HTTP headers.
	ReadHeaderTimeout = time.Minute
)

// ProxyServer is a server that proxies requests to Coinbase WaaS servers.
type ProxyServer struct {
	blockchainpb.UnimplementedBlockchainServiceServer
	protocolspb.UnimplementedProtocolServiceServer
	poolspb.UnimplementedPoolServiceServer
	mpckeyspb.UnimplementedMPCKeyServiceServer
	mpcwalletspb.UnimplementedMPCWalletServiceServer
	mpctransactionspb.UnimplementedMPCTransactionServiceServer
	longrunningpb.UnimplementedOperationsServer

	log *zap.Logger

	// Proxy Implementations
	blockchainProxy      blockchain.BlockchainProxy
	mpcKeysProxy         mpckeys.MPCKeysProxy
	mpcTransactionsProxy mpctransactions.MPCTransactionsProxy
	mpcWalletsProxy      mpcwallets.MPCWalletsProxy
	poolsProxy           pools.PoolsProxy
	protocolsProxy       protocols.ProtocolsProxy
	operationsProxy      operations.OperationsProxy
}

// NewProxyServer creates a new ProxyServer.
func NewProxyServer(ctx context.Context, log *zap.Logger) (*ProxyServer, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file", zap.Error(err))

		return nil, err
	}

	apiKeyName, apiKeyPrivateKey, err := getAPIKey(log)
	if err != nil {
		return nil, err
	}

	blockchainClient, err := clientsv1.NewBlockchainServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating blockchain client", zap.Error(err))

		return nil, err
	}

	protocolsClient, err := clientsv1.NewProtocolServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating protocols client", zap.Error(err))

		return nil, err
	}

	poolsClient, err := clientsv1.NewPoolServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating pools client", zap.Error(err))

		return nil, err
	}

	mpcKeysClient, err := clientsv1.NewMPCKeyServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating mpc keys client", zap.Error(err))

		return nil, err
	}

	mpcWalletsClient, err := clientsv1.NewMPCWalletServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating mpc wallets client", zap.Error(err))

		return nil, err
	}

	mpcTransactionsClient, err := clientsv1.NewMPCTransactionServiceClient(ctx, clients.WithAPIKey(&auth.APIKey{
		Name:       apiKeyName,
		PrivateKey: apiKeyPrivateKey,
	}))
	if err != nil {
		log.Error("Error creating mpc transactions client", zap.Error(err))

		return nil, err
	}

	operationsMap := make(map[string]operations.OperationType)

	blockchainProxy, err := blockchain.NewBlockchainProxy(ctx, log, blockchainClient)
	if err != nil {
		log.Error("Error instantiating blockchain proxy implementation", zap.Error(err))

		return nil, err
	}

	mpcKeysProxy, err := mpckeys.NewMPCKeysProxy(ctx, log, mpcKeysClient, operationsMap)
	if err != nil {
		log.Error("Error instantiating mpc keys proxy implementation", zap.Error(err))

		return nil, err
	}

	mpcTransactionsProxy, err := mpctransactions.NewMPCTransactionsProxy(ctx, log, mpcTransactionsClient, operationsMap)
	if err != nil {
		log.Error("Error instantiating mpc transactions proxy implementation", zap.Error(err))

		return nil, err
	}

	mpcWalletsProxy, err := mpcwallets.NewMPCWalletsProxy(ctx, log, mpcWalletsClient, operationsMap)
	if err != nil {
		log.Error("Error instantiating mpc wallets proxy implementation", zap.Error(err))

		return nil, err
	}

	poolsProxy, err := pools.NewPoolsProxy(ctx, log, poolsClient)
	if err != nil {
		log.Error("Error instantiating pools proxy implementation", zap.Error(err))

		return nil, err
	}

	protocolsProxy, err := protocols.NewProtocolsProxy(ctx, log, protocolsClient)
	if err != nil {
		log.Error("Error instantiating protocols proxy implementation", zap.Error(err))

		return nil, err
	}

	operationsProxy, err := operations.NewOperationsProxy(
		ctx,
		log,
		mpcKeysClient,
		mpcWalletsClient,
		mpcTransactionsClient,
		operationsMap,
	)
	if err != nil {
		log.Error("Error instantiating operations proxy implementation", zap.Error(err))

		return nil, err
	}

	return &ProxyServer{
		log:                  log,
		blockchainProxy:      blockchainProxy,
		mpcKeysProxy:         mpcKeysProxy,
		mpcTransactionsProxy: mpcTransactionsProxy,
		mpcWalletsProxy:      mpcWalletsProxy,
		poolsProxy:           poolsProxy,
		protocolsProxy:       protocolsProxy,
		operationsProxy:      operationsProxy,
	}, nil
}

// getAPIKey returns the API key name and private key from the environment variables.
// This method should be implemented using your secret manager of choice.
func getAPIKey(log *zap.Logger) (string, string, error) {
	apiKeyName := os.Getenv("COINBASE_CLOUD_API_KEY_NAME")
	if apiKeyName == "" {
		log.Error("COINBASE_CLOUD_API_KEY_NAME not set in environment variables")

		return "", "", errors.New("COINBASE_CLOUD_API_KEY_NAME not set in environment variables")
	}

	apiKeyPrivateKey := os.Getenv("COINBASE_CLOUD_API_KEY_PRIVATE_KEY")
	if apiKeyPrivateKey == "" {
		log.Error("COINBASE_CLOUD_API_KEY_PRIVATE_KEY not set in environment variables")

		return "", "", errors.New("COINBASE_CLOUD_API_KEY_PRIVATE_KEY not set in environment variables")
	}

	return apiKeyName, apiKeyPrivateKey, nil
}

// Serve starts the proxy server.
func (p *ProxyServer) Serve(ctx context.Context) error {
	var group run.Group

	startGRPCServer, stopGRPCServer, grpcAddress, err := p.getGRPCActor()
	if err != nil {
		return err
	}

	group.Add(startGRPCServer, stopGRPCServer)

	startHTTPServer, stopHTTPServer, err := p.getHTTPActor(ctx, grpcAddress)
	if err != nil {
		return err
	}

	group.Add(startHTTPServer, stopHTTPServer)

	return group.Run()
}

// getGRPCActor returns the execute and interrupt functions for the gRPC server to be added
// to the run.Group. It also returns the address of the gRPC server.
func (p *ProxyServer) getGRPCActor() (func() error, func(error), string, error) {
	server := grpc.NewServer()

	grpcAddress := os.Getenv("GRPC_ADDRESS")
	if grpcAddress == "" {
		grpcAddress = defaultGRPCAddress
	}

	log := p.log.With(zap.String("grpc_address", grpcAddress))

	grpcListener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Error("Error listening on gRPC address", zap.Error(err))

		return nil, nil, "", err
	}

	p.registerGRPCServer(server)

	startGRPCServer := func() error {
		log.Info("Starting gRPC server")

		if err = server.Serve(grpcListener); err != nil {
			log.Error("Error serving gRPC", zap.Error(err))

			return err
		}

		log.Info("gRPC server stopped")

		return nil
	}

	stopGRPCServer := func(err error) {
		log.Info("Stopping gRPC server", zap.Error(err))

		server.GracefulStop()
	}

	return startGRPCServer, stopGRPCServer, grpcAddress, nil
}

// getHTTPActor returns the execute and interrupt functions for the HTTP server to be added
// to the run.Group.
func (p *ProxyServer) getHTTPActor(
	ctx context.Context,
	grpcAddress string,
) (func() error, func(error), error) {
	mux := runtime.NewServeMux()

	// Transport security must be set on the server.
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	httpAddress := os.Getenv("HTTP_ADDRESS")
	if httpAddress == "" {
		httpAddress = defaultHTTPAddress
	}

	log := p.log.With(
		zap.String("grpc_address", grpcAddress),
		zap.String("http_address", httpAddress),
	)

	httpListener, err := net.Listen("tcp", httpAddress)
	if err != nil {
		log.Error("Error listening on HTTP address", zap.Error(err))
	}

	if err := p.registerHTTPServer(ctx, mux, grpcAddress, opts); err != nil {
		log.Error("Error registering HTTP server", zap.Error(err))

		return nil, nil, err
	}

	httpServer := &http.Server{
		Handler:           mux,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	startHTTPServer := func() error {
		log.Info("Starting HTTP server")

		if err = httpServer.Serve(httpListener); err != nil {
			log.Error("Error serving HTTP", zap.Error(err))

			return err
		}

		log.Info("HTTP server stopped")

		return nil
	}

	stopHTTPServer := func(err error) {
		log.Info("Stopping HTTP server", zap.Error(err))

		shutdownErr := httpServer.Shutdown(ctx)
		if shutdownErr != nil {
			log.Error("Error shutting down HTTP server", zap.Error(shutdownErr))
		}
	}

	return startHTTPServer, stopHTTPServer, nil
}

// registerGRPCServer registers the services of this ProxyServer with the given gRPC server.
func (p *ProxyServer) registerGRPCServer(s *grpc.Server) {
	blockchainpb.RegisterBlockchainServiceServer(s, p)
	protocolspb.RegisterProtocolServiceServer(s, p)
	poolspb.RegisterPoolServiceServer(s, p)
	mpckeyspb.RegisterMPCKeyServiceServer(s, p)
	mpcwalletspb.RegisterMPCWalletServiceServer(s, p)
	mpctransactionspb.RegisterMPCTransactionServiceServer(s, p)
	longrunningpb.RegisterOperationsServer(s, p)

	// Reflection is required for gRPCurl.
	reflection.Register(s)
}

// registerHTTPServer registers the HTTP handlers for the services of this ProxyServer to the given mux.
// The handlers forward requests to the specified gRPC address.
func (p *ProxyServer) registerHTTPServer(
	ctx context.Context,
	mux *runtime.ServeMux,
	grpcAddress string,
	opts []grpc.DialOption,
) error {
	err := blockchainpb.RegisterBlockchainServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = protocolspb.RegisterProtocolServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = poolspb.RegisterPoolServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = mpckeyspb.RegisterMPCKeyServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = mpcwalletspb.RegisterMPCWalletServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = mpctransactionspb.RegisterMPCTransactionServiceHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	err = longrunninggw.RegisterOperationsHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		return err
	}

	return nil
}
