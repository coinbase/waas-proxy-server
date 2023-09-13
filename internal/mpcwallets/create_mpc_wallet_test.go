package mpcwallets_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

func (s *ts) Test_CreateMPCWallet() {
	var (
		metadata = &mpcwalletspb.CreateMPCWalletMetadata{
			DeviceGroup: "pools/test-pool/deviceGroups/test-device-group",
		}

		metadataAny, _ = anypb.New(metadata)

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		createMPCWalletReq = &mpcwalletspb.CreateMPCWalletRequest{
			Parent: "pools/test-pool/wallets/test-wallet",
			Device: "devices/test-device",
		}

		newRequestFn = func() *mpcwalletspb.CreateMPCWalletRequest {
			return createMPCWalletReq
		}

		validMutation = func(req *mpcwalletspb.CreateMPCWalletRequest) *longrunningpb.Operation {
			s.CreatesMPCWallet(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *mpcwalletspb.CreateMPCWalletRequest,
			err error,
		) *longrunningpb.Operation {
			s.CreatesMPCWallet(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *mpcwalletspb.CreateMPCWalletRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcWalletsProxy.CreateMPCWallet(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
