package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIGet(t *testing.T) {

	t.Run("TestTrending", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/trending", nil)
		response := mockRequest(req)

		assert.Equal(t, http.StatusOK, response.Code)

		if body := response.Body.String(); body == "" {
			t.Errorf("Expected response . Got %s", body)
		}

	})
	t.Run("TestTrendingByLanguage", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/trending/Go", nil)
		response := mockRequest(req)

		assert.Equal(t, http.StatusOK, response.Code)

		body := response.Body
		if body.String() == "" {
			t.Errorf("Expected response . Got %s", body)
		}
	})
	t.Run("TestErrorHandling", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/trending/", nil)
		response := mockRequest(req)

		assert.Equal(t, http.StatusNotFound, response.Code)

	})

}
func mockRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	service := NewService()
	service.SetupService()
	service.Router.ServeHTTP(rr, req)

	return rr
}
