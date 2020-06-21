package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testRouterOk(method, path string, t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// TestRouter 测试路由
func TestRouter(t *testing.T) {
	testRouterOk("GET", "/welcome", t)
	// testRouterOk("GET", "/register", t)
}
