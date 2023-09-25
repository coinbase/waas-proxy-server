package proxyserver

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// ConstructTransaction calls ConstructTransaction on the protocols proxy.
func (p *ProxyServer) ConstructTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransactionRequest,
) (*typespb.Transaction, error) {
	return p.protocolsProxy.ConstructTransaction(ctx, req)
}

// ConstructTransferTransaction calls ConstructTransferTransaction on the protocols proxy.
func (p *ProxyServer) ConstructTransferTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransferTransactionRequest,
) (*typespb.Transaction, error) {
	return p.protocolsProxy.ConstructTransferTransaction(ctx, req)
}

// BraodcastTransaction calls BroadcastTransaction on the protocols proxy.
func (p *ProxyServer) BroadcastTransaction(
	ctx context.Context,
	req *protocolspb.BroadcastTransactionRequest,
) (*typespb.Transaction, error) {
	return p.protocolsProxy.BroadcastTransaction(ctx, req)
}

// EstimateFee calls EstimateFee on the protocols proxy.
func (p *ProxyServer) EstimateFee(
	ctx context.Context,
	req *protocolspb.EstimateFeeRequest,
) (*protocolspb.EstimateFeeResponse, error) {
	return p.protocolsProxy.EstimateFee(ctx, req)
}
