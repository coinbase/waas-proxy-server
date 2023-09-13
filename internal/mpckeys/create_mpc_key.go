package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// CreateMPCKey proxies the CreateMPCKey method.
func (m *mpcKeysProxy) CreateMPCKey(
	ctx context.Context,
	req *v1.CreateMPCKeyRequest,
) (*v1.MPCKey, error) {
	m.log.Info("CreateMPCKey")

	return m.mpcKeysClient.CreateMPCKey(ctx, req)
}
