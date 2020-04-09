package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kucoin-api-demo/sdk"
	"log"
	"os"
	"reflect"
	"strings"
)

var (
	kucoin *sdk.ApiService
	PATH   = "./config.json"
)

func main() {
	//1. get api authentication info from user input or configuration file.
	var apiURI, apiKey, apiSecret, apiPassphrase string
	var apiSkipVerifyTls bool

	// read config from json file
	var config ApiConfig
	file, er := ioutil.ReadFile(PATH)
	if er != nil {
		log.Println("Failed to read configuration from json file:", er)
		log.Println("Now please input api information manually.")
	}else{
		// unmarshal config data
		err := json.Unmarshal(file, &config)
		if err != nil {
			log.Fatal("error when unmarshal config from json file:", err)
			return
		}
	}
	// if we have configuration from json file
	if len(config.ApiBaseURI) > 0 {
		apiURI = config.ApiBaseURI
	} else {
		log.Println("please input api base URI,such as:https://api.kucoin.com")
		fmt.Scanf("%s", &apiURI)
	}

	if len(config.ApiKey) > 0 {
		apiURI = config.ApiKey
	} else {
		log.Println("please input your api key...")
		fmt.Scanf("%s", &apiKey)
	}

	if len(config.ApiSecret) > 0 {
		apiURI = config.ApiSecret
	} else {
		log.Println("please input your api secret...")
		fmt.Scanf("%s", &apiSecret)
	}

	if len(config.ApiPassphrase) > 0 {
		apiURI = config.ApiPassphrase
	} else {
		log.Println("please input your api passphrase...")
		fmt.Scanf("%s", &apiPassphrase)
	}

	apiSkipVerifyTls = config.ApiSkipVerifyTls

	initApiService(apiURI, apiKey, apiSecret, apiPassphrase, apiSkipVerifyTls)

	// 2. api function map to user input.
	apiMap := initApiMap()
	// 3. get api info from user input.
	var apiName string
	log.Println("please input the api name which you want to test...")
	fmt.Scanf("%s", &apiName)
	log.Println("please input the api parameters,use ',' to spilt parameters.")
	log.Println("if it is no parameter api,just enter to skip...")

	reader := bufio.NewReader(os.Stdin)
	inputs, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("input parameter error:", err)
	}

	var params []string
	if len(strings.TrimSpace(inputs)) > 0 {
		params = strings.Split(strings.TrimSpace(inputs), ",")
	}
	// 4. call api function which user choose to test.
	call(apiMap, apiName, params)
}

// initial api configuration information from json file.
type ApiConfig struct {
	ApiBaseURI       string `json:ApiBaseURI`
	ApiKey           string `json:ApiKey`
	ApiSecret        string `json:ApiSecret`
	ApiPassphrase    string `json:ApiPassphrase`
	ApiSkipVerifyTls bool   `json:ApiSkipVerifyTls`
}

// initApiService init an api service
func initApiService(apiURI, apiKey, apiSecret, apiPassphrase string, apiSkipVerifyTls bool) {
	kucoin = sdk.NewApiService(
		sdk.ApiBaseURIOption(apiURI),
		sdk.ApiKeyOption(apiKey),
		sdk.ApiSecretOption(apiSecret),
		sdk.ApiPassPhraseOption(apiPassphrase),
		sdk.ApiSkipVerifyTlsOption(apiSkipVerifyTls),
	)
	sdk.DebugMode = true
}

// call according to string input to call function which named the input.
func call(m map[string]interface{}, name string, params []string) {
	// if string input is a function
	f := reflect.ValueOf(m[name])
	if f.Kind() != reflect.Func {
		log.Fatal("input error,your input is not a name of api.")
		return
	}

	in := make([]reflect.Value, len(params))
	// if the function has no parameters
	if len(params) > 0 {
		for k, v := range params {
			fmt.Println(v)
			in[k] = reflect.ValueOf(v)
		}
	} else {
		in = nil
	}

	f.Call(in)
}

// initApiMap init api function mapping
func initApiMap() map[string]interface{} {
	return map[string]interface{}{
		"CreateAccount":   kucoin.CreateAccount,
		"Accounts":        kucoin.Accounts,
		"Account":         kucoin.Account,
		"SubAccountUsers": kucoin.SubAccountUsers,
		"SubAccounts":     kucoin.SubAccounts,
		"SubAccount":      kucoin.SubAccount,
		"AccountLedgers":  kucoin.AccountLedgers,
		"AccountHolds":    kucoin.AccountHolds,

		"InnerTransferV2":      kucoin.InnerTransferV2,
		"SubTransfer":          kucoin.SubTransfer,
		"CreateDepositAddress": kucoin.CreateDepositAddress,
		"DepositAddresses":     kucoin.DepositAddresses,
		"V1Deposits":           kucoin.V1Deposits,
		"Deposits":             kucoin.Deposits,

		"Fills":       kucoin.Fills,
		"RecentFills": kucoin.RecentFills,

		"CreateOrder":      kucoin.CreateOrder,
		"CreateMultiOrder": kucoin.CreateMultiOrder,
		"CancelOrder":      kucoin.CancelOrder,
		"CancelOrders":     kucoin.CancelOrders,
		"Orders":           kucoin.Orders,
		"Order":            kucoin.Order,
		"RecentOrders":     kucoin.RecentOrders,

		"WebSocketPublicToken":  kucoin.WebSocketPublicToken,
		"WebSocketPrivateToken": kucoin.WebSocketPrivateToken,
		"NewWebSocketClient":    kucoin.NewWebSocketClient,

		"WithdrawalQuotas": kucoin.WithdrawalQuotas,
		"Withdrawals":      kucoin.Withdrawals,
		"ApplyWithdrawal":  kucoin.ApplyWithdrawal,
		"CancelWithdrawal": kucoin.CancelWithdrawal,

		"Currencies": kucoin.Currencies,
		"Currency":   kucoin.Currency,
		"Prices":     kucoin.Prices,

		"Symbols":                 kucoin.Symbols,
		"TickerLevel1":            kucoin.TickerLevel1,
		"Tickers":                 kucoin.Tickers,
		"AggregatedPartOrderBook": kucoin.AggregatedPartOrderBook,
		"AggregatedFullOrderBook": kucoin.AggregatedFullOrderBook,
		"AtomicFullOrderBook":     kucoin.AtomicFullOrderBook,
		"TradeHistories":          kucoin.TradeHistories,
		"KLines":                  kucoin.KLines,
		"Stats24hr":               kucoin.Stats24hr,
		"Markets":                 kucoin.Markets,

		"ServerTime": kucoin.ServerTime,

		"ServiceStatus": kucoin.ServiceStatus,
	}
}
