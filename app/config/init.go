package config

import (
	"errors"
	"fmt"
	"log"
	"one-dock/pkgs/console"
	"os"

	"github.com/spf13/viper"
)

func Init() *Cfg {
	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("config")
	config.SetConfigType("yaml")

	readErr := config.ReadInConfig()
	if readErr != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(readErr, &configFileNotFoundError) {
			if err := createDefaultConfig(); err != nil {
				log.Fatalf("创建默认配置文件失败：%v", err)
			}
			if err := config.ReadInConfig(); err != nil {
				log.Fatalf("重新读取配置文件失败：%v", err)
			}
		}
	}

	var cfg Cfg
	if err := config.Unmarshal(&cfg); err != nil {
		log.Fatalf("解析配置文件失败：%v", err)
	}

	if err := validateConfig(&cfg); err != nil {
		log.Fatalf("%v", err)
	}

	config.WatchConfig()

	console.Info("配置文件加载成功!")

	cfg.Viper = config
	return &cfg
}

func (cfg *Cfg) Reload() error {
	console.Info("重新加载配置文件...")
	if err := cfg.Viper.ReadInConfig(); err != nil {
		return fmt.Errorf("重新读取配置文件失败：%w", err)
	}

	viper := cfg.Viper

	if err := cfg.Viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("解析重新配置文件失败：%w", err)
	}

	cfg.Viper = viper

	return nil
}

func createDefaultConfig() error {
	console.Info("配置文件不存在，创建默认配置文件...")

	file, err := os.Create("./config.yaml")
	if err != nil {
		return fmt.Errorf("创建配置文件失败：%w", err)
	}

	if _, err := file.WriteString(configTemplate); err != nil {
		return fmt.Errorf("写入配置文件失败：%w", err)
	}
	defer file.Close()

	return nil
}
