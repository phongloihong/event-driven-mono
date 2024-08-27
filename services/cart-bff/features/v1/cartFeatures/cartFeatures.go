package cartFeatures

import "github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"

type cartFeature struct {
	SvCtx *contexts.ServiceContext
}

func NewCartFeature(svCtx *contexts.ServiceContext) cartFeature {
	return cartFeature{
		SvCtx: svCtx,
	}
}
