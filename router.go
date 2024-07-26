package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "github.com/tundrawork/powcha/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.GET("/challenge", handler.Challenge)
	r.POST("/validate", handler.Validate)
}
