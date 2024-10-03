/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-28 21:28:50
 * @LastEditTime: 2024-10-04 00:23:22
 */
package system

import (
	"context"
	"errors"

	g "github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/initialize"
	"github.com/zhljt/gin-webserver/model/system"
	"gorm.io/gorm"
)

//	type UserMigrateDao interface {
//		MigrateTable() error
//		InsertInitData() error
//	}
const migrateUserOlder = initialize.MigrateOrderBusiness + 1

type userDaoImpl struct{}

func init() {
	initialize.RegisterInitModelSeq(&userDaoImpl{})
}

func (u *userDaoImpl) TableName() string {
	return system.User{}.TableName()
}
func (u *userDaoImpl) Older() uint {
	return migrateUserOlder
}

func (u *userDaoImpl) MigrateTable(ctx context.Context) (context.Context, error) {
	// TODO: implement this method
	db, ok := ctx.Value(g.DB_OBJ_KEY).(*gorm.DB)
	if !ok {
		return ctx, errors.New("db not found in context")
	}

	if db.Migrator().HasTable(&system.User{}) {
		// table already exists, skip migration
	} else {
		if err := db.AutoMigrate(&system.User{}); err != nil {
			return ctx, err
		}
	}

	return ctx, nil
}
func (u *userDaoImpl) InsertData(ctx context.Context) (context.Context, error) {
	// TODO: implement this method
	db, ok := ctx.Value(g.DB_OBJ_KEY).(*gorm.DB)
	if !ok {
		return ctx, errors.New("db not found in context")
	}
	datas := []system.User{ // 假数据
		{
			ID:       1,
			Account:  "sysadmin",
			Name:     "管理员", // 假数据
			Password: "1",
			Phone:    "13800138000",
			Status:   1,
			RoleID:   1,
		},
		{
			ID:       2,
			Account:  "test",
			Name:     "测试用户",
			Password: "1",
			Phone:    "13800138001",
			Status:   1,
			RoleID:   2,
		},
	}
	return ctx, db.Create(&datas).Error
}

func (u *userDaoImpl) ValidateData(ctx context.Context) bool {
	db, ok := ctx.Value(g.DB_OBJ_KEY).(*gorm.DB)
	if !ok {
		return false
	}
	var records []system.User
	result := db.Where("id = ?", 1).First(&records)

	return result.RowsAffected != 0
}

func (u *userDaoImpl) GetUserForId(id int) (user *system.User, err error) {
	// TODO: implement this method
	return nil, nil
}
