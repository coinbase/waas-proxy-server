package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// Test_ListAddresses tests ListAddresses with a range of scenarios.
func (s *ts) Test_ListAddresses() {
	var (
		address1 = &v1.Address{
			Name:    "pools/test-pool/mpcWallets/test-wallet/addresses/address-1",
			Address: "test-address-1",
		}

		address2 = &v1.Address{
			Name:    "pools/test-pool/mpcWallets/test-wallet/addresses/address-2",
			Address: "test-address-2",
		}

		listAddressesReq = &v1.ListAddressesRequest{
			PageSize:  5,
			PageToken: "",
		}

		listAddressesResp = &v1.ListAddressesResponse{
			Addresses:     []*v1.Address{address1, address2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *v1.ListAddressesRequest {
			return listAddressesReq
		}

		validMutation = func(req *v1.ListAddressesRequest) *v1.ListAddressesResponse {
			s.ListsAddresses(req, listAddressesResp, nil)

			return listAddressesResp
		}

		errorMutation = func(
			req *v1.ListAddressesRequest,
			err error,
		) *v1.ListAddressesResponse {
			s.ListsAddresses(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListAddressesRequest,
		) (*v1.ListAddressesResponse, error) {
			return s.mpcWalletsProxy.ListAddresses(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
