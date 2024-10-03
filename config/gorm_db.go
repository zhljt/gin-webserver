/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-16 15:43:03
 * @LastEditTime: 2024-10-01 21:29:10
 */
/*
 * @Author: linjinting@gs
 * @Email: 840359545@qq.com
 * @Date: 2024-09-18 13:04:34
 * @Last Modified by: linjinting@gs
 * @Last Modified time: 2024-09-18 17:31:33
 * @FilePath: \gin-web\config\gorm_db.go
 * @Description:
 * Copyright (c) 2024 by ljt930@gmail.com, All Rights Reserved.
 */

package config

type BaseDB struct {
	// 数据库路径
	Path string `json:"path" yaml:"path"`
	// 数据库主机名
	Host string `json:"host" yaml:"host"`
	// 数据库端口
	Port string `json:"port" yaml:"port"`
	// 数据库用户名
	User string `json:"user" yaml:"user"`
	//	数据库密码
	Password string `json:"password" yaml:"password"`
	// 数据库名称
	Database string `json:"database" yaml:"database"`
	// 最大连接数
	MaxOpenConns int `json:"maxOpenConns" yaml:"maxOpenConns"`
	//	最大空闲连接数
	MaxIdleConns int `json:"maxIdleConns" yaml:"maxIdleConns"`
	// 连接最大存活时间
	ConnMaxLifeTime int `json:"connMaxLifeTime" yaml:"connMaxLifeTime"`
	// 表名前缀
	Prefix string `json:"prefix" yaml:"prefix"`
	// 是否开启全局禁用复数，true表示开启
	Singular bool `json:"singular" yaml:"singular"`
	// 是否启用日志
	EnableLog bool `json:"enableLog" yaml:"enableLog"`
	// 日志模式
	LogMode string `json:"logMode" yaml:"logMode"`
	// 是否启用缓存
	// EnableCache bool `json:"enableCache" yaml:"enableCache"`

}

type GormDBConfig struct {
	// mysql数据库配置
	MysqlDB MysqlDBConfig `json:"mysqlDB" yaml:"mysqlDB"`
	// sqlLite数据库配置
	SqlLiteDB SqlLiteDBConfig `json:"sqlLiteDB" yaml:"sqlLiteDB"`
}

// MysqlDBConfig 数据库配置
type MysqlDBConfig struct {
	// Gorm 基本配置
	BaseDB `json:"baseDB" yaml:",inline"  mapstructure:",squash"`
	// 数据库驱动
	Driver string `json:"driver" yaml:"driver"`
}

func (m *MysqlDBConfig) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *MysqlDBConfig) DsnOmitDatabase() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *MysqlDBConfig) GetlogMode() string {
	if m.EnableLog {
		return m.LogMode
	}
	return "silent"
}

// SqlLiteDBConfig 数据库配置
type SqlLiteDBConfig struct {
	// Gorm 基本配置
	BaseDB `json:"baseDB" yaml:"baseDB,inline" mapstructure:",squash"`
	// 数据库驱动
	Driver string `json:"driver" yaml:"driver"`
}

func (s *SqlLiteDBConfig) Dsn() string {
	return s.Path + "?cache=shared&mode=rwc"
}

func (s *SqlLiteDBConfig) DsnOmitDatabase() string {
	return s.Path + "?cache=shared&mode=rwc"
}

func (s *SqlLiteDBConfig) GetlogMode() string {
	if s.EnableLog {
		return s.LogMode
	}
	return "silent"
}
