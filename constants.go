package rpc

const (
	MIME_TYPE_JSON = "application/json"

	JSON_RPC_VERSION = "2.0"

	ERROR_CODE_PARSE_ERROR         = -32700
	ERROR_MESSAGE_PARSE_ERROR      = "Parse error"
	ERROR_CODE_INVALID_REQUEST     = -32600
	ERROR_MESSAGE_INVALID_REQUEST  = "invalid request"
	ERROR_CODE_METHOD_NOT_FOUND    = -32601
	ERROR_MESSAGE_METHOD_NOT_FOUND = "method not found"
	ERROR_CODE_SERVER_ERROR        = -32000
	ERROR_MESSAGE_SERVER_ERROR     = "server error"
)
