package blockchain

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// GetNetwork proxies the GetNetwork method.
func (b *blockchainProxy) GetNetwork(
	ctx context.Context,
	req *blockchainpb.GetNetworkRequest,
) (*blockchainpb.Network, error) {
	b.log.Info("GetNetwork")

	return b.blockchainClient.GetNetwork(ctx, req)
}
