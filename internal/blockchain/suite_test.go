package blockchain_test

import (
	"context"
	"testing"

	"github.com/coinbase/waas-proxy-server/internal/blockchain"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	blockchainpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/blockchain/v1"
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

	blockchainClient *mocks.BlockchainServiceClient
	networkIterator  *mocks.NetworkIterator
	assetIterator    *mocks.AssetIterator
	blockchainProxy  blockchain.BlockchainProxy
}

func (s *ts) SetupTest() {
	var err error

	s.blockchainClient = mocks.NewBlockchainServiceClient(s.T())

	s.networkIterator = mocks.NewNetworkIterator(s.T())

	s.assetIterator = mocks.NewAssetIterator(s.T())

	s.blockchainProxy, err = blockchain.NewBlockchainProxy(context.Background(), zap.NewNop(), s.blockchainClient)
	if err != nil {
		s.FailNow("failed to initialize pools proxy", err)
	}
}

func TestBlockchainProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// GetsNetwork implements the mocks/assertions for getting a network.
func (s *ts) GetsNetwork(
	req *blockchainpb.GetNetworkRequest,
	resp *blockchainpb.Network,
	err error,
) *mock.Call {
	call := s.blockchainClient.On(
		"GetNetwork",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ListsNetworks implements the mocks/assertions for listing networks.
func (s *ts) ListsNetworks(
	req *blockchainpb.ListNetworksRequest,
	resp *blockchainpb.ListNetworksResponse,
	err error,
) *mock.Call {
	call := s.blockchainClient.On(
		"ListNetworks",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.networkIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.networkIterator)
	}

	s.networkIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.networkIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.networkIterator)
}

// GetsAsset implements the mocks/assertions for getting an asset.
func (s *ts) GetsAsset(
	req *blockchainpb.GetAssetRequest,
	resp *blockchainpb.Asset,
	err error,
) *mock.Call {
	call := s.blockchainClient.On(
		"GetAsset",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ListsAssets implements the mocks/assertions for listing assets.
func (s *ts) ListsAssets(
	req *blockchainpb.ListAssetsRequest,
	resp *blockchainpb.ListAssetsResponse,
	err error,
) *mock.Call {
	call := s.blockchainClient.On(
		"ListAssets",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		s.assetIterator.On(
			"Next",
		).Return(nil, err)

		return call.Return(s.assetIterator)
	}

	s.assetIterator.On(
		"Next",
	).Return(nil, iterator.Done)

	s.assetIterator.On(
		"Response",
	).Return(resp)

	return call.Return(s.assetIterator)
}

// BatchesGetAssets implements the mocks/assertions for batch getting an asset.
func (s *ts) BatchesGetAssets(
	req *blockchainpb.BatchGetAssetsRequest,
	resp *blockchainpb.BatchGetAssetsResponse,
	err error,
) *mock.Call {
	call := s.blockchainClient.On(
		"BatchGetAssets",
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
