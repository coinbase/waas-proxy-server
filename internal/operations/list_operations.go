package operations

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListOperations proxies the ListOperations method.
func (o *operationsProxy) ListOperations(context.Context, *longrunningpb.ListOperationsRequest) (*longrunningpb.ListOperationsResponse, error) {
	o.log.Info("ListOperations")

	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
