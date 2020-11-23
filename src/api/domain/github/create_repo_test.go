package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest {
		Name: "GolangIntro",
		Description: "GolangIntro",
		Homepage: "https://github.com",
		Private: true,
		HasIssues: true,
		HasProjects: true,
		HasWiki: true,
	}
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	var expected CreateRepoRequest
	err = json.Unmarshal(bytes, &expected)
	assert.Nil(t, err)
	assert.EqualValues(t, expected, request)
	// assert.EqualValues(t, `{"name":"GolangIntro","description":"GolangIntro","homepage":"https://github.com","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))
}