package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// Test_GetAddress tests GetAddress with a range of scenarios.
func (s *ts) Test_GetAddress() {
	var (
		address = &v1.Address{
			Name:    "pools/test-pool/mpcWallets/test-wallet/addresses/address-1",
			Address: "test-address-1",
		}

		getAddressReq = &v1.GetAddressRequest{
			Name: address.GetName(),
		}

		newRequestFn = func() *v1.GetAddressRequest {
			return getAddressReq
		}

		validMutation = func(req *v1.GetAddressRequest) *v1.Address {
			s.GetsAddress(req, address, nil)
			return address
		}

		errorMutation = func(
			req *v1.GetAddressRequest,
			err error,
		) *v1.Address {
			s.GetsAddress(req, nil, err)
			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetAddressRequest,
		) (*v1.Address, error) {
			return s.mpcWalletsProxy.GetAddress(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
