package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
)

const (
	MPC_KEY_OPERATION         OperationType = "MPC_KEY_OPERATION"
	MPC_WALLET_OPERATION      OperationType = "MPC_WALLET_OPERATION"
	MPC_TRANSACTION_OPERATION OperationType = "MPC_TRANSACTION_OPERATION"
)

// OperationType indicates the WaaS longrunning operation type, segmented by service.
type OperationType string

// OperationsProxy is an interface with methods to proxy requests to the operations service.
type OperationsProxy interface {
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest) (*longrunningpb.ListOperationsResponse, error)
	GetOperation(context.Context, *longrunningpb.GetOperationRequest) (*longrunningpb.Operation, error)
	DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest) (*emptypb.Empty, error)
	CancelOperation(context.Context, *longrunningpb.CancelOperationRequest) (*emptypb.Empty, error)
	WaitOperation(context.Context, *longrunningpb.WaitOperationRequest) (*longrunningpb.Operation, error)
}

// operationsProxy implements the OperationsProxy interface.
type operationsProxy struct {
	log *zap.Logger

	mpcKeysClient         clientsv1.MPCKeyServiceClient
	mpcWalletsClient      clientsv1.MPCWalletServiceClient
	mpcTransactionsClient clientsv1.MPCTransactionServiceClient

	operationMap map[string]OperationType
}

// NewOperationsProxy instantiates a new operationsProxy.
func NewOperationsProxy(
	ctx context.Context,
	log *zap.Logger,
	mpcKeysClient clientsv1.MPCKeyServiceClient,
	mpcWalletsClient clientsv1.MPCWalletServiceClient,
	mpcTransactionsClient clientsv1.MPCTransactionServiceClient,
	operationMap map[string]OperationType,
) (*operationsProxy, error) {
	return &operationsProxy{
		log:                   log,
		mpcKeysClient:         mpcKeysClient,
		mpcWalletsClient:      mpcWalletsClient,
		mpcTransactionsClient: mpcTransactionsClient,
		operationMap:          operationMap,
	}, nil
}
