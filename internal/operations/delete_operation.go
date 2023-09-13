package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteOperation proxies the DeleteOperation method.
func (o *operationsProxy) DeleteOperation(context.Context, *longrunningpb.DeleteOperationRequest) (*emptypb.Empty, error) {
	o.log.Info("DeleteOperation")

	return nil, status.Errorf(codes.Unimplemented, "method DeleteOperation not implemented")
}
