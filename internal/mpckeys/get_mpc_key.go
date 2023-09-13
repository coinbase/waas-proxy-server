package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// GetMPCKey proxies the GetMPCKey method.
func (m *mpcKeysProxy) GetMPCKey(
	ctx context.Context,
	req *v1.GetMPCKeyRequest,
) (*v1.MPCKey, error) {
	m.log.Info("GetMPCKey")

	return m.mpcKeysClient.GetMPCKey(ctx, req)
}
