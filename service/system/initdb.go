/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-22 23:01:22
 * @LastEditTime: 2024-10-04 00:44:55
 */
package system

import (
	"context"
	"errors"
	"sort"

	_ "github.com/zhljt/gin-webserver/dao/system"
	g "github.com/zhljt/gin-webserver/global"
	"github.com/zhljt/gin-webserver/initialize"
	"github.com/zhljt/gin-webserver/model/system/request"
	"github.com/zhljt/gin-webserver/service/system/impls"
	"go.uber.org/zap"
)

type InitDBService interface {
	InitDB(ctx context.Context, dbReq *request.InitDBRequest) (next context.Context, err error)
	InitTables(ctx context.Context, moders initialize.MigrateModelOlderSlice) (next context.Context, err error)
	InitData(ctx context.Context, moders initialize.MigrateModelOlderSlice) (next context.Context, err error)
	SaveConfig(ctx context.Context) error
}

func InitDB(ctx context.Context, dbReq *request.InitDBRequest) error {
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	if len(initialize.MigrateModelOlderS) == 0 {
		log.DPanic("没有可迁移模型表，请检查初始化是否已经完成或未进行配置 ")
		return errors.New("没有可迁移模型表，请检查初始化是否已经完成或未进行配置 ")
	}

	sort.Sort(&initialize.MigrateModelOlderS)
	var impl InitDBService
	switch dbReq.DBType {
	case "mysql":
		impl = &impls.IinitDBMysqlImpl{}
	case "sqllite":
		impl = &impls.InitSqlliteImpl{}
	case "postgres":
		return nil
	default:
		impl = &impls.IinitDBMysqlImpl{}
	}
	// 1. create database
	next, err := impl.InitDB(ctx, dbReq)
	if err != nil {
		log.DPanic("初始化数据库失败", zap.String("err", err.Error()))
		return errors.New("初始化数据库失败")
	}
	// 2. create tables
	next, err = impl.InitTables(next, initialize.MigrateModelOlderS)
	if err != nil {
		log.DPanic("初始化表失败", zap.String("err", err.Error()))
		return errors.New("初始化表失败")
	}
	// 3. insert initial data
	next, err = impl.InitData(next, initialize.MigrateModelOlderS)
	if err != nil {
		log.DPanic("初始化数据失败", zap.String("err", err.Error()))
		return errors.New("初始化数据失败")
	}
	// 4. save config
	err = impl.SaveConfig(next)
	if err != nil {
		log.DPanic("保存配置失败", zap.String("err", err.Error()))
		return errors.New("保存配置失败")
	}

	// 清空initialize.MigrateModelOlderS
	initialize.MigrateModelOlderS = initialize.MigrateModelOlderSlice{}
	initialize.CheckOlderCache = initialize.MigrateModelOlderSMap{}
	log.Info("初始化完成")
	return nil
}
