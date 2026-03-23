package database

import (
	"fmt"
	"log"
	"one-dock/app/config"
	"one-dock/pkgs/console"

	"gorm.io/gorm"
)

func Init(cfg config.DBConfig) *gorm.DB {

	console.Info("数据库连接中...")
	db, err := dbConnect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func dbConnect(cfg config.DBConfig) (db *gorm.DB, err error) {

	driver := cfg.Driver

	switch driver {
	case "mysql":
		db, err = InitMysql(CoonMysqlParams{
			Host:     cfg.Connections.Mysql.Host,
			Port:     cfg.Connections.Mysql.Port,
			User:     cfg.Connections.Mysql.User,
			Password: cfg.Connections.Mysql.Password,
			Database: cfg.Connections.Mysql.Database,
			Prefix:   cfg.Prefix,
		})
		if err == nil {
			console.Info("MySQL数据库连接成功！")
		}
	case "postgres":
		db, err = InitPostgres(CoonPgParams{
			Host:     cfg.Connections.Postgres.Host,
			Port:     cfg.Connections.Postgres.Port,
			User:     cfg.Connections.Postgres.User,
			Password: cfg.Connections.Postgres.Password,
			Database: cfg.Connections.Postgres.Database,
		})
		if err == nil {
			console.Info("Postgres数据库连接成功！")
		}
	case "sqlite":
		prefix := cfg.Prefix
		db, err = InitSqlite(cfg.Connections.Sqlite.Database, prefix)
		if err == nil {
			console.Info("SQLite数据库连接成功！")
		}
	default:
		err = fmt.Errorf("暂不不支持的数据库驱动%s", driver)
	}

	return
}
