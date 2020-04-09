package sdk

import (
	"net/http"
)

// A WithdrawalModel represents a withdrawal.
type WithdrawalModel struct {
	Id         string `json:"id"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Currency   string `json:"currency"`
	Amount     string `json:"amount"`
	Fee        string `json:"fee"`
	WalletTxId string `json:"walletTxId"`
	IsInner    bool   `json:"isInner"`
	Status     string `json:"status"`
	Remark     string `json:"remark"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

// A WithdrawalsModel is the set of *WithdrawalModel.
type WithdrawalsModel []*WithdrawalModel

// Withdrawals returns a list of withdrawals.
func (as *ApiService) Withdrawals() (*ApiResponse, error) {
	params := map[string]string{}
	pagination := &PaginationParam{CurrentPage: 1, PageSize: 10}
	pagination.ReadParam(params)
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals", params)
	return as.Call(req)
}

// A WithdrawalQuotasModel represents the quotas for a currency.
type WithdrawalQuotasModel struct {
	Currency            string `json:"currency"`
	AvailableAmount     string `json:"availableAmount"`
	RemainAmount        string `json:"remainAmount"`
	WithdrawMinSize     string `json:"withdrawMinSize"`
	LimitBTCAmount      string `json:"limitBTCAmount"`
	InnerWithdrawMinFee string `json:"innerWithdrawMinFee"`
	UsedBTCAmount       string `json:"usedBTCAmount"`
	IsWithdrawEnabled   bool   `json:"isWithdrawEnabled"`
	WithdrawMinFee      string `json:"withdrawMinFee"`
	Precision           uint8  `json:"precision"`
}

// WithdrawalQuotas returns the quotas of withdrawal.
func (as *ApiService) WithdrawalQuotas(currency string) (*ApiResponse, error) {
	params := map[string]string{"currency": currency}
	req := NewRequest(http.MethodGet, "/api/v1/withdrawals/quotas", params)
	return as.Call(req)
}

// ApplyWithdrawalResultModel represents the result of ApplyWithdrawal().
type ApplyWithdrawalResultModel struct {
	WithdrawalId string `json:"withdrawalId"`
}

// ApplyWithdrawal applies a withdrawal.
func (as *ApiService) ApplyWithdrawal(currency, address, amount string) (*ApiResponse, error) {
	p := map[string]string{
		"currency": currency,
		"address":  address,
		"amount":   amount,
	}
	options := map[string]string{}
	for k, v := range options {
		p[k] = v
	}
	req := NewRequest(http.MethodPost, "/api/v1/withdrawals", p)
	return as.Call(req)
}

// CancelWithdrawalResultModel represents the result of CancelWithdrawal().
type CancelWithdrawalResultModel struct {
	CancelledWithdrawIds []string `json:"cancelledWithdrawIds"`
}

// CancelWithdrawal cancels a withdrawal by withdrawalId.
func (as *ApiService) CancelWithdrawal(withdrawalId string) (*ApiResponse, error) {
	req := NewRequest(http.MethodDelete, "/api/v1/withdrawals/"+withdrawalId, nil)
	return as.Call(req)
}
