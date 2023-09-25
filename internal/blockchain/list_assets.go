package blockchain

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListAssets proxies the ListAssets method.
func (b *blockchainProxy) ListAssets(
	ctx context.Context,
	req *blockchainpb.ListAssetsRequest,
) (*blockchainpb.ListAssetsResponse, error) {
	b.log.Info("ListAssets")

	iter := b.blockchainClient.ListAssets(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		b.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
