package mpckeys

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
)

// MPCKeysProxy is an interface with methods to proxy requests to the mpc key service.
type MPCKeysProxy interface {
	RegisterDevice(context.Context, *v1.RegisterDeviceRequest) (*v1.Device, error)
	GetDevice(context.Context, *v1.GetDeviceRequest) (*v1.Device, error)
	CreateDeviceGroup(context.Context, *v1.CreateDeviceGroupRequest) (*longrunningpb.Operation, error)
	GetDeviceGroup(context.Context, *v1.GetDeviceGroupRequest) (*v1.DeviceGroup, error)
	ListMPCOperations(context.Context, *v1.ListMPCOperationsRequest) (*v1.ListMPCOperationsResponse, error)
	CreateMPCKey(context.Context, *v1.CreateMPCKeyRequest) (*v1.MPCKey, error)
	GetMPCKey(context.Context, *v1.GetMPCKeyRequest) (*v1.MPCKey, error)
	CreateSignature(context.Context, *v1.CreateSignatureRequest) (*longrunningpb.Operation, error)
	PrepareDeviceArchive(context.Context, *v1.PrepareDeviceArchiveRequest) (*longrunningpb.Operation, error)
	PrepareDeviceBackup(context.Context, *v1.PrepareDeviceBackupRequest) (*longrunningpb.Operation, error)
	AddDevice(context.Context, *v1.AddDeviceRequest) (*longrunningpb.Operation, error)
	RevokeDevice(context.Context, *v1.RevokeDeviceRequest) (*emptypb.Empty, error)
}

// mpcKeysProxy implements the MPCKeysProxy interface.
type mpcKeysProxy struct {
	log *zap.Logger

	mpcKeysClient clientsv1.MPCKeyServiceClient

	operationMap map[string]operations.OperationType
}

// NewMPCKeysProxy instantiates a new mpcKeysProxy.
func NewMPCKeysProxy(
	ctx context.Context,
	log *zap.Logger,
	mpcKeysClient clientsv1.MPCKeyServiceClient,
	operationMap map[string]operations.OperationType,
) (*mpcKeysProxy, error) {
	return &mpcKeysProxy{
		log:           log,
		mpcKeysClient: mpcKeysClient,
		operationMap:  operationMap,
	}, nil
}
