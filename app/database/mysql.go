package database

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type CoonMysqlParams struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Prefix   string
}

func InitMysql(params CoonMysqlParams) (db *gorm.DB, err error) {
	db, err = connectMysql(params.Host,
		params.Port,
		params.User,
		params.Database,
		params.Password,
		"utf8mb4",
		params.Prefix)
	return
}

func connectMysql(host, port, username, database, password, charset, prefix string) (*gorm.DB, error) {
	conn, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=5s",
				username,
				password,
				host,
				port,
				database,
				charset),
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   cast.ToString(prefix), // 表名前缀
				SingularTable: true,                  // 使用单数表名
			},
			Logger: logger.Default.LogMode(logger.Silent), // 关闭终端显示查询信息
		})
	if err != nil {
		err = fmt.Errorf("mysql数据库连接失败: %w", err)
		return nil, err
	}

	sqlDB, _ := conn.DB()
	sqlDB.SetMaxIdleConns(10)           //  空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)          //  打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 连接可复用的最大时间

	return conn, err

}
