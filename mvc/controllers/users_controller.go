package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/ag3ntsc4rn/golang-microservices/mvc/services"
	"github.com/ag3ntsc4rn/golang-microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}
	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}