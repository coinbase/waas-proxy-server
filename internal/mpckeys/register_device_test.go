package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_RegisterDevice tests RegisterDevice with a range of scenarios.
func (s *ts) Test_RegisterDevice() {
	var (
		bytes  = make([]byte, 32)
		device = &v1.Device{
			Name: "devices/test-device",
		}

		registerDeviceReq = &v1.RegisterDeviceRequest{
			RegistrationData: bytes,
		}

		newRequestFn = func() *v1.RegisterDeviceRequest {
			return registerDeviceReq
		}

		validMutation = func(req *v1.RegisterDeviceRequest) *v1.Device {
			s.RegistersDevice(req, device, nil)

			return device
		}

		errorMutation = func(
			req *v1.RegisterDeviceRequest,
			err error,
		) *v1.Device {
			s.RegistersDevice(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.RegisterDeviceRequest,
		) (*v1.Device, error) {
			return s.mpcKeysProxy.RegisterDevice(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
