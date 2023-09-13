package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// GetDeviceGroup proxies the GetDeviceGroup method.
func (m *mpcKeysProxy) GetDeviceGroup(
	ctx context.Context,
	req *v1.GetDeviceGroupRequest,
) (*v1.DeviceGroup, error) {
	m.log.Info("GetDeviceGroup")

	return m.mpcKeysClient.GetDeviceGroup(ctx, req)
}
