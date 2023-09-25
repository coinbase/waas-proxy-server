package blockchain_test

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// Test_BatchGetAssets tests BatchGetAssets with a range of scenarios.
func (s *ts) Test_BatchGetAssets() {
	var (
		asset1 = &blockchainpb.Asset{
			Name:             "networks/test-network/assets/test-asset-1",
			AdvertisedSymbol: "TEST-1",
		}

		asset2 = &blockchainpb.Asset{
			Name:             "networks/test-network/assets/test-asset-2",
			AdvertisedSymbol: "TEST-2",
		}

		batchGetAssetsReq = &blockchainpb.BatchGetAssetsRequest{
			Parent: "networks/test-network",
			Names:  []string{asset1.GetName(), asset2.GetName()},
		}

		batchGetAssetsResp = &blockchainpb.BatchGetAssetsResponse{
			Assets: []*blockchainpb.Asset{asset1, asset2},
		}

		newRequestFn = func() *blockchainpb.BatchGetAssetsRequest {
			return batchGetAssetsReq
		}

		validMutation = func(req *blockchainpb.BatchGetAssetsRequest) *blockchainpb.BatchGetAssetsResponse {
			s.BatchesGetAssets(req, batchGetAssetsResp, nil)

			return batchGetAssetsResp
		}

		errorMutation = func(
			req *blockchainpb.BatchGetAssetsRequest,
			err error,
		) *blockchainpb.BatchGetAssetsResponse {
			s.BatchesGetAssets(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *blockchainpb.BatchGetAssetsRequest,
		) (*blockchainpb.BatchGetAssetsResponse, error) {
			return s.blockchainProxy.BatchGetAssets(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
