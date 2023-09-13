package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// GetAddress proxies the GetAddress method.
func (m *mpcWalletsProxy) GetAddress(
	ctx context.Context,
	req *mpcwalletspb.GetAddressRequest,
) (*mpcwalletspb.Address, error) {
	m.log.Info("GetAddress")

	return m.mpcWalletsClient.GetAddress(ctx, req)
}
