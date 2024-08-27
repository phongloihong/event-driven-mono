package httpV1

import (
	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/features/v1/cartFeatures"
)

func newCartRoutes(group *gin.RouterGroup, serviceCtx *contexts.ServiceContext) {
	cartFeatures := cartFeatures.NewCartFeature(serviceCtx)
	group.POST("/cart", cartFeatures.CreateCart)
}
