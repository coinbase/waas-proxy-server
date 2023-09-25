package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CancelOperation proxies the CancelOperation method.
func (o *operationsProxy) CancelOperation(context.Context, *longrunningpb.CancelOperationRequest) (*emptypb.Empty, error) {
	o.log.Info("CancelOperation")

	return nil, status.Errorf(codes.Unimplemented, "method CancelOperation not implemented")
}
