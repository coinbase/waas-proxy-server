package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListBalanceDetails proxies the ListBalanceDetails method.
func (m *mpcWalletsProxy) ListBalanceDetails(
	ctx context.Context,
	req *mpcwalletspb.ListBalanceDetailsRequest,
) (*mpcwalletspb.ListBalanceDetailsResponse, error) {
	m.log.Info("ListBalanceDetails")

	iter := m.mpcWalletsClient.ListBalanceDetails(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		m.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
