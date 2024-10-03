/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-09-19 22:19:55
 * @LastEditTime: 2024-10-01 22:44:41
 */
package config

import (
	"reflect"

	"github.com/spf13/viper"
)

type SystemConfig struct {
	ZapLog ZapLogConfig `json:"zapLog" yaml:"zapLog" mapstructure:"zapLog"`

	// gorm
	GormDB GormDBConfig `json:"gormDB" yaml:"gormDB" mapstructure:"gormDB"`
}

func (sc *SystemConfig) ToViper(v *viper.Viper) {
	flattenStruct(v, reflect.ValueOf(sc).Elem(), "")
}

// 递归方法，将结构体的字段平铺到 viper 中
func flattenStruct(v *viper.Viper, value reflect.Value, prefix string) {
	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i) // 获取字段类型信息
		fieldValue := value.Field(i)   // 获取字段值

		key := field.Tag.Get("mapstructure") // 获取 mapstructure 标签
		if key == "" {
			key = field.Tag.Get("yaml") // 如果没有 mapstructure 标签，尝试使用 yaml 标签
		}
		if key == "" {
			key = field.Name
		}

		// 如果当前字段是结构体，递归处理
		if fieldValue.Kind() == reflect.Struct {
			if key == ",squash" { // 如果字段有 ,squash 标签，则拍平字段,推到上
				key = prefix
			} else if prefix != "" {
				key = prefix + "." + key // 组合键
			}
			flattenStruct(v, fieldValue, key) // 继续递归，使用当前字段的 key 作为前缀
		} else {
			if prefix != "" {
				key = prefix + "." + key // 组合键
			}
			// 处理基础类型字段
			v.Set(key, fieldValue.Interface()) // 使用前缀组合键

		}
	}
}
