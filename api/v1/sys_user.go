/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-10-02 22:32:35
 */
/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-14 22:14:10
 * @LastEditTime: 2024-09-25 22:48:43
 */
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	g "github.com/zhljt/gin-webserver/global"
	model_comm "github.com/zhljt/gin-webserver/model/common"
	model_sys "github.com/zhljt/gin-webserver/model/system"
	model_sys_req "github.com/zhljt/gin-webserver/model/system/request"
	service_sys "github.com/zhljt/gin-webserver/service/system"
)

type UserApi struct{}

func (uapi *UserApi) Login(c *gin.Context) {
	var lreq model_sys_req.LoginRequest
	// 绑定参数json数据
	err := c.ShouldBindJSON(&lreq)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 创建用户对象
	userLogin := &model_sys.User{
		Account:  lreq.Account,
		Password: lreq.Password,
	}
	// 调用service层
	service := service_sys.NewUserService()
	user, err := service.Login(*userLogin)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.LOGIN_AUTH_ERROR, "登录失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithComplete("登录成功", user))
}

func (uapi *UserApi) Register(c *gin.Context) {
	var rres model_sys_req.RegisterRequest
	err := c.ShouldBindJSON(&rres)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.ARG_BIND_ERROR, "绑定参数错误"))
	}
	// 验证参数
	// ...

	// 创建用户对象
	userRegister := &model_sys.User{
		Account:  rres.Account,
		Password: rres.Password,
	}
	// 调用service层
	service := service_sys.NewUserService()
	user, err := service.Register(*userRegister)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.REGISTER_USER_EXIST_ERROR, "注册失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithComplete("注册成功", user))
}

func (uapi *UserApi) Logout(c *gin.Context) {
	// 调用service层
	// ...
	c.JSON(http.StatusOK, model_comm.SuccessWithMessage("退出成功"))
}

func (uapi *UserApi) ChangePassword(c *gin.Context) {
	// 绑定参数json数据
	var cprq model_sys_req.ChangePasswordRequest
	err := c.ShouldBindJSON(&cprq)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 调用service层
	service := service_sys.NewUserService()
	err = service.ChangePassword(cprq.UserId, cprq.OldPassword, cprq.NewPassword)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.CHANGE_PWD_ERROR, "修改密码失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithMessage("修改密码成功"))
}

func (uapi *UserApi) ResetPassword(c *gin.Context) {
	// 绑定参数json数据
	var rprq model_sys_req.ResetPasswordRequest
	err := c.ShouldBindJSON(&rprq)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.ARG_BIND_ERROR, "绑定参数错误"))

	}
	// 验证参数
	// ...

	// 调用service层
	service := service_sys.NewUserService()
	err = service.ResetPassword(rprq.UserId)
	if err != nil {
		c.JSON(http.StatusOK, model_comm.ErrorWithCodeAndMessage(g.RESET_PWD_ERROR, "重置密码失败"))
		return
	}
	c.JSON(http.StatusOK, model_comm.SuccessWithMessage("重置密码成功"))
}
