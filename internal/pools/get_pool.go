package pools

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// GetPool proxies the GetPool method.
func (p *poolsProxy) GetPool(
	ctx context.Context,
	req *poolspb.GetPoolRequest,
) (*poolspb.Pool, error) {
	p.log.Info("GetPool")

	return p.poolsClient.GetPool(ctx, req)
}
