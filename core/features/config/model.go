package config

import (
	"fmt"
	"one-dock/core/comm"
	"one-dock/pkgs/utils"

	"gorm.io/gorm"
)

type ConfigModel struct {
	ID        int64  `gorm:"primaryKey;autoIncrement;comment:主键" json:"id"`
	ConfigKey string `gorm:"size:256;comment:配置键;unique" json:"configKey"`
	Name      string `gorm:"size:256;comment:配置名称" json:"name"`
	Value     string `gorm:"type:text;comment:配置值" json:"value"`
	Public    bool   `gorm:"comment:是否公开" json:"public"`
	comm.BaseModel
}

func (ConfigModel) TableName() string {
	return "config"
}

func InitConfig(db *gorm.DB) error {
	err := db.AutoMigrate(&ConfigModel{})
	if err != nil {
		return fmt.Errorf("数据表Config迁移失败: %w", err)
	}
	return createDefaultConfig(db)
}

func createDefaultConfig(db *gorm.DB) error {
	defaultConfigs := []ConfigModel{
		{
			ConfigKey: "SYSTEM-CONFIG",
			Name:      "系统配置",
			Public:    true,
			Value: utils.Json.Encode(map[string]any{
				"title":       "OneDock",
				"description": "工作与生活，都有一处归栈。",
				"logo":        "/static/logo/logo.svg",
				"favicon":     "/favicon.svg",
			}),
		},
		{
			ConfigKey: "SYSTEM-EMAIL",
			Name:      "邮件配置",
			Public:    false,
			Value: utils.Json.Encode(map[string]any{
				"from":     "",
				"nickname": "",
				"username": "",
				"host":     "",
				"port":     "",
				"password": "",
			}),
		},
	}

	for _, cfg := range defaultConfigs {
		// 通过 Where 查找 ConfigKey，如果不存在，则用 cfg 的数据进行 Create。
		err := db.Where(ConfigModel{ConfigKey: cfg.ConfigKey}).FirstOrCreate(&cfg).Error
		if err != nil {
			// 如果创建或查询失败，向上抛出错误，终止程序或记录日志
			return fmt.Errorf("初始化默认配置 [%s] 失败: %w", cfg.ConfigKey, err)
		}
	}

	return nil
}
