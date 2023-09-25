package protocols_test

import (
	"context"

	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// Test_ConstructTransferTransaction tests ConstructTransferTransaction with a range of scenarios.
func (s *ts) Test_ConstructTransferTransaction() {
	var (
		v1Tx        = s.newTxForConstruct()
		networkName = "networks/test-network"

		constructTransferTransactionReq = &protocolspb.ConstructTransferTransactionRequest{
			Network:   networkName,
			Asset:     networkName + "/assets/test-asset",
			Sender:    "0xsender",
			Recipient: "0xrecipient",
			Amount:    "0xdeadbeef",
			Nonce:     1,
			Fee: &typespb.TransactionFee{
				Fee: &typespb.TransactionFee_EthereumFee{
					EthereumFee: &ethereumpb.DynamicFeeInput{
						MaxPriorityFeePerGas: "0x100",
						MaxFeePerGas:         "0x10000",
					},
				},
			},
		}

		newRequestFn = func() *protocolspb.ConstructTransferTransactionRequest {
			return constructTransferTransactionReq
		}

		validMutation = func(req *protocolspb.ConstructTransferTransactionRequest) *typespb.Transaction {
			s.ConstructsTransferTransaction(req, v1Tx, nil)

			return v1Tx
		}

		errorMutation = func(
			req *protocolspb.ConstructTransferTransactionRequest,
			err error,
		) *typespb.Transaction {
			s.ConstructsTransferTransaction(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *protocolspb.ConstructTransferTransactionRequest,
		) (*typespb.Transaction, error) {
			return s.protocolsProxy.ConstructTransferTransaction(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
