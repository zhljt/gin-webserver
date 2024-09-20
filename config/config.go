/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-19 22:19:55
 * @LastEditTime: 2024-09-20 17:24:15
 */
package config

type SystemConfig struct {
	ZapLog ZapLog `json:"zapLog" mapstructure:"zapLog"`

	// gorm
	GormDB GormDB `json:"gormDB" mapstructure:"gormDB"`
}
