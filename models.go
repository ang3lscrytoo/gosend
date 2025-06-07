package gosend

import "time"

type CreateInvoiceOptions struct {
	CurrencyType   string `json:"currency_type,omitempty"`
	Asset          string `json:"asset,omitempty"`
	Fiat           string `json:"fiat,omitempty"`
	AcceptedAssets string `json:"accepted_assets,omitempty"`
	Amount         string `json:"amount"`
	Description    string `json:"description,omitempty"`
	HiddenMessage  string `json:"hidden_message,omitempty"`
	PaidBtnName    string `json:"paid_btn_name,omitempty"`
	PaidBtnUrl     string `json:"paid_btn_url,omitempty"`
	Payload        string `json:"payload,omitempty"`
	AllowComments  bool   `json:"allow_comments,omitempty"`
	AllowAnonymous bool   `json:"allow_anonymous,omitempty"`
	ExpiresIn      int    `json:"expires_in,omitempty"`
}

type DeleteInvoiceOptions struct {
	InvoiceId int64 `json:"invoice_id"`
}

type CreateCheckOptions struct {
	Asset         string `json:"asset"`
	Amount        string `json:"amount"`
	PinToUserId   int64  `json:"pin_to_user_id,omitempty"`
	PinToUsername string `json:"pin_to_username,omitempty"`
}

type DeleteCheckOptions struct {
	CheckId int64 `json:"check_id"`
}

type TransferOptions struct {
	UserId                  int64  `json:"user_id"`
	Asset                   string `json:"asset"`
	Amount                  string `json:"amount"`
	SpendId                 string `json:"spend_id"`
	Comment                 string `json:"comment,omitempty"`
	DisableSendNotification bool   `json:"disable_send_notification,omitempty"`
}

type GetInvoicesOptions struct {
	Asset      string   `json:"asset,omitempty"`
	Fiat       string   `json:"fiat,omitempty"`
	InvoiceIds []string `json:"invoice_ids,omitempty"`
	Status     string   `json:"status,omitempty"`
	Offset     int      `json:"offset,omitempty"`
	Count      int      `json:"count,omitempty"`
}

type GetTransfersOptions struct {
	Asset       string `json:"asset,omitempty"`
	TransferIds string `json:"transfer_ids,omitempty"`
	SpendId     string `json:"spend_id,omitempty"`
	Offset      int    `json:"offset,omitempty"`
	Count       int    `json:"count,omitempty"`
}

type GetChecksOptions struct {
	Asset    string   `json:"asset,omitempty"`
	CheckIds []string `json:"check_ids,omitempty"`
	Status   string   `json:"status,omitempty"`
	Offset   int      `json:"offset,omitempty"`
	Count    int      `json:"count,omitempty"`
}

type GetStatsOptions struct {
	StartAt string `json:"start_at,omitempty"`
	EndAt   string `json:"end_at,omitempty"`
}

type Invoice struct {
	InvoiceId         int64     `json:"invoice_id"`
	Hash              string    `json:"hash"`
	CurrencyType      string    `json:"currency_type"`
	Asset             string    `json:"asset,omitempty"`
	Fiat              string    `json:"fiat,omitempty"`
	Amount            string    `json:"amount"`
	PaidAsset         string    `json:"paid_asset,omitempty"`
	PaidAmount        string    `json:"paid_amount,omitempty"`
	PaidFiatRate      string    `json:"paid_fiat_rate,omitempty"`
	AcceptedAssets    []string  `json:"accepted_assets,omitempty"`
	FeeAsset          string    `json:"fee_asset,omitempty"`
	FeeAmount         string    `json:"fee_amount,omitempty"`
	Fee               string    `json:"fee,omitempty"`
	PayUrl            string    `json:"pay_url,omitempty"`
	BotInvoiceUrl     string    `json:"bot_invoice_url"`
	MiniAppInvoiceUrl string    `json:"mini_app_invoice_url,omitempty"`
	WebAppInvoiceUrl  string    `json:"web_app_invoice_url,omitempty"`
	Description       string    `json:"description,omitempty"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	PaidUsdRate       string    `json:"paid_usd_rate,omitempty"`
	UsdRate           string    `json:"usd_rate,omitempty"`
	AllowComments     bool      `json:"allow_comments"`
	AllowAnonymous    bool      `json:"allow_anonymous"`
	ExpirationDate    time.Time `json:"expiration_date,omitempty"`
	PaidAt            time.Time `json:"paid_at,omitempty"`
	PaidAnonymously   bool      `json:"paid_anonymously,omitempty"`
	Comment           string    `json:"comment,omitempty"`
	HiddenMessage     string    `json:"hidden_message,omitempty"`
	Payload           string    `json:"payload,omitempty"`
	PaidBtnName       string    `json:"paid_btn_name,omitempty"`
	PaidBtnUrl        string    `json:"paid_btn_url,omitempty"`
}

type Transfer struct {
	TransferId  int64     `json:"transfer_id"`
	SpendId     string    `json:"spend_id"`
	UserId      int64     `json:"user_id"`
	Asset       string    `json:"asset"`
	Amount      string    `json:"amount"`
	Status      string    `json:"status"`
	CompletedAt time.Time `json:"completed_at"`
	Comment     string    `json:"comment,omitempty"`
}

type Check struct {
	CheckId     int64     `json:"check_id"`
	Hash        string    `json:"hash"`
	Asset       string    `json:"asset"`
	Amount      string    `json:"amount"`
	BotCheckUrl string    `json:"bot_check_url"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ActivatedAt time.Time `json:"activated_at,omitempty"`
}

