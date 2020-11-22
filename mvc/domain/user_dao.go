package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ag3ntsc4rn/golang-microservices/mvc/utils"
)

type userDao struct {
}

var (
	users = map[int64]*User{
		123: {
			ID:        123,
			FirstName: "Manish",
			LastName:  "Menon",
			Email:     "manish@example.com",
		},
	}
	UserDao userDaoInterface
)

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func init() {
	UserDao = &userDao{}
}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	log.Println("Connecting to database")
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
