package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Test_RevokeDevice tests RevokeDevice with a range of scenarios.
func (s *ts) Test_RevokeDevice() {
	var (
		device = &v1.Device{
			Name: "devices/test-device",
		}

		revokeDeviceReq = &v1.RevokeDeviceRequest{
			Name: device.GetName(),
		}

		newRequestFn = func() *v1.RevokeDeviceRequest {
			return revokeDeviceReq
		}

		validMutation = func(req *v1.RevokeDeviceRequest) *emptypb.Empty {
			s.RevokesDevice(req, nil)

			return &emptypb.Empty{}
		}

		errorMutation = func(
			req *v1.RevokeDeviceRequest,
			err error,
		) *emptypb.Empty {
			s.RevokesDevice(req, err)

			return &emptypb.Empty{}
		}

		testFn = func(
			ctx context.Context,
			req *v1.RevokeDeviceRequest,
		) (*emptypb.Empty, error) {
			return s.mpcKeysProxy.RevokeDevice(ctx, req)
		}
	)

	scenarios := map[string]struct {
		mutate             func(req *v1.RevokeDeviceRequest) *emptypb.Empty
		expectedStatusCode codes.Code
	}{
		"success": {
			mutate:             validMutation,
			expectedStatusCode: codes.OK,
		},
		"client error": {
			mutate: func(req *v1.RevokeDeviceRequest) *emptypb.Empty {
				return errorMutation(req, status.Errorf(codes.Internal, "boom"))
			},
			expectedStatusCode: codes.Internal,
		},
	}

	for name, scenario := range scenarios {
		scenario := scenario

		s.Run(name, func() {
			s.SetupTest()

			req := newRequestFn()

			var expectedResponse *emptypb.Empty

			if scenario.mutate != nil {
				expectedResponse = scenario.mutate(req)
			}

			resp, err := testFn(context.Background(), req)

			if scenario.expectedStatusCode == codes.OK {
				s.NoError(err)
				s.ProtoEqual(expectedResponse, resp)
			} else {
				s.StatusCodeEqual(scenario.expectedStatusCode, err)
				s.ProtoEqual(expectedResponse, resp)
			}
		})
	}
}
