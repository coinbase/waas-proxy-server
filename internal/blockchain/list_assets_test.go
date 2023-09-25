package blockchain_test

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// Test_ListAssets tests ListAssets with a range of scenarios.
func (s *ts) Test_ListAssets() {
	var (
		asset1 = &blockchainpb.Asset{
			Name:             "networks/test-network/assets/test-asset-1",
			AdvertisedSymbol: "TEST-1",
		}

		asset2 = &blockchainpb.Asset{
			Name:             "networks/test-network/assets/test-asset-2",
			AdvertisedSymbol: "TEST-2",
		}

		listAssetsReq = &blockchainpb.ListAssetsRequest{
			Parent:    "networks/test-network",
			PageSize:  5,
			PageToken: "",
		}

		listAssetsResp = &blockchainpb.ListAssetsResponse{
			Assets:        []*blockchainpb.Asset{asset1, asset2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *blockchainpb.ListAssetsRequest {
			return listAssetsReq
		}

		validMutation = func(req *blockchainpb.ListAssetsRequest) *blockchainpb.ListAssetsResponse {
			s.ListsAssets(req, listAssetsResp, nil)

			return listAssetsResp
		}

		errorMutation = func(
			req *blockchainpb.ListAssetsRequest,
			err error,
		) *blockchainpb.ListAssetsResponse {
			s.ListsAssets(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *blockchainpb.ListAssetsRequest,
		) (*blockchainpb.ListAssetsResponse, error) {
			return s.blockchainProxy.ListAssets(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
