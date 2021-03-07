package rpc

import (
	"encoding/json"
	"fmt"
)

type ApiType interface {
	Json() (data []byte, err error)
	Unjson() (data []byte, err error)
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ApiRawType struct {
	DebugMode bool `json:"-"`
}

type ApiRpc struct {
	ApiRawType
	initComplete bool
	Id           RequestId `json:"id"`
	Jsonrpc      string    `json:"jsonrpc"`
	Signature    string    `json:"sign,omitempty"`
	ApiKey       string    `json:"-"`
}

type ApiRequest struct {
	ApiRpc
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params,omitempty"`
}

type ApiResponse struct {
	ApiRpc
	Error  *RpcError       `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type ApiMethodProcessor func(ctx *ApiContext, request *ApiRequest, response *ApiResponse) (err error)

func (r *ApiRequest) String() string {
	b, _ := json.MarshalIndent(r, "", " ")
	return string(b)
}

type RequestId string

func (id RequestId) String() string {
	return string(id)
}

func (id RequestId) MarshalJSON() ([]byte, error) {
	if id == "" {
		return []byte("0"), nil
	}
	out := fmt.Sprintf("%s", id)
	return []byte(out), nil
}

func (id *RequestId) UnmarshalJSON(data []byte) error {
	var dc []byte
	for _, b := range data {
		fc := true
		if b < 58 && b > 47 {
			if fc && b != 48 {
				fc = false
			}
			if !fc {
				dc = append(dc, b)
			}
		}
	}
	if len(dc) == 0 {
		*id = RequestId("0")
		return nil
	}
	*id = RequestId(dc)
	return nil
}
