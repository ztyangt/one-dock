package install

type Service interface {
	InstallBySqlite(req *BySqliteRequest) error
	InstallByMysql(req *ByMysqlRequest) error
	InstallByPostgres(req *ByPostgresRequest) error
}

type service struct{}

func newService() Service {
	return &service{}
}

// InstallBySqlite 安装BySqlite
func (s *service) InstallBySqlite(req *BySqliteRequest) error {
	return nil
}

// InstallByMysql 安装ByMysql
func (s *service) InstallByMysql(req *ByMysqlRequest) error {
	return nil
}

// InstallByPostgres 安装ByPostgres
func (s *service) InstallByPostgres(req *ByPostgresRequest) error {
	return nil
}
