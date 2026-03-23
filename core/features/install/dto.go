package install

// BySqliteRequest 安装BySqlite请求参数
type BySqliteRequest struct {
	DBName string `json:"db_name" validate:"required,alphanum" alias:"数据库名称"`
}

// ByMysqlRequest 安装ByMysql请求参数
type ByMysqlRequest struct {
}

// ByPostgresRequest 安装ByPostgres请求参数
type ByPostgresRequest struct{}
