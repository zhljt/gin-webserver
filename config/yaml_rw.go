/*
 * @Author: Lin Jin Ting
 * @LastEditors: Lin Jin Ting
 * @Email: ljt930@gmail.com
 * @Description:
 * @Date: 2024-10-01 23:34:19
 * @LastEditTime: 2024-10-02 00:16:47
 */
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYaml(filePath string) (*SystemConfig, error) {
	// TODO: implement this function
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	// 解析YAML内容到结构体中
	cfg := &SystemConfig{}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func WriteYaml(filePath string, config *SystemConfig) error {
	// 将结构体编码为YAML格式
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	// 将YAML数据写入文件
	return os.WriteFile(filePath, data, os.ModePerm)
}
