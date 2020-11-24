package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigConstants(t *testing.T) {
	assert.EqualValues(t, "SECRET_GITHUB_ACCESS_TOKEN", apiGithubAccessToken)
}

func TestGithubAccessToken(t *testing.T){
	assert.EqualValues(t, "", GetGithubAccessToken())
}
