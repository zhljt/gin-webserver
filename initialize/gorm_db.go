package initialize

import (
	"fmt"

	"github.com/zhljt/gin-webserver/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// dsn := "root:123456@tcp(139.198.115.192:8085)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(139.198.115.192:8085)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	// db.Exec("CREATE DATABASE IF NOT EXISTS mytest DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;")
	// 创建表时添加后缀
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&system.User{}, &system.Role{}, &system.UserRole{})

	entdata := []system.User{
		{
			ID:       123456,
			Account:  "ljt",
			Name:     "LINJINTING",
			Password: "1",
			Status:   1,
			RoleID:   1234,
			Roles: []system.Role{
				{
					ID:   1234,
					Code: "ptgyl",
					Name: "ptgyl",
				},
				{
					ID:   1235,
					Code: "admin",
					Name: "admin",
					Type: 1,
				},
			},
		},
	}
	if err = db.Create(&entdata).Error; err != nil {
		fmt.Print(err, system.User{}.TableName()+"表数据初始化失败!")
	}
	// tx := db.Raw("show databases").Scan(&databases)
	// tx.Exec("use at500edb_v3")
	// tx.Raw("show tables").Scan(&tables)

	// fmt.Println(databases)
	// fmt.Println(tables)
}

func GormMysql() *gorm.DB {
	// m := global.GVA_CONFIG.Mysql
	// if m.Dbname == "" {
	// 	return nil
	// }
	MaxIdleConns := 10
	MaxOpenConns := 100
	Engine := "InnoDB"
	// Prefix := ""
	// Singular := true

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	dsn := "root:123456@tcp(139.198.115.192:8085)/mytest?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		SkipInitializeWithVersion: false, // 根据版本自动配置 林鸣
	}
	// if db, err := gorm.Open(mysql.New(mysqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
	// 	return nil
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{}); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MaxIdleConns)
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		return db
	}
}
