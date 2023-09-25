package pools

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	"go.uber.org/zap"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
)

// PoolsProxy is an interface with methods to proxy requests to the pool service.
type PoolsProxy interface {
	CreatePool(context.Context, *poolspb.CreatePoolRequest) (*poolspb.Pool, error)
	GetPool(context.Context, *poolspb.GetPoolRequest) (*poolspb.Pool, error)
	ListPools(context.Context, *poolspb.ListPoolsRequest) (*poolspb.ListPoolsResponse, error)
}

// poolsProxy implements the PoolsProxy interface.
type poolsProxy struct {
	log *zap.Logger

	poolsClient clientsv1.PoolServiceClient
}

// NewPoolsProxy instantiates a new poolsProxy.
func NewPoolsProxy(
	ctx context.Context,
	log *zap.Logger,
	poolsClient clientsv1.PoolServiceClient,
) (*poolsProxy, error) {
	return &poolsProxy{
		log:         log,
		poolsClient: poolsClient,
	}, nil
}
