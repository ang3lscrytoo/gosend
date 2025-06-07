
# gosend

gosend - High-performance [FastHTTP][FastHTTP] based client for CryptoPay API with full [API][CryptoPay] implementation (one-to-one)

[CryptoPay]: https://help.send.tg/en/articles/10279948-crypto-pay-api
[FastHTTP]: https://github.com/valyala/fasthttp
[FastHTTPRouter]: github.com/fasthttp/router
## Install



```shell
go get github.com/ang3lscrytoo/gosend
```
    
## Features

- High-performance
- One-to-one implementation
- Long-polling support
- Webhook support (for now only for FastHTTP)
- Clear methods


## Getting Started

```go
CryptoPayClient = gosend.NewClient(gosend.TestNet, "12345:XXXX")

me, _ := CryptoPayClient.GetMe()
log.Println(me.Name)
```

## Configure NewClient
- network - you can choose two networks: gosend.MainNet (for main) or gosend.TestNet (for testing)
- token - your application token
## Handling Invoices

To handle payments, you can use a simple handler that is installed for all types of receiving updates (Long-polling or Webhook)

```go
CryptoPayClient.OnInvoicePaid(func(invoice *gosend.Invoice) {
       log.Println(fmt.Sprintf("Payment received %d for the amount of %s", invoice.InvoiceId, invoice.Amount)  
    })
```
## Polling example

Before you start polling, make sure you have configured the OnInvoicePaid handler.

```go
_ = CryptoPayClient.StartPolling()
defer CryptoPayClient.StopPolling()
```
## Webhook example
WebhookFastHTTP returns the default RequestHandler func(ctx *RequestCtx) for FastHTTP Server

```go
srv := fasthttp.Server{
		Handler: CryptoPayClient.WebhookFastHTTP,
	}

go func() { _ = srv.ListenAndServe(":80") }()
```

You can also use your endpoints with [router][FastHTTPRouter]
```go
r := router.New()
r.POST("/webhook", CryptoPayClient.WebhookFastHTTP)

go func() { _ = srv.ListenAndServe(":80", r.Handler) }()
```
## Methods
gosend presents all methods available in CryptoPay API
- GetMe()
- CreateInvoice()
- DeleteInvoice()
- CreateCheck()
- DeleteCheck()
- Transfer()
- GetInvoices()
- GetTransfers()
- GetChecks()
- GetBalance()
- GetExchangeRates()
- GetCurrencies()
- GetStats()
## CreateInvoice Example
This method creates an Invoice for 100 RUB with payment in all available cryptocurrencies. All parameters are described through a structure with the pattern MethodNameOptions{} - in this case CreateInvoiceOptions{}
```go
CryptoPayClient.CreateInvoice(gosend.CreateInvoiceOptions{
		CurrencyType: gosend.CurrencyTypeFiat, // "fiat" or "crypto"
		Fiat:         gosend.FiatRUB,
		Amount:       "100",
		ExpiresIn:    1200, // expires in seconds
		Payload:      fmt.Sprintf("%d", telegramId), // your custom payload. In this example - telegram id
	})
```
