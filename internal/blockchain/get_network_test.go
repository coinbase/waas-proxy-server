package blockchain_test

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// Test_GetNetwork tests GetNetwork with a range of scenarios.
func (s *ts) Test_GetNetwork() {
	var (
		network = &blockchainpb.Network{
			Name:           "networks/test-network",
			DisplayName:    "test-network",
			NativeAsset:    "networks/test-network/assets/test-asset",
			ProtocolFamily: "protocolFamilies/evm",
		}

		getNetworkReq = &blockchainpb.GetNetworkRequest{
			Name: "networks/test-network",
		}

		newRequestFn = func() *blockchainpb.GetNetworkRequest {
			return getNetworkReq
		}

		validMutation = func(req *blockchainpb.GetNetworkRequest) *blockchainpb.Network {
			s.GetsNetwork(req, network, nil)

			return network
		}

		errorMutation = func(
			req *blockchainpb.GetNetworkRequest,
			err error,
		) *blockchainpb.Network {
			s.GetsNetwork(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *blockchainpb.GetNetworkRequest,
		) (*blockchainpb.Network, error) {
			return s.blockchainProxy.GetNetwork(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
