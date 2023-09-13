package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_CreateMPCKey tests CreateMPCKey with a range of scenarios.
func (s *ts) Test_CreateMPCKey() {
	var (
		mpcKey = &v1.MPCKey{
			Name: "pools/test-pool/deviceGroups/test-device-group/mpcKeys/test-mpc-key",
		}

		createMPCKeyReq = &v1.CreateMPCKeyRequest{
			Parent: "pools/test-pool/deviceGroups/test-device-group",
			MpcKey: mpcKey,
		}

		newRequestFn = func() *v1.CreateMPCKeyRequest {
			return createMPCKeyReq
		}

		validMutation = func(req *v1.CreateMPCKeyRequest) *v1.MPCKey {
			s.CreatesMPCKey(req, mpcKey, nil)

			return mpcKey
		}

		errorMutation = func(
			req *v1.CreateMPCKeyRequest,
			err error,
		) *v1.MPCKey {
			s.CreatesMPCKey(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.CreateMPCKeyRequest,
		) (*v1.MPCKey, error) {
			return s.mpcKeysProxy.CreateMPCKey(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
