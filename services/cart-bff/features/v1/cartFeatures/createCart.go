package cartFeatures

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type createCartReq struct {
	Name       string             `json:"name" binding:"required"`
	CustomerId primitive.ObjectID `json:"customerId" binding:"required"`
	Products   []models.CartItem  `json:"products" binding:"required,min=1"`
	Price      int                `json:"price" binding:"required,min=0"`
}

func shouldBindJson[T any](c *gin.Context) (T, error) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, err
	}
	return req, nil
}

func (cf *cartFeature) CreateCart(c *gin.Context) {
	req, err := shouldBindJson[createCartReq](c)
	if err != nil {
		cf.SvCtx.Log.Error(c.Request.Context(), "Failed to bind JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx := c.Request.Context()
	newCart := models.CartModel{
		ID:         primitive.NewObjectID(),
		Name:       req.Name,
		CustomerId: req.CustomerId,
		Products:   req.Products,
		Price:      req.Price,
	}

	_, err = cf.SvCtx.CartRepo.InsertOne(ctx, newCart)
	if err != nil {
		cf.SvCtx.Log.Error(ctx, "Failed to insert cart", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCart})
}
