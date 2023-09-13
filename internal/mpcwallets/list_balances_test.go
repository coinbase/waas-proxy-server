package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

func (s *ts) Test_ListBalances() {
	var (
		balance1 = &v1.Balance{
			Name:   "pools/test-pool/mpcWallets/test-wallet/balanceDetails/test-balance-detail-1",
			Asset:  "assets/test-asset-1",
			Amount: "5",
		}

		balance2 = &v1.Balance{
			Name:   "pools/test-pool/mpcWallets/test-wallet/balanceDetails/test-balance-detail-2",
			Asset:  "assets/test-asset-2",
			Amount: "8",
		}

		listBalancesReq = &v1.ListBalancesRequest{
			PageSize:  5,
			PageToken: "",
		}

		listBalancesResp = &v1.ListBalancesResponse{
			Balances:      []*v1.Balance{balance1, balance2},
			NextPageToken: "test-next-page-token",
		}

		newRequestFn = func() *v1.ListBalancesRequest {
			return listBalancesReq
		}

		validMutation = func(req *v1.ListBalancesRequest) *v1.ListBalancesResponse {
			s.ListsBalances(req, listBalancesResp, nil)

			return listBalancesResp
		}

		errorMutation = func(
			req *v1.ListBalancesRequest,
			err error,
		) *v1.ListBalancesResponse {
			s.ListsBalances(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListBalancesRequest,
		) (*v1.ListBalancesResponse, error) {
			return s.mpcWalletsProxy.ListBalances(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
