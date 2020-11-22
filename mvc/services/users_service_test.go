package services

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ag3ntsc4rn/golang-microservices/mvc/utils"
	"github.com/ag3ntsc4rn/golang-microservices/mvc/domain"
)

type userDaoMock struct{
}

func (u *userDaoMock) GetUser(userID int64) (*domain.User, *utils.ApplicationError){
	return getUserFunction(userID)
}

var (
	UserDaoMock userDaoMock
	getUserFunction func (userID int64) (*domain.User, *utils.ApplicationError)
)

func init(){
	domain.UserDao = &userDaoMock{}
}

func TestGetUserNoUserInDatabase(t *testing.T) {
	getUserFunction = func (userID int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
		Message:    "User 0 was not found",
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
		}
	}
	user, err := UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "User 0 was not found", err.Message)
}

func TestGetUserFoundInDatabase(t *testing.T) {
	getUserFunction = func (userID int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
			FirstName: "Manish",
			LastName: "Menon",
			Email: "manish@example.com",
		},nil
	}
	user, err := UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}