package protocols

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
)

// EstimateFee proxies the EstimateFee method.
func (p *protocolsProxy) EstimateFee(
	ctx context.Context,
	req *protocolspb.EstimateFeeRequest,
) (*protocolspb.EstimateFeeResponse, error) {
	p.log.Info("EstimateFee")

	return p.protocolsClient.EstimateFee(ctx, req)
}
