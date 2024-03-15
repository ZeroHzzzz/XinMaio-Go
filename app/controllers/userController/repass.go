package usercontroller

import (
	"fmt"
	"xinmiao/app/apiException"
	"xinmiao/app/services/userServices"
	"xinmiao/app/utils"

	"github.com/gin-gonic/gin"
)

type ResetPassForm struct {
	UserID  string `json:"userID"`
	OldPass string `json:"oldpass"`
	NewPass string `json:"newpass"`
}

func ResetPass(c *gin.Context) {
	// 获取表单数据
	var postForm ResetPassForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	user, err := userServices.GetUserByUserID(postForm.UserID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.UserNotFind)
		return
	}

	status, err := userServices.GetCode(fmt.Sprintf("%s-%s", user.ID, apiException.ResetPassword))
	if err != nil || status != "Accepted" {
		_ = c.AbortWithError(200, apiException.WrongVerificationCode)
		return
	}

	userServices.DelCode(fmt.Sprintf("%s-%s", user.ID, apiException.ResetPassword)) // 过期验证码

	utils.JsonSuccessResponse(c, nil)
}
