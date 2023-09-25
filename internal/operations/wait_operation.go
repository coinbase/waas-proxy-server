package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// WaitOperation proxies the WaitOperation method.
func (o *operationsProxy) WaitOperation(context.Context, *longrunningpb.WaitOperationRequest) (*longrunningpb.Operation, error) {
	o.log.Info("WaitOperation")

	return nil, status.Errorf(codes.Unimplemented, "method WaitOperation not implemented")
}
