package initialize

import (
	"context"
	"fmt"

	"github.com/zhljt/gin-webserver/model/system"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MigrateOrderSystem   = 10
	MigrateOrderBusiness = 1000
)

type MigrateModel interface {
	TableName() string
	Older() uint
	// CheckTable(ctx context.Context) bool
	MigrateTable(ctx context.Context) (context.Context, error)

	// CheckData(ctx context.Context) bool
	InsertData(ctx context.Context) (context.Context, error)
}

type _migrateModelOlder struct {
	order uint
	MigrateModel
}
type MigrateModelOlderSlice []*_migrateModelOlder
type MigrateModelOlderSMap map[string]*_migrateModelOlder

var (
	MigrateModelOlderS MigrateModelOlderSlice
	CheckOlderCache    MigrateModelOlderSMap
)

// Len implements sort.Interface.
func (ms MigrateModelOlderSlice) Len() int {
	return len(ms)
}

// Less implements sort.Interface.
func (ms MigrateModelOlderSlice) Less(i int, j int) bool {
	return ms[i].order < ms[j].order
}

// Swap implements sort.Interface.
func (ms MigrateModelOlderSlice) Swap(i int, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

func RegisterInitModelSeq(i MigrateModel) {
	order := i.Older()
	if MigrateModelOlderS == nil {
		MigrateModelOlderS = make(MigrateModelOlderSlice, 0)
	}

	if CheckOlderCache == nil {
		CheckOlderCache = make(MigrateModelOlderSMap)
	}
	name := i.TableName()
	if _, ok := CheckOlderCache[name]; ok {
		panic(fmt.Sprintf("model %s already registered", name))
	}
	MigrateModelOlderS = append(MigrateModelOlderS, &_migrateModelOlder{order: order, MigrateModel: i})

	CheckOlderCache[name] = &_migrateModelOlder{order: order, MigrateModel: i}
}

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
		SkipInitializeWithVersion: false, // 根据版本自动配置
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
