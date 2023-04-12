package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSocketIOJSFile(t *testing.T) {
	r := RunAPIServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/socket.io/socket.io.js", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, "get", w.Body.String())
}
