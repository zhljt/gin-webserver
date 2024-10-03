/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-22 21:12:54
 * @LastEditTime: 2024-10-02 23:03:24
 */
package impls

import (
	"context"
	"errors"

	"github.com/zhljt/gin-webserver/config"
	g "github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/initialize"
	"github.com/zhljt/gin-webserver/model/system/request"
	"go.uber.org/zap"
)

type IinitDBMysqlImpl struct{}

func (dbi *IinitDBMysqlImpl) InitDB(ctx context.Context, dbReq *request.InitDBRequest) (next context.Context, err error) {
	// TODO: implement database initialization
	// 0. 创建数据库 // 只有mysql pgsql 才需要创建数据库
	dsn := dbReq.MysqlDsn()
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	log.Info("init db", zap.String("dsn", dsn))
	if createDatabase(ctx, dsn, dbReq.DBType, dbReq.DBName) != nil {
		return nil, err
	}
	// 1. 连接数据库
	gormDB, err := connDB(dsn, dbReq.DBType, "")
	if err != nil {
		return ctx, err
	}
	// gormDB.InstanceSet("gorm:table_options", fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s ", dbReq.Engine, dbReq.Charset))
	g.G_GormDB = gormDB
	next = context.WithValue(ctx, g.DB_OBJ_KEY, gormDB)
	return next, err

}

func (dbi *IinitDBMysqlImpl) InitTables(ctx context.Context, models initialize.MigrateModelOlderSlice) (next context.Context, err error) {
	// TODO: implement table creation
	for _, model := range models {
		next, err := model.MigrateTable(ctx)
		if err != nil {
			return ctx, err
		}
		ctx = next
	}
	return ctx, nil
}

func (dbi *IinitDBMysqlImpl) InitData(ctx context.Context, models initialize.MigrateModelOlderSlice) (next context.Context, err error) {
	// TODO: implement initial data insertion
	for _, model := range models {
		next, err := model.InsertData(ctx)
		if err != nil {
			return ctx, err
		}
		ctx = next
	}
	return ctx, nil
}

func (dbi *IinitDBMysqlImpl) SaveConfig(ctx context.Context) error {
	// TODO: implement config saving
	c, ok := ctx.Value(g.DB_CONFIG_KEY).(config.MysqlDBConfig)
	if !ok {
		return errors.New("config not found in context")
	}
	// 更新全局变量配置
	g.G_SystemConfig.GormDB.MysqlDB = c

	// 保存到配置文件中
	config.WriteYaml("config/config.yaml", g.G_SystemConfig)
	return nil
}
