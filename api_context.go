package rpc

import (
	"github.com/valyala/fasthttp"
)

// NewApiContext create new ApiContext based on  *fasthttp.RequestCtx
// SessionData may be used for transfer http session info
func NewApiContext(ctx *fasthttp.RequestCtx) (context *ApiContext) {
	context = &ApiContext{
		Ctx:         ctx,
		SessionData: make(map[string]interface{}),
	}
	return
}

//ApiContext based on  *fasthttp.RequestCtx
type ApiContext struct {
	Ctx         *fasthttp.RequestCtx
	SessionData map[string]interface{}
}
