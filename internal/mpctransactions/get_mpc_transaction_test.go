package mpctransactions_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	cryptotypespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/crypto/types/v1"
)

// Test_GetMPCTransaction tests GetMPCTransaction with a range of scenarios.
func (s *ts) Test_GetMPCTransaction() {
	var (
		bytes          = make([]byte, 32)
		mpcTransaction = &v1.MPCTransaction{
			Name:          "pools/test-pool/mpcWallets/test-wallet/mpcTransactions/test-mpc-transaction",
			Network:       "networks/test-network",
			FromAddresses: []string{"0xsender"},
			State:         0,
			Transaction: &typespb.Transaction{
				Input: &typespb.TransactionInput{
					Input: &typespb.TransactionInput_EthereumRlpInput{
						EthereumRlpInput: &ethereumpb.RLPTransaction{
							TransactionRlp: []byte("tx"),
						},
					},
				},
				RequiredSignatures: []*typespb.RequiredSignature{
					{
						Payload: []byte("payload"),
						Signature: &cryptotypespb.Signature{
							Signature: &cryptotypespb.Signature_EcdsaSignature{
								EcdsaSignature: &cryptotypespb.ECDSASignature{
									R: bytes,
									S: bytes,
									V: 1,
								},
							},
						},
					},
				},
			},
		}

		getMPCTransactionReq = &v1.GetMPCTransactionRequest{
			Name: mpcTransaction.GetName(),
		}

		newRequestFn = func() *v1.GetMPCTransactionRequest {
			return getMPCTransactionReq
		}

		validMutation = func(req *v1.GetMPCTransactionRequest) *v1.MPCTransaction {
			s.GetsMPCTransaction(req, mpcTransaction, nil)

			return mpcTransaction
		}

		errorMutation = func(
			req *v1.GetMPCTransactionRequest,
			err error,
		) *v1.MPCTransaction {
			s.GetsMPCTransaction(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.GetMPCTransactionRequest,
		) (*v1.MPCTransaction, error) {
			return s.mpcTransactionProxy.GetMPCTransaction(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
