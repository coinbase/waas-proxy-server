package proxyserver

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
)

// CreateMPCTransaction calls CreateMPCTransaction on the mpc transactions proxy.
func (p *ProxyServer) CreateMPCTransaction(
	ctx context.Context,
	req *v1.CreateMPCTransactionRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcTransactionsProxy.CreateMPCTransaction(ctx, req)
}

// GetMPCTransaction calls GetMPCTransaction on the mpc transactions proxy.
func (p *ProxyServer) GetMPCTransaction(
	ctx context.Context,
	req *v1.GetMPCTransactionRequest,
) (*v1.MPCTransaction, error) {
	return p.mpcTransactionsProxy.GetMPCTransaction(ctx, req)
}

// ListMPCTransactions calls ListMPCTransactions on the mpc transactions proxy.
func (p *ProxyServer) ListMPCTransactions(
	ctx context.Context,
	req *v1.ListMPCTransactionsRequest,
) (*v1.ListMPCTransactionsResponse, error) {
	return p.mpcTransactionsProxy.ListMPCTransactions(ctx, req)
}
