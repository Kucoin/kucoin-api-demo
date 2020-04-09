# Test Tool for KuCoin API

## Choose environment

| Environment  | BaseUri                                                      |
| ------------ | ------------------------------------------------------------ |
| *Production* | `https://api.kucoin.com(DEFAULT)` `https://openapi-v2.kucoin.com` `https://api.kcs.top` |
| *Sandbox*    | `https://openapi-sandbox.kucoin.com`                         |

## Example

* First,you should wirte your api information into config.json.

  ```json
  {
    "ApiBaseURI":"",
    "ApiKey": "",
    "ApiSecret": "",
    "ApiPassphrase": "",
    "ApiSkipVerifyTls": false
  }
  ```

* If not, you can run this script and follow the guide,then input your api information, such as:

  ```shell
  # please input api base URI,such as:https://api.kucoin.com
  % https://api.kucoin.com
  # please input your api key...
  % input your api key
  # please input your api secret...
  % input your api secret
  # please input your api passphrase...
  % input your api passphrase
  ```

*  input the api name and parameter information to choose api which you want to test.

  ```shell
  # please input the api name which you want to test
  % Accounts
  # please input the api parameters,use ',' to spilt parameters.
  # if it is no parameter api,just enter to skip...
  % BTC,main
  ```

* Then,it will show the request and response body to you.

## Api map list

<details>
<summary>Account</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| CreateAccount | POST /api/v1/accounts | type, currency | main,BTC | https://docs.kucoin.com/#create-an-account |
| Accounts | GET /api/v1/accounts | currency,type | BTC,main | https://docs.kucoin.com/#list-accounts |
| Account | GET /api/v1/accounts/{accountId}         | accountId | {account ID} | https://docs.kucoin.com/#get-an-account |
| SubAccountUsers | GET /api/v1/sub/user |  |  | https://docs.kucoin.com/#get-user-info-of-all-sub-accounts |
| SubAccounts | GET /api/v1/sub-accounts                 |  |  | https://docs.kucoin.com/#get-the-aggregated-balance-of-all-sub-accounts |
| SubAccount | GET /api/v1/sub-accounts/{subUserId} | subUserId |{sub-account ID} | https://docs.kucoin.com/#get-account-balance-of-a-sub-account |
| AccountLedgers | GET /api/v1/accounts/{accountId}/ledgers | accountId |{account ID} | https://docs.kucoin.com/#get-account-ledgers |
| AccountHolds | GET /api/v1/accounts/{accountId}/holds   | accountId | {account ID} | https://docs.kucoin.com/#get-holds |
| InnerTransferV2 | POST /api/v2/accounts/inner-transfer | Currency,from,to,amount | KCS,main,trade,2 | https://docs.kucoin.com/#inner-transfer |

</details>

<details>
<summary>Deposit</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| CreateDepositAddress | POST /api/v1/deposit-addresses | currency | BTC | https://docs.kucoin.com/#create-deposit-address |
| DepositAddresses | GET /api/v1/deposit-addresses | currency | BTC | https://docs.kucoin.com/#get-deposit-address |
| V1Deposits | GET /api/v1/hist-deposits | currency | KCS | https://docs.kucoin.com/#get-v1-historical-deposits-list |
| Deposits | GET /api/v1/deposits | currency | KCS | https://docs.kucoin.com/#get-deposit-list |

</details>

<details>
<summary>Fill</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| Fills | GET /api/v1/fills | tradeType | TRADE |https://docs.kucoin.com/#list-fills|
| RecentFills | GET /api/v1/limit/fills | |  |https://docs.kucoin.com/#recent-fills|

</details>

