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

type MysqlDB struct {
	// Gorm 基本配置
	BaseDB `json:"baseDB" yaml:"baseDB"`
	// 数据库驱动
	Driver string `json:"driver" yaml:"driver"`
}

func (m *MysqlDB) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func (m *MysqlDB) DsnOmitDatabase() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + "?charset=utf8mb4&parseTime=True&loc=Local"
}
