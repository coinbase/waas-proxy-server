package mpckeys

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

// CreateSignature proxies the CreateSignature method.
func (m *mpcKeysProxy) CreateSignature(
	ctx context.Context,
	req *v1.CreateSignatureRequest,
) (*longrunningpb.Operation, error) {
	m.log.Info("CreateSignature")

	wrappedOp, err := m.mpcKeysClient.CreateSignature(ctx, req)
	if err != nil {
		return nil, err
	}

	m.operationMap[wrappedOp.Name()] = operations.MPC_KEY_OPERATION

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
