/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-10-01 10:56:21
 * @LastEditTime: 2024-10-01 22:11:17
 */
package config

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func TestToViper(t *testing.T) {
	tests := []struct {
		name     string
		input    SystemConfig
		expected map[string]interface{}
	}{
		{
			name: "TestToViper",
			input: SystemConfig{
				ZapLog: ZapLogConfig{
					Development:       true,
					DisableCaller:     true,
					DisableStacktrace: true,
					ZapCores: []ZapCore{
						{
							Level: zap.NewAtomicLevel(),
							Name:  "console",
							CustomConfig: CustomConfig{
								MessageKey:    "msg",
								TimeKey:       "ts",
								LevelKey:      "level",
								NameKey:       "logger",
								CallerKey:     "caller",
								FunctionKey:   "func",
								StacktraceKey: "stacktrace",
							},
						},
					},
				},
				GormDB: GormDBConfig{
					MysqlDB: MysqlDBConfig{
						Driver: "mysql",
						BaseDB: BaseDB{
							Host:     "localhost",
							Port:     "3306",
							User:     "root",
							Password: "root",
							Path:     "test",
						},
					},
				},
			},
			expected: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			viper.Reset() // Reset viper before the test
			v := viper.New()
			tt.input.ToViper(v)
			v.WriteConfigAs("config-test.yaml")
			for key, expectedValue := range tt.expected {
				actualValue := viper.Get(key)
				if !reflect.DeepEqual(actualValue, expectedValue) {
					t.Errorf("For key %s, expected %v but got %v", key, expectedValue, actualValue)
				}
			}
		})
	}

}