<details>
<summary>Order</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| CreateOrder | POST /api/v1/orders | side,symbol,price,size | sell,BTC-USDT,7330,0.1 |https://docs.kucoin.com/#place-a-new-order|
| CreateMultiOrder | POST /api/v1/orders/multi | symbol,side,price,size | BTC-USDT,sell,7350,0.01 |https://docs.kucoin.com/#place-bulk-orders|
| CancelOrder | DELETE /api/v1/orders/{orderId} | orderId | {order ID} |https://docs.kucoin.com/#cancel-an-order|
| CancelOrders | DELETE /api/v1/orders | symbol | BTC-USDT |https://docs.kucoin.com/#cancel-all-orders|
| Orders | GET /api/v1/orders | symbol | BTC-USDT,sell |https://docs.kucoin.com/#list-orders|
| Order | GET /api/v1/orders/{order-id} | orderId | {order ID} |https://docs.kucoin.com/#get-an-order|
| RecentOrders | GET /api/v1/limit/orders | |  |https://docs.kucoin.com/#recent-orders|

</details>

<details>
<summary>WebSocket Feed</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| WebSocketPublicToken | POST /api/v1/bullet-public | | | https://docs.kucoin.com/#apply-connect-token |
| WebSocketPrivateToken | POST /api/v1/bullet-private | | | https://docs.kucoin.com/#apply-connect-token |
| NewWebSocketClient |  | | | https://docs.kucoin.com/#websocket-feed |

</details>

<details>
<summary>Withdrawal</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| WithdrawalQuotas | GET /api/v1/withdrawals/quotas | currency | BTC | https://docs.kucoin.com/#get-withdrawal-quotas |
| Withdrawals | GET /api/v1/withdrawals | | | https://docs.kucoin.com/#get-withdrawals-list |
| ApplyWithdrawal | POST /api/v1/withdrawals | currency, address,amount | KCS,{address},0.1 | https://docs.kucoin.com/#apply-withdraw-2 |
| CancelWithdrawal | DELETE /api/v1/withdrawals/{withdrawalId} | withdrawalId | {withdrawal ID} | https://docs.kucoin.com/#cancel-withdrawal |

</details>

<details>
<summary>Currency</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| Currencies | GET /api/v1/currencies | | | https://docs.kucoin.com/#get-currencies |
| Currency | GET /api/v1/currencies/BTC | currency | BTC | https://docs.kucoin.com/#get-currency-detail |
| Prices | GET /api/v1/prices | base, currencies | USD,KCS | https://docs.kucoin.com/#get-fiat-price |

</details>

<details>
<summary>Symbol</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| Symbols | GET /api/v1/symbols                         | | | https://docs.kucoin.com/#get-symbols-list |
| TickerLevel1 | GET /api/v1/market/orderbook/level1 | symbol | BTC-USDT | https://docs.kucoin.com/#get-ticker |
| Tickers | GET /api/v1/market/allTickers | | |  https://docs.kucoin.com/#get-all-tickers |
| AggregatedPartOrderBook | GET /api/v1/market/orderbook/level2_{depth} | symbol, depth | BTC-USDT,20 | https://docs.kucoin.com/#get-part-order-book-aggregated |
| AggregatedFullOrderBook | GET /api/v2/market/orderbook/level2 | symbol | BTC-USDT |https://docs.kucoin.com/#get-full-order-book-aggregated|
| AtomicFullOrderBook | GET GET /api/v1/market/orderbook/level3 | symbol | BTC-USDT | https://docs.kucoin.com/#get-full-order-book-atomic |
| TradeHistories | GET /api/v1/market/histories | symbol | BTC-USDT | https://docs.kucoin.com/#get-trade-histories |
| KLines | GET /api/v1/market/candles | symbol, type | BTC-USDT,30min | https://docs.kucoin.com/#get-klines |
| Stats24hr | GET /api/v1/market/stats | symbol | BTC-USDT | https://docs.kucoin.com/#get-24hr-stats |
| Markets | GET /api/v1/markets | | | https://docs.kucoin.com/#get-market-list |

</details>

<details>
<summary>Time</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| ServerTime | GET /api/v1/timestamp | | | https://docs.kucoin.com/#server-time |

</details>

<details>
<summary>Service Status</summary>

| API Name | Request | Parameter | Parameter Example | Description |
| -------- | -------- | -------- | -------- | -------- |
| ServiceStatus | GET /api/v1/status | | | https://docs.kucoin.com/#service-status |

</details>