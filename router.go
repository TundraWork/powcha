package main

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	cors "github.com/tundrawork/hertz-cors"
	handler "github.com/tundrawork/powcha/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	cors := cors.New(cors.Config{
		//AllowOrigins:     config.Conf.CORSAllowOrigins,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           time.Hour,
	})

	r.NoMethod(cors)

	api := r.Group("/api")
	api.Use(cors)

	api.GET("/challenge", handler.Challenge)
	api.POST("/validate", handler.Validate)
}
