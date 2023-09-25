package mpctransactions

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
)

// GetMPCTransaction proxies the GetMPCTransaction method.
func (m *mpcTransactionsProxy) GetMPCTransaction(
	ctx context.Context,
	req *v1.GetMPCTransactionRequest,
) (*v1.MPCTransaction, error) {
	m.log.Info("GetMPCTransaction")

	return m.mpcTransactionsClient.GetMPCTransaction(ctx, req)
}
