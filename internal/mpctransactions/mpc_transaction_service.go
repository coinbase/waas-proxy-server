package mpctransactions

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	"go.uber.org/zap"
)

// MPCTransactionsProxy is an interface with methods to proxy requests to the mpc transaction service.
type MPCTransactionsProxy interface {
	CreateMPCTransaction(context.Context, *v1.CreateMPCTransactionRequest) (*longrunningpb.Operation, error)
	GetMPCTransaction(context.Context, *v1.GetMPCTransactionRequest) (*v1.MPCTransaction, error)
	ListMPCTransactions(context.Context, *v1.ListMPCTransactionsRequest) (*v1.ListMPCTransactionsResponse, error)
}

// mpcTransactionsProxy implements the MPCTransactionsProxy interface.
type mpcTransactionsProxy struct {
	log *zap.Logger

	mpcTransactionsClient clientsv1.MPCTransactionServiceClient

	operationMap map[string]operations.OperationType
}

// NewMPCTransactionsProxy instantiates a new mpcTransactionsProxy.
func NewMPCTransactionsProxy(
	ctx context.Context,
	log *zap.Logger,
	mpcTransactionsClient clientsv1.MPCTransactionServiceClient,
	operationMap map[string]operations.OperationType,
) (*mpcTransactionsProxy, error) {
	return &mpcTransactionsProxy{
		log:                   log,
		mpcTransactionsClient: mpcTransactionsClient,
		operationMap:          operationMap,
	}, nil
}
