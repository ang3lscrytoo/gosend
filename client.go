package gosend

import (
	"context"
	"fmt"
)

const (
	headerSecretToken string = "Crypto-Pay-API-Token"
	getMe             string = "api/getMe"
	createInvoice     string = "api/createInvoice"
	deleteInvoice     string = "api/deleteInvoice"
	createCheck       string = "api/createCheck"
	deleteCheck       string = "api/deleteCheck"
	transfer          string = "api/transfer"
	getInvoices       string = "api/getInvoices"
	getChecks         string = "api/getChecks"
	getTransfers      string = "api/getTransfers"
	getBalance        string = "api/getBalance"
	getExchangeRates  string = "api/getExchangeRates"
	getCurrencies     string = "api/getCurrencies"
	getStats          string = "api/getStats"
)

type Client struct {
	Network        CryptoPayNetwork
	Token          string
	internal       *Core
	pollingManager *PollingManager
	webhookManager *WebhookManager
}

type APIError struct {
	Code    int
	Message string
	Method  string
}

func (e APIError) Error() string {
	return fmt.Sprintf("%s | Ошибка #%d - %s", e.Method, e.Code, e.Message)
}

func NewClient(
	network CryptoPayNetwork,
	token string,
) *Client {
	baseUrl := string(network)

	internal := NewCore(baseUrl, token)
	ctx, cancel := context.WithCancel(context.Background())

	return &Client{
		Network:  network,
		Token:    token,
		internal: internal,
		pollingManager: &PollingManager{
			trackedInvoices: make(map[int64]*TrackedInvoice),
			ctx:             ctx,
			cancel:          cancel,
			Period:          3,
		},
		webhookManager: NewWebhookManager(token),
	}
}

func (client *Client) GetMe() (*Me, error) {
	var response APIResponse[Me]
	err := client.internal.Post(getMe, nil, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getMe)
	}

	return &response.Result, nil
}

func (client *Client) CreateInvoice(opt CreateInvoiceOptions) (*Invoice, error) {
	var response APIResponse[Invoice]
	err := client.internal.Post(createInvoice, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(createInvoice)
	}

	if client.pollingManager.pollingActive {
		client.pollingManager.TrackInvoice(&response.Result)
	}

	return &response.Result, nil
}

func (client *Client) DeleteInvoice(opt DeleteInvoiceOptions) (bool, error) {
	var response APIResponse[bool]
	err := client.internal.Post(deleteInvoice, opt, &response)

	if err != nil {
		return false, err
	}
	if !response.Ok {
		return false, response.Err(deleteInvoice)
	}

	return response.Result, nil
}

func (client *Client) CreateCheck(opt CreateCheckOptions) (*Check, error) {
	var response APIResponse[Check]
	err := client.internal.Post(createCheck, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(createCheck)
	}

	return &response.Result, nil
}

func (client *Client) DeleteCheck(opt DeleteCheckOptions) (bool, error) {
	var response APIResponse[bool]
	err := client.internal.Post(deleteCheck, opt, &response)

	if err != nil {
		return false, err
	}
	if !response.Ok {
		return false, response.Err(deleteCheck)
	}

	return response.Result, nil
}

func (client *Client) Transfer(opt TransferOptions) (*Transfer, error) {
	var response APIResponse[Transfer]
	err := client.internal.Post(transfer, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(transfer)
	}

	return &response.Result, nil
}

func (client *Client) GetInvoices(opt GetInvoicesOptions) ([]*Invoice, error) {
	var response APIResponse[InvoicesResult]
	err := client.internal.Post(getInvoices, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getInvoices)
	}

	return response.Result.Items, nil
}

func (client *Client) GetTransfers(opt GetTransfersOptions) ([]*Transfer, error) {
	var response APIResponse[TransfersResult]
	err := client.internal.Post(getTransfers, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getTransfers)
	}

	return response.Result.Items, nil
}

func (client *Client) GetChecks(opt GetChecksOptions) ([]*Check, error) {
	var response APIResponse[ChecksResult]
	err := client.internal.Post(getChecks, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getChecks)
	}

	return response.Result.Items, nil
}

func (client *Client) GetBalance() ([]*Balance, error) {
	var response APIResponse[[]*Balance]
	err := client.internal.Post(getBalance, nil, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getBalance)
	}

	return response.Result, nil
}

func (client *Client) GetExchangeRates() ([]*ExchangeRate, error) {
	var response APIResponse[[]*ExchangeRate]
	err := client.internal.Post(getExchangeRates, nil, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getExchangeRates)
	}

	return response.Result, nil
}

func (client *Client) GetCurrencies() ([]*Currency, error) {
	var response APIResponse[[]*Currency]
	err := client.internal.Post(getCurrencies, nil, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getCurrencies)
	}

	return response.Result, nil
}

func (client *Client) GetStats(opt GetStatsOptions) (*AppStats, error) {
	var response APIResponse[*AppStats]
	err := client.internal.Post(getStats, opt, &response)

	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, response.Err(getStats)
	}

	return response.Result, nil
}

func (client *Client) OnInvoicePaid(handler InvoiceHandler) {
	client.pollingManager.invoiceHandler = handler
	client.webhookManager.invoiceHandler = handler
}
