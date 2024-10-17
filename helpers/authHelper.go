package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func MathcUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uid := c.GetString("uid")

	if userType == "USER" && uid != userId {
		err = errors.New("unauthorized to access this resource")
		return
	}

	err = CheckUserType(c, userType)
	return
}

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	if userType != role {
		err = errors.New("unauthorized to access this resource")
		return err
	}
	return
}
