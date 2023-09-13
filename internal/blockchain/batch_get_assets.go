package blockchain

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// BatchGetAssets proxies the BatchGetAssets method.
func (b *blockchainProxy) BatchGetAssets(
	ctx context.Context,
	req *blockchainpb.BatchGetAssetsRequest,
) (*blockchainpb.BatchGetAssetsResponse, error) {
	b.log.Info("BatchGetAssets")

	return b.blockchainClient.BatchGetAssets(ctx, req)
}
