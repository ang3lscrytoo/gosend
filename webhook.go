package gosend

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type WebhookManager struct {
	invoiceHandler InvoiceHandler
	tokenHash      []byte
}

func NewWebhookManager(token string) *WebhookManager {
	hash := sha256.New()
	hash.Write([]byte(token))
	return &WebhookManager{tokenHash: hash.Sum(nil)}
}

func (client *Client) WebhookFastHTTP(ctx *fasthttp.RequestCtx) {
	data := ctx.PostBody()
	signature, _ := hex.DecodeString(string(ctx.Request.Header.Peek("crypto-pay-api-signature")))

	if !client.webhookManager.checkSignature(data, signature) {
		ctx.SetStatusCode(fasthttp.StatusForbidden)
		return
	}

	var webhookUpdate WebhookUpdate
	if err := json.Unmarshal(data, &webhookUpdate); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)

	go client.webhookManager.invoiceHandler(&webhookUpdate.Payload)
}

func (wm WebhookManager) checkSignature(requestBody, requestSignature []byte) bool {
	mac := hmac.New(sha256.New, wm.tokenHash)
	mac.Write(requestBody)
	return hmac.Equal(mac.Sum(nil), requestSignature)
}
