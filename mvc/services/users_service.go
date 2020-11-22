package services

import (
	"github.com/ag3ntsc4rn/golang-microservices/mvc/domain"
	"github.com/ag3ntsc4rn/golang-microservices/mvc/utils"
)

type usersService struct {
}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	
	return user, nil
}
