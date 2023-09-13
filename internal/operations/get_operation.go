package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetOperation proxies the GetOperation method.
func (o *operationsProxy) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
) (*longrunningpb.Operation, error) {
	o.log.Info("GetOperation")

	// We use an in-memory map to keep track of the operation type for each operation name.
	// In production, we recommend using a persistent store.
	operationType, ok := o.operationMap[req.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "operation not found")
	}

	switch operationType {
	case MPC_KEY_OPERATION:
		return o.mpcKeysClient.GetOperation(ctx, req)
	case MPC_WALLET_OPERATION:
		return o.mpcWalletsClient.GetOperation(ctx, req)
	case MPC_TRANSACTION_OPERATION:
		return o.mpcTransactionsClient.GetOperation(ctx, req)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "unknown operation type")
	}
}
