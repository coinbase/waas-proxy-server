package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// Test_GenerateAddress tests GenerateAddress with a range of scenarios.
func (s *ts) Test_GenerateAddress() {
	var (
		address = &v1.Address{
			Name:    "pools/test-pool/mpcWallets/test-wallet/addresses/address-generated",
			Address: "test-generated-address",
		}

		generateAddressReq = &v1.GenerateAddressRequest{
			MpcWallet: "pools/test-pool/mpcWallets/test-wallet",
			Network:   "networks/test-network",
		}

		newRequestFn = func() *v1.GenerateAddressRequest {
			return generateAddressReq
		}

		validMutation = func(req *v1.GenerateAddressRequest) *v1.Address {
			s.GeneratesAddress(req, address, nil)
			return address
		}

		errorMutation = func(
			req *v1.GenerateAddressRequest,
			err error,
		) *v1.Address {
			s.GeneratesAddress(req, nil, err)
			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GenerateAddressRequest,
		) (*v1.Address, error) {
			return s.mpcWalletsProxy.GenerateAddress(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
