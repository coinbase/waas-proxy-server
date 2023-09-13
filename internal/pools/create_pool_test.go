package pools_test

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// Test_CreatePool tests CreatePool with a range of scenarios.
func (s *ts) Test_CreatePool() {
	var (
		pool = &poolspb.Pool{
			Name:        "pools/test-pool",
			DisplayName: "test-pool",
		}

		createPoolReq = &poolspb.CreatePoolRequest{
			Pool:   pool,
			PoolId: "test-pool",
		}

		newRequestFn = func() *poolspb.CreatePoolRequest {
			return createPoolReq
		}

		validMutation = func(req *poolspb.CreatePoolRequest) *poolspb.Pool {
			s.CreatesPool(req, pool, nil)

			return pool
		}

		errorMutation = func(
			req *poolspb.CreatePoolRequest,
			err error,
		) *poolspb.Pool {
			s.CreatesPool(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *poolspb.CreatePoolRequest,
		) (*poolspb.Pool, error) {
			return s.poolsProxy.CreatePool(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
