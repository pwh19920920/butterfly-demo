package test

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

func test(context *gin.Context) {
	response.BuildResponseBadRequest(context, "test")
}

func init() {
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/", HandlerFunc: test})
	server.RegisterRoute("/test", route)
}
