package mpckeys_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

// Test_AddDevice tests AddDevice with a range of scenarios.
func (s *ts) Test_AddDevice() {
	var (
		deviceGroup = &v1.DeviceGroup{
			Name:    "pools/test-pool/deviceGroups/test-device-group",
			Devices: []string{"devices/test-device"},
		}

		metadata = &v1.AddDeviceMetadata{
			DeviceGroup:          deviceGroup.GetName(),
			ParticipatingDevices: deviceGroup.GetDevices(),
		}

		metadataAny, _ = anypb.New(metadata)

		addDeviceReq = &v1.AddDeviceRequest{
			DeviceGroup: deviceGroup.GetName(),
			Device:      deviceGroup.GetDevices()[0],
		}

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		newRequestFn = func() *v1.AddDeviceRequest {
			return addDeviceReq
		}

		validMutation = func(req *v1.AddDeviceRequest) *longrunningpb.Operation {
			s.AddsDevice(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *v1.AddDeviceRequest,
			err error,
		) *longrunningpb.Operation {
			s.AddsDevice(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.AddDeviceRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcKeysProxy.AddDevice(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
