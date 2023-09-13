package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// ListMPCOperations proxies the ListMPCOperations method.
func (m *mpcKeysProxy) ListMPCOperations(
	ctx context.Context,
	req *v1.ListMPCOperationsRequest,
) (*v1.ListMPCOperationsResponse, error) {
	m.log.Info("ListMPCOperations")

	return m.mpcKeysClient.ListMPCOperations(ctx, req)
}
