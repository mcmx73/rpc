package rpc

import "errors"

var (
	ErrParamNotFound             = errors.New("param not found")
	ErrInvalidParamType          = errors.New("invalid param type")
	ErrMethodProcessorDuplicated = errors.New("method processor duplicated")
	ErrMethodNotFound            = errors.New("method not found")
	ErrRequestIsNil              = errors.New("request is nil")
)
