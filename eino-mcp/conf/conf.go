package conf

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	// 全局配置实例
	AppConfig Config
	// 命令行参数
	configPath *string
)

// Config 配置
type Config struct {
	DeekSeek struct {
		APIKey  string `mapstructure:"api_key"`  // 模型 API Key
		ModelID string `mapstructure:"model_id"` // 模型 ID
		BaseURL string `mapstructure:"base_url"` // 模型 API Base URL
	} `mapstructure:"DeekSeek"`
}

// Init 初始化配置
func Init() error {
	// 解析命令行参数, 配置文件路径
	configPath = flag.String("config", "", "config file path")

	// 加载配置文件
	_, err := loadConfig(*configPath)
	if err != nil {
		return err
	}
	return nil
}

// GetCfg 获取配置
func GetCfg() *Config {
	return &AppConfig
}

// loadConfig 加载配置
func loadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// 设置配置文件路径
	v.SetConfigType("yaml")

	// 指定配置文件路径
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.AddConfigPath(".")
	}

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Printf("loadConfig failed, read config err: %v", err)
		return nil, fmt.Errorf("loadConfig failed, read config err: %v", err)
	}

	// 监听配置变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(&AppConfig); err != nil {
			log.Printf("loadConfig failed, unmarshal config err: %v", err)
		}
	})

	// 解析配置
	if err := v.Unmarshal(&AppConfig); err != nil {
		log.Printf("loadConfig failed, unmarshal config err: %v", err)
		return nil, fmt.Errorf("loadConfig failed, unmarshal config err: %v", err)
	}
	return &AppConfig, nil
}
