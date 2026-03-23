package database

import (
	"fmt"

	"github.com/spf13/cast"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitSqlite(database, prefix string) (db *gorm.DB, err error) {

	// 创建数据库文件
	db, err = gorm.Open(sqlite.Open(database), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cast.ToString(prefix), // 表名前缀
			SingularTable: true,                  // 使用单数表名
		},
		Logger: logger.Default.LogMode(logger.Silent), // 关闭终端显示查询信息
	})

	if err != nil {
		err = fmt.Errorf("sqlite数据库连接失败: %w", err)
		return nil, err
	}
	return
}
