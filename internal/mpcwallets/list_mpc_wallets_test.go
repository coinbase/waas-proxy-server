package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

func (s *ts) Test_ListMPCWallets() {
	var (
		mpcWallet1 = &v1.MPCWallet{
			Name:        "pools/test-pool/mpcWallets/test-wallet-1",
			DeviceGroup: "deviceGroups/test-device-group-1",
		}

		mpcWallet2 = &v1.MPCWallet{
			Name:        "pools/test-pool/mpcWallets/test-wallet-2",
			DeviceGroup: "deviceGroups/test-device-group-2",
		}

		listMPCWalletsReq = &v1.ListMPCWalletsRequest{
			PageSize:  5,
			PageToken: "",
		}

		listMPCWalletsResp = &v1.ListMPCWalletsResponse{
			MpcWallets:    []*v1.MPCWallet{mpcWallet1, mpcWallet2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *v1.ListMPCWalletsRequest {
			return listMPCWalletsReq
		}

		validMutation = func(req *v1.ListMPCWalletsRequest) *v1.ListMPCWalletsResponse {
			s.ListsMPCWallets(req, listMPCWalletsResp, nil)

			return listMPCWalletsResp
		}

		errorMutation = func(
			req *v1.ListMPCWalletsRequest,
			err error,
		) *v1.ListMPCWalletsResponse {
			s.ListsMPCWallets(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListMPCWalletsRequest,
		) (*v1.ListMPCWalletsResponse, error) {
			return s.mpcWalletsProxy.ListMPCWallets(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
