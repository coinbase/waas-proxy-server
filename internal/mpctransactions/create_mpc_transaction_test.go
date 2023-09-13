package mpctransactions_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	"google.golang.org/protobuf/types/known/anypb"
)

// Test_CreateMPCTransaction tests CreatemPCTransaction with a range of scenarios.
func (s *ts) Test_CreateMPCTransaction() {
	var (
		metadata = &v1.CreateMPCTransactionMetadata{
			DeviceGroup:    "pools/test-pool/deviceGroups/test-device-group",
			MpcTransaction: "pools/test-pool/mpcWallets/test-wallet/mpcTransactions/test-mpc-transaction",
		}

		metadataAny, _ = anypb.New(metadata)

		operation = &longrunningpb.Operation{
			Name:     "operations/test-operation",
			Metadata: metadataAny,
			Done:     false,
			Result:   nil,
		}

		createMPCTransactionReq = &v1.CreateMPCTransactionRequest{
			Parent: "pools/test-pool/wallets/test-wallet",
			Input: &typespb.TransactionInput{
				Input: &typespb.TransactionInput_EthereumRlpInput{
					EthereumRlpInput: &ethereumpb.RLPTransaction{
						TransactionRlp: []byte("tx"),
					},
				},
			},
			OverrideNonce: false,
		}

		newRequestFn = func() *v1.CreateMPCTransactionRequest {
			return createMPCTransactionReq
		}

		validMutation = func(req *v1.CreateMPCTransactionRequest) *longrunningpb.Operation {
			s.CreatesMPCTransaction(req, metadata, metadataAny, nil)

			return operation
		}

		errorMutation = func(
			req *v1.CreateMPCTransactionRequest,
			err error,
		) *longrunningpb.Operation {
			s.CreatesMPCTransaction(req, nil, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.CreateMPCTransactionRequest,
		) (*longrunningpb.Operation, error) {
			return s.mpcTransactionProxy.CreateMPCTransaction(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
