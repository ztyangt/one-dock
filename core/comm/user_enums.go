package comm

// UserGender Gender 性别
type UserGender int64

const (
	MALE    UserGender = iota + 1 // 男
	FEMALE                        // 女
	UNKNOWN                       // 保密
)

// UserStatus 用户状态
type UserStatus int64

const (
	UserStatusActive UserStatus = iota + 1 //正常
	UserStatusBanned                       // 封禁
)

// UserRole 用户角色
type UserRole int64

const (
	UserRoleAdmin UserRole = iota + 1 // 管理员
	UserRoleUser                      // 普通用户
)
