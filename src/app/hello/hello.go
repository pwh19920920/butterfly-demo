package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

func hello(context *gin.Context) {
	response.BuildResponseBadRequest(context, "hello")
}

func init() {
	var route []server.RouteInfo
	route = append(route, server.RouteInfo{HttpMethod: server.HttpGet, Path: "/", HandlerFunc: hello})
	server.RegisterRoute("/hello", route)
}
