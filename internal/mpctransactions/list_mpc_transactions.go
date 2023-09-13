package mpctransactions

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
)

// ListMPCTransactions proxies the ListMPCTransactions method.
func (m *mpcTransactionsProxy) ListMPCTransactions(
	ctx context.Context,
	req *v1.ListMPCTransactionsRequest,
) (*v1.ListMPCTransactionsResponse, error) {
	m.log.Info("ListMPCTransactions")

	iter := m.mpcTransactionsClient.ListMPCTransactions(ctx, req)

	if _, err := iter.Next(); err != nil && err != iterator.Done {
		m.log.Error("Error iterating over assets", zap.Error(err))

		return nil, err
	}

	return iter.Response(), nil
}
