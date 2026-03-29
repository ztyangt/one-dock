package config

import (
	"fmt"
)

func validateConfig(cfg *Cfg) error {
	if err := validateAppConfig(&cfg.App); err != nil {
		return fmt.Errorf("应用配置验证失败：%w", err)
	}

	if err := validateDBConfig(&cfg.DB); err != nil {
		return fmt.Errorf("数据库配置验证失败：%w", err)
	}

	if err := validateLogConfig(&cfg.Log); err != nil {
		return fmt.Errorf("日志配置验证失败：%w", err)
	}

	return nil
}

func validateAppConfig(cfg *AppConfig) error {
	if cfg.Port == "" {
		return fmt.Errorf("端口不能为空")
	}
	if cfg.Version == "" {
		return fmt.Errorf("应用版本不能为空")
	}
	return nil
}

func validateDBConfig(cfg *DBConfig) error {
	if cfg.Driver == "" {
		return fmt.Errorf("数据库驱动不能为空")
	}

	validDrivers := map[string]bool{
		"mysql":    true,
		"postgres": true,
		"sqlite":   true,
	}

	if !validDrivers[cfg.Driver] {
		return fmt.Errorf("不支持的数据库驱动：%s", cfg.Driver)
	}

	switch cfg.Driver {
	case "mysql":
		if cfg.Connections.Mysql.Host == "" {
			return fmt.Errorf("MySQL 主机地址不能为空")
		}
		if cfg.Connections.Mysql.Database == "" {
			return fmt.Errorf("MySQL 数据库名不能为空")
		}
	case "postgres":
		if cfg.Connections.Postgres.Host == "" {
			return fmt.Errorf("PostgreSQL 主机地址不能为空")
		}
		if cfg.Connections.Postgres.Database == "" {
			return fmt.Errorf("PostgreSQL 数据库名不能为空")
		}
	case "sqlite":
		if cfg.Connections.Sqlite.Database == "" {
			return fmt.Errorf("SQLite 数据库路径不能为空")
		}
	}

	return nil
}

func validateLogConfig(cfg *LogConfig) error {
	if cfg.Size <= 0 {
		return fmt.Errorf("日志文件大小必须大于 0")
	}
	if cfg.Age <= 0 {
		return fmt.Errorf("日志保留天数必须大于 0")
	}
	if cfg.Backups < 0 {
		return fmt.Errorf("日志备份数量不能为负数")
	}
	return nil
}
