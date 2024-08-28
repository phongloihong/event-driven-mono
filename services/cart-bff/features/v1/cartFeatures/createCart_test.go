package cartFeatures

import (
	"errors"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/log"
	"github.com/phongloihong/event-driven-mono/libs/mocks"
	"github.com/phongloihong/event-driven-mono/libs/models"
	"github.com/phongloihong/event-driven-mono/libs/test/apitest"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CartTestSuite struct {
	apitest.TestSuite
	mockRepo *mocks.IRepository[models.CartModel]
	cf       cartFeature // Change this from *cartFeature to cartFeature
}

func (s *CartTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	s.Router = gin.New()
	s.mockRepo = mocks.NewIRepository[models.CartModel](s.T())

	// Initialize your cartFeature with the mock repository
	s.cf = NewCartFeature(&contexts.ServiceContext{
		CartRepo: s.mockRepo,
		Log:      log.NewLogger(),
	})

	// Set up the route
	s.Router.POST("/v1/cart", s.cf.CreateCart)
}

func (s *CartTestSuite) TestCreateCart() {
	testCases := []apitest.APITestCase{
		{
			Name:   "Success",
			Method: http.MethodPost,
			URL:    "/v1/cart",
			Body: createCartReq{
				Name:       "Test Cart",
				CustomerId: primitive.NewObjectID(),
				Products: []models.CartItem{
					{Product: models.ProductModel{ID: primitive.NewObjectID()}, Quantity: 2},
				},
				Price: 100,
			},
			SetupMock: func() {
				s.mockRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("models.CartModel"), mock.Anything).
					Return(&mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}, nil).Once()
			},
			ExpectedStatus: http.StatusCreated,
			CheckResponse: func(t *testing.T, response map[string]interface{}) {
				assert.NotNil(t, response["data"], "Response data should not be nil")
				data, ok := response["data"].(map[string]interface{})
				assert.True(t, ok, "Response data should be a map")
				assert.NotEmpty(t, data["id"])
				assert.Equal(t, "Test Cart", data["name"])
				assert.NotEmpty(t, data["customer_id"])
				assert.Equal(t, float64(100), data["price"])

				products, ok := data["products"].([]interface{})
				assert.True(t, ok, "Products should be a slice")
				assert.Len(t, products, 1)
				product, ok := products[0].(map[string]interface{})
				assert.True(t, ok, "Product should be a map")
				assert.NotEmpty(t, product["product"].(map[string]interface{})["id"])
				assert.Equal(t, float64(2), product["quantity"])
			},
		},
		{
			Name:   "InvalidRequest",
			Method: http.MethodPost,
			URL:    "/v1/cart",
			Body:   "invalid json",
			SetupMock: func() {
			},
			ExpectedStatus: http.StatusBadRequest,
			CheckResponse: func(t *testing.T, response map[string]interface{}) {
				assert.Contains(t, response, "error")
				assert.Equal(t, "Invalid request payload", response["error"])
			},
		},
		{
			Name:   "DatabaseError",
			Method: http.MethodPost,
			URL:    "/v1/cart",
			Body: createCartReq{
				Name:       "Test Cart",
				CustomerId: primitive.NewObjectID(),
				Products: []models.CartItem{
					{Product: models.ProductModel{ID: primitive.NewObjectID()}, Quantity: 2},
				},
				Price: 100,
			},
			SetupMock: func() {
				s.mockRepo.On("InsertOne", mock.Anything, mock.AnythingOfType("models.CartModel"), mock.Anything).
					Return(nil, errors.New("database error")).Once()
			},
			ExpectedStatus: http.StatusInternalServerError,
			CheckResponse: func(t *testing.T, response map[string]interface{}) {
				assert.Contains(t, response, "error")
				assert.Equal(t, "Failed to create cart", response["error"])
			},
		},
	}

	s.RunAPITests(testCases)
}

func TestCartSuite(t *testing.T) {
	suite.Run(t, new(CartTestSuite))
}
