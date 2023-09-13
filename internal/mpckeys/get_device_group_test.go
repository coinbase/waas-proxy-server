package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_GetDeviceGroup tests GetDeviceGroup with a range of scenarios.
func (s *ts) Test_GetDeviceGroup() {
	var (
		deviceGroup = &v1.DeviceGroup{
			Name:    "pools/test-pool/deviceGroups/test-device-group",
			Devices: []string{"devices/test-device"},
		}

		getDeviceGroupReq = &v1.GetDeviceGroupRequest{
			Name: deviceGroup.GetName(),
		}

		newRequestFn = func() *v1.GetDeviceGroupRequest {
			return getDeviceGroupReq
		}

		validMutation = func(req *v1.GetDeviceGroupRequest) *v1.DeviceGroup {
			s.GetsDeviceGroup(req, deviceGroup, nil)

			return deviceGroup
		}

		errorMutation = func(
			req *v1.GetDeviceGroupRequest,
			err error,
		) *v1.DeviceGroup {
			s.GetsDeviceGroup(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetDeviceGroupRequest,
		) (*v1.DeviceGroup, error) {
			return s.mpcKeysProxy.GetDeviceGroup(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
