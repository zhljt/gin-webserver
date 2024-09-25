/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-22 23:01:22
 * @LastEditTime: 2024-09-25 22:03:29
 */
package system

import (
	"context"

	"github.com/zhljt/gin-webserver/model/system/request"
	"github.com/zhljt/gin-webserver/service/system/impls"
)

type InitDBService interface {
	InitDB(dbReq *request.InitDBRequest) error
	InitTables(ctx context.Context) error
	InitData(ctx context.Context) error
	SaveConfig(ctx context.Context) error
}

type InitModelSeq interface {
	TableName() string

	CheckTable(ctx context.Context) bool
	MigrateTable(ctx context.Context) error

	CheckData(ctx context.Context) bool
	InsertData(ctx context.Context) error
}

func GetDBIniter(dbType string) InitDBService {
	switch dbType {
	case "mysql":
		return &impls.MysqlDBInitImpl{}
	case "sqlite3":
		return nil
	case "postgres":
		return nil
	default:
		return &impls.MysqlDBInitImpl{}
	}

}
