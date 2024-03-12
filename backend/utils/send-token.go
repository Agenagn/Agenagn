package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"

	globaltypes "example.com/agenagn/global-types"
)

func SendToken(ctx *gin.Context, role string, response globaltypes.AuthResponse){
	// generate jwt token
	token, err := GetToken(response.ID, role)

	if err != nil {
		fmt.Println(err.Error(), "when generating token")
		ctx.JSON(400, gin.H{"message": "something went wrong when creating token"})
		return
	}
	response.Role = role
	response.Token = token
	ctx.JSON(200, response)
}