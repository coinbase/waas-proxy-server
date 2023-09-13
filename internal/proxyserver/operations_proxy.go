package proxyserver

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

// ListOperations calls ListOperations on the operations proxy.
func (p *ProxyServer) ListOperations(
	ctx context.Context,
	req *longrunningpb.ListOperationsRequest,
) (*longrunningpb.ListOperationsResponse, error) {
	return p.operationsProxy.ListOperations(ctx, req)
}

// GetOperation calls GetOperation on the operations proxy.
func (p *ProxyServer) GetOperation(
	ctx context.Context,
	req *longrunningpb.GetOperationRequest,
) (*longrunningpb.Operation, error) {
	return p.operationsProxy.GetOperation(ctx, req)
}

// DeleteOperation calls DeleteOperation on the operations proxy.
func (p *ProxyServer) DeleteOperation(
	ctx context.Context,
	req *longrunningpb.DeleteOperationRequest,
) (*emptypb.Empty, error) {
	return p.operationsProxy.DeleteOperation(ctx, req)
}

// CancelOperation calls CancelOperation on the operations proxy.
func (p *ProxyServer) CancelOperation(
	ctx context.Context,
	req *longrunningpb.CancelOperationRequest,
) (*emptypb.Empty, error) {
	return p.operationsProxy.CancelOperation(ctx, req)
}

// WaitOperation calls WaitOperation on the operations proxy.
func (p *ProxyServer) WaitOperation(
	ctx context.Context,
	req *longrunningpb.WaitOperationRequest,
) (*longrunningpb.Operation, error) {
	return p.operationsProxy.WaitOperation(ctx, req)
}
