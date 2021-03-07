package rpc

import (
	"github.com/valyala/fasthttp"
)

func NewApiContext(ctx *fasthttp.RequestCtx) (context *ApiContext) {
	context = &ApiContext{
		Ctx:         ctx,
		SessionData: make(map[string]interface{}),
	}
	return
}

type ApiContext struct {
	Ctx         *fasthttp.RequestCtx
	SessionData map[string]interface{}
}
