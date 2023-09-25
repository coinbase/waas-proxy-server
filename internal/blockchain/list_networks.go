package blockchain

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListNetworks proxies the ListNetworks method.
func (b *blockchainProxy) ListNetworks(
	ctx context.Context,
	req *blockchainpb.ListNetworksRequest,
) (*blockchainpb.ListNetworksResponse, error) {
	b.log.Info("ListNetworks")

	iter := b.blockchainClient.ListNetworks(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		b.log.Error("Error iterating over networks", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
