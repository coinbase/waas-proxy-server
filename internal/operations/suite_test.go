package operations_test

import (
	"context"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

type ts struct {
	testutils.TestSuite

	mpcKeysClient         *mocks.MPCKeyServiceClient
	mpcWalletsClient      *mocks.MPCWalletServiceClient
	mpcTransactionsClient *mocks.MPCTransactionServiceClient
	operationsProxy       operations.OperationsProxy
	operationMap          map[string]operations.OperationType
}

func (s *ts) SetupTest() {
	var err error

	s.mpcKeysClient = mocks.NewMPCKeyServiceClient(s.T())

	s.mpcWalletsClient = mocks.NewMPCWalletServiceClient(s.T())

	s.mpcTransactionsClient = mocks.NewMPCTransactionServiceClient(s.T())

	s.operationMap = map[string]operations.OperationType{
		"operations/test-mpckeys-operation":         operations.MPC_KEY_OPERATION,
		"operations/test-mpcwallets-operation":      operations.MPC_WALLET_OPERATION,
		"operations/test-mpctransactions-operation": operations.MPC_TRANSACTION_OPERATION,
	}

	s.operationsProxy, err = operations.NewOperationsProxy(context.Background(), zap.NewNop(), s.mpcKeysClient, s.mpcWalletsClient, s.mpcTransactionsClient, s.operationMap)
	if err != nil {
		s.FailNow("failed to initialize operations proxy", err)
	}
}

func TestOperationsProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// GetsOperationMPCKeysClient implements the mocks/assertions for getting an operation
// using mpc keys client.
func (s *ts) GetsOperationMPCKeysClient(
	req *longrunningpb.GetOperationRequest,
	resp *longrunningpb.Operation,
	err error,
) *mock.Call {
	call := s.mpcKeysClient.On(
		"GetOperation",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsOperationMPCWalletsClient implements the mocks/assertions for getting an operation
// using mpc wallets client.
func (s *ts) GetsOperationMPCWalletsClient(
	req *longrunningpb.GetOperationRequest,
	resp *longrunningpb.Operation,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"GetOperation",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsOperationMPCTransactionsClient implements the mocks/assertions for getting an operation
// using mpc transactions client.
func (s *ts) GetsOperationMPCTransactionsClient(
	req *longrunningpb.GetOperationRequest,
	resp *longrunningpb.Operation,
	err error,
) *mock.Call {
	call := s.mpcTransactionsClient.On(
		"GetOperation",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// Test_ListOperations tests that ListOperations is unimplemented.
func (s *ts) Test_ListOperations() {
	var (
		ctx                = context.Background()
		req                = &longrunningpb.ListOperationsRequest{}
		expectedStatusCode = codes.Unimplemented
	)

	resp, err := s.operationsProxy.ListOperations(ctx, req)
	s.StatusCodeEqual(expectedStatusCode, err)
	s.Nil(resp)
}

// Test_DeleteOperation tests that DeleteOperation is unimplemented.
func (s *ts) Test_DeleteOperation() {
	var (
		ctx                = context.Background()
		req                = &longrunningpb.DeleteOperationRequest{}
		expectedStatusCode = codes.Unimplemented
	)

	resp, err := s.operationsProxy.DeleteOperation(ctx, req)
	s.StatusCodeEqual(expectedStatusCode, err)
	s.Nil(resp)
}

// Test_CancelOperation tests that CancelOperation is unimplemented.
func (s *ts) Test_CancelOperation() {
	var (
		ctx                = context.Background()
		req                = &longrunningpb.CancelOperationRequest{}
		expectedStatusCode = codes.Unimplemented
	)

	resp, err := s.operationsProxy.CancelOperation(ctx, req)
	s.StatusCodeEqual(expectedStatusCode, err)
	s.Nil(resp)
}

// Test_WaitOperation tests that WaitOperation is unimplemented.
func (s *ts) Test_WaitOperation() {
	var (
		ctx                = context.Background()
		req                = &longrunningpb.WaitOperationRequest{}
		expectedStatusCode = codes.Unimplemented
	)

	resp, err := s.operationsProxy.WaitOperation(ctx, req)
	s.StatusCodeEqual(expectedStatusCode, err)
	s.Nil(resp)
}
