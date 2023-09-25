package mpcwallets_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

func (s *ts) Test_ListBalanceDetails() {
	var (
		balanceDetail1 = &v1.BalanceDetail{
			Name:   "pools/test-pool/mpcWallets/test-wallet/balanceDetails/test-balance-detail-1",
			Asset:  "assets/test-asset-1",
			Amount: "5",
		}

		balanceDetail2 = &v1.BalanceDetail{
			Name:   "pools/test-pool/mpcWallets/test-wallet/balanceDetails/test-balance-detail-2",
			Asset:  "assets/test-asset-2",
			Amount: "8",
		}

		listBalanceDetailsReq = &v1.ListBalanceDetailsRequest{
			PageSize:  5,
			PageToken: "",
		}

		listBalanceDetailsResp = &v1.ListBalanceDetailsResponse{
			BalanceDetails: []*v1.BalanceDetail{balanceDetail1, balanceDetail2},
			NextPageToken:  "test-next-page-token",
		}

		newRequestFn = func() *v1.ListBalanceDetailsRequest {
			return listBalanceDetailsReq
		}

		validMutation = func(req *v1.ListBalanceDetailsRequest) *v1.ListBalanceDetailsResponse {
			s.ListsBalanceDetails(req, listBalanceDetailsResp, nil)

			return listBalanceDetailsResp
		}

		errorMutation = func(
			req *v1.ListBalanceDetailsRequest,
			err error,
		) *v1.ListBalanceDetailsResponse {
			s.ListsBalanceDetails(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListBalanceDetailsRequest,
		) (*v1.ListBalanceDetailsResponse, error) {
			return s.mpcWalletsProxy.ListBalanceDetails(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
