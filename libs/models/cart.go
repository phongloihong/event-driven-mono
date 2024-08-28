package models

import (
	"context"
	"time"

	"github.com/phongloihong/event-driven-mono/libs/database/mongoLoader"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CartModel struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	CustomerId primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	Products   []CartItem         `json:"products" bson:"products"`
	Price      int                `json:"price" bson:"price"`
}

type CartItem struct {
	Product  ProductModel `json:"product" bson:"product"`
	Quantity int          `json:"quantity" bson:"quantity"`
}

func NewCartRepository(mongoConn *mongo.Database, collectionName string) mongoLoader.IRepository[CartModel] {
	collection := mongoConn.Collection(collectionName)
	indexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "customer_id", Value: 1}},
		},
	}

	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := collection.Indexes().CreateMany(context.TODO(), indexes, opts)
	if err != nil {
		logrus.Error(err)
	}

	return mongoLoader.NewRepository[CartModel](collection)
}
