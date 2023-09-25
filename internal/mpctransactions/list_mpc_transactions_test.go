package mpctransactions_test

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	cryptotypespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/crypto/types/v1"
)

// Test_ListMPCTransactions tests ListMPCTransactions with a range of scenarios.
func (s *ts) Test_ListMPCTransactions() {
	var (
		bytes           = make([]byte, 32)
		mpcTransaction1 = &v1.MPCTransaction{
			Name:          "pools/test-pool/mpcWallets/test-wallet/mpcTransactions/test-mpc-transaction-1",
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

		mpcTransaction2 = &v1.MPCTransaction{
			Name:          "pools/test-pool/mpcWallets/test-wallet/mpcTransactions/test-mpc-transaction-2",
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

		listMPCTransactionsReq = &v1.ListMPCTransactionsRequest{
			PageSize:  5,
			PageToken: "",
		}

		listMPCTransactionsResp = &v1.ListMPCTransactionsResponse{
			MpcTransactions: []*v1.MPCTransaction{mpcTransaction1, mpcTransaction2},
			NextPageToken:   "test-next-page-token",
		}

		newRequestFn = func() *v1.ListMPCTransactionsRequest {
			return listMPCTransactionsReq
		}

		validMutation = func(req *v1.ListMPCTransactionsRequest) *v1.ListMPCTransactionsResponse {
			s.ListsMPCTransactions(req, listMPCTransactionsResp, nil)

			return listMPCTransactionsResp
		}

		errorMutation = func(
			req *v1.ListMPCTransactionsRequest,
			err error,
		) *v1.ListMPCTransactionsResponse {
			s.ListsMPCTransactions(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *v1.ListMPCTransactionsRequest,
		) (*v1.ListMPCTransactionsResponse, error) {
			return s.mpcTransactionProxy.ListMPCTransactions(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
