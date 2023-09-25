package protocols

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// BroadcastTransaction proxies the BroadcastTransaction method.
func (p *protocolsProxy) BroadcastTransaction(
	ctx context.Context,
	req *protocolspb.BroadcastTransactionRequest,
) (*typespb.Transaction, error) {
	p.log.Info("BroadcastTransaction")

	return p.protocolsClient.BroadcastTransaction(ctx, req)
}
