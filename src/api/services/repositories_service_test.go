package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/clients/restclient"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	restclient.StartMockups()
	os.Exit(m.Run())
}

func TestCreateRepoInvalidInputName(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	request := repositories.CreateRepoRequest{
		Name: "   ",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid repository name", err.Message())
	restclient.StopMockups()
}

func TestCreateRepoErrorFromGithub(t *testing.T) {
	restclient.StartMockups()
	restclient.FlushMockups()
	restclient.AddMockup(restclient.Mock{
		URL:        "https://api.github.com/user/repos",
		HTTPMethod: http.MethodPost,
		Err:        errors.New("invalid restclient response"),
	})
	request := repositories.CreateRepoRequest{
		Name: "Testing",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.EqualValues(t, "invalid restclient response", err.Message())
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
	request := repositories.CreateRepoRequest{
		Name: "Testing",
	}
	response, err := RepositoryService.CreateRepo(request)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, 123, response.ID)
	assert.EqualValues(t, "", response.Name)
	assert.EqualValues(t, "", response.Owner)
	restclient.StopMockups()
}
