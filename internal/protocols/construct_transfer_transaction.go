package protocols

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// ConstructTransferTransaction proxies the ConstructTransferTransaction method.
func (p *protocolsProxy) ConstructTransferTransaction(
	ctx context.Context,
	req *protocolspb.ConstructTransferTransactionRequest,
) (*typespb.Transaction, error) {
	p.log.Info("ConstructTransferTransaction")

	return p.protocolsClient.ConstructTransferTransaction(ctx, req)
}
