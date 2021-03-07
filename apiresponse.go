package rpc

import "encoding/json"

//NewResponse create new JSON-RPC response object
func NewResponse(requestId RequestId) *ApiResponse {
	response := new(ApiResponse)
	response.Id = requestId
	response.Jsonrpc = JSON_RPC_VERSION
	return response
}

//Json marshal response to []byte slice
func (r *ApiResponse) Json() (data []byte, err error) {
	return json.Marshal(r)
}

//Json marshal response to []byte slice
func (r *ApiResponse) JsonBytes() (data []byte) {
	data, _ = json.Marshal(r)
	return
}

//Unjson unmarshal json bytes to ApiResponse object
func (r *ApiResponse) Unjson(data []byte) (err error) {
	return json.Unmarshal(data, r)
}

//SetError set error response
func (r *ApiResponse) SetError(code int, message string) {
	if r.Error == nil {
		r.Error = new(RpcError)
	}
	r.Error.Code, r.Error.Message = code, message
}

//SetResult set request processing result
func (r *ApiResponse) SetResult(result interface{}) (err error) {
	rawMessage, err := json.Marshal(result)
	r.Result = rawMessage
	return
}

//SetResultBytes set request processing result as json bytes
func (r *ApiResponse) SetResultBytes(result []byte) {
	r.Result = result
	return
}

//ParseResult parse response to Go structure
func (r *ApiResponse) ParseResult(target interface{}) (err error) {
	return json.Unmarshal(r.Result, target)
}
