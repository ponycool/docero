package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config 应用程序的配置结构体
type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
		Mode string `mapstructure:"mode"`
	} `mapstructure:"server"`
	// 添加 Converter 字段
	Converter struct {
		LibreofficePath string `mapstructure:"libreoffice_path"`
		OutputDir       string `mapstructure:"output_dir"`
		UploadDir       string `mapstructure:"upload_dir"`
	} `mapstructure:"converter"`
	Log struct {
		// slog configuration
		Level  string `mapstructure:"level"`
		Format string `mapstructure:"format"` // "text" or "json"

		// lumberjack configuration
		Filename   string `mapstructure:"filename"`
		MaxSize    int    `mapstructure:"max_size"`    // megabytes
		MaxBackups int    `mapstructure:"max_backups"` // number of old log files to retain
		MaxAge     int    `mapstructure:"max_age"`     // days to retain old log files
		Compress   bool   `mapstructure:"compress"`    // whether to compress old log files
	} `mapstructure:"log"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")   // 配置文件名称 (不带扩展名)
	viper.SetConfigType("yaml")     // 配置文件类型
	viper.AddConfigPath("./config") // 查找配置文件的路径
	viper.AddConfigPath(".")        // 也可以在当前目录查找

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
