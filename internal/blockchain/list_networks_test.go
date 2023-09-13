package blockchain_test

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// Test_ListNetworks tests ListNetworks with a range of scenarios.
func (s *ts) Test_ListNetworks() {
	var (
		network1 = &blockchainpb.Network{
			Name:        "networks/test-network-1",
			DisplayName: "test-network-1",
		}

		network2 = &blockchainpb.Network{
			Name:        "networks/test-network-2",
			DisplayName: "test-network-2",
		}

		listNetworksReq = &blockchainpb.ListNetworksRequest{
			PageSize:  5,
			PageToken: "",
		}

		listNetworksResp = &blockchainpb.ListNetworksResponse{
			Networks:      []*blockchainpb.Network{network1, network2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *blockchainpb.ListNetworksRequest {
			return listNetworksReq
		}

		validMutation = func(req *blockchainpb.ListNetworksRequest) *blockchainpb.ListNetworksResponse {
			s.ListsNetworks(req, listNetworksResp, nil)

			return listNetworksResp
		}

		errorMutation = func(
			req *blockchainpb.ListNetworksRequest,
			err error,
		) *blockchainpb.ListNetworksResponse {
			s.ListsNetworks(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *blockchainpb.ListNetworksRequest,
		) (*blockchainpb.ListNetworksResponse, error) {
			return s.blockchainProxy.ListNetworks(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
