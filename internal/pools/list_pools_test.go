package pools_test

import (
	"context"

	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
)

// Test_ListPools tests ListPools with a range of scenarios.
func (s *ts) Test_ListPools() {
	var (
		pool1 = &poolspb.Pool{
			Name:        "pools/test-pool-1",
			DisplayName: "test-pool-1",
		}

		pool2 = &poolspb.Pool{
			Name:        "pools/test-pool-2",
			DisplayName: "test-pool-2",
		}

		listPoolsReq = &poolspb.ListPoolsRequest{
			PageSize:  5,
			PageToken: "",
		}

		listPoolsResp = &poolspb.ListPoolsResponse{
			Pools:         []*poolspb.Pool{pool1, pool2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *poolspb.ListPoolsRequest {
			return listPoolsReq
		}

		validMutation = func(req *poolspb.ListPoolsRequest) *poolspb.ListPoolsResponse {
			s.ListsPools(req, listPoolsResp, nil)

			return listPoolsResp
		}

		errorMutation = func(
			req *poolspb.ListPoolsRequest,
			err error,
		) *poolspb.ListPoolsResponse {
			s.ListsPools(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *poolspb.ListPoolsRequest,
		) (*poolspb.ListPoolsResponse, error) {
			return s.poolsProxy.ListPools(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
