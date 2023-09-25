package mpckeys_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

// Test_CreateDeviceGroup tests CreateDeviceGroup with a range of scenarios.
func (s *ts) Test_CreateDeviceGroup() {
	var (
		deviceGroup = &v1.DeviceGroup{
			Name:    "pools/test-pool/deviceGroups/test-device-group",
			Devices: []string{"devices/test-device"},
		}

		metadata = &v1.CreateDeviceGroupMetadata{
			DeviceGroup: deviceGroup.GetName(),
		}

		metadataAny, _ = anypb.New(metadata)

		createDeviceGroupReq = &v1.CreateDeviceGroupRequest{
			Parent:        "pools/test-pool",
			DeviceGroup:   deviceGroup,
			DeviceGroupId: "test-device-group",
		}

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		newRequestFn = func() *v1.CreateDeviceGroupRequest {
			return createDeviceGroupReq
		}

		validMutation = func(req *v1.CreateDeviceGroupRequest) *longrunningpb.Operation {
			s.CreatesDeviceGroup(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *v1.CreateDeviceGroupRequest,
			err error,
		) *longrunningpb.Operation {
			s.CreatesDeviceGroup(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.CreateDeviceGroupRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcKeysProxy.CreateDeviceGroup(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
