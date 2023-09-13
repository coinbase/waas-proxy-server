package protocols_test

import (
	"context"
	"math/big"

	ethereumpb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/ethereum/v1"
	protocolspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/protocols/v1"
)

// Test_EstimateFee tests EstimateFee with a range of scenarios.
func (s *ts) Test_EstimateFee() {
	var (
		networkName = "networks/test-network"

		estimateFeeReq = &protocolspb.EstimateFeeRequest{
			Network: networkName,
		}

		estimateFeeResp = &protocolspb.EstimateFeeResponse{
			NetworkFeeEstimate: &protocolspb.EstimateFeeResponse_EthereumFeeEstimate{
				EthereumFeeEstimate: &ethereumpb.FeeEstimate{
					GasPrice:             big.NewInt(100).String(),
					MaxFeePerGas:         big.NewInt(150).String(),
					MaxPriorityFeePerGas: big.NewInt(200).String(),
				},
			},
		}

		newRequestFn = func() *protocolspb.EstimateFeeRequest {
			return estimateFeeReq
		}

		validMutation = func(req *protocolspb.EstimateFeeRequest) *protocolspb.EstimateFeeResponse {
			s.EstimatesFee(req, estimateFeeResp, nil)

			return estimateFeeResp
		}

		errorMutation = func(
			req *protocolspb.EstimateFeeRequest,
			err error,
		) *protocolspb.EstimateFeeResponse {
			s.EstimatesFee(req, nil, err)

			return nil
		}

		testFn = func(
			ctx context.Context,
			req *protocolspb.EstimateFeeRequest,
		) (*protocolspb.EstimateFeeResponse, error) {
			return s.protocolsProxy.EstimateFee(ctx, req)
		}
	)

	RunScenarios(s, newRequestFn, validMutation, errorMutation, testFn)
}
