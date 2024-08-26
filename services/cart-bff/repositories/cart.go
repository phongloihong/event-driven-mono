package repositories

import (
	"github.com/phongloihong/event-driven-mono/libs/database/mongoLoader"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
)

func NewCartRepository(svCtx contexts.ServiceContext) mongoLoader.IRepository[models.CartModel] {
	return mongoLoader.NewRepository[models.CartModel](svCtx.MongoConn.Collection(svCtx.Cfg.DatabaseName))
}
