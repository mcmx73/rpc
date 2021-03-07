package main

import (
	"fmt"
	"github.com/mcmx73/rpc"
	"github.com/valyala/fasthttp"
)

const (
	LISTEN_ADDRESS = ":8085"
)

func main() {
	apiRouter := rpc.NewRpcRouter()
	apiRouter.RegisterProcessor("users", func(ctx *rpc.ApiContext, request *rpc.ApiRequest, response *rpc.ApiResponse) (err error) {
		// ... do logic ...
		userList := []string{"Alice", "Bob", "John", "Zorro"}
		response.SetResult(userList)
		return nil
	})
	apiRouter.RegisterProcessor("edit_user", ProcessComplexRequest)
	fmt.Println("Listen address: [", LISTEN_ADDRESS, "]")
	err := fasthttp.ListenAndServe(LISTEN_ADDRESS, apiRouter.Handle)
	if err != nil {
		fmt.Println("Can not serve JSON RPC:", err)
	}
}

type MyRequestParams struct {
	UserId int             `json:"user_id"`
	Flags  map[string]bool `json:"flags"`
}

func ProcessComplexRequest(ctx *rpc.ApiContext, request *rpc.ApiRequest, response *rpc.ApiResponse) (err error) {
	myRequestParams := &MyRequestParams{}
	err = request.ParseParamsToObject(myRequestParams)
	if err != nil {
		response.SetError(rpc.ERROR_CODE_INVALID_REQUEST, rpc.ERROR_MESSAGE_INVALID_REQUEST)
		return err
	}
	// ... do logic ...
	response.SetResult(struct {
		Success bool `json:"success"`
	}{Success: true})
	return nil
}
