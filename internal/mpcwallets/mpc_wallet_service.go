package mpcwallets

import (
	"context"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/coinbase/waas-proxy-server/internal/operations"
	mpcwalletspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_wallets/v1"
	"go.uber.org/zap"

	clientsv1 "github.com/coinbase/waas-client-library-go/clients/v1"
)

// MPCWalletsProxy is an interface with methods to proxy requests to the mpc wallet service.
type MPCWalletsProxy interface {
	CreateMPCWallet(context.Context, *mpcwalletspb.CreateMPCWalletRequest) (*longrunningpb.Operation, error)
	GetMPCWallet(context.Context, *mpcwalletspb.GetMPCWalletRequest) (*mpcwalletspb.MPCWallet, error)
	ListMPCWallets(context.Context, *mpcwalletspb.ListMPCWalletsRequest) (*mpcwalletspb.ListMPCWalletsResponse, error)
	GenerateAddress(context.Context, *mpcwalletspb.GenerateAddressRequest) (*mpcwalletspb.Address, error)
	GetAddress(context.Context, *mpcwalletspb.GetAddressRequest) (*mpcwalletspb.Address, error)
	ListAddresses(context.Context, *mpcwalletspb.ListAddressesRequest) (*mpcwalletspb.ListAddressesResponse, error)
	ListBalances(context.Context, *mpcwalletspb.ListBalancesRequest) (*mpcwalletspb.ListBalancesResponse, error)
	ListBalanceDetails(context.Context, *mpcwalletspb.ListBalanceDetailsRequest) (*mpcwalletspb.ListBalanceDetailsResponse, error)
}

// mpcWalletsProxy implements the MPCWalletsProxy interface.
type mpcWalletsProxy struct {
	log *zap.Logger

	mpcWalletsClient clientsv1.MPCWalletServiceClient

	operationMap map[string]operations.OperationType
}

// NewMPCWalletsProxy instantiates a new mpcWalletsProxy.
func NewMPCWalletsProxy(
	ctx context.Context,
	log *zap.Logger,
	mpcWalletsClient clientsv1.MPCWalletServiceClient,
	operationMap map[string]operations.OperationType,
) (*mpcWalletsProxy, error) {
	return &mpcWalletsProxy{
		log:              log,
		mpcWalletsClient: mpcWalletsClient,
		operationMap:     operationMap,
	}, nil
}
