package domain

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(13)

	assert.Nil(t, user, "We were expecting a nil user.")
	assert.NotNil(t, err, "We were expecting an error.")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 13 was not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err, "Was not expecting an error. User should exist.")
	assert.EqualValues(t, 13, user.ID)
}
