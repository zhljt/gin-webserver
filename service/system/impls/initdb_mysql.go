/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-22 21:12:54
 * @LastEditTime: 2024-09-25 22:12:57
 */
package impls

import (
	"context"

	"github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/model/system/request"
)

type MysqlDBInitImpl struct{}

func (dbi *MysqlDBInitImpl) InitDB(dbReq *request.InitDBRequest) error {
	// TODO: implement db initialization
	ctx := context.Background()

	// 0. 创建数据库 // 只有mysql pgsql 才需要创建数据库
	dsn := dbReq.MysqlDsn()
	createDatabase(dsn, dbReq.DBType, dbReq.DBName)
	// 1. 连接数据库
	gormDB, err := connDB(dsn, dbReq.DBType, "")
	if err != nil {
		return err
	}
	// gormDB.InstanceSet("gorm:table_options", fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s ", dbReq.Engine, dbReq.Charset))
	global.GormDB = gormDB
	sqlDB, _ := gormDB.DB()
	sqlDB.SetMaxIdleConns(global.SystemConfig.GormDB.MysqlDB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(global.SystemConfig.GormDB.MysqlDB.MaxOpenConns)
	// 2. create tables
	dbi.InitTables(ctx)
	// 3. insert initial data
	dbi.InitData(ctx)
	// 4. save config
	dbi.SaveConfig(ctx)
	return nil

}

func (dbi *MysqlDBInitImpl) InitTables(ctx context.Context) error {
	// TODO: implement table creation
	return nil
}

func (dbi *MysqlDBInitImpl) InitData(ctx context.Context) error {
	// TODO: implement initial data insertion
	return nil
}

func (dbi *MysqlDBInitImpl) SaveConfig(ctx context.Context) error {
	// TODO: implement config saving
	return nil
}
