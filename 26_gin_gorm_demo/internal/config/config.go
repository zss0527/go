package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 结构体用于存储应用的配置信息
type Config struct {
	Server struct {
		Address string `yaml:"address"`
	} `yaml:"server"`
	Database struct {
		DSN string `yaml:"dsn"`
	} `yaml:"database"`
}

// LoadConfig 函数从配置文件中加载配置信息，如果文件不存在则提供默认配置
func LoadConfig() Config {
	var conf Config
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("config.yaml 文件不存在，使用默认配置")
			// 设置默认配置
			conf.Server.Address = ":8082"
			conf.Database.DSN = "default.db"
			return conf
		}
		fmt.Printf("读取配置文件错误: %v", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("解析配置文件错误: %v", err)
		os.Exit(1)
	}
	return conf
}
