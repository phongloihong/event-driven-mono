package apitest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// APITestCase defines the structure for API test cases
type APITestCase struct {
	Name           string
	Method         string
	URL            string
	Body           interface{}
	SetupMock      func()
	ExpectedStatus int
	CheckResponse  func(*testing.T, map[string]interface{})
}

// TestSuite is a base test suite with common functionality
type TestSuite struct {
	suite.Suite
	Router *gin.Engine
}

// RunAPITests runs the provided API test cases
func (s *TestSuite) RunAPITests(testCases []APITestCase) {
	for _, tc := range testCases {
		s.Run(tc.Name, func() {
			// Setup mock if provided
			tc.SetupMock()

			// Prepare request
			body, _ := json.Marshal(tc.Body)
			req, _ := http.NewRequest(tc.Method, tc.URL, bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Serve the request
			s.Router.ServeHTTP(w, req)

			// Check status code
			assert.Equal(s.T(), tc.ExpectedStatus, w.Code)

			// Parse and check response
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			if err != nil {
				s.T().Errorf("Failed to parse response body: %v", err)
				return
			}
			tc.CheckResponse(s.T(), response)
		})
	}
}
