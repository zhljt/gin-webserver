package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhljt/gin-webserver/service"
)

type UserReq struct {
	Account  string `json:"account" binding:"required"`  // 账号
	Password string `json:"password" binding:"required"` // 密码
}
type UserApi struct{}

func (ua *UserApi) Login(c *gin.Context) {
	var ur UserReq
	err := c.ShouldBindJSON(&ur)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg": "parse json obj error", "code": "400101", "body": err.Error()})

	}
	info := &service.UserInfo{
		Account:  ur.Account,
		Password: ur.Password,
	}
	user, err := service.ServiceGroupPtr.UserService.Login(info)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg": "登陆失败", "code": "400102", "body": err.Error()})
	}
	c.JSON(http.StatusOK, map[string]string{"msg": "ok", "code": "0", "body": user.UUID.String()})
}

func (ua *UserApi) Register(c *gin.Context) {
	var ur UserReq
	err := c.ShouldBindJSON(&ur)
	if err != nil {
		c.JSON(http.StatusOK, map[string]string{"msg": "parse json obj error", "code": "10001", "body": err.Error()})
		return
	}

	c.JSON(http.StatusOK, map[string]string{"msg": "ok", "code": "0", "body": "登陆成功"})
}
