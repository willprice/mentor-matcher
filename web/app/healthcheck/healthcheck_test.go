package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler(t *testing.T) {
	w := httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(w)

	Handler(ctx)

	// Check if the status code is 200.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	// Check if the response body is "OK".
	expectedBody := "OK"
	if w.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, w.Body.String())
	}
}
