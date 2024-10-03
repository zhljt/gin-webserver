/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-10-01 23:34:19
 * @LastEditTime: 2024-10-02 01:23:38
 */

package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 自定义解码钩子函数
func ecodeHookFuncType() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		switch t {
		case reflect.TypeOf(zap.AtomicLevel{}):
			if f.Kind() == reflect.String {
				if str, ok := data.(string); ok {
					var zapLevel zap.AtomicLevel
					if err := zapLevel.UnmarshalText([]byte(str)); err != nil {
						return nil, err
					}
					return zapLevel, nil
				}
			}
			// 源类型 不为 strint
			return nil, fmt.Errorf("expected string for zap.AtomicLevel, got %T", data)
		default:
			// 对于其他类型，使用默认行为
			return data, nil
		}
	}
}
func TestReadYaml(t *testing.T) {
	tests := []struct {
		name    string
		want    *SystemConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestReadYaml",
			want:    &SystemConfig{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadYaml("testA.yaml")
			fmt.Printf("read yaml file: %v\n", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteYaml(t *testing.T) {
	type args struct {
		filePath string
		config   *SystemConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "TestWriteYaml",
			args: args{
				filePath: "testA.yaml",
				config:   &SystemConfig{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteYaml(tt.args.filePath, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("WriteYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadAndWriteYaml(t *testing.T) {
	tests := []struct {
		name    string
		want    *SystemConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestReadAndWriteYaml",
			want:    &SystemConfig{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadYaml("config.yaml")
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadYaml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadYaml() = %v, want %v", got, tt.want)
			}

			got.GormDB.MysqlDB.BaseDB.Password = "12345678"
			got.GormDB.MysqlDB.BaseDB.Host = "127.0.0.1"
			got.GormDB.MysqlDB.BaseDB.Port = "23306"

			if err := WriteYaml("testA.yaml", got); (err != nil) != tt.wantErr {
				t.Errorf("WriteYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadAndWriteViper(t *testing.T) {
	tests := []struct {
		name    string
		want    *SystemConfig
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TestReadAndWriteYaml",
			want:    &SystemConfig{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := viper.New()
			v.SetConfigFile("testA.yaml")
			v.SetConfigType("yaml")
			// viper.AutomaticEnv()
			if err := v.ReadInConfig(); err != nil {
				fmt.Printf("Error reading config file : %s", "testA.yaml")
				panic(err)
			}
			got := &SystemConfig{}
			if err := v.Unmarshal(got, viper.DecodeHook(ecodeHookFuncType())); err != nil {
				panic(err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadYaml() = %v, want %v", got, tt.want)
			}

			got.GormDB.MysqlDB.BaseDB.Password = "7777777777777"
			got.GormDB.MysqlDB.BaseDB.Host = "127.0.0.1"
			got.GormDB.MysqlDB.BaseDB.Port = "23306"

			if err := WriteYaml("testA.yaml", got); (err != nil) != tt.wantErr {
				t.Errorf("WriteYaml() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
