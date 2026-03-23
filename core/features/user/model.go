package user

import (
	"fmt"

	"one-dock/core/comm"
	"one-dock/pkgs/console"
	"one-dock/pkgs/utils"

	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type UserModel struct {
	Id          any             `gorm:"type:bigint; size:64; primaryKey; comment:主键; not null" json:"id"` // 主键
	Nickname    string          `gorm:"size:50; comment:昵称;not null" json:"nickname"`
	Account     string          `gorm:"size:50; comment:账号;unique; not null" json:"account"`
	Email       string          `gorm:"size:50; comment:邮箱;unique;" json:"email"`
	Password    string          `gorm:"size:255; comment:密码;not null" json:"-" `
	Avatar      string          `gorm:"size:255; comment:头像;" json:"avatar"`
	Description string          `gorm:"size:512; comment:描述;" json:"description"`
	Gender      comm.UserGender `gorm:"size:1;comment:性别;default:2;" json:"gender"`
	Status      comm.UserStatus `gorm:"size:1;comment:用户状态;default:1;" json:"status"`
	Role        comm.UserRole   `gorm:"size:1; comment:角色;" json:"role"`
	LastIP      string          `gorm:"size:20; comment:最后登录IP;" json:"last_ip"`
	LoginCount  int64           `gorm:"size:64; comment:登录次数; default:0;" json:"login_count" `
	LoginAt     int64           `gorm:"size:64; comment:登录时间; default:Null;" json:"login_at"`
	comm.BaseModel
}

func (UserModel) TableName() string {
	return "user"
}

func InitUser(db *gorm.DB) error {
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		return fmt.Errorf("数据表User迁移失败: %w", err)
	}
	return createDefaultUser(db)
}

func (u *UserModel) BeforeSave(tx *gorm.DB) (err error) {

	if tx.Statement.Changed("password") {
		newPassword := cast.ToStringMap(tx.Statement.Dest)["password"]
		if newPassword != "" {
			u.Password = utils.Password.Create(newPassword)
		}
	}

	return
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {

	// 生成id
	if u.Id == 0 || u.Id == nil {
		u.Id = comm.SnowNode.Generate().Int64()
	}

	// 检查账号是否已经存在
	var count int64
	tx.Model(&UserModel{}).Unscoped().
		Where("account", u.Account).
		Count(&count)
	if count > 0 {
		err = fmt.Errorf("账号已存在！")
	}

	// 检查邮箱是否已经存在
	var countEmail int64
	tx.Model(&UserModel{}).Unscoped().
		Where("email", u.Email).
		Count(&countEmail)
	if countEmail > 0 {
		err = fmt.Errorf("邮箱已存在！")
	}

	return
}

func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	u.Id = cast.ToString(u.Id)
	return
}

// createDefaultUser 创建默认用户
func createDefaultUser(db *gorm.DB) error {

	// 检查用户数量
	var count int64
	db.Model(&UserModel{}).Unscoped().Count(&count)
	if count > 0 {
		return nil
	}

	pwd := "123456"

	defaultUser := &UserModel{
		Nickname: "admin",
		Account:  "admin",
		Email:    "admin@123.com",
		Role:     comm.UserRoleAdmin,
		Password: utils.Password.Create(pwd),
	}

	err := db.Create(defaultUser).Error
	if err != nil {
		return err
	}

	console.Log(fmt.Sprintf("默认用户创建成功, 账号: %s, 密码: %s", defaultUser.Account, pwd))

	return nil
}
