/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-30 22:52:54
 * @LastEditTime: 2024-10-04 00:38:55
 */
package impls

import (
	"context"

	"github.com/zhljt/gin-webserver/config"
	g "github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/initialize"
	"github.com/zhljt/gin-webserver/model/system/request"
	"go.uber.org/zap"
)

type InitSqlliteImpl struct{}

func (dbi *InitSqlliteImpl) InitDB(ctx context.Context, dbReq *request.InitDBRequest) (next context.Context, err error) {
	// TODO: implement database initialization
	// 0. 创建数据库 // 只有mysql pgsql 才需要创建数据库
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	log.Info("Init sqllite database")
	dsn := dbReq.SqliteDsn()
	// 1. 连接数据库
	log.Info("oepn sqlite database", zap.String("dsn", dsn))
	gormDB, err := connDB(dsn, dbReq.DBType, "")
	if err != nil {
		log.DPanic("open sqlite database failed", zap.Error(err))
		return ctx, err
	}
	// gormDB.InstanceSet("gorm:table_options", fmt.Sprintf("ENGINE=%s DEFAULT CHARSET=%s ", dbReq.Engine, dbReq.Charset))
	g.G_GormDB = gormDB
	// 更新全局变量配置
	g.G_SystemConfig.GormDB.SqlLiteDB = config.SqlLiteDBConfig{
		BaseDB: config.BaseDB{
			Path:     dbReq.DBPath,
			Database: dbReq.DBName,
		},
		Driver: dbReq.DBType,
	}
	next = context.WithValue(ctx, g.DB_OBJ_KEY, gormDB)
	return next, err

}

func (dbi *InitSqlliteImpl) InitTables(ctx context.Context, models initialize.MigrateModelOlderSlice) (next context.Context, err error) {
	// TODO: implement table initialization
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	for _, model := range models {
		log.Debug("Init sqllite db tables", zap.String("table_name", model.TableName()))
		next, err := model.MigrateTable(ctx)
		if err != nil {
			return ctx, err
		}
		ctx = next
	}
	return ctx, nil
}

func (dbi *InitSqlliteImpl) InitData(ctx context.Context, models initialize.MigrateModelOlderSlice) (next context.Context, err error) {
	// TODO: implement data initialization
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	for _, model := range models {
		log.Debug("Init sqllite db data", zap.String("table_name", model.TableName()))
		next, err := model.InsertData(ctx)
		if err != nil {
			return ctx, err
		}
		ctx = next
	}
	return ctx, nil
}

func (dbi *InitSqlliteImpl) SaveConfig(ctx context.Context) error {
	// TODO: implement save config
	// c, ok := ctx.Value(g.DB_CONFIG_KEY).(config.SqlLiteDBConfig)
	// if !ok {
	// 	return errors.New("config not found in context")
	// }

	// g.G_SystemConfig.GormDB.SqlLiteDB = c

	// 保存到配置文件中
	return config.WriteYaml("config/config.yaml", g.G_SystemConfig)

}
