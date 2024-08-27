package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/routes/httpV1"
)

func InitRoutes(router *gin.Engine, serviceCtx *contexts.ServiceContext) {
	httpV1.InitV1Routes(router, serviceCtx)
}
