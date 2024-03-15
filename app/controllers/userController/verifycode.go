package usercontroller

import (
	"xinmiao/app/apiException"
	"xinmiao/app/services/userServices"
	"xinmiao/app/utils"

	"github.com/gin-gonic/gin"
)

type sendcodeForm struct {
	UserID   string `json:"userID"`
	MailTo   string `json:"mailto"`
	Category string `json:"category"` // 业务场景
}
type checkcodeForm struct {
	UserID   string `json:"userID"`
	Code     string `json:"code"`
	Category string `json:"category"`
}

func SendMailVerifyCode(c *gin.Context) {
	var postForm sendcodeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	// 找用户
	user, err := userServices.GetUserByUserID(postForm.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
	}

	err = userServices.SendCodeByMail(user.ID, postForm.MailTo, postForm.Category)
	if err != nil {
		_ = c.AbortWithError(200, apiException.SendVerifyCodeError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

func CheckVerifyCode(c *gin.Context) {
	var postForm checkcodeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	user, err := userServices.GetUserByUserID(postForm.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
	}

	ok, err := userServices.CheckCode(user.UserID, postForm.Category, postForm.Code)
	if err != nil || !ok {
		_ = c.AbortWithError(200, apiException.WrongVerificationCode)
	}

	utils.JsonSuccessResponse(c, nil)
}
