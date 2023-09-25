package protocols

import (
	"context"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	"go.uber.org/zap"
)

// ProtocolsProxy is an interface with methods to proxy requests to the protocol service.
type ProtocolsProxy interface {
	ConstructTransaction(context.Context, *protocolspb.ConstructTransactionRequest) (*typespb.Transaction, error)
	ConstructTransferTransaction(context.Context, *protocolspb.ConstructTransferTransactionRequest) (*typespb.Transaction, error)
	BroadcastTransaction(context.Context, *protocolspb.BroadcastTransactionRequest) (*typespb.Transaction, error)
	EstimateFee(context.Context, *protocolspb.EstimateFeeRequest) (*protocolspb.EstimateFeeResponse, error)
}

// protocolsProxy implements the ProtocolsProxy interface.
type protocolsProxy struct {
	log *zap.Logger

	protocolsClient clientsv1.ProtocolServiceClient
}

// NewProtocolsProxyImpl instantiates a new protocolsProxy.
func NewProtocolsProxy(
	ctx context.Context,
	log *zap.Logger,
	protocolsClient clientsv1.ProtocolServiceClient,
) (*protocolsProxy, error) {
	return &protocolsProxy{
		log:             log,
		protocolsClient: protocolsClient,
	}, nil
}
