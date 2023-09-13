package blockchain_test

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// Test_GetAsset tests GetAsset with a range of scenarios.
func (s *ts) Test_GetAsset() {
	var (
		asset = &blockchainpb.Asset{
			Name:             "networks/test-network/assets/test-asset",
			AdvertisedSymbol: "TEST",
		}

		getAssetReq = &blockchainpb.GetAssetRequest{
			Name: "networks/test-network/assets/test-asset",
		}

		newRequestFn = func() *blockchainpb.GetAssetRequest {
			return getAssetReq
		}

		validMutation = func(req *blockchainpb.GetAssetRequest) *blockchainpb.Asset {
			s.GetsAsset(req, asset, nil)

			return asset
		}

		errorMutation = func(
			req *blockchainpb.GetAssetRequest,
			err error,
		) *blockchainpb.Asset {
			s.GetsAsset(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *blockchainpb.GetAssetRequest,
		) (*blockchainpb.Asset, error) {
			return s.blockchainProxy.GetAsset(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