type Balance struct {
	CurrencyCode string `json:"currency_code"`
	Available    string `json:"available"`
	OnHold       string `json:"onhold"`
}

type ExchangeRate struct {
	IsValid  bool   `json:"is_valid"`
	IsCrypto bool   `json:"is_crypto"`
	IsFiat   bool   `json:"is_fiat"`
	Source   string `json:"source"`
	Target   string `json:"target"`
	Rate     string `json:"rate"`
}

type Currency struct {
	IsBlockchain bool   `json:"is_blockchain"`
	IsStablecoin bool   `json:"is_stablecoin"`
	IsFiat       bool   `json:"is_fiat"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	Url          string `json:"url"`
	Decimals     int    `json:"decimals"`
}

type Me struct {
	AppId                        int    `json:"app_id"`
	Name                         string `json:"name"`
	PaymentProcessingBotUsername string `json:"payment_processing_bot_username"`
}

type AppStats struct {
	Volume              float64   `json:"volume"`
	Conversion          float64   `json:"conversion"`
	UniqueUsersCount    int       `json:"unique_users_count"`
	CreatedInvoiceCount int       `json:"created_invoice_count"`
	PaidInvoiceCount    int       `json:"paid_invoice_count"`
	StartAt             time.Time `json:"start_at"`
	EndAt               time.Time `json:"end_at"`
}

type WebhookUpdate struct {
	UpdateId    int64     `json:"update_id"`
	UpdateType  string    `json:"update_type"`
	RequestDate time.Time `json:"request_date"`
	Payload     Invoice   `json:"payload"`
}

type APIResponse[T any] struct {
	Ok     bool          `json:"ok"`
	Result T             `json:"result,omitempty"`
	Error  ResponseError `json:"error,omitempty"`
}

func (response *APIResponse[T]) Err(method string) error {
	return APIError{
		Code:    response.Error.Code,
		Message: response.Error.Message,
		Method:  method,
	}
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"name"`
}

type InvoicesResult struct {
	Items []*Invoice `json:"items"`
}

type TransfersResult struct {
	Items []*Transfer `json:"items"`
}

type ChecksResult struct {
	Items []*Check `json:"items"`
}

// App представляет информацию о приложении (ответ getMe)
type App struct {
	AppId int64  `json:"app_id"`
	Name  string `json:"name"`
}

// ============= Supported Values Constants =============

const (
	AssetUSDT = "USDT"
	AssetTON  = "TON"
	AssetBTC  = "BTC"
	AssetETH  = "ETH"
	AssetLTC  = "LTC"
	AssetBNB  = "BNB"
	AssetTRX  = "TRX"
	AssetUSDC = "USDC"
	AssetJET  = "JET"

	FiatUSD = "USD"
	FiatEUR = "EUR"
	FiatRUB = "RUB"
	FiatBYN = "BYN"
	FiatUAH = "UAH"
	FiatGBP = "GBP"
	FiatCNY = "CNY"
	FiatKZT = "KZT"
	FiatUZS = "UZS"
	FiatGEL = "GEL"
	FiatTRY = "TRY"
	FiatAMD = "AMD"
	FiatTHB = "THB"
	FiatINR = "INR"
	FiatBRL = "BRL"
	FiatIDR = "IDR"
	FiatAZN = "AZN"
	FiatAED = "AED"
	FiatPLN = "PLN"
	FiatILS = "ILS"

	CurrencyTypeCrypto = "crypto"
	CurrencyTypeFiat   = "fiat"

	InvoiceStatusActive  = "active"
	InvoiceStatusPaid    = "paid"
	InvoiceStatusExpired = "expired"

	CheckStatusActive    = "active"
	CheckStatusActivated = "activated"

	TransferStatusCompleted = "completed"

	PaidBtnNameViewItem    = "viewItem"
	PaidBtnNameOpenChannel = "openChannel"
	PaidBtnNameOpenBot     = "openBot"
	PaidBtnNameCallback    = "callback"

	UpdateTypeInvoicePaid = "invoice_paid"
)
