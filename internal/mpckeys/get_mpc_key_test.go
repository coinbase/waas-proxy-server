package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_GetMPCKey tests GetMPCKey with a range of scenarios.
func (s *ts) Test_GetMPCKey() {
	var (
		mpcKey = &v1.MPCKey{
			Name: "pools/test-pool/deviceGroups/test-device-group/mpcKeys/test-mpc-key",
		}

		getMPCKeyReq = &v1.GetMPCKeyRequest{
			Name: mpcKey.GetName(),
		}

		newRequestFn = func() *v1.GetMPCKeyRequest {
			return getMPCKeyReq
		}

		validMutation = func(req *v1.GetMPCKeyRequest) *v1.MPCKey {
			s.GetsMPCKey(req, mpcKey, nil)

			return mpcKey
		}

		errorMutation = func(
			req *v1.GetMPCKeyRequest,
			err error,
		) *v1.MPCKey {
			s.GetsMPCKey(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetMPCKeyRequest,
		) (*v1.MPCKey, error) {
			return s.mpcKeysProxy.GetMPCKey(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
