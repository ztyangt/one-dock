package user

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

// Repository 定义接口，供外部或Service调用
type Repository interface {
	Create(ctx context.Context, user *UserModel) error
	GetByID(ctx context.Context, id int64) (*UserModel, error)
	GetByCondition(ctx context.Context, conditions map[string]any) (*UserModel, error)
}

// repository 私有结构体，外部无法直接实例化
type repository struct {
	db *gorm.DB
}

// newRepository 私有构造函数，仅在模块内部(router.go)被调用
func newRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Create 创建用户
func (r *repository) Create(ctx context.Context, user *UserModel) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// GetByID 根据ID获取用户
func (r *repository) GetByID(ctx context.Context, id int64) (*UserModel, error) {
	var u UserModel
	err := r.db.WithContext(ctx).First(&u, id).Error

	// 判断错误类型
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在！")
	}
	return &u, err
}

// GetByCondition 根据条件获取用户
func (r *repository) GetByCondition(ctx context.Context, conditions map[string]any) (*UserModel, error) {
	var u UserModel
	err := r.db.WithContext(ctx).Where(conditions).First(&u).Error

	// 判断错误类型
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在！")
	}
	return &u, err
}
