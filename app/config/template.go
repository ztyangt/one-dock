package config

const configTemplate = `# ===========================================
# 系统配置文件
# ===========================================

#【应用配置】
app:
  port: 9527              # 运行端口
  name: OneDock           # 应用名称
  description: 工作与生活，都有一处归栈  # 应用描述
  version: 0.0.0            # 应用版本

#【JWT 配置】
jwt:
  secret: your-secret-key-here      # JWT 密钥，生产环境请修改
  issuer: one-dock          # JWT 签发人
  subject: user-login               # JWT 主题
  expire: 7200                      # JWT 过期时间（秒）

#【数据库配置】
db:
  driver: sqlite          # 数据库驱动：mysql, postgres, sqlite
  connections:
    mysql:
      host: localhost
      port: "3306"
      user: root
      password: your-password
      database: your-database
    postgres:
      host: localhost
      port: "5432"
      user: postgres
      password: your-password
      database: your-database
    sqlite:
      database: ./database.db

#【日志配置】
log:
  on: true      # 是否开启日志
  size: 2       # 单个日志文件大小（MB）
  age: 7        # 日志文件保留天数
  backups: 10   # 日志文件最大备份数量
`
