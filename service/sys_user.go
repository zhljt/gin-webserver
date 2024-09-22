package service

import (
	"fmt"
	"math/rand"

	"github.com/zhljt/gin-webserver/model/system"
	"go.uber.org/zap"
)

type UserService interface {
	Login(info system.User) (userInfo system.User, err error)
	Register(info system.User) (userInfo system.User, err error)
	ChangePassword(id uint, opassword string, npassword string) (err error)
	ResetPassword(ids []uint) (err error)
	// GetUserInfo(id uint) (userInfo system.User, err error)
}

func NewUserService() UserService {
	return &UserServiceImpl{}
}

type UserServiceImpl struct{}

func (US *UserServiceImpl) Login(info system.User) (userInfo system.User, err error) {
	userlog := zap.L().Named("system.user")
	userlog.Info(fmt.Sprintf("user login: %v", info))
	// 查询用户是否已经登陆 判断是否使用缓存查询
	// 验证用户名和密码
	// 数据库返回用户信息
	// 变更用登陆状态

	if MockLogin() {
		userlog.Info("login success")
		return system.User{}, nil
	}
	userlog.Info("login failed")
	return system.User{}, fmt.Errorf("login failed")
}

func (US *UserServiceImpl) Register(info system.User) (userInfo system.User, err error) {
	userlog := zap.L().Named("system.user")
	userlog.Info(fmt.Sprintf("user register: %v", info))
	// 验证用户名是否重复
	// 数据库插入用户信息
	if MockRegister() {
		userlog.Info("register success")
		return system.User{}, nil
	}
	userlog.Info("register failed")
	return system.User{}, fmt.Errorf("register failed")
}

func (US *UserServiceImpl) ChangePassword(id uint, opassword string, npassword string) (err error) {
	userlog := zap.L().Named("system.user")
	userlog.Info(fmt.Sprintf("user change password: id=%d, old password=%s, new password=%s", id, opassword, npassword))
	// 查询用户信息
	// 验证旧密码是否正确
	// 数据库更新密码
	userlog.Info("change password success")
	return nil

}

func (US *UserServiceImpl) ResetPassword(ids []uint) (err error) {
	userlog := zap.L().Named("system.user")
	userlog.Info(fmt.Sprintf("user reset password: ids=%v", ids))
	// 查询用户信息
	// 数据库更新密码

	userlog.Info("reset password failed")
	return fmt.Errorf("reset password failed")
}

func MockRegister() bool {
	// 创建随机数
	if rand.Intn(100) < 50 {
		return true
	}
	return false
}

func MockLogin() bool {
	// 创建随机数
	if rand.Intn(100) < 70 {
		return true
	}
	return false
}
