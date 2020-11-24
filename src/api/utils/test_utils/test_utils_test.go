package test_utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMockContext(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/someendpoint", nil)
	request.Header = http.Header{"X-Mock": {"true"}}
	response := httptest.NewRecorder()
	c := GetMockContext(request, response)
	assert.EqualValues(t, "true", c.Request.Header.Get("X-Mock"))
}
