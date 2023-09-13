package protocols_test

import (
	"context"

	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
)

// Test_BroadcastTransaction tests BroadcastTransaction with a range of scenarios.
func (s *ts) Test_BroadcastTransaction() {
	var (
		v1Tx        = s.newTxForBroadcast()
		networkName = "networks/test-network"

		broadcastReq = &protocolspb.BroadcastTransactionRequest{
			Network:     networkName,
			Transaction: v1Tx,
		}

		newRequestFn = func() *protocolspb.BroadcastTransactionRequest {
			return broadcastReq
		}

		validMutation = func(req *protocolspb.BroadcastTransactionRequest) *typespb.Transaction {
			s.BroadcastsTransaction(req, v1Tx, nil)

			return v1Tx
		}

		errorMutation = func(
			req *protocolspb.BroadcastTransactionRequest,
			err error,
		) *typespb.Transaction {
			s.BroadcastsTransaction(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *protocolspb.BroadcastTransactionRequest,
		) (*typespb.Transaction, error) {
			return s.protocolsProxy.BroadcastTransaction(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
