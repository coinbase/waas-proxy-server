package mpckeys_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

// Test_ListMPCOperation tests ListMPCOperation with a range of scenarios.
func (s *ts) Test_ListMPCOperation() {
	var (
		bytes                = make([]byte, 32)
		listMPCOperationResp = &v1.ListMPCOperationsResponse{
			MpcOperations: []*v1.MPCOperation{{
				Name:    "pools/test-pool/deviceGroups/test-device-group/mpcOperations/test-mpc-operation",
				MpcData: bytes,
			}},
		}

		listMPCOperationsReq = &v1.ListMPCOperationsRequest{
			Parent: "pools/test-pool/deviceGroups/test-device-group",
		}

		newRequestFn = func() *v1.ListMPCOperationsRequest {
			return listMPCOperationsReq
		}

		validMutation = func(req *v1.ListMPCOperationsRequest) *v1.ListMPCOperationsResponse {
			s.ListsMPCOperations(req, listMPCOperationResp, nil)

			return listMPCOperationResp
		}

		errorMutation = func(
			req *v1.ListMPCOperationsRequest,
			err error,
		) *v1.ListMPCOperationsResponse {
			s.ListsMPCOperations(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListMPCOperationsRequest,
		) (*v1.ListMPCOperationsResponse, error) {
			return s.mpcKeysProxy.ListMPCOperations(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
