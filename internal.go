package gosend

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"net"
	"time"
)

type Core struct {
	httpClient *fasthttp.Client
	baseUrl    string
	token      string
}

func NewCore(baseUrl string, token string) *Core {
	client := &fasthttp.Client{
		DialTimeout: func(addr string, timeout time.Duration) (net.Conn, error) {
			return net.DialTimeout("tcp", addr, timeout)
		},
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	return &Core{
		httpClient: client,
		baseUrl:    baseUrl,
		token:      token,
	}
}

func (core *Core) Post(endpoint string, data any, target any) error {
	url := fmt.Sprintf("%s%s", core.baseUrl, endpoint)
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("POST")
	req.Header.Add(headerSecretToken, core.token)
	req.SetRequestURI(url)

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return err
		}
		req.SetBody(jsonData)
		req.Header.Set("Content-Type", "application/json")
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := core.httpClient.Do(req, resp)
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Body(), &target)
}
