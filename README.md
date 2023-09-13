# WaaS Proxy Server

**waas-proxy-server** is a [proxy server](https://en.wikipedia.org/wiki/Proxy_server) for the WaaS API. It is designed to be easy to use and act as an intermediary between a client application and the WaaS API, while providing security for a developer’s API keys.

## Prerequisites

- [Golang 1.21+](https://go.dev/learn/)
- [make](https://www.gnu.org/software/make/)
- [protoc](https://grpc.io/docs/protoc-installation/)

## Setup

Open the `.env` file.
Replace the `COINBASE_CLOUD_API_KEY_NAME` and `COINBASE_CLOUD_API_KEY_PRIVATE_KEY` variables with your own API credentials provided by Coinbase.

## Running the Proxy Server

You can build and run the proxy server using [make](https://www.gnu.org/software/make/), including targets for [Docker](https://www.docker.com/).

### Using `make`

Build the proxy server:

```bash
make waas-proxy-server
```

Run the proxy server:

```bash
make run
```

#### Docker

Build the proxy server:

```bash
make docker/build
```

Run the proxy server:

```bash
make docker/run
```

## Calling the Proxy Server

Once the proxy server is running, you can make requests to it using your desired endpoint.

***Note:** The proxy server is set to listen on `localhost:8091` by default. You can specify your desired host and port by modifying `HTTP_ADDRESS` in the `.env` file.*

### Using `curl`

```bash
# Calls ListNetworks
curl -X GET -d '{}' localhost:8091/v1/networks
```

### Using JavaScript

```bash
/* Calls ListNetworks */
fetch('http://localhost:8091/v1/networks', {
    method: 'GET',
    headers: { 'Content-Type': 'application/json' }
})
.then(response => {
    if (!response.ok) {
        throw new Error()
    }
    return response.json()
})
```

## Supported Methods

Below is a list of WaaS API methods supported by the proxy server and their corresponding endpoints.

### Blockchain Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [GetNetwork](https://docs.cloud.coinbase.com/waas/reference/blockchainservice_getnetwork) | GET | `/v1/{networkName}` |
| [ListNetworks](https://docs.cloud.coinbase.com/waas/reference/blockchainservice_listnetworks) | GET | `/v1/networks`  |
| [GetAsset](https://docs.cloud.coinbase.com/waas/reference/blockchainservice_getasset) | GET | `/v1/{assetName}` |
| [ListAssets](https://docs.cloud.coinbase.com/waas/reference/blockchainservice_listassets) | GET | `/v1/{networkName}/assets` |
| [BatchGetAssets](https://docs.cloud.coinbase.com/waas/reference/blockchainservice_batchgetassets) | GET | `/v1/{networkName}/assets:batchGet` |

### MPC Key Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [RegisterDevice](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_registerdevice) | POST | `/v1/device:register` |
| [GetDevice](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_getdevice) | GET | `/v1/{deviceName}` |
| [CreateDeviceGroup](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_createdevicegroup) | POST | `/v1/{poolName}/deviceGroups` |
| [GetDeviceGroup](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_getdevicegroup) | GET | `/v1/{deviceGroupName}` |
| [ListMPCOperations](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_listmpcoperations) | GET | `/v1/{deviceGroupName}/mpcOperations` |
| [CreateMPCKey](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_creatempckey) | POST | `/v1/{deviceGroupName}/mpcKeys` |
| [GetMPCKey](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_getmpckey) | GET | `/v1/{mpcKeyName}` |
| [CreateSignature](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_createsignature) | POST | `/v1/{mpcKeyName}/signatures` |
| [PrepareDeviceArchive](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_preparedevicearchive) | POST | `/v1/{deviceGroupName}:prepareDeviceArchive` |
| [PrepareDeviceBackup](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_preparedevicebackup) | POST | `/v1/{deviceGroupName}:prepareDeviceBackup` |
| [AddDevice](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_adddevice) | POST | `/v1/{deviceGroupName}:addDevice` |
| [RevokeDevice](https://docs.cloud.coinbase.com/waas/reference/mpckeyservice_revokedevice) | POST | `/v1/device:revoke` |

### MPC Transaction Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [CreateMPCTransaction](https://docs.cloud.coinbase.com/waas/reference/mpctransactionservice_creatempctransaction) | POST | `/v1/{mpcWalletName}/mpcTransactions` |
| [GetMPCTransaction](https://docs.cloud.coinbase.com/waas/reference/mpctransactionservice_getmpctransaction) | GET | `/v1/{mpcTransactionName}` |
| [ListMPCTransactions](https://docs.cloud.coinbase.com/waas/reference/mpctransactionservice_listmpctransactions) | GET | `/v1/{mpcWalletName}/mpcTransactions` |

### MPC Wallet Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [CreateMPCWallet](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_creatempcwallet) | POST | `/v1/{poolName}/mpcWallets` |
| [GetMPCWallet](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_getmpcwallet) | GET | `/v1/{mpcWalletName}` |
| [ListMPCWallets](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_listmpcwallets) | GET | `/v1/{poolName}/mpcWallets` |
| [GenerateAddress](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_generateaddress) | POST | `/v1/{mpcWalletName}:generateAddress` |
| [GetAddress](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_getaddress) | GET | `/v1/{networkName}` |
| [ListAddresses](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_listaddresses) | GET | `/v1/{networkName}/addresses` |
| [ListBalances](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_listbalances) | GET | `/v1/{addressName}/balances` |
| [ListBalanceDetails](https://docs.cloud.coinbase.com/waas/reference/mpcwalletservice_listbalancedetails) | GET | `/v1/{balanceName}/balanceDetails` |

### Operation Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| GetOperation | GET | `/v1/{operationName}` |

### Pool Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [CreatePool](https://docs.cloud.coinbase.com/waas/reference/poolservice_createpool) | POST | `/v1/pools` |
| [GetPool](https://docs.cloud.coinbase.com/waas/reference/poolservice_getpool) | GET | `/v1/{poolName}` |
| [ListPools](https://docs.cloud.coinbase.com/waas/reference/poolservice_listpools) | GET | `/v1/pools` |

### Protocol Service

| Name | Method | Endpoint |
| :--- | :--- | :--- |
| [ConstructTransaction](https://docs.cloud.coinbase.com/waas/reference/protocolservice_constructtransaction) | POST | `/v1/{networkName}:constructTransaction` |
| [ConstructTransferTransaction](https://docs.cloud.coinbase.com/waas/reference/protocolservice_constructtransfertransaction) | POST | `/v1/{networkName}:constructTransferTransaction` |
| [BroadcastTransaction](https://docs.cloud.coinbase.com/waas/reference/protocolservice_broadcasttransaction) | POST | `/v1/{networkName}:broadcastTransaction` |
| [EstimateFee](https://docs.cloud.coinbase.com/waas/reference/protocolservice_estimatefee) | GET | `/v1/{networkName}:estimateFee` |

## Documentation

For full documentation, refer to [docs.cloud.coinbase.com/waas](http://docs.cloud.coinbase.com/waas).

## License

```
Copyright © 2023 Coinbase, Inc. <https://www.coinbase.com/>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
