package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/clients/restclient"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/github"
)

const (
	headerAuthorization       = "Authorization"
	headerAuthorizationFormat = "token %s"
	urlCreateRepo             = "https://api.github.com/user/repos"
)

func getAuthorizationHeader(accessToken string) string {
	return fmt.Sprintf(headerAuthorizationFormat, accessToken)
}

func CreateRepo(accessToken string, request github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GitHubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuthorization, getAuthorizationHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, request, headers)
	if err != nil {
		log.Println(fmt.Sprintf("error when trying to create new repo in github: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "invalid response body",
		}
	}
	defer response.Body.Close()
	
	if response.StatusCode > 299 {
		var errResponse github.GitHubErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &github.GitHubErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    "invalid json response body",
			}
		}
		errResponse.StatusCode = response.StatusCode
		return nil, &errResponse
	}
	var result github.CreateRepoResponse
	if err := json.Unmarshal(bytes, &result); err != nil {
		log.Println(fmt.Sprintf("error when trying to unmarshal create new repo successful response: %s", err.Error()))
		return nil, &github.GitHubErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "error when trying to unmarshal github create repo response",
		}
	}
	return &result, nil
}
