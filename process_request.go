package rpc

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
)

type ApiRpcClient struct {
	Url         *RpcURL
	RpcUser     string
	RpcPassword string
	auth        string
	fullUrl     string
	client      *fasthttp.Client
}

func NewApiRpcClient(url *RpcURL, user, password string) (client *ApiRpcClient) {
	client = &ApiRpcClient{
		Url:         url,
		RpcUser:     user,
		RpcPassword: password,
	}
	client.Prepare()
	return
}

func (c *ApiRpcClient) Prepare() {
	c.fullUrl = c.Url.String()
	c.prepareAuth()
}

func (c *ApiRpcClient) SetUrl(url string) {
	c.fullUrl = url
}

func (c *ApiRpcClient) prepareAuth() {
	if c.RpcUser != "" {
		c.auth = base64.StdEncoding.EncodeToString([]byte(c.RpcUser + ":" + c.RpcPassword))
	}
}

func (c *ApiRpcClient) QueryRaw(data []byte) (result []byte, err error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)
	req.SetRequestURI(c.fullUrl)
	req.Header.SetContentType(MIME_TYPE_JSON)
	if c.auth != "" {
		req.Header.Set("Authorization", "Basic "+c.auth)
	}
	req.Header.SetMethod("POST")
	req.SetBody(data)
	err = c.client.Do(req, resp)
	if err != nil {
		return
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		err = fmt.Errorf("invalid server response: %v", resp.StatusCode())
	}
	result = resp.Body()
	return
}

func (c *ApiRpcClient) Query(request *ApiRequest, target interface{}) (err error) {
	body, err := request.Json()
	if err != nil {
		return
	}
	responseBody, err := c.QueryRaw(body)
	if err != nil {
		if len(responseBody) == 0 {
			return
		}
	}
	if target == nil {
		return
	}
	jsonResponse := NewResponse()
	err = jsonResponse.Unjson(responseBody)
	if err != nil {
		return
	}
	if jsonResponse.Error != nil {
		return fmt.Errorf(jsonResponse.Error.Message)
	}
	if target != nil {
		err = json.Unmarshal(jsonResponse.Result, target)
	}
	return
}
