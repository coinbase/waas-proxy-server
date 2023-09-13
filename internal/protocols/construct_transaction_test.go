package protocols_test

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// Test_ConstructTransaction tests ConstructTransaction with a range of scenarios.
func (s *ts) Test_ConstructTransaction() {
	var (
		v1Tx        = s.newTxForConstruct()
		networkName = "networks/test-network"

		constructTransactionReq = &protocolspb.ConstructTransactionRequest{
			Network: networkName,
			Input:   v1Tx.GetInput(),
		}

		newRequestFn = func() *protocolspb.ConstructTransactionRequest {
			return constructTransactionReq
		}

		validMutation = func(req *protocolspb.ConstructTransactionRequest) *typespb.Transaction {
			s.ConstructsTransaction(req, v1Tx, nil)

			return v1Tx
		}

		errorMutation = func(
			req *protocolspb.ConstructTransactionRequest,
			err error,
		) *typespb.Transaction {
			s.ConstructsTransaction(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *protocolspb.ConstructTransactionRequest,
		) (*typespb.Transaction, error) {
			return s.protocolsProxy.ConstructTransaction(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
