package mpckeys_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

// Test_PrepareDeviceArchive tests PrepareDeviceArchive with a range of scenarios.
func (s *ts) Test_PrepareDeviceArchive() {
	var (
		deviceGroup = &v1.DeviceGroup{
			Name:    "pools/test-pool/deviceGroups/test-device-group",
			Devices: []string{"devices/test-device"},
		}

		metadata = &v1.PrepareDeviceArchiveMetadata{
			DeviceGroup: deviceGroup.GetName(),
		}

		metadataAny, _ = anypb.New(metadata)

		prepareDeviceArchiveReq = &v1.PrepareDeviceArchiveRequest{
			DeviceGroup: deviceGroup.GetName(),
			Device:      "devices/test-device",
		}

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		newRequestFn = func() *v1.PrepareDeviceArchiveRequest {
			return prepareDeviceArchiveReq
		}

		validMutation = func(req *v1.PrepareDeviceArchiveRequest) *longrunningpb.Operation {
			s.PreparesDeviceArchive(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *v1.PrepareDeviceArchiveRequest,
			err error,
		) *longrunningpb.Operation {
			s.PreparesDeviceArchive(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.PrepareDeviceArchiveRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcKeysProxy.PrepareDeviceArchive(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
