package config

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	DeekSeek struct {
		APIKey  string `mapstructure:"api_key"`  // 模型 API Key
		ModelID string `mapstructure:"model_id"` // 模型 ID
		BaseURL string `mapstructure:"base_url"` // 模型 API Base URL
	} `mapstructure:"DeekSeek"`
}

var (
	// 全局配置实例
	AppConfig Config
	// 命令行参数
	configPath *string
)

// Init 初始化配置系统
func Init() error {
	// 定义命令行参数
	configPath = flag.String("config", "", "配置文件路径")

	// 如果程序已经调用过 flag.Parse() 则不再重复调用
	if !flag.Parsed() {
		flag.Parse()
	}

	// 加载配置
	_, err := loadConfig(*configPath)
	if err != nil {
		return err
	}
	return nil
}

// GetCfg 获取全局配置实例
func GetCfg() *Config {
	return &AppConfig
}

// 内部方法：加载配置文件
func loadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// 设置配置文件类型
	v.SetConfigType("yaml")

	// 如果指定了配置路径，使用指定的配置文件
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// 否则从当前目录加载 config.yaml
		v.SetConfigName("config")
		v.AddConfigPath(".")
	}

	// 支持环境变量覆盖
	v.SetEnvPrefix("EINO")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("loadConfig failed, read fail err: %v", err)
	}
	log.Printf("loadConfig debug, success load file from %s", v.ConfigFileUsed())

	// 监听配置文件变化并热重载
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("loadConfig debug, file has changeded: %s", e.Name)
		if err := v.Unmarshal(&AppConfig); err != nil {
			log.Printf("loadConfig failed, load file err: %v", err)
		}
	})

	// 解析到结构体
	if err := v.Unmarshal(&AppConfig); err != nil {
		return nil, fmt.Errorf("loadConfig failed, Unmarshal file err: %v", err)
	}
	return &AppConfig, nil
}
