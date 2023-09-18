package rpc_util

import (
	gin "github.com/baytan0720/gin-grpc"
)

type RpcFunctions interface {
	Functions() Functions
}

type Function struct {
	Name    string
	Hooks   gin.HandlersChain
	Handler gin.HandlerFunc
}

type Functions []Function

func AddFunctions(e *gin.Engine, api RpcFunctions) {
	functions := api.Functions()
	if functions == nil {
		panic("invalid rpc functions")
	}

	for _, function := range functions {
		if function.Name == "" {
			panic("invalid rpc function name")
		} else {
			e.Handle(function.Name, append(function.Hooks, function.Handler)...)
		}
	}
}
