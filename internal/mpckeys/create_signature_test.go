package mpckeys_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

// Test_CreateSignature tests CreateSignature with a range of scenarios.
func (s *ts) Test_CreateSignature() {
	var (
		signature = &v1.Signature{
			Name:    "pools/test-pool/deviceGroups/test-device-group/mpcKeys/test-mpc-key/signatures/test-signature",
			Payload: []byte("test-payload"),
		}

		metadata = &v1.CreateSignatureMetadata{
			DeviceGroup: signature.GetName(),
		}

		metadataAny, _ = anypb.New(metadata)

		createSignatureReq = &v1.CreateSignatureRequest{
			Parent:    "pools/test-pool/deviceGroups/test-device-group/mpcKeys/test-mpc-key",
			Signature: signature,
		}

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		newRequestFn = func() *v1.CreateSignatureRequest {
			return createSignatureReq
		}

		validMutation = func(req *v1.CreateSignatureRequest) *longrunningpb.Operation {
			s.CreatesSignature(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *v1.CreateSignatureRequest,
			err error,
		) *longrunningpb.Operation {
			s.CreatesSignature(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.CreateSignatureRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcKeysProxy.CreateSignature(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
