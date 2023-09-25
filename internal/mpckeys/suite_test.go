package mpckeys_test

import (
	"context"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/mpckeys"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type ts struct {
	testutils.TestSuite

	mpcKeysClient                        *mocks.MPCKeyServiceClient
	mpcKeysProxy                         mpckeys.MPCKeysProxy
	operationsMap                        map[string]operations.OperationType
	createDeviceGroupWrappedOperation    *mocks.ClientWrappedCreateDeviceGroupOperation
	createSignatureWrappedOperation      *mocks.ClientWrappedCreateSignatureOperation
	prepareDeviceArchiveWrappedOperation *mocks.ClientWrappedPrepareDeviceArchiveOperation
	prepareDeviceBackupWrappedOperation  *mocks.ClientWrappedPrepareDeviceBackupOperation
	addDeviceWrappedOperation            *mocks.ClientWrappedAddDeviceOperation
}

func (s *ts) SetupTest() {
	var err error

	s.mpcKeysClient = mocks.NewMPCKeyServiceClient(s.T())

	s.operationsMap = make(map[string]operations.OperationType)

	s.createDeviceGroupWrappedOperation = mocks.NewClientWrappedCreateDeviceGroupOperation(s.T())

	s.createSignatureWrappedOperation = mocks.NewClientWrappedCreateSignatureOperation(s.T())

	s.prepareDeviceArchiveWrappedOperation = mocks.NewClientWrappedPrepareDeviceArchiveOperation(s.T())

	s.prepareDeviceBackupWrappedOperation = mocks.NewClientWrappedPrepareDeviceBackupOperation(s.T())

	s.addDeviceWrappedOperation = mocks.NewClientWrappedAddDeviceOperation(s.T())

	s.mpcKeysProxy, err = mpckeys.NewMPCKeysProxy(context.Background(), zap.NewNop(), s.mpcKeysClient, s.operationsMap)
	if err != nil {
		s.FailNow("failed to initialize mpc keys proxy", err)
	}
}

func TestMPCKeysProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// RegistersDevice implements the mocks/assertions for registering a device.
func (s *ts) RegistersDevice(
	req *v1.RegisterDeviceRequest,
	resp *v1.Device,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"RegisterDevice",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsDevice implements the mocks/assertions for getting a device.
func (s *ts) GetsDevice(
	req *v1.GetDeviceRequest,
	resp *v1.Device,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"GetDevice",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// CreatesDeviceGroup implements the mocks/assertions for creating a device group.
func (s *ts) CreatesDeviceGroup(
	req *v1.CreateDeviceGroupRequest,
	metadata *v1.CreateDeviceGroupMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"CreateDeviceGroup",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.createDeviceGroupWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.createDeviceGroupWrappedOperation.On(
		"Name",
	).Return(operation.GetName())

	s.createDeviceGroupWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.createDeviceGroupWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.createDeviceGroupWrappedOperation, nil)
}

// GetsDeviceGroup implements the mocks/assertions for getting a device group.
func (s *ts) GetsDeviceGroup(
	req *v1.GetDeviceGroupRequest,
	resp *v1.DeviceGroup,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"GetDeviceGroup",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ListsMPCOperations implements the mocks/assertions for listing mpc operaions.
func (s *ts) ListsMPCOperations(
	req *v1.ListMPCOperationsRequest,
	resp *v1.ListMPCOperationsResponse,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"ListMPCOperations",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// CreatesMPCKey implements the mocks/assertions for creating an mpc key.
func (s *ts) CreatesMPCKey(
	req *v1.CreateMPCKeyRequest,
	resp *v1.MPCKey,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"CreateMPCKey",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsMPCKey implements the mocks/assertions for getting an mpc key.
func (s *ts) GetsMPCKey(
	req *v1.GetMPCKeyRequest,
	resp *v1.MPCKey,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"GetMPCKey",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// CreatesSignature implements the mocks/assertions for creating a signature.
func (s *ts) CreatesSignature(
	req *v1.CreateSignatureRequest,
	metadata *v1.CreateSignatureMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"CreateSignature",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.createSignatureWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.createSignatureWrappedOperation.On(
		"Name",
	).Return(operation.GetName())

	s.createSignatureWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.createSignatureWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.createSignatureWrappedOperation, nil)
}

// PreparesDeviceArchive implements the mocks/assertions for preparing device archive.
func (s *ts) PreparesDeviceArchive(
	req *v1.PrepareDeviceArchiveRequest,
	metadata *v1.PrepareDeviceArchiveMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"PrepareDeviceArchive",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.prepareDeviceArchiveWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.prepareDeviceArchiveWrappedOperation.On(
		"Name",
	).Return(operation.GetName())

	s.prepareDeviceArchiveWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.prepareDeviceArchiveWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.prepareDeviceArchiveWrappedOperation, nil)
}

// PreparesDeviceBackup implements the mocks/assertions for preparing device backup.
func (s *ts) PreparesDeviceBackup(
	req *v1.PrepareDeviceBackupRequest,
	metadata *v1.PrepareDeviceBackupMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"PrepareDeviceBackup",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.prepareDeviceBackupWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.prepareDeviceBackupWrappedOperation.On(
		"Name",
	).Return(operation.GetName())

	s.prepareDeviceBackupWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.prepareDeviceBackupWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.prepareDeviceBackupWrappedOperation, nil)
}

// AddsDevice implements the mocks/assertions for adding a device.
func (s *ts) AddsDevice(
	req *v1.AddDeviceRequest,
	metadata *v1.AddDeviceMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"AddDevice",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.addDeviceWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.addDeviceWrappedOperation.On(
		"Name",
	).Return(operation.GetName())

	s.addDeviceWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.addDeviceWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.addDeviceWrappedOperation, nil)
}

// RevokesDevice implements the mocks/assertions for revoking a device.
func (s *ts) RevokesDevice(
	req *v1.RevokeDeviceRequest,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"RevokeDevice",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(err)
	}

	return call.Return(nil)
}

func RunScenarios[
	RequestType protoreflect.ProtoMessage,
	ResponseType protoreflect.ProtoMessage,
	NewRequestFn func() RequestType,
	MutationFn func(RequestType) ResponseType,
	ErrorMutationFn func(RequestType, error) ResponseType,
	TestFn func(context.Context, RequestType) (ResponseType, error),
](
	s *ts,
	newRequest NewRequestFn,
	validMutation MutationFn,
	errorMutation ErrorMutationFn,
	testFunction TestFn,
) {
	scenarios := map[string]struct {
		mutate             MutationFn
		expectedStatusCode codes.Code
	}{
		"success": {
			mutate:             validMutation,
			expectedStatusCode: codes.OK,
		},
		"client error": {
			mutate: func(req RequestType) ResponseType {
				return errorMutation(req, status.Errorf(codes.Internal, "boom"))
			},
			expectedStatusCode: codes.Internal,
		},
	}

	for name, scenario := range scenarios {
		scenario := scenario

		s.Run(name, func() {
			s.SetupTest()

			req := newRequest()

			var expectedResponse ResponseType

			if scenario.mutate != nil {
				expectedResponse = scenario.mutate(req)
			}

			resp, err := testFunction(context.Background(), req)

			if scenario.expectedStatusCode == codes.OK {
				s.NoError(err)
				s.ProtoEqual(expectedResponse, resp)
			} else {
				s.StatusCodeEqual(scenario.expectedStatusCode, err)
				s.Nil(resp)
			}
		})
	}
}
