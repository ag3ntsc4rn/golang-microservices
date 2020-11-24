package polo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}

func TestPolo(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/marco", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	Polo(c)
	assert.EqualValues(t, http.StatusOK, response.Code)
	assert.EqualValues(t, "polo", response.Body.String())
}
