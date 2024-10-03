package impls

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	g "github.com/zhljt/gin-webserver/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func createDatabase(ctx context.Context, dsn string, driver string, dbName string) error {
	log := ctx.Value(g.LOG_KEY).(*zap.Logger)
	log.Info("Creating database...")

	var query string
	// 生成创建数据库SQL语句
	switch driver {
	case "mysql":
		query = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName)

	// case "postgres":
	// 	_dialector = postgres.Open(dsn)
	default:
		query = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName)

	}

	var sqlDB *sql.DB
	defer func() {
		if sqlDB != nil { // 判断 sqlDB 是否为 nil
			if err := sqlDB.Close(); err != nil {
				log.Error("Failed to close database connection", zap.Error(err))
			}
		}
	}()

	if db, err := connDB(dsn, driver, ""); err != nil {
		log.DPanic("Failed to connect to database", zap.Error(err), zap.String("dsn", dsn), zap.String("driver", driver))
		return err
	} else {
		sqlDB, _ := db.DB()
		_, err = sqlDB.Exec(query) // 执行创建数据库的SQL语句
		if err != nil {
			log.DPanic("Failed to create database", zap.Error(err))
			return err
		}
	}
	return nil
}

func connDB(dsn string, driver string, prefix string) (*gorm.DB, error) {
	var _dialector gorm.Dialector
	switch driver {
	case "mysql":
		_dialector = mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			SkipInitializeWithVersion: false, // 根据版本自动配置表结构
		})
		return gorm.Open(_dialector, getConfig(prefix))
	// case "postgres":
	// 	_dialector = postgres.New(postgres.Config{
	// 		DSN:                  dsn,
	// 		PreferSimpleProtocol: false, // 禁用复用连接
	// 	})
	// 	return gorm.Open(postgres.Open(dsn), getConfig(prefix))
	case "sqllite":
		return gorm.Open(sqlite.Open(dsn), getConfig(prefix))
	default:
		_dialector = mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			SkipInitializeWithVersion: false, // 根据版本自动配置表结构
		})
		//gorm.Open(_dialector, getConfig(prefix))
		if db, err := gorm.Open(_dialector, getConfig(prefix)); err != nil {
			return nil, err
		} else {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxIdleConns(g.G_SystemConfig.GormDB.MysqlDB.MaxIdleConns)
			sqlDB.SetMaxOpenConns(g.G_SystemConfig.GormDB.MysqlDB.MaxOpenConns)
			return db, nil
		}

	}
}

func getConfig(prefix string) *gorm.Config {
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束
		SkipDefaultTransaction:                   true, // 禁用默认事务

		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix, // 表名前缀
			SingularTable: true,   // 使用单数表名
			NoLowerCase:   true,   // 禁用小写表名
		},
	}
	// 日志配置
	defaultLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 260 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})

	config.Logger = defaultLogger
	return config
}
