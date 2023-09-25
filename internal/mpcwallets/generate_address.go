package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// GenerateAddress proxies the GenerateAddress method.
func (m *mpcWalletsProxy) GenerateAddress(
	ctx context.Context,
	req *mpcwalletspb.GenerateAddressRequest,
) (*mpcwalletspb.Address, error) {
	m.log.Info("GenerateAddress")

	return m.mpcWalletsClient.GenerateAddress(ctx, req)
}
