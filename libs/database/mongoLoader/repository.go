package mongoLoader

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRepository[T any] interface {
	FindOne(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*T, error)
	Find(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]T, error)
	InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error)
	InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter bson.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter bson.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type Repository[T any] struct {
	Collection *mongo.Collection
}

func NewRepository[T any](collection *mongo.Collection) IRepository[T] {
	return &Repository[T]{Collection: collection}
}

// FindOne implements IRepository.
func (c *Repository[T]) FindOne(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*T, error) {
	var result T
	if err := c.Collection.FindOne(ctx, filter, opts...).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

// Find implements IRepository.
func (c *Repository[T]) Find(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]T, error) {
	cursor, err := c.Collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	result := make([]T, 0)
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// InsertMany implements IRepository.
func (c *Repository[T]) InsertMany(ctx context.Context, data []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	result, err := c.Collection.InsertMany(ctx, data, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// InsertOne implements IRepository.
func (c *Repository[T]) InsertOne(ctx context.Context, data interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	result, err := c.Collection.InsertOne(ctx, data, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateOne implements IRepository.
func (c *Repository[T]) UpdateOne(ctx context.Context, filter bson.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := c.Collection.UpdateOne(ctx, filter, data, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateMany implements IRepository.
func (c *Repository[T]) UpdateMany(ctx context.Context, filter bson.M, data interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	result, err := c.Collection.UpdateMany(ctx, filter, data, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteOne implements IRepository.
func (c *Repository[T]) DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	result, err := c.Collection.DeleteOne(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}

	return result, nil
}
