package operations_test

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ts) Test_GetOperation() {
	var (
		ctx = context.Background()
		req = &longrunningpb.GetOperationRequest{}
	)

	scenarios := map[string]struct {
		mutate             func(*longrunningpb.GetOperationRequest) *longrunningpb.Operation
		expectedStatusCode codes.Code
	}{
		"success mpc keys operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpckeys-operation"

				resp := &longrunningpb.Operation{
					Name: req.GetName(),
					Done: false,
				}

				s.GetsOperationMPCKeysClient(req, resp, nil)

				return resp
			},
			expectedStatusCode: codes.OK,
		},
		"success mpc wallets operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpcwallets-operation"

				resp := &longrunningpb.Operation{
					Name: req.GetName(),
					Done: false,
				}

				s.GetsOperationMPCWalletsClient(req, resp, nil)

				return resp
			},
			expectedStatusCode: codes.OK,
		},
		"success mpc transactions operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpctransactions-operation"

				resp := &longrunningpb.Operation{
					Name: req.GetName(),
					Done: false,
				}

				s.GetsOperationMPCTransactionsClient(req, resp, nil)

				return resp
			},
			expectedStatusCode: codes.OK,
		},
		"client error mpc keys operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpckeys-operation"

				s.GetsOperationMPCKeysClient(req, nil, status.Errorf(codes.Internal, "boom"))

				return nil
			},
			expectedStatusCode: codes.Internal,
		},
		"client error mpc wallets operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpcwallets-operation"

				resp := &longrunningpb.Operation{
					Name: req.GetName(),
					Done: false,
				}

				s.GetsOperationMPCWalletsClient(req, nil, status.Errorf(codes.Internal, "boom"))

				return resp
			},
			expectedStatusCode: codes.Internal,
		},
		"client error mpc transactions operation": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/test-mpctransactions-operation"

				resp := &longrunningpb.Operation{
					Name: req.GetName(),
					Done: false,
				}

				s.GetsOperationMPCTransactionsClient(req, nil, status.Errorf(codes.Internal, "boom"))

				return resp
			},
			expectedStatusCode: codes.Internal,
		},
		"operation not found": {
			mutate: func(req *longrunningpb.GetOperationRequest) *longrunningpb.Operation {
				req.Name = "operations/invalid-operation"

				return nil
			},
			expectedStatusCode: codes.NotFound,
		},
	}

	for name, scenario := range scenarios {
		scenario := scenario

		s.Run(name, func() {
			s.SetupTest()

			expectedResponse := scenario.mutate(req)

			resp, err := s.operationsProxy.GetOperation(ctx, req)

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
