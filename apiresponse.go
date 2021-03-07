package rpc

import "encoding/json"

func NewResponse(requestId ...RequestId) *ApiResponse {
	response := new(ApiResponse)
	if len(requestId) != 0 {
		response.Id = requestId[0]
	}
	response.Jsonrpc = JSON_RPC_VERSION
	return response
}

func (r *ApiResponse) Json() (data []byte, err error) {
	return json.Marshal(r)
}

func (r *ApiResponse) JsonBytes() (data []byte) {
	data, _ = json.Marshal(r)
	return
}

func (r *ApiResponse) Unjson(data []byte) (err error) {
	return json.Unmarshal(data, r)
}

func (r *ApiResponse) SetError(code int, message string) {
	if r.Error == nil {
		r.Error = new(RpcError)
	}
	r.Error.Code, r.Error.Message = code, message
}

func (r *ApiResponse) SetResult(result interface{}) (err error) {
	rawMessage, err := json.Marshal(result)
	r.Result = rawMessage
	return
}

func (r *ApiResponse) SetResultBytes(result []byte) {
	r.Result = result
	return
}

func (r *ApiResponse) GetResult(target interface{}) (err error) {
	return json.Unmarshal(r.Result, target)
}
