package proxyserver

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// RegisterDevice calls RegisterDevice on the mpc keys proxy.
func (p *ProxyServer) RegisterDevice(
	ctx context.Context,
	req *v1.RegisterDeviceRequest,
) (*v1.Device, error) {
	return p.mpcKeysProxy.RegisterDevice(ctx, req)
}

// GetDevice calls GetDevice on the mpc keys proxy.
func (p *ProxyServer) GetDevice(
	ctx context.Context,
	req *v1.GetDeviceRequest,
) (*v1.Device, error) {
	return p.mpcKeysProxy.GetDevice(ctx, req)
}

// CreateDeviceGroup calls CreateDeviceGroup on the mpc keys proxy.
func (p *ProxyServer) CreateDeviceGroup(
	ctx context.Context,
	req *v1.CreateDeviceGroupRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcKeysProxy.CreateDeviceGroup(ctx, req)
}

// GetDeviceGroup calls GetDeviceGroup on the mpc keys proxy.
func (p *ProxyServer) GetDeviceGroup(
	ctx context.Context,
	req *v1.GetDeviceGroupRequest,
) (*v1.DeviceGroup, error) {
	return p.mpcKeysProxy.GetDeviceGroup(ctx, req)
}

// ListMPCOperations calls ListMPCOperations on the mpc keys proxy.
func (p *ProxyServer) ListMPCOperations(
	ctx context.Context,
	req *v1.ListMPCOperationsRequest,
) (*v1.ListMPCOperationsResponse, error) {
	return p.mpcKeysProxy.ListMPCOperations(ctx, req)
}

// CreateMPCKey calls CreateMPCKey on the mpc keys proxy.
func (p *ProxyServer) CreateMPCKey(
	ctx context.Context,
	req *v1.CreateMPCKeyRequest,
) (*v1.MPCKey, error) {
	return p.mpcKeysProxy.CreateMPCKey(ctx, req)
}

// GetMPCKey calls GetMPCKey on the mpc keys proxy.
func (p *ProxyServer) GetMPCKey(
	ctx context.Context,
	req *v1.GetMPCKeyRequest,
) (*v1.MPCKey, error) {
	return p.mpcKeysProxy.GetMPCKey(ctx, req)
}

// CreateSignature calls CreateSignature on the mpc keys proxy.
func (p *ProxyServer) CreateSignature(
	ctx context.Context,
	req *v1.CreateSignatureRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcKeysProxy.CreateSignature(ctx, req)
}

// PrepareDeviceArchive calls PrepareDeviceArchive on the mpc keys proxy.
func (p *ProxyServer) PrepareDeviceArchive(
	ctx context.Context,
	req *v1.PrepareDeviceArchiveRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcKeysProxy.PrepareDeviceArchive(ctx, req)
}

// PrepareDeviceBackup calls PrepareDeviceBackup on the mpc keys proxy.
func (p *ProxyServer) PrepareDeviceBackup(
	ctx context.Context,
	req *v1.PrepareDeviceBackupRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcKeysProxy.PrepareDeviceBackup(ctx, req)
}

// AddDevice calls AddDevice on the mpc keys proxy.
func (p *ProxyServer) AddDevice(
	ctx context.Context,
	req *v1.AddDeviceRequest,
) (*longrunningpb.Operation, error) {
	return p.mpcKeysProxy.AddDevice(ctx, req)
}

// RevokeDevice calls RevokeDevice on the mpc keys proxy.
func (p *ProxyServer) RevokeDevice(
	ctx context.Context,
	req *v1.RevokeDeviceRequest,
) (*emptypb.Empty, error) {
	return p.mpcKeysProxy.RevokeDevice(ctx, req)
}
