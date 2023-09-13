package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// GetDevice proxies the GetDevice method.
func (m *mpcKeysProxy) GetDevice(
	ctx context.Context,
	req *v1.GetDeviceRequest,
) (*v1.Device, error) {
	m.log.Info("GetDevice")

	return m.mpcKeysClient.GetDevice(ctx, req)
}
