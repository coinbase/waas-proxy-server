package mpcwallets

import (
	"context"

	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// GetMPCWallet proxies the GetMPCWallet method.
func (m *mpcWalletsProxy) GetMPCWallet(
	ctx context.Context,
	req *mpcwalletspb.GetMPCWalletRequest,
) (*mpcwalletspb.MPCWallet, error) {
	m.log.Info("GetMPCWallet")

	return m.mpcWalletsClient.GetMPCWallet(ctx, req)
}
