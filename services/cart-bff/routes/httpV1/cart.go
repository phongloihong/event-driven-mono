package httpV1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/repositories"
)

func newCartRoutes(group *gin.RouterGroup, serviceCtx contexts.ServiceContext) {
	group.POST("/cart", func(c *gin.Context) {
		ctx := c.Request.Context()
		newCart := models.CartModel{
			Name: "Cart 1",
			Products: []models.CartItem{{
				Product: models.ProductModel{
					Name:  "Product 1",
					Price: 100,
				},
				Quantity: 1,
			}},
			Price: 100,
		}

		cartRepo := repositories.NewCartRepository(serviceCtx)
		insertCartResult, err := cartRepo.InsertOne(ctx, newCart)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		newCart.ID = fmt.Sprintf("%v", insertCartResult.InsertedID)
		c.JSON(http.StatusOK, gin.H{"data": newCart})
	})
}
