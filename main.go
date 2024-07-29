package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tundrawork/powcha/config"
)

func main() {
	config.Init()
	h := server.Default(
		server.WithHandleMethodNotAllowed(true),
	)
	register(h)
	h.Spin()
}
