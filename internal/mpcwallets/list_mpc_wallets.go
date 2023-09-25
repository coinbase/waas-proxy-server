package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListMPCWallets proxies the ListMPCWallets method.
func (m *mpcWalletsProxy) ListMPCWallets(
	ctx context.Context,
	req *mpcwalletspb.ListMPCWalletsRequest,
) (*mpcwalletspb.ListMPCWalletsResponse, error) {
	m.log.Info("ListMPCWallets")

	iter := m.mpcWalletsClient.ListMPCWallets(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		m.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
