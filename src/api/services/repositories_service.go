package services

import (
	"strings"

	"github.com/ag3ntsc4rn/golang-microservices/src/api/config"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/github"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/domain/repositories"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/providers/github_provider"
	"github.com/ag3ntsc4rn/golang-microservices/src/api/utils/errors"
)

type reposService struct{}

type reposServiceInterface interface {
	CreateRepo(request repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repository name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Private:     false,
		Description: input.Description,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		ID:    response.ID,
		Name:  response.Name,
		Owner: response.Owner.Login,
	}

	return &result, nil
}
