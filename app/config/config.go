package config

import "github.com/spf13/viper"

type Cfg struct {
	Viper *viper.Viper `json:"-"`
	App   AppConfig    `mapstructure:"app" json:"app"`
	JWT   JWTConfig    `mapstructure:"jwt" json:"jwt"`
	DB    DBConfig     `mapstructure:"db" json:"db"`
	Log   LogConfig    `mapstructure:"log" json:"log"`
}

type AppConfig struct {
	Port string `mapstructure:"port" json:"port"`
	//Name        string `mapstructure:"name" json:"name"`
	//Description string `mapstructure:"description" json:"description"`
	Version string `mapstructure:"version" json:"version"`
}

type JWTConfig struct {
	Secret  string `mapstructure:"secret" json:"secret"`
	Issuer  string `mapstructure:"issuer" json:"issuer"`
	Subject string `mapstructure:"subject" json:"subject"`
	Expire  int    `mapstructure:"expire" json:"expire"`
}

type DBConfig struct {
	Driver      string        `mapstructure:"driver" json:"driver"`
	Prefix      string        `mapstructure:"prefix" json:"prefix"`
	Connections DBConnections `mapstructure:"connections" json:"connections"`
}

type DBConnections struct {
	Mysql    MysqlConfig    `mapstructure:"mysql" json:"mysql"`
	Postgres PostgresConfig `mapstructure:"postgres" json:"postgres"`
	Sqlite   SqliteConfig   `mapstructure:"sqlite" json:"sqlite"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     string `mapstructure:"port" json:"port"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
	Database string `mapstructure:"database" json:"database"`
}

type SqliteConfig struct {
	Database string `mapstructure:"database" json:"database"`
}

type LogConfig struct {
	On      bool `mapstructure:"on" json:"on"`
	Size    int  `mapstructure:"size" json:"size"`
	Age     int  `mapstructure:"age" json:"age"`
	Backups int  `mapstructure:"backups" json:"backups"`
}
