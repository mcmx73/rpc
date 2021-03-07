package main

import (
	"fmt"
	"github.com/mcmx73/rpc"
	"github.com/valyala/fasthttp"
)

func main() {
	apiRouter := rpc.NewRpcRouter()
	apiRouter.RegisterProcessor("users", func(ctx *rpc.ApiContext, request *rpc.ApiRequest, response *rpc.ApiResponse) (err error) {
		// ... do logic ...
		userList := []string{"Alice", "Bob", "John", "Zorro"}
		response.SetResult(userList)
		return nil
	})
	err := fasthttp.ListenAndServe(":8085", apiRouter.Handle)
	if err != nil {
		fmt.Println("Can not serve JSON RPC:", err)
	}
}
