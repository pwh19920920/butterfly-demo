package main

import (
	_ "butterfly-demo/src/app/hello"
	_ "butterfly-demo/src/app/test"
	"github.com/gin-gonic/gin"
	"github.com/pwh19920920/butterfly"
	"github.com/pwh19920920/butterfly/response"
	"github.com/pwh19920920/butterfly/server"
)

func init() {
	server.Register404Route(func(context *gin.Context) {
		response.Response(context, 404, "404", nil)
	})

	server.Register500Route(func(context *gin.Context) {
		response.Response(context, 500, "500", nil)
	})

	server.RegisterMiddleware()
}

func main() {
	butterfly.Run()
}
