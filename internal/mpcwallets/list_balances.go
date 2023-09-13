package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListBalances proxies the ListBalances method.
func (m *mpcWalletsProxy) ListBalances(
	ctx context.Context,
	req *mpcwalletspb.ListBalancesRequest,
) (*mpcwalletspb.ListBalancesResponse, error) {
	m.log.Info("ListBalances")

	iter := m.mpcWalletsClient.ListBalances(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		m.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
