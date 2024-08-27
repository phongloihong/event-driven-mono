package httpV1

import (
	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
)

func InitV1Routes(route *gin.Engine, serviceCtx *contexts.ServiceContext) {
	v1 := route.Group("/v1")
	newCartRoutes(v1, serviceCtx)
}
