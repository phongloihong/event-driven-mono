package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/http/ginServer"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/routes"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
	defer cancel()
	servicesCtx := contexts.NewServiceContext(ctx)

	ginCfg := ginServer.GinConfig{
		Port: "8080",
	}
	r := ginServer.NewHttpServer(&ginCfg)
	r.Use(gin.Recovery())
	routes.InitRoutes(r, servicesCtx)

	r.Run()
}
