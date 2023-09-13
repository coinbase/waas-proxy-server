package blockchain

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// GetAsset proxies the GetAsset method.
func (b *blockchainProxy) GetAsset(
	ctx context.Context,
	req *blockchainpb.GetAssetRequest,
) (*blockchainpb.Asset, error) {
	b.log.Info("GetAsset")

	return b.blockchainClient.GetAsset(ctx, req)
}
