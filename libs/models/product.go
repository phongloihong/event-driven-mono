package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductModel struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Price int                `json:"price" bson:"price"`
}
