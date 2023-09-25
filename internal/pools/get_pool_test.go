package pools_test

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// Test_GetPool tests GetPool with a range of scenarios.
func (s *ts) Test_GetPool() {
	var (
		pool = &poolspb.Pool{
			Name:        "pools/test-pool",
			DisplayName: "test-pool",
		}

		getPoolReq = &poolspb.GetPoolRequest{
			Name: pool.GetName(),
		}

		newRequestFn = func() *poolspb.GetPoolRequest {
			return getPoolReq
		}

		validMutation = func(req *poolspb.GetPoolRequest) *poolspb.Pool {
			s.GetsPool(req, pool, nil)

			return pool
		}

		errorMutation = func(
			req *poolspb.GetPoolRequest,
			err error,
		) *poolspb.Pool {
			s.GetsPool(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *poolspb.GetPoolRequest,
		) (*poolspb.Pool, error) {
			return s.poolsProxy.GetPool(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
