package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-sso/models"
)

func DoLogin(c *gin.Context,user models.Users)  error{
	customClaims :=CustomClaims{
		UserId:         user.Id,
	}
	accessToken, err :=customClaims.MakeToken()
	if err != nil {
		return err
	}
	fmt.Println(accessToken)
	c.Header(HEADER_TOKEN,accessToken)
	return nil
}