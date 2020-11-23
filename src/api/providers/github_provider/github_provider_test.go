package github_provider

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/clients/restclient"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/github"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M){
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestGetAuthorizationHeader(t *testing.T){
	header := getAuthorizationHeader("abc123")
	assert.EqualValues(t, "token abc123", header)
}

func TestCreateRepoErrorFromRestClient(t *testing.T){
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err: errors.New("invalid restclient response"),
	})
	
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	restclient.StopMockups()
}

func TestCreateRepoInvalidResponseBody(t *testing.T){
	restclient.StartMockups()
	restclient.FlushMockups()
	invalidCloser, _ := os.Open("-asf3")
	restclient.AddMockup(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: invalidCloser,
		},
	})
	
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	restclient.StopMockups()
}

func TestCreateRepoUnauthorizedInvalidResponseBody(t *testing.T){
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusUnauthorized,
			Body: ioutil.NopCloser(strings.NewReader(`{"message": 123}`)),
		},
	})
	
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid json response body", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	restclient.StopMockups()
}

func TestCreateRepoSuccessInvalidResponseBody(t *testing.T){
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": "123"}`)),
		},
	})
	
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "error when trying to unmarshal github create repo response", err.Message)
	assert.EqualValues(t, http.StatusInternalServerError, err.StatusCode)
	restclient.StopMockups()
}

func TestCreateRepoSuccess(t *testing.T){
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL: "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Response: &http.Response{
			StatusCode: http.StatusCreated,
			Body: ioutil.NopCloser(strings.NewReader(`{"id": 123}`)),
		},
	})
	
	response, err := CreateRepo("", github.CreateRepoRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.ID)
	restclient.StopMockups()
}