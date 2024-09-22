/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-22 11:13:38
 */
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	. "github.com/zhljt/gin-webserver/global"
	commonModel "github.com/zhljt/gin-webserver/model/common"
	systemModel "github.com/zhljt/gin-webserver/model/system"
	systemRequestModel "github.com/zhljt/gin-webserver/model/system/request"
	"github.com/zhljt/gin-webserver/service"
)

type UserApi struct{}

func (uapi *UserApi) Login(c *gin.Context) {
	var lreq systemRequestModel.LoginRequest
	// 绑定参数json数据
	err := c.ShouldBindJSON(&lreq)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 创建用户对象
	userLogin := &systemModel.User{
		Account:  lreq.Account,
		Password: lreq.Password,
	}
	// 调用service层
	service := service.NewUserService()
	user, err := service.Login(*userLogin)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(LOGIN_AUTH_ERROR, "登录失败"))
		return
	}
	c.JSON(http.StatusOK, commonModel.SuccessWithComplete("登录成功", user))
}

func (uapi *UserApi) Register(c *gin.Context) {
	var rres systemRequestModel.RegisterRequest
	err := c.ShouldBindJSON(&rres)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(ARG_BIND_ERROR, "绑定参数错误"))
	}
	// 验证参数
	// ...

	// 创建用户对象
	userRegister := &systemModel.User{
		Account:  rres.Account,
		Password: rres.Password,
	}
	// 调用service层
	service := service.NewUserService()
	user, err := service.Register(*userRegister)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(REGISTER_USER_EXIST_ERROR, "注册失败"))
		return
	}
	c.JSON(http.StatusOK, commonModel.SuccessWithComplete("注册成功", user))
}

func (uapi *UserApi) Logout(c *gin.Context) {
	// 调用service层
	// ...
	c.JSON(http.StatusOK, commonModel.SuccessWithMessage("退出成功"))
}

func (uapi *UserApi) ChangePassword(c *gin.Context) {
	// 绑定参数json数据
	var cprq systemRequestModel.ChangePasswordRequest
	err := c.ShouldBindJSON(&cprq)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 调用service层
	service := service.NewUserService()
	err = service.ChangePassword(cprq.UserId, cprq.OldPassword, cprq.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(CHANGE_PWD_ERROR, "修改密码失败"))
		return
	}
	c.JSON(http.StatusOK, commonModel.SuccessWithMessage("修改密码成功"))
}

func (uapi *UserApi) ResetPassword(c *gin.Context) {
	// 绑定参数json数据
	var rprq systemRequestModel.ResetPasswordRequest
	err := c.ShouldBindJSON(&rprq)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 调用service层
	service := service.NewUserService()
	err = service.ResetPassword(rprq.UserId)
	if err != nil {
		c.JSON(http.StatusOK, commonModel.ErrorWithCodeAndMessage(RESET_PWD_ERROR, "重置密码失败"))
		return
	}
	c.JSON(http.StatusOK, commonModel.SuccessWithMessage("重置密码成功"))
}
