package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type CoonPgParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func InitPostgres(params CoonPgParams) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			params.Host,
			params.User,
			params.Password,
			params.Database,
			params.Port,
		),
	}),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent), // 关闭终端显示查询信息
		})
	if err != nil {
		err = fmt.Errorf("postgres数据库连接失败: %w", err)
		return
	}
	return
}
