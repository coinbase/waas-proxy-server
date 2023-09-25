package mpckeys

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// RevokeDevice proxies the RevokeDevice method.
func (m *mpcKeysProxy) RevokeDevice(
	ctx context.Context,
	req *v1.RevokeDeviceRequest,
) (*emptypb.Empty, error) {
	m.log.Info("RevokeDevice")

	return &emptypb.Empty{}, m.mpcKeysClient.RevokeDevice(ctx, req)
}
