package domain

import (
	"fmt"
)

var users = map[int64]*User{
	123: {
		ID: 1,
		FirstName: "Manish",
		LastName: "Menon",
		Email: "manish@example.com",
	},
}


func GetUser(userID int64) (*User, error){
	if user := users[userID]; user != nil{
		return user, nil
	}
	return nil, fmt.Errorf("User %v was not found", userID)
}