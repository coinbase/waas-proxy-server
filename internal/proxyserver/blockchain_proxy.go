package proxyserver

import (
	"context"

	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
)

// GetNetwork calls GetNetwork on the blockchain proxy.
func (p *ProxyServer) GetNetwork(
	ctx context.Context,
	req *blockchainpb.GetNetworkRequest,
) (*blockchainpb.Network, error) {
	return p.blockchainProxy.GetNetwork(ctx, req)
}

// ListNetworks calls ListNetworks on the blockchain proxy.
func (p *ProxyServer) ListNetworks(
	ctx context.Context,
	req *blockchainpb.ListNetworksRequest,
) (*blockchainpb.ListNetworksResponse, error) {
	return p.blockchainProxy.ListNetworks(ctx, req)
}

// GetAsset calls GetAsset on the blockchain proxy.
func (p *ProxyServer) GetAsset(
	ctx context.Context,
	req *blockchainpb.GetAssetRequest,
) (*blockchainpb.Asset, error) {
	return p.blockchainProxy.GetAsset(ctx, req)
}

// ListAssets calls ListAssets on the blockchain proxy.
func (p *ProxyServer) ListAssets(
	ctx context.Context,
	req *blockchainpb.ListAssetsRequest,
) (*blockchainpb.ListAssetsResponse, error) {
	return p.blockchainProxy.ListAssets(ctx, req)
}

// BatchGetAssets calls BatchGetAssets on the blockchain proxy.
func (p *ProxyServer) BatchGetAssets(
	ctx context.Context,
	req *blockchainpb.BatchGetAssetsRequest,
) (*blockchainpb.BatchGetAssetsResponse, error) {
	return p.blockchainProxy.BatchGetAssets(ctx, req)
}
