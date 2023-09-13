package mpctransactions

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

// CreateMPCTransaction proxies the CreateMPCTransaction method.
func (m *mpcTransactionsProxy) CreateMPCTransaction(
	ctx context.Context,
	req *v1.CreateMPCTransactionRequest,
) (*longrunningpb.Operation, error) {
	m.log.Info("CreateMPCTransaction")

	wrappedOp, err := m.mpcTransactionsClient.CreateMPCTransaction(ctx, req)
	if err != nil {
		return nil, err
	}

	m.operationMap[wrappedOp.Name()] = operations.MPC_TRANSACTION_OPERATION

	metadata, err := wrappedOp.Metadata()
	if err != nil {
		return nil, err
	}

	metadataAny, err := anypb.New(metadata)
	if err != nil {
		return nil, err
	}

	result, pollErr := wrappedOp.Poll(ctx)

	op := &longrunningpb.Operation{
		Name:     wrappedOp.Name(),
		Done:     wrappedOp.Done(),
		Metadata: metadataAny,
	}

	if result != nil {
		responseAny, err := anypb.New(result)
		if err != nil {
			return nil, err
		}

		op.Result = &longrunningpb.Operation_Response{
			Response: responseAny,
		}
	} else if pollErr != nil {
		op.Result = &longrunningpb.Operation_Error{
			Error: status.Convert(pollErr).Proto(),
		}
	}

	return op, nil
}
