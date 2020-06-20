package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestWelcome
func TestWelcome(t *testing.T) {
	router := gin.New()
	router.Handle("GET", "/welcome", Welcome)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/welcome", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":200,\"data\":{\"content\":\"hello welcome\"},\"message\":\"success\"}", w.Body.String())
}
