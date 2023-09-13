package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_GetDevice tests GetDevice with a range of scenarios.
func (s *ts) Test_GetDevice() {
	var (
		device = &v1.Device{
			Name: "devices/test-device",
		}

		getDeviceReq = &v1.GetDeviceRequest{
			Name: "devices/test-device",
		}

		newRequestFn = func() *v1.GetDeviceRequest {
			return getDeviceReq
		}

		validMutation = func(req *v1.GetDeviceRequest) *v1.Device {
			s.GetsDevice(req, device, nil)

			return device
		}

		errorMutation = func(
			req *v1.GetDeviceRequest,
			err error,
		) *v1.Device {
			s.GetsDevice(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetDeviceRequest,
		) (*v1.Device, error) {
			return s.mpcKeysProxy.GetDevice(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
