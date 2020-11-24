package repositories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/clients/restclient"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/repositories"
	utilserror "github.com/ag3ntsc4rn/golang-microservices/src/api/utils/errors"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateRepoInvalidJsonRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(``))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)
	assert.EqualValues(t, http.StatusBadRequest, response.Code)
	apiErr, err := utilserror.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusBadRequest, apiErr.Status())
	assert.EqualValues(t, "invalid json body", apiErr.Message())
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing","description":"test description"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)
	assert.EqualValues(t, http.StatusInternalServerError, response.Code)
	apiErr, err := utilserror.NewApiErrorFromBytes(response.Body.Bytes())
	assert.Nil(t, err)
	assert.NotNil(t, apiErr)
	assert.EqualValues(t, http.StatusInternalServerError, apiErr.Status())
	assert.EqualValues(t, "invalid restclient response", apiErr.Message())
	restclient.StopMockups()
}

func TestCreateRepoNoError(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body:       ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})
	request, _ := http.NewRequest(http.MethodPost, "/repositories", strings.NewReader(`{"name":"testing","description":"test description"}`))
	response := httptest.NewRecorder()
	c := test_utils.GetMockContext(request, response)

	CreateRepo(c)
	assert.EqualValues(t, http.StatusCreated, response.Code)
	var result repositories.CreateRepoResponse
	err := json.Unmarshal(response.Body.Bytes(), &result)
	assert.Nil(t, err)
	assert.EqualValues(t, 123, result.ID)
	restclient.StopMockups()
}
