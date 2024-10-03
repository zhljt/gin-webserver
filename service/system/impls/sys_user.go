package impls

import (
	"fmt"
	"math/rand"

	"github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/model/system"
)

type UserServiceImpl struct{}

func (US *UserServiceImpl) Login(info system.User) (userInfo system.User, err error) {
	userlog := global.G_ZapLogger.Named("system.user")
	userlog.Debug(fmt.Sprintf("user login: %T", info))
	// 查询用户是否已经登陆 判断是否使用缓存查询
	// 验证用户名和密码
	// 数据库返回用户信息
	// 变更用登陆状态

	if MockLogin() {
		userlog.Debug("login success")
		return system.User{}, nil
	}
	userlog.Debug("login failed")
	return system.User{}, fmt.Errorf("login failed")
}

func (US *UserServiceImpl) Register(info system.User) (userInfo system.User, err error) {
	userlog := global.G_ZapLogger.Named("system.user")
	userlog.Debug(fmt.Sprintf("user register: %v", info))
	// 验证用户名是否重复
	// 数据库插入用户信息
	if MockRegister() {
		userlog.Debug("register success")
		return system.User{}, nil
	}
	userlog.Debug("register failed")
	return system.User{}, fmt.Errorf("register failed")
}

func (US *UserServiceImpl) ChangePassword(id uint, opassword string, npassword string) (err error) {
	userlog := global.G_ZapLogger.Named("system.user")
	userlog.Debug(fmt.Sprintf("user change password: id=%d, old password=%s, new password=%s", id, opassword, npassword))
	// 查询用户信息
	// 验证旧密码是否正确
	// 数据库更新密码
	userlog.Debug("change password success")
	return nil

}

func (US *UserServiceImpl) ResetPassword(ids []uint) (err error) {
	userlog := global.G_ZapLogger.Named("system.user")
	userlog.Debug(fmt.Sprintf("user reset password: ids=%v", ids))
	// 查询用户信息
	// 数据库更新密码
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
