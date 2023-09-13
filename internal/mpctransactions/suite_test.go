package mpctransactions_test

import (
	"context"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/mpctransactions"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_transactions/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type ts struct {
	testutils.TestSuite

	mpcTransactionsClient                *mocks.MPCTransactionServiceClient
	mpcTransactionIterator               *mocks.MPCTransactionIterator
	createMPCTransactionWrappedOperation *mocks.ClientWrappedCreateMPCTransactionOperation
	mpcTransactionProxy                  mpctransactions.MPCTransactionsProxy
	operationMap                         map[string]operations.OperationType
}

func (s *ts) SetupTest() {
	var err error

	s.mpcTransactionsClient = mocks.NewMPCTransactionServiceClient(s.T())

	s.mpcTransactionIterator = mocks.NewMPCTransactionIterator(s.T())

	s.createMPCTransactionWrappedOperation = mocks.NewClientWrappedCreateMPCTransactionOperation(s.T())

	s.operationMap = make(map[string]operations.OperationType)

	s.mpcTransactionProxy, err = mpctransactions.NewMPCTransactionsProxy(context.Background(), zap.NewNop(), s.mpcTransactionsClient, s.operationMap)
	if err != nil {
		s.FailNow("failed to initialize mpc transactions proxy", err)
	}
}

func TestMPCTransactionsProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// CreatesMPCTransaction implements the mocks/assertions for creating an mpc transaction.
func (s *ts) CreatesMPCTransaction(
	req *v1.CreateMPCTransactionRequest,
	metadata *v1.CreateMPCTransactionMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcTransactionsClient.On(
		"CreateMPCTransaction",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	operation := &longrunningpb.Operation{
		Name:     "operations/test-operation",
		Metadata: metadataAny,
		Done:     false,
	}

	s.createMPCTransactionWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.createMPCTransactionWrappedOperation.On(
		"Name",
	).Return(operation.Name)

	s.createMPCTransactionWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.createMPCTransactionWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.createMPCTransactionWrappedOperation, nil)
}

// GetsMPCTransaction implements the mocks/assertions for getting an mpc transaction.
func (s *ts) GetsMPCTransaction(
	req *v1.GetMPCTransactionRequest,
	resp *v1.MPCTransaction,
	err error,
) *mock.Call {
	call := s.mpcTransactionsClient.On(
		"GetMPCTransaction",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ListsMPCTransactions implements the mocks/assertions for listing mpc transactions.
func (s *ts) ListsMPCTransactions(
	req *v1.ListMPCTransactionsRequest,
	resp *v1.ListMPCTransactionsResponse,
	err error,
) *mock.Call {
	call := s.mpcTransactionsClient.On(
		"ListMPCTransactions",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.mpcTransactionIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.mpcTransactionIterator)
	}

	s.mpcTransactionIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.mpcTransactionIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.mpcTransactionIterator)
}

func RunScenarios[
	RequestType protoreflect.ProtoMessage,
	ResponseType protoreflect.ProtoMessage,
	NewRequestFn func() RequestType,
	MutationFn func(RequestType) ResponseType,
	ErrorMutationFn func(RequestType, error) ResponseType,
	TestFn func(context.Context, RequestType) (ResponseType, error),
](
	s *ts,
	newRequest NewRequestFn,
	validMutation MutationFn,
	errorMutation ErrorMutationFn,
	testFunction TestFn,
) {
	scenarios := map[string]struct {
		mutate             MutationFn
		expectedStatusCode codes.Code
	}{
		"success": {
			mutate:             validMutation,
			expectedStatusCode: codes.OK,
		},
		"client error": {
			mutate: func(req RequestType) ResponseType {
				return errorMutation(req, status.Errorf(codes.Internal, "boom"))
			},
			expectedStatusCode: codes.Internal,
		},
	}

	for name, scenario := range scenarios {
		scenario := scenario

		s.Run(name, func() {
			s.SetupTest()

			req := newRequest()

			var expectedResponse ResponseType

			if scenario.mutate != nil {
				expectedResponse = scenario.mutate(req)
			}

			resp, err := testFunction(context.Background(), req)

			if scenario.expectedStatusCode == codes.OK {
				s.NoError(err)
				s.ProtoEqual(expectedResponse, resp)
			} else {
				s.StatusCodeEqual(scenario.expectedStatusCode, err)
				s.Nil(resp)
			}
		})
	}
}
