package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListAddresses proxies the ListAddresses method.
func (m *mpcWalletsProxy) ListAddresses(
	ctx context.Context,
	req *mpcwalletspb.ListAddressesRequest,
) (*mpcwalletspb.ListAddressesResponse, error) {
	m.log.Info("ListAddresses")

	iter := m.mpcWalletsClient.ListAddresses(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		m.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
