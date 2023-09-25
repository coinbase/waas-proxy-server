package blockchain

import (
	"context"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"

	"go.uber.org/zap"
)

// BlockchainProxy is an interface with methods to proxy requests to the blockchain service.
type BlockchainProxy interface {
	GetNetwork(context.Context, *blockchainpb.GetNetworkRequest) (*blockchainpb.Network, error)
	ListNetworks(context.Context, *blockchainpb.ListNetworksRequest) (*blockchainpb.ListNetworksResponse, error)
	GetAsset(context.Context, *blockchainpb.GetAssetRequest) (*blockchainpb.Asset, error)
	ListAssets(context.Context, *blockchainpb.ListAssetsRequest) (*blockchainpb.ListAssetsResponse, error)
	BatchGetAssets(context.Context, *blockchainpb.BatchGetAssetsRequest) (*blockchainpb.BatchGetAssetsResponse, error)
}

// BlockchainProxy implements the BlockchainProxy interface.
type blockchainProxy struct {
	log *zap.Logger

	blockchainClient clientsv1.BlockchainServiceClient
}

// NewBlockchainProxy instantiates a new blockchainProxy.
func NewBlockchainProxy(
	ctx context.Context,
	log *zap.Logger,
	blockchainClient clientsv1.BlockchainServiceClient,
) (*blockchainProxy, error) {
	return &blockchainProxy{
		log:              log,
		blockchainClient: blockchainClient,
	}, nil
}
