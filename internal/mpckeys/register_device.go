package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// RegisterDevice proxies the RegisterDevice method.
func (m *mpcKeysProxy) RegisterDevice(
	ctx context.Context,
	req *v1.RegisterDeviceRequest,
) (*v1.Device, error) {
	m.log.Info("RegisterDevice")

	return m.mpcKeysClient.RegisterDevice(ctx, req)
}
