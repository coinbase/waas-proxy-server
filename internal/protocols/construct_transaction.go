package protocols

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// ConstructTransaction proxies the ConstructTransaction method.
func (p *protocolsProxy) ConstructTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransactionRequest,
) (*typespb.Transaction, error) {
	p.log.Info("ConstructTransaction")

	return p.protocolsClient.ConstructTransaction(ctx, req)
}
