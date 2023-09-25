package pools

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// CreatePool proxies the CreatePool method.
func (p *poolsProxy) CreatePool(
	ctx context.Context,
	req *poolspb.CreatePoolRequest,
) (*poolspb.Pool, error) {
	p.log.Info("CreatePool")

	return p.poolsClient.CreatePool(ctx, req)
}
