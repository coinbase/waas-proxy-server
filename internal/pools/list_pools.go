package pools

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListPools proxies the ListPools method.
func (p *poolsProxy) ListPools(
	ctx context.Context,
	req *poolspb.ListPoolsRequest,
) (*poolspb.ListPoolsResponse, error) {
	p.log.Info("ListPools")

	iter := p.poolsClient.ListPools(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		p.log.Error("Error iterating over pools", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
