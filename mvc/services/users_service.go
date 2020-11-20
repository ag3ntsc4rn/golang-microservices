package services

import (
	"github.com/ag3ntsc4rn/golang-microservices/mvc/domain"
)

func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}