package cartFeatures

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/database/mongoLoader"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockCartRepo is a mock implementation of the mongoLoader.IRepository[models.CartModel] interface
type MockCartRepo struct {
	mock.Mock
}

var _ mongoLoader.IRepository[models.CartModel] = (*MockCartRepo)(nil)

func (m *MockCartRepo) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockCartRepo) DeleteOne(ctx context.Context, filter primitive.M, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func (m *MockCartRepo) Find(ctx context.Context, filter primitive.M, opts ...*options.FindOptions) ([]models.CartModel, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).([]models.CartModel), args.Error(1)
}

func (m *MockCartRepo) FindOne(ctx context.Context, filter primitive.M, opts ...*options.FindOneOptions) (*models.CartModel, error) {
	args := m.Called(ctx, filter, opts)
	return args.Get(0).(*models.CartModel), args.Error(1)
}

func (m *MockCartRepo) UpdateOne(ctx context.Context, filter primitive.M, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func (m *MockCartRepo) InsertMany(ctx context.Context, documents []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	args := m.Called(ctx, documents, opts)
	return args.Get(0).(*mongo.InsertManyResult), args.Error(1)
}

func (m *MockCartRepo) UpdateMany(ctx context.Context, filter primitive.M, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update, opts)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func TestCreateCart(t *testing.T) {
	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Create a mock CartRepo
	mockRepo := new(MockCartRepo)

	// Create a mock ServiceContext
	mockSvCtx := &contexts.ServiceContext{
		CartRepo: mockRepo,
	}

	// Create a new Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a sample request body
	customerId := primitive.NewObjectID()
	reqBody := createCartReq{
		Name:       "Test Cart",
		CustomerId: customerId,
		Products: []models.CartItem{
			{Product: models.ProductModel{ID: primitive.NewObjectID()}, Quantity: 2},
		},
		Price: 100,
	}

	// Convert request body to JSON
	jsonBody, _ := json.Marshal(reqBody)
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/cart", bytes.NewBuffer(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	// Set up expectations
	mockRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("models.CartModel"), mock.Anything).Return(&mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil)

	// Create a cartFeature instance
	cf := NewCartFeature(mockSvCtx)

	// Call the CreateCart function
	cf.CreateCart(c)

	// Assert the response
	assert.Equal(t, http.StatusCreated, w.Code)

	// Verify that our expectations were met
	mockRepo.AssertExpectations(t)
}
