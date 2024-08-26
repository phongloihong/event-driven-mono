package models

type CartModel struct {
	ID       string     `json:"id" bson:"_id"`
	Name     string     `json:"name" bson:"name"`
	Products []CartItem `json:"products" bson:"products"`
	Price    int        `json:"price" bson:"price"`
}

type CartItem struct {
	Product  ProductModel `json:"product" bson:"product"`
	Quantity int          `json:"quantity" bson:"quantity"`
}
