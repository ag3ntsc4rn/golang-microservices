package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ag3ntsc4rn/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("user_id must be a number"))
		return
	}
	user, err := services.GetUser(userID)
	if err != nil {
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte(err.Error()))
		// Hanlde the error and return to the client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}