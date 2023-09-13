package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// Test_GetMPCWallet tests GetMPCWallet with a range of scenarios.
func (s *ts) Test_GetMPCWallet() {
	var (
		mpcWallet = &v1.MPCWallet{
			Name:        "pools/test-pool/mpcWallets/test-wallet",
			DeviceGroup: "device-groups/test-device-group",
		}

		getMPCWalletReq = &v1.GetMPCWalletRequest{
			Name: mpcWallet.GetName(),
		}

		newRequestFn = func() *v1.GetMPCWalletRequest {
			return getMPCWalletReq
		}

		validMutation = func(req *v1.GetMPCWalletRequest) *v1.MPCWallet {
			s.GetsMPCWallet(req, mpcWallet, nil)
			return mpcWallet
		}

		errorMutation = func(
			req *v1.GetMPCWalletRequest,
			err error,
		) *v1.MPCWallet {
			s.GetsMPCWallet(req, nil, err)
			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetMPCWalletRequest,
		) (*v1.MPCWallet, error) {
			return s.mpcWalletsProxy.GetMPCWallet(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
