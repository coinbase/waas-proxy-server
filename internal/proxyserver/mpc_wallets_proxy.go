package proxyserver

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
)

// CreateMPCWallet calls CreateMPCWallet on the mpc wallets proxy.
func (p *ProxyServer) CreateMPCWallet(
	ctx context.Context,
	req *mpcwalletspb.CreateMPCWalletRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcWalletsProxy.CreateMPCWallet(ctx, req)
}

// GetMPCWallet calls GetMPCWallet on the mpc wallets proxy.
func (p *ProxyServer) GetMPCWallet(
	ctx context.Context,
	req *mpcwalletspb.GetMPCWalletRequest,
) (*mpcwalletspb.MPCWallet, error) {
	return p.mpcWalletsProxy.GetMPCWallet(ctx, req)
}

// ListMPCWallets calls ListMPCWallets on the mpc wallets proxy.
func (p *ProxyServer) ListMPCWallets(
	ctx context.Context,
	req *mpcwalletspb.ListMPCWalletsRequest,
) (*mpcwalletspb.ListMPCWalletsResponse, error) {
	return p.mpcWalletsProxy.ListMPCWallets(ctx, req)
}

// GenerateAddress calls GenerateAddress on the mpc wallets proxy.
func (p *ProxyServer) GenerateAddress(
	ctx context.Context,
	req *mpcwalletspb.GenerateAddressRequest,
) (*mpcwalletspb.Address, error) {
	return p.mpcWalletsProxy.GenerateAddress(ctx, req)
}

// GetAddress calls GetAddress on the mpc wallets proxy.
func (p *ProxyServer) GetAddress(
	ctx context.Context,
	req *mpcwalletspb.GetAddressRequest,
) (*mpcwalletspb.Address, error) {
	return p.mpcWalletsProxy.GetAddress(ctx, req)
}

// ListAddresses calls ListAddresses on the mpc wallets proxy.
func (p *ProxyServer) ListAddresses(
	ctx context.Context,
	req *mpcwalletspb.ListAddressesRequest,
) (*mpcwalletspb.ListAddressesResponse, error) {
	return p.mpcWalletsProxy.ListAddresses(ctx, req)
}

// ListBalances calls ListBalances on the mpc wallets proxy.
func (p *ProxyServer) ListBalances(
	ctx context.Context,
	req *mpcwalletspb.ListBalancesRequest,
) (*mpcwalletspb.ListBalancesResponse, error) {
	return p.mpcWalletsProxy.ListBalances(ctx, req)
}

// ListBalanceDetails calls ListBalanceDetails on the mpc wallets proxy.
func (p *ProxyServer) ListBalanceDetails(
	ctx context.Context,
	req *mpcwalletspb.ListBalanceDetailsRequest,
) (*mpcwalletspb.ListBalanceDetailsResponse, error) {
	return p.mpcWalletsProxy.ListBalanceDetails(ctx, req)
}
