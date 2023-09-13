package pools_test

import (
	"context"
	"testing"

	"github.com/coinbase/waas-proxy-server/internal/pools"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	poolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/pools/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ts struct {
	testutils.TestSuite

	poolsClient  *mocks.PoolServiceClient
	poolIterator *mocks.PoolIterator
	poolsProxy   pools.PoolsProxy
}

func (s *ts) SetupTest() {
	var err error

	s.poolsClient = mocks.NewPoolServiceClient(s.T())

	s.poolIterator = mocks.NewPoolIterator(s.T())

	s.poolsProxy, err = pools.NewPoolsProxy(context.Background(), zap.NewNop(), s.poolsClient)
	if err != nil {
		s.FailNow("failed to initialize pools proxy", err)
	}
}

func TestPoolsProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// CreatesPool implements the mocks/assertions for creating a pool.
func (s *ts) CreatesPool(
	req *poolspb.CreatePoolRequest,
	resp *poolspb.Pool,
	err error,
) *mock.Call {
	call := s.poolsClient.On(
		"CreatePool",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsPool implements the mocks/assertions for getting a pool.
func (s *ts) GetsPool(
	req *poolspb.GetPoolRequest,
	resp *poolspb.Pool,
	err error,
) *mock.Call {
	call := s.poolsClient.On(
		"GetPool",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ListsPools implements the mocks/assertions for listing pools.
func (s *ts) ListsPools(
	req *poolspb.ListPoolsRequest,
	resp *poolspb.ListPoolsResponse,
	err error,
) *mock.Call {
	call := s.poolsClient.On(
		"ListPools",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.poolIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.poolIterator)
	}

	s.poolIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.poolIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.poolIterator)
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
