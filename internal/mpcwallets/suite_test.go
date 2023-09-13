package mpcwallets_test

import (
	"context"
	"testing"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/mpcwallets"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
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

	mpcWalletsClient                *mocks.MPCWalletServiceClient
	mpcWalletIterator               *mocks.MPCWalletIterator
	addressIterator                 *mocks.AddressIterator
	balanceDetailIterator           *mocks.BalanceDetailIterator
	balanceIterator                 *mocks.BalanceIterator
	createMPCWalletWrappedOperation *mocks.ClientWrappedCreateMPCWalletOperation
	mpcWalletsProxy                 mpcwallets.MPCWalletsProxy
	operationsMap                   map[string]operations.OperationType
}

func (s *ts) SetupTest() {
	var err error

	s.mpcWalletsClient = mocks.NewMPCWalletServiceClient(s.T())

	s.mpcWalletIterator = mocks.NewMPCWalletIterator(s.T())

	s.addressIterator = mocks.NewAddressIterator(s.T())

	s.balanceDetailIterator = mocks.NewBalanceDetailIterator(s.T())

	s.balanceIterator = mocks.NewBalanceIterator(s.T())

	s.createMPCWalletWrappedOperation = mocks.NewClientWrappedCreateMPCWalletOperation(s.T())

	s.operationsMap = make(map[string]operations.OperationType)

	s.mpcWalletsProxy, err = mpcwallets.NewMPCWalletsProxy(context.Background(), zap.NewNop(), s.mpcWalletsClient, s.operationsMap)
	if err != nil {
		s.FailNow("failed to initialize mpc wallets proxy", err)
	}
}

func TestMPCWalletsProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// CreatesMPCWallet implements the mocks/assertions for creating an mpc wallet.
func (s *ts) CreatesMPCWallet(
	req *mpcwalletspb.CreateMPCWalletRequest,
	metadata *mpcwalletspb.CreateMPCWalletMetadata,
	metadataAny *anypb.Any,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"CreateMPCWallet",
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

	s.createMPCWalletWrappedOperation.On(
		"Metadata",
	).Return(metadata, nil)

	s.createMPCWalletWrappedOperation.On(
		"Name",
	).Return(operation.Name)

	s.createMPCWalletWrappedOperation.On(
		"Poll",
		mock.MatchedBy(s.MatchAnyContext),
	).Return(nil, nil)

	s.createMPCWalletWrappedOperation.On(
		"Done",
	).Return(false)

	return call.Return(s.createMPCWalletWrappedOperation, nil)
}

// ListsAddresses implements the mocks/assertions for listing mpc addresses.
func (s *ts) ListsAddresses(
	req *mpcwalletspb.ListAddressesRequest,
	resp *mpcwalletspb.ListAddressesResponse,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"ListAddresses",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.addressIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.addressIterator)
	}

	s.addressIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.addressIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.addressIterator)
}

// ListsBalanceDetails implements the mocks/assertions for listing mpc balance details.
func (s *ts) ListsBalanceDetails(
	req *mpcwalletspb.ListBalanceDetailsRequest,
	resp *mpcwalletspb.ListBalanceDetailsResponse,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"ListBalanceDetails",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.balanceDetailIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.balanceDetailIterator)
	}

	s.balanceDetailIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.balanceDetailIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.balanceDetailIterator)
}

// ListsBalances implements the mocks/assertions for listing mpc balances.
func (s *ts) ListsBalances(
	req *mpcwalletspb.ListBalancesRequest,
	resp *mpcwalletspb.ListBalancesResponse,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"ListBalances",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.balanceIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.balanceIterator)
	}

	s.balanceIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.balanceIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.balanceIterator)
}

// ListsMPCWallets implements the mocks/assertions for listing mpc wallets.
func (s *ts) ListsMPCWallets(
	req *mpcwalletspb.ListMPCWalletsRequest,
	resp *mpcwalletspb.ListMPCWalletsResponse,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"ListMPCWallets",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.mpcWalletIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.mpcWalletIterator)
	}

	s.mpcWalletIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.mpcWalletIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.mpcWalletIterator)
}

// GetsMPCWallet implements the mocks/assertions for getting an MPC wallet.
func (s *ts) GetsMPCWallet(
	req *mpcwalletspb.GetMPCWalletRequest,
	resp *mpcwalletspb.MPCWallet,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"GetMPCWallet",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GeneratesAddress implements the mocks/assertions for generating an address.
func (s *ts) GeneratesAddress(
	req *mpcwalletspb.GenerateAddressRequest,
	resp *mpcwalletspb.Address,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"GenerateAddress",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// GetsAddress implements the mocks/assertions for getting an address.
func (s *ts) GetsAddress(
	req *mpcwalletspb.GetAddressRequest,
	resp *mpcwalletspb.Address,
	err error,
) *mock.Call {
	call := s.mpcWalletsClient.On(
		"GetAddress",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
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
