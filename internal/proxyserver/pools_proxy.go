package proxyserver

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// CreatePool calls CreatePool on the pools proxy.
func (p *ProxyServer) CreatePool(
	ctx context.Context,
	req *poolspb.CreatePoolRequest,
) (*poolspb.Pool, error) {
	return p.poolsProxy.CreatePool(ctx, req)
}

// GetPool calls GetPool on the pools proxy.
func (p *ProxyServer) GetPool(
	ctx context.Context,
	req *poolspb.GetPoolRequest,
) (*poolspb.Pool, error) {
	return p.poolsProxy.GetPool(ctx, req)
}

// ListPools calls ListPools on the pools proxy.
func (p *ProxyServer) ListPools(
	ctx context.Context,
	req *poolspb.ListPoolsRequest,
) (*poolspb.ListPoolsResponse, error) {
	return p.poolsProxy.ListPools(ctx, req)
}
