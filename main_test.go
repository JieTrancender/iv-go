package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPingRoute 测试路由/ping
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

// TestWelcomeRoute 测试路由/welcome?firstname=xx&lastname=xx
func TestWelcomeRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/welcome?lastname=Jestin", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello iv Jestin", w.Body.String())
}

// TestFormPost 测试路由/form_post
func TestFormPost(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/form_post", strings.NewReader("message=hello world"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"code\":200,\"message\":\"hello world\",\"nick\":\"iv\"}", w.Body.String())
}

// TestMapPost 测试路由/map_post
func TestMapPost(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/map_post?ids[a]=1234&ids[b]=hello", strings.NewReader("names[first]=jane&names[second]=jason"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ids: map[a:1234 b:hello], names: map[first:jane second:jason]", w.Body.String())
}

// TestUpload 测试上传文件
func TestUpload(t *testing.T) {
	// router := setupRouter()

	// w := httptest.NewRecorder()
	// req, _ := http.NewRequest("POST", "/upload", nil)

	// router.ServeHTTP(w, req)

	// assert.Equal(t, http.StatusOK, w.Code)

}
