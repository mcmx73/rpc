package rpc

import (
	"github.com/valyala/fasthttp"
	"sync"
)

type RpcRouter struct {
	mux        sync.RWMutex
	processors map[string]ApiMethodProcessor
}

func NewRpcRouter() *RpcRouter {
	router := new(RpcRouter)
	router.Init()
	return router
}

// process fasthttp http request
func (r *RpcRouter) Handle(ctx *fasthttp.RequestCtx) {
	if string(ctx.Request.Header.Method()) == "POST" {
		request := NewRequest()
		err := request.Unjson(ctx.Request.Body())
		if err != nil {
			response := NewResponse("")
			response.SetError(ERROR_CODE_INVALID_REQUEST, ERROR_MESSAGE_INVALID_REQUEST)
			ctx.Response.Header.SetContentType(MIME_TYPE_JSON)
			ctx.Response.SetBody(response.JsonBytes())
			return
		}
		apiCtx := NewApiContext(ctx)
		response, err := r.Process(apiCtx, request)
		ctx.Response.Header.SetContentType(MIME_TYPE_JSON)
		ctx.Response.SetBody(response.JsonBytes())
	} else {
		ctx.Error("JSON-RPC 2.0", 405)
	}
}

func (r *RpcRouter) Init() {
	r.mux.Lock()
	defer r.mux.Unlock()
	r.processors = make(map[string]ApiMethodProcessor)
}

//Register RPC Request method
func (r *RpcRouter) RegisterProcessor(method string, processor ApiMethodProcessor) (err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if _, ok := r.processors[method]; ok {
		err = ErrMethodProcessorDuplicated
	} else {
		r.processors[method] = processor
	}
	return
}

// Route and  process JSON RPC Request
func (r *RpcRouter) Process(context *ApiContext, request *ApiRequest) (response *ApiResponse, err error) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	response = NewResponse(request.Id)
	//check BaseAuth
	if processor, ok := r.processors[request.Method]; ok {
		err = processor(context, request, response)
		response.Id = request.Id
	} else {
		err = ErrMethodNotFound
		response.SetError(ERROR_CODE_METHOD_NOT_FOUND, ERROR_MESSAGE_METHOD_NOT_FOUND)
	}
	return
}
