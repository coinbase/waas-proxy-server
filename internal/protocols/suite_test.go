package protocols_test

import (
	"context"
	"testing"

	"github.com/coinbase/waas-proxy-server/internal/protocols"
	"github.com/coinbase/waas-proxy-server/internal/testutils"
	"github.com/coinbase/waas-client-library-go/clients/v1/mocks"
	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
	typespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/types/v1"
	cryptotypespb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/crypto/types/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type ts struct {
	testutils.TestSuite

	protocolsClient *mocks.ProtocolServiceClient
	protocolsProxy  protocols.ProtocolsProxy
}

func (s *ts) SetupTest() {
	var err error

	s.protocolsClient = mocks.NewProtocolServiceClient(s.T())

	s.protocolsProxy, err = protocols.NewProtocolsProxy(context.Background(), zap.NewNop(), s.protocolsClient)
	if err != nil {
		s.FailNow("failed to initialize protocols proxy", err)
	}
}

func TestProtocolsProxy(t *testing.T) {
	suite.Run(t, new(ts))
}

// ConstructsTransaction implements the mocks/assertions for constructing a transaction.
func (s *ts) ConstructsTransaction(
	req *protocolspb.ConstructTransactionRequest,
	resp *typespb.Transaction,
	err error) *mock.Call {
	call := s.protocolsClient.On(
		"ConstructTransaction",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// ConstructsTransferTransaction implements the mocks/assertions for constructing a transfer transaction.
func (s *ts) ConstructsTransferTransaction(
	req *protocolspb.ConstructTransferTransactionRequest,
	resp *typespb.Transaction,
	err error) *mock.Call {
	call := s.protocolsClient.On(
		"ConstructTransferTransaction",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// BroadcastsTransaction implements the mocks/assertions for broadcasting a transaction.
func (s *ts) BroadcastsTransaction(
	req *protocolspb.BroadcastTransactionRequest,
	resp *typespb.Transaction,
	err error) *mock.Call {
	call := s.protocolsClient.On(
		"BroadcastTransaction",
		mock.MatchedBy(s.MatchAnyContext),
		req,
	)

	if err != nil {
		return call.Return(nil, err)
	}

	return call.Return(resp, nil)
}

// EstimateFee implements the mocks/assertions for estimating fee.
func (s *ts) EstimatesFee(
	req *protocolspb.EstimateFeeRequest,
	resp *protocolspb.EstimateFeeResponse,
	err error) *mock.Call {
	call := s.protocolsClient.On(
		"EstimateFee",
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

// newTxForConstruct returns a transaction to use for testing ConstructTransaction.
func (s *ts) newTxForConstruct() *typespb.Transaction {
	return &typespb.Transaction{
		Input: &typespb.TransactionInput{
			Input: &typespb.TransactionInput_EthereumRlpInput{
				EthereumRlpInput: &ethereumpb.RLPTransaction{
					TransactionRlp: []byte("tx"),
				},
			},
		},
	}
}

// newTxForBroadcast returns a transaction to use for testing BroadcastTransaction.
func (s *ts) newTxForBroadcast() *typespb.Transaction {
	// ECDSA Signature R and S values must have 32 bytes exactly.
	bytes := make([]byte, 32)

	for i := 0; i < 32; i++ {
		bytes[i] = byte('a')
	}

	return &typespb.Transaction{
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
	}
}
